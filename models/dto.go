package models

type Config struct {
	Uri      string
	Database string
}

type Seq struct {
	ID  string `bson:"_id"`
	Seq int
}

type ShortURL struct {
	ID      string `bson:"_id"`
	LongURL string
	Hash    string
}

type HashRequest struct {
	Seq string `json:"seq"`
}

type HashResponse struct {
	Hash string `json:"hash"`
}
