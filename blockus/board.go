package blockus

import "strconv"
import "errors"

type board struct {
	field [14][14]int
}

func NewBoard() *board {

	board := new(board)

	return board

}

func (board *board) IsAllowed(block *block, x int, y int) bool {
		
		if x+len(block.shape[0])>13 {
			return false
		}
		if y+len(block.shape)>13 {
			return false	
		}
		
		for j:=0;j<len(block.shape);j++{
			for i:=0;i<len(block.shape[j]);i++{
				if board.field[y+j][x+i]>0 {
					return false	
			}
			
		}
		
		
	
}

	return true

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
	for j:=0;j<len(block.shape);j++{
		for i:=0;i<len(block.shape[j]);i++{
	
			board.field[y+j][x+i]=block.shape[j][i]	
			
		}
		
	}
	return nil
}

