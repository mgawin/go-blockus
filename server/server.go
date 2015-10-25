package server

import (
	"blockus_game/blockus"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"log"
	"strconv"
	"time"
)

type BlockusAPI struct {
}

var manager *blockus.Manager
var db *AppengineStore

func init() {

	api, err := endpoints.RegisterService(BlockusAPI{}, "blockus", "v1", "blockus api", true)
	if err != nil {
		log.Fatal(err)
	}

	info := api.MethodByName("Create").Info()
	info.Name, info.HTTPMethod, info.Path = "NewGame", "GET", "new"

	info = api.MethodByName("GetAllowedMoves").Info()
	info.Name, info.HTTPMethod, info.Path = "GetAllowedMoves", "GET", "moves"

	info = api.MethodByName("DoMove").Info()
	info.Name, info.HTTPMethod, info.Path = "DoMove", "POST", "move"

	info = api.MethodByName("Status").Info()
	info.Name, info.HTTPMethod, info.Path = "Status", "GET", "status"

	db = new(AppengineStore)
	manager = blockus.Init(db)

	endpoints.HandleHTTP()

}

type NewGameRes struct {
	Pid  *int          `json:"pid"`
	GID  *string       `json:"gid"`
	Game *blockus.Game `json:"game"`
}

func (BlockusAPI) Create(c context.Context) (*NewGameRes, error) {
	db.SetContext(c)
	res := NewGameRes{}
	var err error
	res.GID, res.Game, res.Pid, err = manager.DispatchPlayer()
	if err != nil {
		return nil, err
	}
	return &res, nil
}

type MovesReq struct {
	PID     string `json:"pid"`
	GID     string `json:"gid"`
	BID     string `json:"bid"`
	Rotates string `json:"rotates"`
}

type MovesRes struct {
	Moves [][2]int `json:"moves"`
}

func (BlockusAPI) GetAllowedMoves(c context.Context, r *MovesReq) (*MovesRes, error) {

	if len(r.BID) == 0 || len(r.GID) == 0 || len(r.Rotates) == 0 || len(r.PID) == 0 {

		return nil, endpoints.NewBadRequestError("Missing parameters")

	}

	db.SetContext(c)

	t0 := time.Now()
	defer log.Println(time.Since(t0).String())

	res := MovesRes{}

	game, err := manager.GetGame(&r.GID)

	switch {

	case err == datastore.ErrNoSuchEntity:
		return nil, endpoints.NewNotFoundError("game not found")

	case err == datastore.ErrInvalidKey:
		return nil, endpoints.NewNotFoundError("Invalid key")

	case err == ErrGenericKeyError:
		return nil, endpoints.NewNotFoundError("Invalid key")

	case err != nil:
		return nil, endpoints.NewInternalServerError("server error")

	}

	player := game.PlayerB
	if i, _ := strconv.Atoi(r.PID); i > 0 {
		player = game.PlayerA
	}

	j, _ := strconv.Atoi(r.BID)
	if j >= len(player.Blocks) {
		return nil, endpoints.NewNotFoundError("game not found")
	}
	l, _ := strconv.Atoi(r.Rotates)
	if l >= 4 || l < 0 {
		return nil, endpoints.NewNotFoundError("game not found")
	}

	block := player.Blocks[j]
	for k := 0; k < block.Get_offset(l); k++ {
		block.Rotate()

	}
	res.Moves = game.Get_allowed_moves(block)
	return &res, nil

}

type DoMoveReq struct {
	PID     string `json:"pid"`
	GID     string `json:"gid"`
	BID     string `json:"bid"`
	Rotates string `json:"rotates"`
	X       string `json:"x"`
	Y       string `json:"y"`
}

func (BlockusAPI) DoMove(c context.Context, r *DoMoveReq) error {

	if len(r.BID) == 0 || len(r.GID) == 0 || len(r.Rotates) == 0 || len(r.PID) == 0 || len(r.X) == 0 || len(r.Y) == 0 {

		return endpoints.NewBadRequestError("Missing parameters")

	}

	db.SetContext(c)

	t0 := time.Now()
	defer log.Println(time.Since(t0).String())

	game, err := manager.GetGame(&r.GID)

	switch {

	case err == datastore.ErrNoSuchEntity:
		return endpoints.NewNotFoundError("game not found")

	case err == datastore.ErrInvalidKey:
		return endpoints.NewNotFoundError("Invalid key")

	case err == ErrGenericKeyError:
		return endpoints.NewNotFoundError("Invalid key")

	case err != nil:
		return endpoints.NewInternalServerError("server error")

	}

	player := game.PlayerB
	if i, _ := strconv.Atoi(r.PID); i > 0 {
		player = game.PlayerA
	}

	j, _ := strconv.Atoi(r.BID)
	if j >= len(player.Blocks) {
		return endpoints.NewNotFoundError("game not found")
	}
	l, _ := strconv.Atoi(r.Rotates)
	if l >= 4 || l < 0 {
		return endpoints.NewNotFoundError("game not found")
	}

	x, _ := strconv.Atoi(r.X)

	if x >= 14 || x < 0 {
		return endpoints.NewNotFoundError("game not found")
	}

	y, _ := strconv.Atoi(r.Y)

	if y >= 14 || y < 0 {
		return endpoints.NewNotFoundError("game not found")
	}

	block := player.Blocks[j]
	for k := 0; k < block.Get_offset(l); k++ {
		block.Rotate()

	}

	if game.Check(block, x, y) {

		game.Move(block, x, y)
		err := manager.SaveGame(game, &r.GID)
		if err != nil {
			log.Println(err)
			return endpoints.NewInternalServerError("server error")

		}

	}
	return nil
}

type StatusReq struct {
	PID string `json:"pid"`
	GID string `json:"gid"`
}

type StatusRes struct {
	Code     string   `json:"code"`
	LastMove [][2]int `json:"lastmove"`
}

func (BlockusAPI) Status(c context.Context, r *StatusReq) (*StatusRes, error) {

	db.SetContext(c)

	t0 := time.Now()
	defer log.Println(time.Since(t0).String())

	if len(r.GID) == 0 || len(r.PID) == 0 {

		return nil, endpoints.NewBadRequestError("Missing parameters")

	}

	game, err := manager.GetGame(&r.GID)

	switch {

	case err == datastore.ErrNoSuchEntity:
		return nil, endpoints.NewNotFoundError("game not found")

	case err == datastore.ErrInvalidKey:
		return nil, endpoints.NewNotFoundError("Invalid key")

	case err == ErrGenericKeyError:
		return nil, endpoints.NewNotFoundError("Invalid key")

	case err != nil:
		return nil, endpoints.NewInternalServerError("server error")

	}

	res := StatusRes{Code: strconv.Itoa(int(game.State)), LastMove: game.LastMove}

	return &res, nil
}
