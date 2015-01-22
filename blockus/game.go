package blockus

type game struct {
	PlayerA     player
	PlayerB     player
	board       *board
	moves_taken int
}

func NewGame(name1 string, name2 string) *game {
	game := new(game)
	game.PlayerB = NewPlayer(name1, 1)
	game.PlayerA = NewPlayer(name2, 2)
	game.board = NewBoard()
	return game

}

func (game *game) ToString() string {

	str := "Blockus Game\n"
	str += game.PlayerA.name + " vs. " + game.PlayerB.name
	str += "\n\n" + game.board.ToString() + "\n\n"
	str += game.PlayerA.ToString()
	return str
}

func (game *game) Move(player *player, block *block, x int, y int) {

	game.board.Put(block, x, y)

}
