package blockus

type Game struct {
	PlayerA     *player
	PlayerB     *player
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

func (game *Game) Move(block *block, x int, y int) {

	game.Board.Put(block, x, y)

}

func (game *Game) Check(block *block, x int, y int) bool {

	if game.Board.is_allowed(block, x, y) {
		return true
	}
	return false
}

func (game *Game) Get_allowed_moves(block *block) [][2]int {

	result := [][2]int{}

	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {

			if game.Board.is_allowed(block, i, j) {
				coords := [2]int{i, j}
				result = append(result, coords)
			}

		}

	}
	return result

}
