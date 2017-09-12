package bitbar

import (
	"fmt"
	"strings"
)

// Menu TODO
type Menu struct {
	MenuBar  MenuBar
	DropDown *DropDown
}

// Line TODO
type Line struct {
	text          string
	href          string
	color         string
	font          string
	size          int
	terminal      *bool
	refresh       *bool
	dropDown      *bool
	length        int
	trim          *bool
	alternate     *bool
	emojize       *bool
	ansi          *bool
	bash          string
	param         []string
	templateImage string
	image         string
}

// Style TODO
type Style struct {
	Color string
	Font  string
	Size  int
}

// Options TODO
type Options struct {
	Terminal  bool
	Refresh   bool
	DropDown  bool
	Length    int
	Trim      bool
	Alternate bool
	Emojize   bool
	Ansi      bool
}

// Cmd TODO
type Cmd struct {
	Bash  string
	Param []string
}

// MenuBar TODO
type MenuBar struct {
	Line []*Line
}

// DropDownItem TODO
type DropDownItem interface{}

// DropDown TODO
type DropDown struct {
	Level int
	Line  []DropDownItem
}

// New returns an empty Bitbar menu without any context
func New() Menu {
	return Menu{}
}

// AddMenuBar TODO
func (b *Menu) AddMenuBar(s string) *Line {
	l := new(Line)
	l.text = s
	b.MenuBar.Line = append(b.MenuBar.Line, l)
	return l
}

// AddDropDown TODO
func (b *Menu) AddDropDown() *DropDown {
	b.DropDown = new(DropDown)
	b.DropDown.Level = 0
	return b.DropDown
}

// Add TODO
func (d *DropDown) Add(s string) *Line {
	l := new(Line)
	l.text = s
	d.Line = append(d.Line, l)
	return l
}

// AddDropDown TODO
func (d *DropDown) AddDropDown() *DropDown {
	newDropDown := new(DropDown)
	newDropDown.Level = d.Level + 1
	d.Line = append(d.Line, newDropDown)
	return newDropDown
}

// Style TODO
func (l *Line) Style(s Style) *Line {
	l.color = s.Color
	l.font = s.Font
	l.size = s.Size
	return l
}

// Options TODO
func (l *Line) Options(o Options) *Line {
	l.terminal = &o.Terminal
	l.refresh = &o.Refresh
	l.dropDown = &o.DropDown
	l.length = o.Length
	l.trim = &o.Trim
	l.alternate = &o.Alternate
	l.emojize = &o.Emojize
	l.ansi = &o.Ansi
	return l
}

// Command TODO
func (l *Line) Command(c Cmd) *Line {
	l.bash = c.Bash
	l.param = c.Param
	return l
}

// Href TODO
func (l *Line) Href(s string) *Line {
	l.href = s
	return l
}

// Color TODO
func (l *Line) Color(s string) *Line {
	l.color = s
	return l
}

// Font TODO
func (l *Line) Font(s string) *Line {
	l.font = s
	return l
}

// Size TODO
func (l *Line) Size(i int) *Line {
	l.size = i
	return l
}

// Bash TODO
func (l *Line) Bash(s string) *Line {
	l.bash = s
	return l
}

// Param TODO
func (l *Line) Param(s []string) *Line {
	l.param = s
	return l
}

// Terminal TODO
func (l *Line) Terminal(b bool) *Line {
	l.terminal = &b
	return l
}

// Refresh TODO
func (l *Line) Refresh(b bool) *Line {
	l.refresh = &b
	return l
}

// DropDown TODO
func (l *Line) DropDown(b bool) *Line {
	l.dropDown = &b
	return l
}

// Length TODO
func (l *Line) Length(i int) *Line {
	l.length = i
	return l
}

// Trim TODO
func (l *Line) Trim(b bool) *Line {
	l.trim = &b
	return l
}

// Alternate TODO
func (l *Line) Alternate(b bool) *Line {
	l.alternate = &b
	return l
}

// TemplateImage TODO
func (l *Line) TemplateImage(s string) *Line {
	l.templateImage = s
	return l
}

// Image TODO
func (l *Line) Image(s string) *Line {
	l.image = s
	return l
}

// Emojize TODO
func (l *Line) Emojize(b bool) *Line {
	l.emojize = &b
	return l
}

// Ansi TODO
func (l *Line) Ansi(b bool) *Line {
	l.ansi = &b
	return l
}

// Render TODO
func (b *Menu) Render() string {
	var output string
	for _, line := range b.MenuBar.Line {
		output = output + fmt.Sprintf("%s\n", renderLine(line))
	}
	output = output + "---\n"
	output = output + renderDropDown(b.DropDown)
	return output
}

func renderDropDown(d *DropDown) string {
	var output string
	var prefix string
	if d.Level > 0 {
		prefix = strings.Repeat("--", d.Level) + " "
	}
	for _, line := range d.Line {
		switch v := line.(type) {
		case *Line:
			output = output + fmt.Sprintf("%s%s\n", prefix, renderLine(v))
		case *DropDown:
			output = output + renderDropDown(v)
		}
	}
	return output
}

func renderLine(l *Line) string {
	result := []string{l.text}
	options := []string{"|"}
	if l.href != "" {
		options = append(options, fmt.Sprintf("href=%s", l.href))
	}
	if l.color != "" {
		options = append(options, fmt.Sprintf("color=%s", l.color))
	}
	if l.font != "" {
		options = append(options, fmt.Sprintf("font=%s", l.font))
	}
	if l.size > 0 {
		options = append(options, fmt.Sprintf("size=%d", l.size))
	}
	if l.bash != "" {
		options = append(options, fmt.Sprintf("bash=\"%s\"", l.bash))
	}
	if len(l.param) > 0 {
		for i, param := range l.param {
			options = append(options, fmt.Sprintf("param%d=%s", i+1, param))
		}
	}
	if l.terminal != nil {
		options = append(options, fmt.Sprintf("terminal=%t", *l.terminal))
	}
	if l.refresh != nil {
		options = append(options, fmt.Sprintf("refresh=%t", *l.refresh))
	}
	if l.dropDown != nil {
		options = append(options, fmt.Sprintf("dropdown=%t", *l.dropDown))
	}
	if l.length > 0 {
		options = append(options, fmt.Sprintf("length=%d", l.length))
	}
	if l.trim != nil {
		options = append(options, fmt.Sprintf("trim=%t", *l.trim))
	}
	if l.alternate != nil {
		options = append(options, fmt.Sprintf("alternate=%t", *l.alternate))
	}
	if len(l.templateImage) > 0 {
		options = append(options, fmt.Sprintf("templateImage=%s", l.templateImage))
	}
	if len(l.image) > 0 {
		options = append(options, fmt.Sprintf("image=%s", l.image))
	}
	if l.emojize != nil {
		options = append(options, fmt.Sprintf("emojize=%t", *l.emojize))
	}
	if l.ansi != nil {
		options = append(options, fmt.Sprintf("ansi=%t", *l.ansi))
	}

	if len(options) > 1 {
		result = append(result, options...)
	}

	return strings.Join(result, " ")
}
