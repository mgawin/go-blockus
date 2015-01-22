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

func (block *block) is_corner(i int, j int) bool {
	c1 := 0
	c2 := 0
	c3 := 0
	c4 := 0
	if j > 0 {
		c1 = block.shape[j-1][i]
	}
	if len(block.shape[0])-1 > i {
		c2 = block.shape[j][i+1]
	}
	if len(block.shape)-1 > j {
		c3 = block.shape[j+1][i]
	}
	if i > 0 {
		c4 = block.shape[j][i-1]
	}

	if ((c1 + c2) == 0) || ((c2 + c3) == 0) || ((c3 + c4) == 0) || ((c4 + c1) == 0) {
		return true
	}
	return false
}
