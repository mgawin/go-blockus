package main

import "blockus_game/server"

func main() {

	server.Start()
	// game := blockus.NewGame("Jack", "Jason")

	// //fmt.Println(game.ToString())

	// game.Move(&game.PlayerA, &game.PlayerA.Blocks[2], 0, 0)
	// //game.Move(&game.PlayerA, &game.PlayerA.Blocks[20], 3, 3)

	// block := &game.PlayerA.Blocks[14]
	// block.Rotate()
	// block.Rotate()
	// game.Move(&game.PlayerA, block, 0, 2)
	// block = &game.PlayerA.Blocks[9]
	// block.Flip()
	// game.Move(&game.PlayerA, block, 2, 1)

	// fmt.Println(game.ToString())

}
