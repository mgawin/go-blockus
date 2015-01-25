package blockus

type Game struct {
	PlayerA     player
	PlayerB     player
	Board       *board
	moves_taken int
}

func NewGame(name1 string, name2 string) *Game {
	game := new(Game)
	game.PlayerB = NewPlayer(name1, 1)
	game.PlayerA = NewPlayer(name2, 2)
	game.Board = NewBoard()
	return game

}

func (game *Game) ToString() string {

	str := "Blockus Game\n"
	str += game.PlayerA.Name + " vs. " + game.PlayerB.Name
	str += "\n\n" + game.Board.ToString() + "\n\n"
	str += game.PlayerA.ToString()
	return str
}

func (game *Game) Move(player *player, block *block, x int, y int) {

	game.Board.Put(block, x, y)

}
