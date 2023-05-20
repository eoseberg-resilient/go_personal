package main

import (
	character "PERSONAL/txt_game/player"
	"PERSONAL/txt_game/system"
	"os"

	//"PERSONAL/txt_game/locations"

	"fmt"
)


func main() {
	fmt.Println("[ initializing ... ]")

	game 					:= system.NewGame(250, system.StateViewMenu)
	input_handler 			:= system.NewInputHandler()
	text_output_handler 	:= system.NewTextOutputHandler(100, 5)
	menu_output_handler 	:= system.NewMenuOutputHandler(100, 2)
	player					:= character.NewPlayer()

	//location				:= locations.NewLocationGroup([]locations.Location{
	//	locations.NewLocation(0, "kitchen", "kitchen description"),
	//	locations.NewLocation(0, "living room", "living room description"),
	//})

	player.SetInventory(map[string]string{
		"key": "An old, rusty key. Doesn't seem to work...",
		"soda": "Lovely stuff, this is...",
		"loaded m1911 handgun": "Huh.",
	})

	menu_output_handler.SetOutputText(player.Inventory().Display())

	for {
		if game.Tick() {
			if game.State() == system.StatePrintText {
				isDone := text_output_handler.WriteOutput()
				if isDone {
					game.SetGameState(system.StateUserInput)
				}
			} else if game.State() == system.StateViewMenu {
				isDone := menu_output_handler.WriteOutput()
				if isDone {
					game.SetGameState(system.StateUserInput)
				}
			}
		}

		if game.State() == system.StateUserInput && !game.IsWaiting() {
			input_handler.InitInput()
			for {
				input_handler.Display()
				input_handler.CheckForKeypress()
				if input_handler.InputDone() {
					if input_handler.GetInput() == "exit" { //TODO fix
						input_handler.TearDown()
						fmt.Println()
						os.Exit(0)
					}
					break
				}
			}
			game.Wait(3)
		}
		if game.State() == system.StateBetween {
			game.SetGameState(system.StateViewMenu)
		}
	}
}
