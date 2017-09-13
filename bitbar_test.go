package bitbar

import (
	"fmt"
)

// Example Bitbar plugin resulting in the following output:
//  MenuItem 1 | color=red href=http://localhost:8080 dropdown=false
//  MenuItem 2 | dropdown=false
//  ---
//  DropDown Level 1 A | color=red font=UbuntuMono-Bold size=12
//  -- DropDown Level 2 A
//  -- DropDown Level 1 B | bash="/path/to/cmd" param1=arg1 param2=arg2
//  ---- DropDown Level 3 A
//  DropDown Level 1 B | color=red font=UbuntuMono-Bold size=12
func Example() {
	b := New()
	s := Style{
		Font:  "UbuntuMono-Bold",
		Color: "red",
		Size:  12,
	}
	c := Cmd{
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
