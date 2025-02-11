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

	universeID, err := rbx.Games.GetUniverseID(17872901145)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Universe ID: %d\n", universeID)

	game, err := rbx.Games.Get(universeID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Game: %+v\n", game)
}
