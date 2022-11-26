//package main is to show how multiple menus can be used with eachother
package main

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
)

func main() {
	menu := wmenu.NewMenu("Select option")

	menu.Action(func(opts []wmenu.Opt) error { fmt.Printf(opts[0].Text + " default response"); return nil })
	menu.Option("A", nil, true, nil)
	menu.Option("B", nil, false, nil)
	menu.Option("C", nil, false, func(opt wmenu.Opt) error {
		fmt.Printf("single response")
		return nil
	})
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
