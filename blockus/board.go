package blockus

import ("strconv"
		"log"
"errors")


type board struct {
	Field [14][14]int
	Count int
}

func NewBoard() board {

	board := board{}
	board.Count = 0
	return board

}

func (board *board) is_allowed(block *block, x int, y int) bool {

	if x+len(block.Shape[0])-1 > 13 {
		return false
	}
	if y+len(block.Shape)-1 > 13 {
		return false
	}

	for j := 0; j < len(block.Shape); j++ {
		for i := 0; i < len(block.Shape[j]); i++ {

			if (block.Shape[j][i] > 0) && (board.Field[y+j][x+i] > 0 || board.touch_edge(x+i, y+j, block.Shape[j][i])) {
				return false
			}
		}
	}

	if board.Count < 2 {
		cov := false
		for j := 0; j < len(block.Shape); j++ {
			for i := 0; i < len(block.Shape[j]); i++ {

				if ((y+j == 9) && (x+i == 4)) || ((y+j == 4) && (x+i == 9)) {
					if block.Shape[j][i] > 0 {

						cov = true

					}
				}
			}
		}

		return cov
	}
	for j := 0; j < len(block.Shape); j++ {
		for i := 0; i < len(block.Shape[0]); i++ {

			if block.Shape[j][i] > 0 {
				if block.is_corner(i, j) {
					if board.touch_corner(x+i, y+j, block.Shape[j][i]) {

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
			res += strconv.Itoa(board.Field[i][j])

		}
		res += "\n"
	}
	return res

}

func (board *board) Put(block *block, x int, y int) ([][2]int, error) {

	if !board.is_allowed(block, x, y) {
		return nil, errors.New("Illegal move.")
	}
	filledItems := [][2]int{}
	for j := 0; j < len(block.Shape); j++ {
		for i := 0; i < len(block.Shape[j]); i++ {

			board.Field[y+j][x+i] = block.Shape[j][i]
			log.Println(block.Shape[j][i])
		
			if block.Shape[j][i] != 0 {
				coords := [2]int{x + i, y + j}
				filledItems = append(filledItems, coords)
			}
		}

	}
	board.Count++
	return filledItems, nil
}

func (board *board) touch_corner(i int, j int, value int) bool {

	c1 := 0
	c2 := 0
	c3 := 0
	c4 := 0

	if j > 0 && i < 13 {
		c1 = board.Field[j-1][i+1]
	}
	if j < 13 && i < 13 {
		c2 = board.Field[j+1][i+1]
	}
	if j < 13 && i > 0 {
		c3 = board.Field[j+1][i-1]
	}
	if i > 0 && j > 0 {
		c4 = board.Field[j-1][i-1]
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
		c1 = board.Field[j-1][i]
	}
	if i < 13 {
		c2 = board.Field[j][i+1]
	}
	if j < 13 {
		c3 = board.Field[j+1][i]
	}
	if i > 0 {
		c4 = board.Field[j][i-1]
	}

	if (c4 == value) || (c3 == value) || (c2 == value) || (c1 == value) {
		return true
	}
	return false

}
