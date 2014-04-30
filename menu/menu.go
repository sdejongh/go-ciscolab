/*
MENU package
Generates and display the command-line menu
*/

package menu

const (
	// Some characters used to draw frames around menus
	CHAR_CORNER                   = "+"
	CHAR_HLINE                    = "-"
	CHAR_VLINE                    = "|"
	CHAR_SEPARATOR                = "-"
	CHAR_EMPTY                    = " "
	TXTALIGN_LEFT   TextAlignment = iota
	TXTALIGN_CENTER TextAlignment = iota
	TXTALIGN_RIGHT  TextAlignment = iota
)

// Text alignment type definition (center/left/right)
type TextAlignment int

// Generic menu line definition
type MenuLine interface {
	GetText()
	GetLength()
	GetAlignment()
}

// Command line definition
type CommandLine struct {
	key     int
	text    string
	command func()
	align   TextAlignment
}

// Simple text line definition
type TextLine struct {
	content string
	align   TextAlignment
}

// Separator line definition
type SeparatorLine string

// Empty line definition
type EmptyLine string

// Menu type definition
type Menu struct {
	content   []MenuLine
	width     int
	border    bool
	padding_x int
	padding_y int
	margin_x  int
	margin_y  int
}

/* MENU RELATED FUNCTIONS */

// Prints the menu to the standard output
func (m Menu) Display() {}

// Returns menu width
func (m Menu) GetWidth() int {
	return m.width
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

/* TEXTLINE RELATED FUNCTIONS */

// Returns TextLine text contend
func (t TextLine) GetText() string {
	return t.GetText()
}

// Returns TextLine Alignment
func (t TextLine) GetAlignment() TextAlignment {
	return t.align
}

// Returns Textline text length
func (t TextLine) GetLength() int {
	return len(t.content)
}

/* COMMANDLINE RELATED FUNCTIONS */

// Returns CommandLine text description
func (c CommandLine) GetText() string {
	return c.GetText()
}

// Returns CommandLine text length
func (t CommandLine) GetLength() int {
	return len(t.text)
}

// Returns CommandLine text alignment
func (t CommandLine) GetAlignment() TextAlignment {
	return t.align
}

/* SEPARATOR LINE RELATED FUNCTIONS */

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

/* EMPTY LINE RELATED FUNCTIONS */

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
