package main

import (
	"fmt"

	"github.com/johnmccabe/bitbar"
)

func main() {
	b := bitbar.New()
	s := bitbar.Style{
		Font:  "UbuntuMono-Bold",
		Color: "red",
		Size:  12,
	}
	c := bitbar.Cmd{
		Bash:   "/path/to/cmd",
		Params: []string{"arg1", "arg2"},
	}
	b.StatusLine("MenuItem 1").Href("http://localhost:8080").Color("red").DropDown(false)
	b.StatusLine("MenuItem 2").DropDown(false)
	menu := b.NewSubMenu()
	menu.Line("DropDown Level 1 A").Style(s)
	submenu := b.SubMenu.NewSubMenu()
	submenu.Line("DropDown Level 2 A")
	submenu.Line("DropDown Level 1 B").Command(c)
	subsubmenu := submenu.NewSubMenu()
	subsubmenu.Line("DropDown Level 3 A")
	menu.Line("DropDown Level 1 B").Style(s)

	fmt.Print(b.Render())
}
