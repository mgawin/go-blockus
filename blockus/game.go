package blockus

import (
	"encoding/json"
)

type Game struct {
	PlayerA     player
	PlayerB     player
	Board       board
	moves_taken int
	State       gameStatus
	LastMove    [][2]int
}

type gameStatus int

const (
	NULL gameStatus = iota
	WAITING
	AMOVE
	BMOVE
	FINISHED
)

func (game *Game) ToString() string {

	str := "Blockus Game\n"
	str += game.PlayerA.Name + " vs. " + game.PlayerB.Name
	str += "\n\n" + game.Board.ToString() + "\n\n"
	str += game.PlayerA.ToString()
	return str
}

func (game *Game) ToJSON() []byte {
	str, _ := json.Marshal(game)

	return str

}

func FromJSON(s []byte) (Game, error) {
	var game Game
	err := json.Unmarshal(s, &game)

	return game, err

}

func (game *Game) Move(block *block, x int, y int) {

	game.LastMove, _ = game.Board.Put(block, x, y)
	if game.State == AMOVE {
		game.State = BMOVE
	} else {
		game.State = AMOVE

	}
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
