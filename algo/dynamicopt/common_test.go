package dynamicopt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fibrecursion(t *testing.T) {
	// 0,1,1,2,3,5
	assert.Equal(t, 5, fibrecursion(5))
	assert.Equal(t, 5, fibdynamic(5))
}

func Test_ClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=6 => want 13",
			n:    6,
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ClimbingStairsRecursion(tt.n))
			assert.Equal(t, tt.want, ClimbingStairsDynamic(tt.n))
			assert.Equal(t, tt.want, ClimbingStairsRecWithMemo(tt.n))
		})
	}
}

func Benchmark_ClimbingStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbingStairsDynamic(i)
	}
}
