package main

import (
	"fmt"
	"strings"
)

func main() {
	b := BitBarMenu{}
	s := TextStyle{
		font:  "UbuntuMono-Bold",
		color: "red",
		size:  12,
	}
	c := Cmd{
		bash:  "/path/to/cmd",
		param: []string{"arg1", "arg2"},
	}
	b.AddMenuBar("MenuItem 1").Href("http://localhost:8080").Color("red").DropDown(false)
	b.AddMenuBar("MenuItem 2").DropDown(false)
	dropdown := b.AddDropDown()
	dropdown.Add("DropDown Level 1 A").Style(s)
	subdrop := b.DropDown.AddDropDown()
	subdrop.Add("DropDown Level 2 A")
	dropdown.Add("DropDown Level 1 B").Command(c)
	subsubdrop := subdrop.AddDropDown()
	subsubdrop.Add("DropDown Level 3 A")
	b.Render()
}

type BitBarMenu struct {
	MenuBar  MenuBar
	DropDown *DropDown
}

type Line struct {
	text string
	href string
	TextStyle
	Options
	Cmd
	templateImage string
	image         string
}

type TextStyle struct {
	color string
	font  string
	size  int
}

type Options struct {
	terminal  *bool
	refresh   *bool
	dropDown  *bool
	length    int
	trim      *bool
	alternate *bool
	emojize   *bool
	ansi      *bool
}

type Cmd struct {
	bash  string
	param []string
}

type MenuBar struct {
	Line []*Line
}

type DropDownItem interface{}

type DropDown struct {
	Level int
	Line  []DropDownItem
}

func (b *BitBarMenu) AddMenuBar(s string) *Line {
	l := new(Line)
	l.text = s
	b.MenuBar.Line = append(b.MenuBar.Line, l)
	return l
}

func (b *BitBarMenu) AddDropDown() *DropDown {
	b.DropDown = new(DropDown)
	b.DropDown.Level = 0
	return b.DropDown
}

func (d *DropDown) Add(s string) *Line {
	l := new(Line)
	l.text = s
	d.Line = append(d.Line, l)
	return l
}

func (parent *DropDown) AddDropDown() *DropDown {
	child := new(DropDown)
	child.Level = parent.Level + 1
	parent.Line = append(parent.Line, child)
	return child
}

func (b *BitBarMenu) Render() {
	for _, line := range b.MenuBar.Line {
		fmt.Printf("%s\n", RenderLine(line))
	}
	fmt.Println("---")
	// fmt.Printf("%+v\n", b.DropDown)
	RenderDropDown(b.DropDown)
}

func RenderDropDown(d *DropDown) {
	for _, line := range d.Line {
		switch v := line.(type) {
		case *Line:
			fmt.Printf("%s %s\n", strings.Repeat("--", d.Level), RenderLine(v))
		case *DropDown:
			RenderDropDown(v)
		}
	}
}

func (l *Line) Style(s TextStyle) *Line {
	l.TextStyle = s
	return l
}

func (l *Line) Command(c Cmd) *Line {
	l.Cmd = c
	return l
}

func (l *Line) Href(s string) *Line {
	l.href = s
	return l
}

func (l *Line) Color(s string) *Line {
	l.color = s
	return l
}

func (l *Line) Font(s string) *Line {
	l.font = s
	return l
}

func (l *Line) Size(i int) *Line {
	l.size = i
	return l
}

func (l *Line) Bash(s string) *Line {
	l.bash = s
	return l
}

func (l *Line) Param(s []string) *Line {
	l.param = s
	return l
}

func (l *Line) Terminal(b bool) *Line {
	l.terminal = &b
	return l
}

func (l *Line) Refresh(b bool) *Line {
	l.refresh = &b
	return l
}

func (l *Line) DropDown(b bool) *Line {
	l.dropDown = &b
	return l
}

func (l *Line) Length(i int) *Line {
	l.length = i
	return l
}

func (l *Line) Trim(b bool) *Line {
	l.trim = &b
	return l
}

func (l *Line) Alternate(b bool) *Line {
	l.alternate = &b
	return l
}

func (l *Line) TemplateImage(s string) *Line {
	l.templateImage = s
	return l
}

func (l *Line) Image(s string) *Line {
	l.image = s
	return l
}

func (l *Line) Emojize(b bool) *Line {
	l.emojize = &b
	return l
}

func (l *Line) Ansi(b bool) *Line {
	l.ansi = &b
	return l
}

func RenderLine(l *Line) string {
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
