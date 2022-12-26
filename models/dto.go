package models

type HashType int

const (
	BASE62 HashType = iota + 1
	XXH3
)

type Config struct {
	Uri      string
	Database string
}

type Seq struct {
	ID  string `bson:"_id"`
	Seq int
}

type URLMap struct {
	// ID      string `bson:"_id"`
	LongURL string
	Hash    string
}

type NewHashRequest struct {
	URL    string   `json:"url"`
	Method HashType `json:"method"`
}

type NewHashResponse struct {
	URL  string `json:"url"`
	Hash string `json:"hash"`
}

type GetURLRequest struct {
	Hash   string   `json:"hash"`
	Method HashType `json:"method"`
}
