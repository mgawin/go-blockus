package blockus

import "strconv"

type player struct {
	name   string
	id     int
	Blocks []block
}

func NewPlayer(name string, id int) player {

	player := player{name: name}
	player.id = id
	player.Blocks = make([]block, 21)
	player.Blocks[0] = block{value: 1, shape: [][]int{{id}}, flippable: false, rotatable: false}
	player.Blocks[1] = block{value: 2, shape: [][]int{{id, id}}, flippable: false, rotatable: true}
	player.Blocks[2] = block{value: 3, shape: [][]int{{id, id}, {0, id}}, flippable: true, rotatable: true}
	player.Blocks[3] = block{value: 3, shape: [][]int{{id, id, id}}, flippable: false, rotatable: true}
	player.Blocks[4] = block{value: 4, shape: [][]int{{id, id}, {id, id}}, flippable: false, rotatable: false}
	player.Blocks[5] = block{value: 4, shape: [][]int{{0, id, 0}, {id, id, id}}, flippable: true, rotatable: true}
	player.Blocks[6] = block{value: 4, shape: [][]int{{id, id, id, id}}, flippable: false, rotatable: true}
	player.Blocks[7] = block{value: 4, shape: [][]int{{0, 0, id}, {id, id, id}}, flippable: true, rotatable: true}
	player.Blocks[8] = block{value: 4, shape: [][]int{{0, id, id}, {id, id, 0}}, flippable: true, rotatable: true}
	player.Blocks[9] = block{value: 5, shape: [][]int{{id, 0, 0, 0}, {id, id, id, id}}, flippable: true, rotatable: true}
	player.Blocks[10] = block{value: 5, shape: [][]int{{0, id, 0}, {0, id, 0}, {id, id, id}}, flippable: true, rotatable: true}
	player.Blocks[11] = block{value: 5, shape: [][]int{{id, 0, 0}, {id, 0, 0}, {id, id, id}}, flippable: true, rotatable: true}
	player.Blocks[12] = block{value: 5, shape: [][]int{{0, id, id, id}, {id, id, 0, 0}}, flippable: true, rotatable: true}
	player.Blocks[13] = block{value: 5, shape: [][]int{{0, 0, id}, {id, id, id}, {id, 0, 0}}, flippable: true, rotatable: true}
	player.Blocks[14] = block{value: 5, shape: [][]int{{id}, {id}, {id}, {id}, {id}}, flippable: false, rotatable: true}
	player.Blocks[15] = block{value: 5, shape: [][]int{{id, 0}, {id, id}, {id, id}}, flippable: true, rotatable: true}
	player.Blocks[16] = block{value: 5, shape: [][]int{{0, id, id}, {id, id, 0}, {id, 0, 0}}, flippable: true, rotatable: true}
	player.Blocks[17] = block{value: 5, shape: [][]int{{id, id}, {id, 0}, {id, id}}, flippable: true, rotatable: true}
	player.Blocks[18] = block{value: 5, shape: [][]int{{0, id, id}, {id, id, 0}, {0, id, 0}}, flippable: true, rotatable: true}
	player.Blocks[19] = block{value: 5, shape: [][]int{{0, id, 0}, {id, id, id}, {0, id, 0}}, flippable: false, rotatable: false}
	player.Blocks[20] = block{value: 5, shape: [][]int{{0, id, 0, 0, 0}, {id, id, id, id, id}}, flippable: true, rotatable: true}
	return player

}

func (player *player) ToString() string {
	str := player.name + "\n"

	for _, v := range player.Blocks {
		str += v.ToString() + "\n"

	}
	str += "Remaining value:" + strconv.Itoa(player.remaining_value())

	return str
}

func (player *player) remaining_value() int {
	val := 0
	for _, v := range player.Blocks {

		val += v.value
	}
	return val

}

func (player *player) Delete_block(pos int) {

	s := player.Blocks
	s = s[:pos+copy(s[pos:], s[pos+1:])]
	player.Blocks = s

}