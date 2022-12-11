package shortnr

import (
	"log"
	"testing"
)

// TestHashing is a test that given an integer will use it to generate a hash
// that will be used as the key for a URL
func TestHashing(t *testing.T) {
	cases := []struct {
		name     string
		seq      int
		expected string
	}{
		{
			name:     "hash ok",
			seq:      999,
			expected: "",
		},
	}

	for i := range cases {
		tt := &cases[i]

		t.Run(tt.name, func(t *testing.T) {
			hash, err := NewHash(tt.seq)
			if err != nil {
				t.Fatalf("failed to generate new hash: %+v", err)
			}

			log.Printf("got hash: %s", hash)
		})
	}
}
