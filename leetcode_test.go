package main

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			"7 -> 0 -> 8 = (2 -> 4 -> 3) + (5 -> 6 -> 4)",
			&ListNode{
				2,
				&ListNode{
					4,
					&ListNode{
						3,
						nil,
					},
				},
			},
			&ListNode{
				5,
				&ListNode{
					6,
					&ListNode{
						4,
						nil,
					},
				},
			},
			&ListNode{
				7,
				&ListNode{
					0,
					&ListNode{
						8,
						nil,
					},
				},
			},
		},
		{
			"0 -> 3 -> 1 = (2 -> 4) + (8 -> 8)",
			&ListNode{
				2,
				&ListNode{
					4,
					nil,
				},
			},
			&ListNode{
				8,
				&ListNode{
					8,
					nil,
				},
			},
			&ListNode{
				0,
				&ListNode{
					3,
					&ListNode{
						1,
						nil,
					},
				},
			},
		},
		{
			"2 -> 4  = (2 -> 4) + (0)",
			&ListNode{
				2,
				&ListNode{
					4,
					nil,
				},
			},
			&ListNode{
				0,
				nil,
			},
			&ListNode{
				2,
				&ListNode{
					4,
					nil,
				},
			},
		},
		{
			"2 -> 4  = (0) + (2 -> 4)",
			&ListNode{
				0,
				nil,
			},
			&ListNode{
				2,
				&ListNode{
					4,
					nil,
				},
			},
			&ListNode{
				2,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddTwoNumbers(tt.l1, tt.l2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_LengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		expect  int
	}{
		{
			"abcabcbb => abc => 3",
			"abcabcbb",
			3,
		},
		{
			"bbbbbbb => b => 1",
			"bbbbbbb",
			1,
		},
		{
			"abacdef => bacdef => 6",
			"abacdef",
			6,
		},
		{
			"SPACE => SPACE => 1",
			" ",
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tt.payload)
			assert.Equal(t, tt.expect, got)
		})
	}
}

func Test_LongestPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		expect  string
	}{
		{
			"ababa => ababa",
			"ababa",
			"ababa",
		},
		{
			"aaa => aaa",
			"aaa",
			"aaa",
		},
		{
			"a => a",
			"a",
			"a",
		},
		{
			"abba => abba",
			"abba",
			"abba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, LongestPalindrome(tt.payload))
		})
	}
}

func Test_findPalindromeByMark(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		mark    int
		expect  string
	}{
		{
			"ababa => 2 => ababa",
			"ababa",
			2,
			"ababa",
		},
		{
			"ababa => 1 => aba",
			"ababa",
			1,
			"aba",
		},
		{
			"#a#b#b#a# => 4 => #a#b#b#a#",
			"#a#b#b#a#",
			4,
			"#a#b#b#a#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, findPalindromeByMark(tt.payload, tt.mark))
		})
	}
}

func Test_Zconvert(t *testing.T) {
	tests := []struct {
		name    string
		origin  string
		numRows int
		expect  string
	}{
		{
			`
abcdefg && 3 =>

a   e
b d f => aebdfcg
c   g
			`,
			"abcdefg",
			3,
			"aebdfcg",
		},
		{
			`
abcdefgh && 4 =>

a   g
b f h => agbfhced
c e
d  
			`,
			"abcdefgh",
			4,
			"agbfhced",
		},
		{
			`
ab && 1 => ab
			`,
			"ab",
			1,
			"ab",
		},
		{
			`
ab && 5 => ab
			`,
			"ab",
			5,
			"ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, Zconvert(tt.origin, tt.numRows))
		})
	}
}

func Test_newLinkList(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		f      func(node *linkNode) []string
		expect []string
	}{
		{
			"list all: from head",
			3,
			func(head *linkNode) []string {
				result := []string{}
				for node := head.next; node != head; node = node.next {
					result = append(result, node.data)
				}
				return result
			},
			[]string{"0", "1", "2"},
		},
		{
			"list all: reverse",
			4,
			func(head *linkNode) []string {
				result := []string{}
				for node := head.prior; node != head; node = node.prior {
					result = append(result, node.data)
				}
				return result
			},
			[]string{"3", "2", "1", "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := newLinkList(tt.n)
			assert.Equal(t, tt.expect, tt.f(list))
		})
	}
}

func Test_IntReverse(t *testing.T) {
	tests := []struct {
		name    string
		payload int
		expect  int
	}{
		{
			"1234 => 4321",
			1234,
			4321,
		},
		{
			"1200 => 21",
			1200,
			21,
		},
		{
			"1534236469 => 9646324351",
			1534236469,
			0,
		},
		{
			"-1234 => -4321",
			-1234,
			-4321,
		},
		{
			"-1200 => -21",
			-1200,
			-21,
		},
		{
			"max 64 => 0",
			math.MaxInt64,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, IntReverse(tt.payload))
		})
	}
}

