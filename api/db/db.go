package db

type Db struct{}

const (
	ItemsCol = "items"
)

func New() *Db {
	return &Db{}
}

type Item struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}
