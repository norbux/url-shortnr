package shortnr

import (
	"log"
	"testing"
)

func TestXxh3Hash(t *testing.T) {
	cases := []struct {
		name string
		seq  string
	}{
		{name: "hash 1", seq: "1"},
		{name: "hash 2", seq: "2"},
		{name: "hash 10", seq: "10"},
		{name: "hash 98", seq: "98"},
		{name: "hash 245", seq: "245"},
		{name: "hash 980", seq: "980"},
		{name: "hash 4235", seq: "4235"},
		{name: "hash 4987453", seq: "4987453"},
	}

	for i := range cases {
		tt := &cases[i]

		t.Run(tt.name, func(t *testing.T) {
			hash, err := Xxh3Hash(tt.seq)
			if err != nil {
				t.Fatalf("failed to hash with xxh3: %v", err)
			}
			log.Println(hash)
		})
	}
}

// TestHashing is a test that given an integer will use it to generate a hash
// that will be taken as the key for a URL
func TestHashing(t *testing.T) {
	cases := []struct {
		name     string
		seq      int
		expected string
	}{
		{name: "hash 1", seq: 1, expected: "1"},
		{name: "hash 2", seq: 2, expected: "2"},
		{name: "hash 10", seq: 10, expected: "a"},
		{name: "hash 98", seq: 98, expected: "1A"},
		{name: "hash 245", seq: 245, expected: "3X"},
		{name: "hash 980", seq: 980, expected: "fO"},
		{name: "hash 4235", seq: 4235, expected: "16j"},
		{name: "hash 4987453", seq: 4987453, expected: "kVsN"},
	}

	for i := range cases {
		tt := &cases[i]

		t.Run(tt.name, func(t *testing.T) {
			hash, err := NewHash(tt.seq)
			if err != nil {
				t.Fatalf("failed to generate new hash: %+v", err)
			}

			if tt.expected != hash {
				t.Fatalf("expected: %v, got: %v", tt.expected, hash)
			}

			log.Printf("got hash: %s", hash)
		})
	}
}
