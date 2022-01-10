package db

import (
	"context"
	"fmt"
)

type AddRequest struct{}
type AddResponse struct{}

func (d *Db) Create(ctx context.Context, i Item) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Disconnect()
	if err := db.Create(ItemsCol, i); err != nil {
		return fmt.Errorf("error creating new item with id %d - %v", i.ID, err)
	}
	return nil
}
