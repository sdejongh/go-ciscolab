package network

import (
	"code.google.com/p/go.crypto/ssh"
	"code.google.com/p/go.crypto/ssh/terminal"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

/*********************/
/* Types definitions */
/*********************/

// SSH Server type definition
type SSHServer struct {
	listeningAddress string
	serverConfig     ssh.ServerConfig
	listener         net.Listener
	clients          []SSHClient
}

// SSH Client linking SSH side with Telnet side
type SSHClient struct {
	clientUser    string
	clientConn    net.Conn
	clientSSHChan ssh.Channel
}

/*******************************/
/* SSHServer related functions */
/*******************************/

// SSHServer type constructor, listeningaddress must be specified as IPADDRESS:PORT (ex: 0.0.0.0:9022)
func NewSSHServer(listeningAddress string, rsaKeyFilePath string) *SSHServer {
	privateKeyBytes := LoadPEMFile(rsaKeyFilePath)
	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatal("Failed to parse private key bytes:", err)
	}
	server := SSHServer{listeningAddress: listeningAddress}
	server.serverConfig = ssh.ServerConfig{

		// Test authentiocation callback
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			// Should use constant-time compare (or better, salt+hash) in
			// a production setting.
			if c.User() == "testuser" && string(pass) == "tiger" {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}
	server.serverConfig.AddHostKey(privateKey)
	return &server
}

// Launches SSH server listening process
func (s *SSHServer) Listen() {

	// Open TCP socket for incoming connections
	listener, err := net.Listen("tcp", s.listeningAddress)
	if err != nil {
		log.Fatal("Could not start listening on"+s.listeningAddress+":", err)
	}
	s.listener = listener
	for {
		// Start accepting incoming TCP connections
		log.Println("Accepting new connection")
		sConn, err := s.listener.Accept()
		if err != nil {
			log.Println("Failed to accept incoming connection:", err)
			continue
		}
		newclient := SSHClient{clientConn: sConn}
		s.clients = append(s.clients, newclient)
		go HandleNewConn(&newclient, &s.serverConfig)
	}

}

// Manage handshake etc.
func HandleNewConn(client *SSHClient, sConfig *ssh.ServerConfig) {
	// Handshake, needed before using the connection
	_, chans, reqs, err := ssh.NewServerConn(client.NetConn(), sConfig)

	if err != nil {
		log.Println("Handshake failed:", err)
		client.NetConn().Close()
	} else {

		// Discards all request made out of the normal process
		go ssh.DiscardRequests(reqs)

		for newChannel := range chans {
			// Only accept "session" type channels for pseudo shell
			if newChannel.ChannelType() != "session" {
				newChannel.Reject(ssh.UnknownChannelType, "Unsupported channel type")
				log.Println("Unknow channel type... rejecting")
				continue
			}
			go HandleClient(client, newChannel)
		}
	}

}

// Handles sshChannel
func HandleClient(c *SSHClient, newChannel ssh.NewChannel) {
	//channel, requests, err := newChannel.Accept()
	channel, _, err := newChannel.Accept()
	c.clientSSHChan = channel
	if err != nil {
		log.Println("Failed to accept channel:", err)
	}
	log.Println("NewClient:", c.clientConn.RemoteAddr())

	/*
		// Handles SSH requests
		go func(in <-chan *ssh.Request) {
			for req := range in {
				// do something ?
				log.Println(req)
			}
		}(requests)
	*/

	term := terminal.NewTerminal(channel, "> ")

	go func() {
		defer channel.Close()
		for {
			line, err := term.ReadLine()
			if err != nil {
				break
			}
			fmt.Println(c.clientConn.RemoteAddr().String()+":", line)
		}
	}()

}

/***********************************/
/* RSAPrivateKey related functions */
/***********************************/

// Loads RSA key from PEM file (openssl generated) and stores bytes in SSHServer.RSAPrivateKey.pemBytes
func LoadPEMFile(filepath string) []byte {
	privateBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Private key file not found:", err)
	}
	return privateBytes
}

/*******************************/
/* SSHCLIENT related functions */
/*******************************/

func (c SSHClient) NetConn() net.Conn {
	return c.clientConn
}

func (c SSHClient) SSHChannel() ssh.Channel {
	return c.clientSSHChan
}
