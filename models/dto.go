package models

type Seq struct {
	ID  string `bson:"_id"`
	Seq int
}

type ShortURL struct {
	ID      string `bson:"_id"`
	LongURL string
	Hash    string
}
