package server

import (
	"appengine/datastore"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"log"
	"blockus_game/blockus"
	"time"
	"strconv"
)
type BlockusAPI struct{
	
	}
	
type GameCache struct{
	Games []blockus.Game	
}

var cache = GameCache{}

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


	endpoints.HandleHTTP()

}

type NewGameRes struct{
	Pid string `json:"pid"`
	UID *datastore.Key `json:"gid"`
	Game *blockus.Game  `json:"game"`
	
}

type HibernatedGame struct{
	JSON []byte

}


func (BlockusAPI) Create(c endpoints.Context)  (*NewGameRes,error){
	
	res:= NewGameRes{}
	res.Game = blockus.NewGame("Jack", "Jason")
	res.UID, _ = cache.StoreGame(c, res.Game, nil)
	
	res.Pid = "0"
	return &res,nil
	
}

func (cache *GameCache) StoreGame(c endpoints.Context, game *blockus.Game, k *datastore.Key) (*datastore.Key, error){
	
	str:=HibernatedGame{JSON: game.ToJSON()}
	key:=k
	if k==nil {
		key = datastore.NewIncompleteKey(c, "Game", nil)
	} 
	
	key, err := datastore.Put(c, key, &str)
	if err != nil {
		return nil, err
	}
	return key, nil
}


type MovesReq struct{
	PID string `json:"pid"`
	UID *datastore.Key `json:"gid"`
	BID string `json:"bid"`
	Rotates string `json:"rotates"`
}



type MovesRes struct{
	Moves [][2]int `json:"moves"`
	
}

func (BlockusAPI) GetAllowedMoves(c endpoints.Context,r *MovesReq)  (*MovesRes,error){
	t0:=time.Now()
	defer log.Println(time.Since(t0).String())

	res:=MovesRes{}


	var storedGame HibernatedGame
	
	if err := datastore.Get(c, r.UID, &storedGame); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("game not found")
	} else if err != nil {
		return nil, err
	}
	game,_:=blockus.FromJSON(storedGame.JSON)
	
	player := game.PlayerB
	if i,_:=strconv.Atoi(r.PID); i > 0 {
		player = game.PlayerA
	}

	j,_:=strconv.Atoi(r.BID);
	if  j>= len(player.Blocks) {
		return nil, endpoints.NewNotFoundError("game not found")
	}
	l,_:=strconv.Atoi(r.Rotates);
	if  l >= 4 || l < 0 {
		return nil, endpoints.NewNotFoundError("game not found")
	}

	block := player.Blocks[j]
	for k := 0; k < block.Get_offset(l); k++ {
		block.Rotate()

	}
	res.Moves = game.Get_allowed_moves(block)
	return &res,nil

}

type DoMoveReq struct{
	PID string `json:"pid"`
	UID *datastore.Key `json:"gid"`
	BID string `json:"bid"`
	Rotates string `json:"rotates"`
	X string `json:"x"`
	Y string `json:"y"`
}


func (BlockusAPI) DoMove(c endpoints.Context,r *DoMoveReq)  error {

	var storedGame HibernatedGame
	
	if err := datastore.Get(c, r.UID, &storedGame); err == datastore.ErrNoSuchEntity {
		return  endpoints.NewNotFoundError("game not found")
	} else if err != nil {
		return  err
	}
	game,_:=blockus.FromJSON(storedGame.JSON)
	
	player := game.PlayerB
	if i,_:=strconv.Atoi(r.PID); i > 0 {
		player = game.PlayerA
	}

	j,_:=strconv.Atoi(r.BID);
	if  j>= len(player.Blocks) {
		return endpoints.NewNotFoundError("game not found")
	}
	l,_:=strconv.Atoi(r.Rotates);
	if  l >= 4 || l < 0 {
		return  endpoints.NewNotFoundError("game not found")
	}
	
	x,_:=strconv.Atoi(r.X);
	
	if x >= 14 || x < 0 {
		return endpoints.NewNotFoundError("game not found")
	}

	y,_:=strconv.Atoi(r.Y);

	if y >= 14 || y < 0 {
		return endpoints.NewNotFoundError("game not found")
	}

	block := player.Blocks[j]
	for k := 0; k < block.Get_offset(l); k++ {
		block.Rotate()

	}
	
	if game.Check(block, x, y) {

		 game.Move(block, x, y)
		 _, err := cache.StoreGame(c, &game, r.UID)
		log.Println(err)
		return nil
		
	} else {
		return endpoints.NewNotFoundError("game not found")
	}
}


