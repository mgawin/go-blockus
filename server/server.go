package server

import (
	"blockus_game/blockus"
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	"fmt"
)

func Start() {


	container := game_container{}
	container.games = make(map[string]*blockus.Game)

	http.HandleFunc("/new", container.create)
	http.HandleFunc("/status", container.status)
	http.HandleFunc("/check", container.is_allowed)
	
	fs := http.FileServer(http.Dir("client"))
  	http.Handle("/client/",http.StripPrefix("/client/", fs))
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
	fmt.Fprint(w, "{id:"+id+"}")
	return
}

func (container *game_container) status(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query()["id"]
	game, ok := container.games[id[0]]
	if !ok {
		http.Error(w, "Game doesn't exist", http.StatusNotFound)

	}
	content, _ := json.Marshal(&game)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(content))
	return
}

func (container *game_container) is_allowed(w http.ResponseWriter, r *http.Request) {

	if (len(r.URL.Query()["gid"]) == 0) || (len(r.URL.Query()["pid"]) == 0) || (len(r.URL.Query()["bid"]) == 0) ||
		(len(r.URL.Query()["x"]) == 0) || (len(r.URL.Query()["y"]) == 0) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	gid := r.URL.Query()["gid"][0]
	bid, _ := strconv.Atoi(r.URL.Query()["bid"][0])
	pid, _ := strconv.Atoi(r.URL.Query()["pid"][0])
	x, _ := strconv.Atoi(r.URL.Query()["x"][0])
	y, _ := strconv.Atoi(r.URL.Query()["y"][0])

	game, ok := container.games[gid]
	if !ok {
		http.Error(w, "Game doesn't exist", http.StatusNotFound)

	}
	if pid > 1 {
		http.Error(w, "Invalid player id", http.StatusNotFound)
	}
	player := game.PlayerB
	if pid > 0 {
		player = game.PlayerA
	}

	if bid >= len(player.Blocks) {
		http.Error(w, "Invalid block id", http.StatusNotFound)
	}
	if x >= 14 || y >= 14 {
		http.Error(w, "Invalid coordinates", http.StatusNotFound)
	}

	block := player.Blocks[bid]
	if game.Check(player, block, x, y) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
