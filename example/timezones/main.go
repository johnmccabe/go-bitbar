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
		Bash:  "/path/to/cmd",
		Param: []string{"arg1", "arg2"},
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

	fmt.Print(b.Render())
}
