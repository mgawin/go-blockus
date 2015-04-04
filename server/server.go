package server

import (
	"blockus_game/blockus"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Start() {

	container := game_container{}
	container.games = make(map[string]*blockus.Game)

	http.HandleFunc("/new", container.create)
	http.HandleFunc("/status", container.status)
	http.HandleFunc("/moves", container.get_allowed_moves)
	http.HandleFunc("/move", container.do_move)

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/client/", http.StripPrefix("/client/", fs))
	log.Println("Blockus server started.")

	http.ListenAndServe(":8080", nil)

}

type game_container struct {
	games map[string]*blockus.Game
}

func (cont *game_container) create(w http.ResponseWriter, r *http.Request) {

	game := blockus.NewGame("Jack", "Jason")
	id := "10" + strconv.Itoa(len(cont.games))
	cont.games[id] = game
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	content, _ := json.Marshal(&game)
	pid := "0"
	fmt.Fprint(w, "{\"gid\":"+id+",\"pid\":"+pid+",\"game\":"+string(content)+"}")
	return
}

func (container *game_container) status(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query()["id"]
	game, ok := container.games[id[0]]
	if !ok {
		http.Error(w, "Game doesn't exist", http.StatusNotFound)
		return
	}
	content, _ := json.Marshal(&game)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, string(content))
	return
}

func (container *game_container) get_allowed_moves(w http.ResponseWriter, r *http.Request) {
	t0:=time.Now()
	defer log.Println(time.Since(t0).String())

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if (len(r.URL.Query()["gid"]) == 0) || (len(r.URL.Query()["pid"]) == 0) || (len(r.URL.Query()["bid"]) == 0) ||
		(len(r.URL.Query()["rotates"]) == 0) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	gid := r.URL.Query()["gid"][0]
	bid, _ := strconv.Atoi(r.URL.Query()["bid"][0])
	pid, _ := strconv.Atoi(r.URL.Query()["pid"][0])
	rotates, _ := strconv.Atoi(r.URL.Query()["rotates"][0])

	game, ok := container.games[gid]
	if !ok {
		http.Error(w, "Game doesn't exist", http.StatusNotFound)
		return
	}
	if pid > 1 {
		http.Error(w, "Invalid player id", http.StatusNotFound)
		return
	}
	player := game.PlayerB
	if pid > 0 {
		player = game.PlayerA
	}

	if bid >= len(player.Blocks) {
		http.Error(w, "Invalid block id", http.StatusNotFound)
		return
	}

	if rotates >= 4 || rotates < 0 {
		http.Error(w, "Invalid rotate parameter", http.StatusNotFound)
		return
	}

	block := player.Blocks[bid]
	for k := 0; k < block.Get_offset(rotates); k++ {
		block.Rotate()

	}
	response, _ := json.Marshal(game.Get_allowed_moves(block))

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(response))

}

func (container *game_container) do_move(w http.ResponseWriter, r *http.Request) {
	t0:=time.Now()
	defer log.Println(time.Since(t0).String())
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if (len(r.URL.Query()["gid"]) == 0) || (len(r.URL.Query()["pid"]) == 0) || (len(r.URL.Query()["bid"]) == 0) ||
		(len(r.URL.Query()["rotates"]) == 0) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	gid := r.URL.Query()["gid"][0]
	bid, _ := strconv.Atoi(r.URL.Query()["bid"][0])
	pid, _ := strconv.Atoi(r.URL.Query()["pid"][0])
	rotates, _ := strconv.Atoi(r.URL.Query()["rotates"][0])
	x, _ := strconv.Atoi(r.URL.Query()["x"][0])
	y, _ := strconv.Atoi(r.URL.Query()["y"][0])

	game, ok := container.games[gid]
	if !ok {
		http.Error(w, "Game doesn't exist", http.StatusNotFound)
		return
	}
	if pid > 1 {
		http.Error(w, "Invalid player id", http.StatusNotFound)
		return
	}
	player := game.PlayerB
	if pid > 0 {
		player = game.PlayerA
	}

	if bid >= len(player.Blocks) {
		http.Error(w, "Invalid block id", http.StatusNotFound)
		return
	}

	if x >= 14 || x < 0 {
		http.Error(w, "Invalid coordinates", http.StatusNotFound)
		return
	}

	if y >= 14 || y < 0 {
		http.Error(w, "Invalid coordinates", http.StatusNotFound)
		return
	}

	if rotates >= 4 || rotates < 0 {
		http.Error(w, "Invalid rotate parameter", http.StatusNotFound)
		return
	}

	block := player.Blocks[bid]
	for k := 0; k < block.Get_offset(rotates); k++ {
		block.Rotate()

	}

	if game.Check(block, x, y) {

		go game.Move(block, x, y)
		w.WriteHeader(http.StatusCreated)
	} else {

		http.Error(w, "Invalid game status", http.StatusInternalServerError)
		return
	}
}
