package functor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewUserWithOptions(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		want *user
	}{
		{
			name: "name is tocatss",
			opts: []Option{WithName("tocatss")},
			want: &user{
				name:     "tocatss",
				sex:      "default",
				password: "default",
				imgURL:   "default",
			},
		},
		{
			name: "name is tocatss, password is 123",
			opts: []Option{WithPassword("123"), WithName("tocatss")},
			want: &user{
				name:     "tocatss",
				sex:      "default",
				password: "123",
				imgURL:   "default",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewUserWithOptions(tt.opts...))
		})
	}
}
