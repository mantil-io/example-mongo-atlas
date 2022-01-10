package db

import (
	"context"
	"fmt"
)

func (d *Db) Get(ctx context.Context, id int) (*Item, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Disconnect()
	i := Item{}
	if err := db.Get(ItemsCol, id, &i); err != nil {
		return nil, fmt.Errorf("error fetching item with id %d - %v", id, err)
	}
	return &i, nil
}
