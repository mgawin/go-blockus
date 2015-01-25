package server

import (
	"blockus_game/blockus"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Start() {

	fmt.Println("Blockus server started.")

	container := game_container{}
	container.games = make(map[string]*blockus.Game)

	http.HandleFunc("/new", container.new_game)
	http.HandleFunc("/status", container.game_status)

	http.ListenAndServe(":8080", nil)

}

type game_container struct {
	games map[string]*blockus.Game
}

func (cont *game_container) new_game(w http.ResponseWriter, r *http.Request) {

	game := blockus.NewGame("Jack", "Jason")
	id := "10" + strconv.Itoa(len(cont.games))
	cont.games[id] = game
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{id:"+id+"}")
	return
}

func (container *game_container) game_status(w http.ResponseWriter, r *http.Request) {

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