func Test_shareStack(t *testing.T) {
	left := "left"
	right := "right"
	tests := []struct {
		name    string
		size    int
		do      func(ss *shareStack) error
		expect  []string
		wantErr bool
	}{
		{
			name: "Push 1,2 from left and 5,4,3 from right pop from left and right",
			size: 5,
			do: func(ss *shareStack) error {
				for i := 1; i <= 5; i++ {
					direct := left
					if i >= 3 {
						direct = right
					}
					if err := ss.push(strconv.Itoa(i), direct); err != nil {
						return err
					}
				}
				if _, err := ss.pop(left); err != nil {
					return err
				}
				if _, err := ss.pop(right); err != nil {
					return err
				}
				return nil
			},
			expect: []string{"1", "", "", "4", "3"},
		},
		{
			name: "error: push 1 from left 2 from right",
			size: 1,
			do: func(ss *shareStack) error {
				for i := 1; i <= 2; i++ {
					direct := left
					if i > 1 {
						direct = right
					}
					if err := ss.push(strconv.Itoa(i), direct); err != nil {
						return err
					}
				}
				return nil
			},
			expect:  nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := newShareStack(tt.size)
			err := tt.do(ss)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			if err == nil {
				assert.Equal(t, tt.expect, ss.dump())
			}
		})
	}
}

func Test_FindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name   string
		a      []int
		b      []int
		expect float64
	}{
		{
			"[1,3,5] + [2,4,6] = [1,2,3,4,5,6]",
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			3.5,
		},
		{
			"[1,3,5] + [2,4,6,7] = [1,2,3,4,5,6,7]",
			[]int{1, 3, 5},
			[]int{2, 4, 6, 7},
			4,
		},
		{
			"[1] + [] = []",
			[]int{1},
			nil,
			1,
		},
		{
			"[1] + [1] = [1,1]",
			[]int{1},
			[]int{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, FindMedianSortedArrays(tt.a, tt.b))
		})
	}
}

func Test_FindNthSmallestOf2Slice(t *testing.T) {
	tests := []struct {
		name   string
		k      int
		a      []int
		b      []int
		expect int
	}{
		{
			"4th of [1,3,5] && [2,4,6]  = 4", // 2th of 5 246; 1th 5 46
			4,
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			4,
		},
		{
			"4th of  [2,4,6] && [1,3,5]   = 4", // 2th of 5 246; 1th 5 46
			4,
			[]int{2, 4, 6},
			[]int{1, 3, 5},
			4,
		},
		{
			"4th of [1] && [2,4,6,7,8]  = 6",
			4,
			[]int{1},
			[]int{2, 4, 6, 7, 8},
			6,
		},
		{
			"4th of [2,4,6,7,8] && [1] = 6",
			4,
			[]int{2, 4, 6, 7, 8},
			[]int{1},
			6,
		},
		{
			"3th of [2,4,6,7,8] && [1] = 4",
			3,
			[]int{2, 4, 6, 7, 8},
			[]int{1},
			4,
		},
		{
			"4th of [2,4,6,8] && [1] = 6",
			4,
			[]int{2, 4, 6, 8},
			[]int{1},
			6,
		},
		{
			"4th of [2,4,6,8] && [1] = 1",
			1,
			[]int{2, 4, 6, 8},
			[]int{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, FindNthSmallestOf2Slice(tt.a, tt.b, tt.k))
		})
	}
}

func Test_isIntPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		payload int
		expect  bool
	}{
		{
			"121 == 121",
			121,
			true,
		},
		{
			"1221 == 1221",
			1221,
			true,
		},
		{
			"222 == 222",
			222,
			true,
		},
		{
			"123 != 321",
			123,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, isIntPalindrome(tt.payload))
		})
	}
}

func Test_BinarySearch(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		target  int
		expect  bool
	}{
		{
			"Search 3 from [1,3,5,6,10,100]",
			[]int{1, 3, 5, 6, 10, 100},
			3,
			true,
		},
		{
			"Search 5 from [1, 5, 6, 10, 100]",
			[]int{1, 5, 6, 10, 100},
			5,
			true,
		},
		{
			"Search 1 from [1]",
			[]int{1},
			1,
			true,
		},
		{
			"Cannot find 2 in [1,3,5,6,10,100]",
			[]int{1, 5, 6, 10, 100},
			2,
			false,
		},
		{
			"Cannot find 2 in [1]",
			[]int{1},
			2,
			false,
		},
		{
			"Search 2 from []",
			[]int{},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, binarySearch(tt.payload, tt.target))
		})
	}
}

func Test_MaxArea(t *testing.T) {
	tests := []struct {
		name    string
		expect  int
		payload []int
	}{
		{
			name:    "[1,3,2] => 2",
			expect:  2,
			payload: []int{1, 3, 2},
		},
		{
			name:    "[2,1,3] => 4",
			expect:  4,
			payload: []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, maxArea(tt.payload))
		})
	}
}

func Test_trap(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    int
	}{
		{
			name:    "[0,1,0,2,1,0,1,3,2,1,2,1] => 6 ",
			payload: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:    6,
		},
		{
			name:    "[4,2,3] => 1",
			payload: []int{4, 2, 3},
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, trap(tt.payload))
		})
	}
}

// func Test_threeSum(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		payload []int
// 		want    [][]int
// 	}{
// 		{
// 			name:    "[-1, 0, 1, 2, -1, -4] =>  [-1, 0, 1] and [-1, -1, 2]",
// 			payload: []int{-1, 0, 1, 2, -1, -4},
// 			want: [][]int{
// 				[]int{-1, 0, 1},
// 				[]int{-1, -1, 2},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			assert.ElementsMatch(t, tt.want, threeSum(tt.payload))
// 		})
// 	}
// }
