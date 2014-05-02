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
	LINETYPE_PROMPT    LineType      = iota
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

// Generic menu line definition
type MenuLine struct {
	text     string
	align    TextAlignment
	linetype LineType
}

// Menu type definition
type Menu struct {
	content   []MenuLine
	commands  map[int]func()
	prompt    MenuLine
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
func NewMenu(title string, prompt string, border bool, padding_x int, padding_y int, margin_x int, margin_y int) *Menu {
	menu := Menu{width: 0, border: border, padding_x: padding_x, padding_y: padding_y, margin_x: margin_x, margin_y: margin_y}
	menu.SetTitle(title)
	menu.SetPrompt(prompt)
	menu.commands = make(map[int]func())
	return &menu
}

// Prints the menu to the standard output
func (m Menu) Display() {
	menuWidth := m.GetWidth()
	emptyLine := MenuLine{text: CHAR_EMPTY, align: TXTALIGN_CENTER, linetype: LINETYPE_EMPTY}
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
		// Prints vertical padding (space between border and prompt)
	}

	for i := 0; i < m.GetPaddingY(); i++ {
		fmt.Println(CHAR_EMPTY)
	}
	// Print prompt
	fmt.Print(m.FormatString(m.GetPrompt()))
}

// Returns formated string including padding, border etc...
func (m Menu) FormatString(line MenuLine) string {
	menuWidth := m.GetWidth()
	emptyChars := menuWidth - line.GetLength()
	formatedString := ""

	if line.GetType() == LINETYPE_PROMPT {
		formatedString += strings.Repeat(CHAR_EMPTY, m.margin_x)
		if m.border {
			formatedString += CHAR_EMPTY
		}
		formatedString += strings.Repeat(CHAR_EMPTY, m.padding_x)
		formatedString += line.GetText()
		return formatedString
	}

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

// Returns prompt of the Menu
func (m Menu) GetPrompt() MenuLine {
	return m.prompt
}

/**********************************/
/* MENU CONTENT RELATED FUNCTIONS */
/**********************************/

// Set TitleLine of the Menu
func (m *Menu) SetTitle(title string) {
	newline := MenuLine{text: title, align: TXTALIGN_CENTER, linetype: LINETYPE_TITLE}
	m.content = append(m.content, newline)
}

// Set prompt of the Menu
func (m *Menu) SetPrompt(prompt string) {
	m.prompt = MenuLine{text: prompt, align: TXTALIGN_LEFT, linetype: LINETYPE_PROMPT}
}

// Add a TextLine to the Menu
func (m *Menu) AddTextLine(text string, align TextAlignment) {
	newline := MenuLine{text: text, align: align, linetype: LINETYPE_TEXT}
	m.content = append(m.content, newline)
}

// Add a blank line to the Menu
func (m *Menu) AddEmptyLine() {
	newline := MenuLine{text: CHAR_EMPTY, align: TXTALIGN_CENTER, linetype: LINETYPE_EMPTY}
	m.content = append(m.content, newline)
}

// Add a command line to the Menu
func (m *Menu) AddCommandLine(key int, text string, command func(), align TextAlignment) {
	m.commands[key] = command
	text = fmt.Sprintf("%d %s %s", key, CHAR_SEPARATOR, text)
	newline := MenuLine{text: text, align: align, linetype: LINETYPE_COMMAND}
	m.content = append(m.content, newline)
}

/******************************/
/* MENULINE RELATED FUNCTIONS */
/******************************/

// Returns MenuLine type
func (l MenuLine) GetType() LineType {
	return l.linetype
}

// Returns TextLine text content
func (l MenuLine) GetText() string {
	return l.text
}

// Returns TextLine Alignment
func (l MenuLine) GetAlignment() TextAlignment {
	return l.align
}

// Returns Textline text length
func (l MenuLine) GetLength() int {
	return len(l.text)
}
