package main

import (
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func main() {
	err := client.Login("DISCORD_APP_ID")
	if err != nil {
		panic(err)
	}

	fmt.Println("Logged in as")

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "grupobright.com",
		Details:    "Playing in Bright Cloud",
		// LargeImage: "./bcg_branco.png",
		// LargeText:  "This is the large image :D",
		// SmallImage: "smallimageid",
		// SmallText:  "And this is the small image",
		LargeImage: "bcg_branco",
		SmallImage: "bcg_branco",

		// Party: &client.Party{
		// 	ID:         "-1",
		// 	Players:    15,
		// 	MaxPlayers: 24,
		// },
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "Discord",
				Url:   "https://discord.gg/grupobright",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 59999)
}
