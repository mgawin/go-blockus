package main

import "fmt"
import "blockus_game/blockus"

func main() {

	game := blockus.NewGame("Jack", "Jason")

	fmt.Println(game.ToString())

	
	game.Move(&game.PlayerA,&game.PlayerA.Blocks[2],0,0)
	game.Move(&game.PlayerA,&game.PlayerA.Blocks[20],8,8)
	
	block:=&game.PlayerA.Blocks[14]
//	block.Rotate()
	game.Move(&game.PlayerA,block,4,10)
	
	fmt.Println(game.ToString())

}
