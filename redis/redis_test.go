package redis

import "testing"

func Test_NewRedisClient(t *testing.T) {
	if _, err := NewRedisClient(); err != nil {
		t.Fatal(err)
	}
}
