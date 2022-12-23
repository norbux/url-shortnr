package shortnr

import (
	"encoding/hex"
	"errors"

	"github.com/zeebo/xxh3"
)

var s = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y",
	"z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q",
	"R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func toBase62(d int) string {
	var hash string

	for d > 0 {
		hash = s[d%62] + hash
		d = d / 62
	}

	return hash
}

func B62Hash(seq int) (string, error) {
	if seq < 1 {
		return "", errors.New("sequence number must be greater than zero")
	}

	return toBase62(seq), nil
}

func Xxh3Hash(input string) (hash string, err error) {
	if len(input) <= 0 {
		return "", errors.New("hashing input is empty")
	}

	hasher := xxh3.New()
	_, _ = hasher.WriteString(input)
	hash = hex.EncodeToString(hasher.Sum(nil))

	return string(hash[:6]), nil
}
