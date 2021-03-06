package blockus


type Manager struct {
	gamesCache  map[string]*Game
	storage     Persister
	currentGame *string
}

type Persister interface {
	GetGame(*string) (*Game, error)
	StoreGame(*Game, *string) (string, error)
}

func Init(db Persister) *Manager {

	mgr := new(Manager)
	mgr.storage = db
	mgr.gamesCache = make(map[string]*Game)
	return mgr
}

func (mgr *Manager) DispatchPlayer() (*string, *Game, *int, error) {

	if mgr.currentGame != nil {

		game, err := mgr.GetGame(mgr.currentGame)
		if err != nil {
			ErrorLog.Println("Unable to find current game"+*mgr.currentGame)
			return nil, nil, nil, err

		}
		pid := 2
		game.PlayerB = NewPlayer("PlayerB", pid)
		gid, err := mgr.storage.StoreGame(game, mgr.currentGame)
		if err != nil {
			return nil, nil, nil, err
		}
		mgr.currentGame = nil
		game.State = AMOVE
		return &gid, game, &pid, nil

	} else {

		game := new(Game)
		pid := 1
		game.PlayerA = NewPlayer("Player1", pid)
		game.State = WAITING

		game.Board = NewBoard()

		gid, err := mgr.storage.StoreGame(game, nil)
		if err != nil {
			return nil, nil, nil, err
		}
		mgr.gamesCache[gid] = game
		mgr.currentGame = &gid

		return &gid, game, &pid, nil
	}
}

func (mgr *Manager) GetGame(gid *string) (*Game, error) {

	game, prs := mgr.gamesCache[*gid]
	DebugLog.Println(mgr.gamesCache[*gid])
	if !prs {
		var err error
		DebugLog.Println("Game from storage")
		game, err = mgr.storage.GetGame(gid)
		if err != nil {
			ErrorLog.Println(err)
			return nil, err
		}

	}

	return game, nil

}

func (mgr *Manager) SaveGame(game *Game, gid *string) error {

	g, err := mgr.storage.StoreGame(game, gid)
	if err != nil || *gid != g {
		return err
	}
	if mgr.gamesCache[*gid] != game {
		DebugLog.Println(*gid)
		DebugLog.Println(game)
		mgr.gamesCache[*gid] = game
	}
	return nil
}
