package blockus

import "fmt"

type block struct {
	value     int
	shape     [][]int
	flippable bool
	rotatable bool
}

func (block *block) ToString() string {

	str := ""
	for i := 0; i < len(block.shape); i++ {
		for j := 0; j < len(block.shape[i]); j++ {
			if block.shape[i][j] > 0 {
				str += "x"
			} else {
				str += "0"
			}

		}
		str += "\n"
	}
	return str

}

func (block *block) Rotate() {
	if !block.rotatable {
		fmt.Println(block.ToString())
		return
	}
	j := len(block.shape)
	i := len(block.shape[0])

	nshape := make([][]int, i)
	for k := range nshape {
		nshape[k] = make([]int, j)
	}

	for y := 0; y < j; y++ {

		for x := 0; x < i; x++ {

			nshape[x][j-1-y] = block.shape[y][x]

		}
	}

	block.shape = nshape
	fmt.Println(block.ToString())
}

func (block *block) Flip() {

	if !block.flippable {
		fmt.Println(block.ToString())
		return
	}
	j := len(block.shape)
	i := len(block.shape[0])

	nshape := make([][]int, j)
	for k := range nshape {
		nshape[k] = make([]int, i)
	}

	for y := 0; y < j; y++ {

		for x := 0; x < i; x++ {

			nshape[y][i-1-x] = block.shape[y][x]

		}
	}

	block.shape = nshape
	fmt.Println(block.ToString())
}
