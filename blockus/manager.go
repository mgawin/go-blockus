package blockus




type manager struct {
	Name   string `json:"-"`
	id     int
	Blocks []*block
}





type storage interface{
    
    GetGame() *Game
    StoreGame() 
}
