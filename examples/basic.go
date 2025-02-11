package main

import (
	"fmt"

	"github.com/cleesim/roblox"
)

func main() {
	rbx := roblox.New()

	user, err := rbx.Users.Get(748671568)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("User: %+v\n", user)

	group, err := rbx.Groups.Get(13666720)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Group: %+v\n", group)
}
