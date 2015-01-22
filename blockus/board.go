package blockus

import "strconv"
import "errors"

type board struct {
	field [14][14]int
	count int
}

func NewBoard() *board {

	board := new(board)
	board.count = 0
	return board

}

func (board *board) IsAllowed(block *block, x int, y int) bool {

	if x+len(block.shape[0]) > 13 {
		return false
	}
	if y+len(block.shape) > 13 {
		return false
	}

	for j := 0; j < len(block.shape); j++ {
		for i := 0; i < len(block.shape[j]); i++ {

			if (block.shape[j][i] > 0) && (board.field[y+j][x+i] > 0 || board.touch_edge(x+i, y+j, block.shape[j][i])) {
				return false
			}
		}
	}

	if board.count == 0 {
		return true
	}

	for j := 0; j < len(block.shape); j++ {
		for i := 0; i < len(block.shape[0]); i++ {

			if block.shape[j][i] > 0 {
				if block.is_corner(i, j) {
					if board.touch_corner(x+i, y+j, block.shape[j][i]) {
						return true
					}

				}

			}

		}

	}

	return false

}

func (board *board) ToString() string {

	res := ""
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			res += strconv.Itoa(board.field[i][j])

		}
		res += "\n"
	}
	return res

}

func (board *board) Put(block *block, x int, y int) error {

	if !board.IsAllowed(block, x, y) {
		return errors.New("Illegal move.")
	}
	for j := 0; j < len(block.shape); j++ {
		for i := 0; i < len(block.shape[j]); i++ {

			board.field[y+j][x+i] = block.shape[j][i]

		}

	}
	board.count++
	return nil
}

func (board *board) touch_corner(i int, j int, value int) bool {

	c1 := 0
	c2 := 0
	c3 := 0
	c4 := 0

	if j > 0 && i < 13 {
		c1 = board.field[j-1][i+1]
	}
	if j < 13 && i < 13 {
		c2 = board.field[j+1][i+1]
	}
	if j < 13 && i > 0 {
		c3 = board.field[j+1][i-1]
	}
	if i > 0 && j > 0 {
		c4 = board.field[j-1][i-1]
	}
	if (c4 == value) || (c3 == value) || (c2 == value) || (c1 == value) {

		return true
	}
	return false

}

func (board *board) touch_edge(i int, j int, value int) bool {

	c1 := 0
	c2 := 0
	c3 := 0
	c4 := 0

	if j > 0 {
		c1 = board.field[j-1][i]
	}
	if i < 13 {
		c2 = board.field[j][i+1]
	}
	if j < 13 {
		c3 = board.field[j+1][i]
	}
	if i > 0 {
		c4 = board.field[j][i-1]
	}

	if (c4 == value) || (c3 == value) || (c2 == value) || (c1 == value) {

		return true
	}
	return false

}
