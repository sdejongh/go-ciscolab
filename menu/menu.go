/*
MENU package
Generates and display the command-line menu
*/

package menu

import (
	"fmt"
	"strings"
)

const (
	// Some characters used to draw frames around menus
	LINETYPE_TEXT      LineType      = iota
	LINETYPE_COMMAND   LineType      = iota
	LINETYPE_SEPARATOR LineType      = iota
	LINETYPE_EMPTY     LineType      = iota
	LINETYPE_TITLE     LineType      = iota
	CHAR_CORNER                      = "+"
	CHAR_HLINE                       = "-"
	CHAR_VLINE                       = "|"
	CHAR_SEPARATOR                   = "-"
	CHAR_EMPTY                       = " "
	CHAR_NEWLINE                     = "\n"
	BORDER                           = true
	NOBORDER                         = false
	TXTALIGN_LEFT      TextAlignment = iota
	TXTALIGN_CENTER    TextAlignment = iota
	TXTALIGN_RIGHT     TextAlignment = iota
)

// Line type definition
type LineType int

// Text alignment type definition (center/left/right)
type TextAlignment int

// MenuCommand type
type MenuCommand struct {
	key int
	f   func()
}

// Generic menu line definition
type MenuLine interface {
	GetText() string
	GetLength() int
	GetAlignment() TextAlignment
	GetType() LineType
}

// Command line definition
type CommandLine struct {
	text    string
	command MenuCommand
	align   TextAlignment
}

// Simple text line definition
type TextLine struct {
	content string
	align   TextAlignment
}

// Title line definition
type TitleLine string

// Separator line definition
type SeparatorLine string

// Empty line definition
type EmptyLine string

// Menu type definition
type Menu struct {
	content   []MenuLine
	commands  map[int]CommandLine
	width     int
	border    bool
	padding_x int
	padding_y int
	margin_x  int
	margin_y  int
}

/**************************/
/* MENU RELATED FUNCTIONS */
/**************************/

// Menu Type constructor
func NewMenu(title string, border bool, padding_x int, padding_y int, margin_x int, margin_y int) *Menu {
	menu := Menu{width: 0, border: border, padding_x: padding_x, padding_y: padding_y, margin_x: margin_x, margin_y: margin_y}
	menu.SetTitle(title)
	menu.commands = make(map[int]CommandLine)
	return &menu
}

// Prints the menu to the standard output
func (m Menu) Display() {
	menuWidth := m.GetWidth()
	var emptyLine EmptyLine
	hline := strings.Repeat(CHAR_EMPTY, m.margin_x) + CHAR_CORNER + strings.Repeat(CHAR_HLINE, menuWidth) + CHAR_CORNER

	// Print Vertical margin
	for i := 0; i < m.margin_y; i++ {
		fmt.Println(CHAR_EMPTY)
	}

	// Print menu top border if needed
	if m.border {
		fmt.Println(hline)
	}

	// Prints vertical padding (space between border and next line)
	for i := 0; i < m.GetPaddingY(); i++ {
		fmt.Println(m.FormatString(emptyLine))
	}

	// Print each line added to the menu
	for _, line := range m.content {
		fmt.Println(m.FormatString(line))
		switch line.GetType() {
		case LINETYPE_TITLE:
			for i := 0; i < m.GetPaddingY(); i++ {
				fmt.Println(m.FormatString(emptyLine))
			}
			if m.border {
				fmt.Println(hline)
			}
			for i := 0; i < m.GetPaddingY(); i++ {
				fmt.Println(m.FormatString(emptyLine))
			}
		}
	}

	// Prints vertical padding
	for i := 0; i < m.GetPaddingY(); i++ {
		fmt.Println(m.FormatString(emptyLine))
	}

	// Print menu bottom border if needed
	if m.border {
		fmt.Println(hline)
	}

}

// Returns formated string including padding, border etc...
func (m Menu) FormatString(line MenuLine) string {
	menuWidth := m.GetWidth()
	emptyChars := menuWidth - line.GetLength()
	formatedString := ""
	switch line.GetAlignment() {
	case TXTALIGN_LEFT:
		formatedString = strings.Repeat(CHAR_EMPTY, m.padding_x) + line.GetText() + strings.Repeat(CHAR_EMPTY, emptyChars-m.padding_x)
	case TXTALIGN_CENTER:
		formatedString = strings.Repeat(CHAR_EMPTY, int(emptyChars/2)) + line.GetText()
		formatedString += strings.Repeat(CHAR_EMPTY, menuWidth-len(formatedString))
	case TXTALIGN_RIGHT:
		formatedString = strings.Repeat(CHAR_EMPTY, emptyChars-m.padding_x) + line.GetText() + strings.Repeat(CHAR_EMPTY, m.padding_x)
	}
	if m.border {
		formatedString = CHAR_VLINE + formatedString + CHAR_VLINE
	}
	formatedString = strings.Repeat(CHAR_EMPTY, m.margin_x) + formatedString
	return formatedString
}

