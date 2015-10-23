package server

import (
	"blockus_game/blockus"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"log"
)

type AppengineStore struct {
	c context.Context
}

type HibernatedGame struct {
	JSON []byte
}

var ErrGenericKeyError = errors.New("Invalid key identifier")

func (ds *AppengineStore) SetContext(ctx context.Context) {

	ds.c = ctx

}

func (ds *AppengineStore) GetGame(k *string) (*blockus.Game, error) {

	var storedGame HibernatedGame
	var key *datastore.Key
	var err error

	key, err = datastore.DecodeKey(*k)
	if err != nil {

		log.Println(err)
		log.Println("klucz" + *k)
		return nil, ErrGenericKeyError
	}

	if err := datastore.Get(ds.c, key, &storedGame); err != nil {

		return nil, err
	}

	game, err := blockus.FromJSON(storedGame.JSON)
	if err != nil {

		return nil, err
	}

	return &game, nil
}

func (ds *AppengineStore) StoreGame(game *blockus.Game, k *string) (string, error) {

	str := HibernatedGame{JSON: game.ToJSON()}
	var key *datastore.Key
	var err error

	if k == nil {
		key = datastore.NewIncompleteKey(ds.c, "Game", nil)
	} else {

		key, err = datastore.DecodeKey(*k)
		if err != nil {

			return "", err
		}
	}

	key, err = datastore.Put(ds.c, key, &str)
	if err != nil {
		return "", err
	}
	return key.Encode(), nil
}
