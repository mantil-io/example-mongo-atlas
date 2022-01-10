package db

import (
	"context"
	"fmt"
)

func (d *Db) Delete(ctx context.Context, id int) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Disconnect()
	if err := db.Delete(ItemsCol, id); err != nil {
		return fmt.Errorf("error deleting item with id %d - %v", id, err)
	}
	return nil
}
