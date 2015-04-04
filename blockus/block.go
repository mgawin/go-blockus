package blockus

import "fmt"

type block struct {
	Value     int     `json:"val"`
	Shape     [][]int `json:"shape"`
	flippable bool
	rotatable bool
	rotate    int
}

func (block *block) ToString() string {

	str := ""
	for i := 0; i < len(block.Shape); i++ {
		for j := 0; j < len(block.Shape[i]); j++ {
			if block.Shape[i][j] > 0 {
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
		return
	}
	j := len(block.Shape)
	i := len(block.Shape[0])

	nshape := make([][]int, i)
	for k := range nshape {
		nshape[k] = make([]int, j)
	}

	for y := 0; y < j; y++ {

		for x := 0; x < i; x++ {

			nshape[x][j-1-y] = block.Shape[y][x]

		}
	}
	block.rotate += 1
	if block.rotate > 3 {
		block.rotate = 0
	}
	block.Shape = nshape
}

func (block *block) Flip() {

	if !block.flippable {
		fmt.Println(block.ToString())
		return
	}
	j := len(block.Shape)
	i := len(block.Shape[0])

	nshape := make([][]int, j)
	for k := range nshape {
		nshape[k] = make([]int, i)
	}

	for y := 0; y < j; y++ {

		for x := 0; x < i; x++ {

			nshape[y][i-1-x] = block.Shape[y][x]

		}
	}

	block.Shape = nshape
}

func (block *block) is_corner(i int, j int) bool {
	c1 := 0
	c2 := 0
	c3 := 0
	c4 := 0
	if j > 0 {
		c1 = block.Shape[j-1][i]
	}
	if len(block.Shape[0])-1 > i {
		c2 = block.Shape[j][i+1]
	}
	if len(block.Shape)-1 > j {
		c3 = block.Shape[j+1][i]
	}
	if i > 0 {
		c4 = block.Shape[j][i-1]
	}

	if ((c1 + c2) == 0) || ((c2 + c3) == 0) || ((c3 + c4) == 0) || ((c4 + c1) == 0) {
		return true
	}
	return false
}

func (block *block) Get_offset(new_rotate int) int {

	k := new_rotate - block.rotate
	if k < 0 {
		k = 4 + k
	}
	return k
}
