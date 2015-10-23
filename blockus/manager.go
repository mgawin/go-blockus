package blockus

import (
	"log"
)

type Manager struct {
	gamesCache map[string]*Game
	storage    Persister
}

type Persister interface {
	GetGame(*string) (*Game, error)
	StoreGame(*Game, *string) (string, error)
}

func Init(db Persister) *Manager {

	manager := new(Manager)
	manager.storage = db
	manager.gamesCache = make(map[string]*Game)
	return manager
}

func (manager *Manager) DispatchPlayer() (*string, *Game, *int, error) {
	game := new(Game)
	pid := 1
	game.PlayerA = NewPlayer("Player1", pid)
	game.PlayerB = NewPlayer("Player2", pid)

	game.Board = NewBoard()

	gid, err := manager.storage.StoreGame(game, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	log.Println(gid)
	manager.gamesCache[gid] = game
	return &gid, game, &pid, nil

}

func (manager *Manager) GetGame(gid *string) (*Game, error) {

	game, prs := manager.gamesCache[*gid]
	if !prs {
		var err error
		game, err = manager.storage.GetGame(gid)
		if err != nil {
			log.Println(err)
			return nil, err
		}

	}

	return game, nil

}

func (manager *Manager) SaveGame(game *Game, gid *string) error {

	g, err := manager.storage.StoreGame(game, gid)
	if err != nil || *gid != g {
		return err
	}
	if manager.gamesCache[*gid] != game {
		log.Println(*gid)
		log.Println(game)
		log.Println("Cache inconsistency")
		manager.gamesCache[*gid] = game
	}
	return nil
}