// Returns menu width (including padding etc., but without border)
func (m Menu) GetWidth() int {
	maxLength := m.width
	for _, line := range m.content {
		length := line.GetLength()
		if length > maxLength {
			maxLength = length
		}
	}
	maxLength += 2 * m.GetPaddingX()
	return maxLength
}

// Sets menu width
func (m *Menu) SetWidth(width int) {
	m.width = width
}

// Returns horizontal padding
func (m Menu) GetPaddingX() int {
	return m.padding_x
}

// Sets horizontal padding
func (m *Menu) SetPaddingX(x int) {
	m.padding_x = x
}

// Returns vertical padding
func (m Menu) GetPaddingY() int {
	return m.padding_y
}

// Sets vertical padding
func (m *Menu) SetPaddingY(y int) {
	m.padding_y = y
}

// Returns horizontal margin
func (m Menu) GetMarginX() int {
	return m.margin_x
}

// Sets horizontal margin
func (m *Menu) SetMarginX(x int) {
	m.margin_x = x
}

// Returns vertical margin
func (m Menu) GetMarginY() int {
	return m.margin_y
}

// Sets vertical margin
func (m *Menu) SetMarginY(y int) {
	m.margin_y = y
}

// Set TitleLine of the Menu
func (m *Menu) SetTitle(title string) {
	newline := TitleLine(title)
	m.content = append(m.content, newline)
}

// Add a TextLine to the Menu
func (m *Menu) AddTextLine(text string, align TextAlignment) {
	newline := TextLine{content: text, align: align}
	m.content = append(m.content, newline)
}

// Add a blank line to the Menu
func (m *Menu) AddEmptyLine() {
	var newline EmptyLine
	m.content = append(m.content, newline)
}

// Add a command line to the Menu
func (m *Menu) AddCommandLine(key int, text string, command func(), align TextAlignment) {
	newcmd := MenuCommand{key: key, f: command}
	newline := CommandLine{text: text, command: newcmd, align: align}
	m.content = append(m.content, newline)
	m.commands[key] = newline
}

/******************************/
/* TEXTLINE RELATED FUNCTIONS */
/******************************/

// Returns LineType LINETYPE_TEXT
func (t TextLine) GetType() LineType {
	return LINETYPE_TEXT
}

// Returns TextLine text content
func (t TextLine) GetText() string {
	return t.content
}

// Returns TextLine Alignment
func (t TextLine) GetAlignment() TextAlignment {
	return t.align
}

// Returns Textline text length
func (t TextLine) GetLength() int {
	return len(t.content)
}

/*******************************/
/* TITLELINE RELATED FUNCTIONS */
/*******************************/

// Returns LineType LINETYPE_TITLE
func (t TitleLine) GetType() LineType {
	return LINETYPE_TITLE
}

// Returns TitleLine text content
func (t TitleLine) GetText() string {
	return string(t)
}

// Returns TitleLine Alignment
func (t TitleLine) GetAlignment() TextAlignment {
	return TXTALIGN_CENTER
}

// Returns Textline text length
func (t TitleLine) GetLength() int {
	return len(t)
}

/*********************************/
/* COMMANDLINE RELATED FUNCTIONS */
/*********************************/

// Returns LineType LINETYPE_COMMAND
func (c CommandLine) GetType() LineType {
	return LINETYPE_COMMAND
}

// Returns CommandLine text description
func (c CommandLine) GetText() string {
	return fmt.Sprintf("%d %s %s", c.command.key, CHAR_SEPARATOR, c.text)
}

// Returns CommandLine text length
func (c CommandLine) GetLength() int {
	return len(c.GetText())
}

// Returns CommandLine text alignment
func (c CommandLine) GetAlignment() TextAlignment {
	return c.align
}

// Returns CommandLine key
func (c CommandLine) GetKey() int {
	return c.command.key
}

/************************************/
/* SEPARATOR LINE RELATED FUNCTIONS */
/************************************/

// Returns LineType LINETYPE_SEPARATOR
func (s SeparatorLine) GetType() LineType {
	return LINETYPE_SEPARATOR
}

// Returns a single separator character
func (s SeparatorLine) GetText() string {
	return CHAR_HLINE
}

// Returns 1 (length of a single char)
func (s SeparatorLine) GetLength() int {
	return 1
}

// Returns TXTALIGN_CENTER, Separtor line will always fill the width of the menu
func (s SeparatorLine) GetAlignment() TextAlignment {
	return TXTALIGN_CENTER
}

/********************************/
/* EMPTY LINE RELATED FUNCTIONS */
/********************************/

// Returns LineType LINETYPE_EMPTY
func (e EmptyLine) GetType() LineType {
	return LINETYPE_EMPTY
}

// Returns a single space character
func (e EmptyLine) GetText() string {
	return CHAR_EMPTY
}

// Returns 1 (length of a single char)
func (e EmptyLine) GetLength() int {
	return 1
}

// Returns TXTALIGN_CENTER, empty line will always fill the width of the menu
func (e EmptyLine) GetAlignment() TextAlignment {
	return TXTALIGN_CENTER
}
