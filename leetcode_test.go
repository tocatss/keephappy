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

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    [][]int
	}{
		{
			name: `[-1, 0, 1, 2, -1, -4] 
						=> 
						  [-4, -1, -1, 0, 1, 2]
						=> 
						  [-1,-1,2],[-1,0,1]`,
			payload: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				[]int{-1, 0, 1},
				[]int{-1, -1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, ThreeSum(tt.payload))
		})
	}
}

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    []string
	}{
		{
			name:    `23 => ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]`,
			payload: "23",
			want:    []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:    `123 => ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]`,
			payload: "123",
			want:    []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:    `1234 => ...`,
			payload: "1234",
			want: []string{"adg", "aeg", "afg", "bdg", "beg", "bfg", "cdg", "ceg", "cfg",
				"adh", "aeh", "afh", "bdh", "beh", "bfh", "cdh", "ceh", "cfh",
				"adi", "aei", "afi", "bdi", "bei", "bfi", "cdi", "cei", "cfi",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, LetterCombinations(tt.payload))
		})
	}
}

func Test_IntToRoman(t *testing.T) {
	tests := []struct {
		name    string
		payload int
		want    string
	}{
		{
			name:    "112 => CXII",
			payload: 112,
			want:    "CXII",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IntToRoman(tt.payload))
		})
	}
}

func Test_RomanToInt(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    int
	}{
		{
			name:    "CXII => 112",
			payload: "CXII",
			want:    112,
		},
		{
			name:    "XLI => 41",
			payload: "XLI",
			want:    41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RomanToInt(tt.payload))
		})
	}
}

func Test_LongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name    string
		payload []string
		want    string
	}{
		{
			name:    "[ab, abc] => ab",
			payload: []string{"ab", "abc"},
			want:    "ab",
		},
		{
			name:    "[abc, ab] => ab",
			payload: []string{"abc", "ab"},
			want:    "ab",
		},
		{
			name:    "[ab, abc,c] => ab",
			payload: []string{"ab", "abc", "c"},
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LongestCommonPrefix(tt.payload))
		})
	}
}

func Test_ThreeSumClosest(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		target  int
		want    int
	}{
		{
			name:    `target 5 of [-1, 0, 1, 2]  => [0,1,2]`,
			payload: []int{-1, 0, 1, 2},
			target:  5,
			want:    3,
		},
		{
			name:    `target 1 of [-10,-4,-1,-1,2,1,4,5,9]  => [-1,2,1]`,
			payload: []int{-10, -4, -1, -1, 2, 1, 4, 5, 9},
			target:  17,
			want:    16,
		},
		{
			name:    `target 0 of [0,0,0]  => [0,0,0]`,
			payload: []int{0, 0, 0},
			target:  1,
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ThreeSumClosest(tt.payload, tt.target))
		})
	}
}

func Test_FourSum(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		target  int
		want    [][]int
	}{
		{
			name:    "nums = [1, 0, -1, 0, -2, 2] and target = 0",
			payload: []int{1, 0, -1, -1, 0, -2, 2, -2},
			target:  0,
			want: [][]int{
				[]int{-1, 0, 0, 1},
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-1, -1, 0, 2},
			},
		},
		{
			name:    "nums = [0,0,4,-2,-3,-2,-2,-3] and target = -1",
			payload: []int{0, 0, 4, -2, -3, -2, -2, -3},
			target:  -1,
			want: [][]int{
				[]int{-3, -2, 0, 4},
			},
		},
		{
			name:    "nums = [-1,0,-5,-2,-2,-4,0,1,-2] and target = -9",
			payload: []int{-1, 0, -5, -2, -2, -4, 0, 1, -2},
			target:  -9,
			want: [][]int{
				[]int{-5, -4, -1, 1},
				[]int{-5, -4, 0, 0},
				[]int{-5, -2, -2, 0},
				[]int{-4, -2, -2, -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, FourSum(tt.payload, tt.target))
		})
	}
}

func Test_RemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		n    int
		want *ListNode
	}{
		{
			name: "1->2->3->4->5 and n = 2 => 1->2->3->5",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			n: 2,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveNthFromEnd(tt.head, tt.n))
		})
	}
}

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name    string
		payload []*ListNode
		want    *ListNode
	}{
		{
			name: "merge [1,2,3] [2,3,4] [3,4,5] => [1,2,2,3,3,3,4,4,5]",
			payload: []*ListNode{
				{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 3,
									Next: &ListNode{
										Val: 4,
										Next: &ListNode{
											Val: 4,
											Next: &ListNode{
												Val: 5,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeKLists(tt.payload))
		})
	}
}

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		name    string
		payload *ListNode
		k       int
		want    *ListNode
	}{
		{
			name: "k=2,1->2->3->4 => 2->1->4->3",
			payload: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			k: 2,
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
		{
			name: "k=3,1->2->3->4 => 3->2->1->4",
			payload: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			k: 3,
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseKGroup(tt.payload, tt.k))
		})
	}
}

func Test_reverseKGroupAgain(t *testing.T) {
	tests := []struct {
		name    string
		payload *ListNode
		k       int
		want    *ListNode
	}{
		{
			name: "k=2,1->2->3->4 => 2->1->4->3",
			payload: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			k: 2,
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
		{
			name: "k=3,1->2->3->4 => 3->2->1->4",
			payload: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			k: 3,
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ReverseKGroupAgain(&copy, tt.k))
		})
	}
}
func Test_RemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "112234 => 1234 => 4",
			nums: []int{1, 1, 2, 2, 3, 4},
			want: 4,
		},
		{
			name: "12222 => 12 => 2",
			nums: []int{1, 2, 2, 2, 2},
			want: 2,
		},
		{
			name: "1 => 1 => 1",
			nums: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveDuplicates(tt.nums))
		})
	}
}

func Test_RemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "112234,target=2 => 1134",
			target: 2,
			nums:   []int{1, 1, 2, 2, 3, 4},
			want:   4,
		},
		{
			name:   "12222, target=1 => 2222",
			target: 1,
			nums:   []int{1, 2, 2, 2, 2},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveElement(tt.nums, tt.target))
		})
	}
}

func Test_MoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "011002 =>112",
			nums: []int{0, 1, 1, 0, 0, 2},
			want: []int{1, 1, 2, 0, 0, 0},
		},
		{
			name: "1010010020 =>1112",
			nums: []int{1, 0, 1, 0, 0, 1, 0, 0, 2, 0},
			want: []int{1, 1, 1, 2, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.nums)
			assert.Equal(t, tt.want, tt.nums)
		})
	}
}

func Test_RemoveDuplicates2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "00011123 =>001123 => 6",
			nums: []int{0, 0, 1, 1, 2, 3},
			want: 6,
		},
		// 1,0 2,0 3,0 4,1 4,1 5,1 6,1 7,2 7,2 8,3
		// 1,0 2,0 2,0 2,0 2,1 3,0 4,1 4,1 4,2 5,1
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RemoveDuplicates2(tt.nums))
		})
	}
}

func Test_Divide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{
			name:     "10 / 3 => 3",
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			name:     "14 / 3 => 3",
			dividend: 14,
			divisor:  3,
			want:     4,
		},
		{
			name:     "3 / 1 => 3",
			dividend: 3,
			divisor:  1,
			want:     3,
		},
		{
			name:     "2 / 1 => 2",
			dividend: 2,
			divisor:  1,
			want:     2,
		},
		{
			name:     "4 / 1 => 4",
			dividend: 4,
			divisor:  1,
			want:     4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Divide(tt.dividend, tt.divisor))
		})
	}
}

func Test_FindSubstring(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  []int
	}{
		{
			name:  "Index [aa] in aaa",
			s:     "aaa",
			words: []string{"a", "a"},
			want:  []int{0, 1},
		},
		{
			name:  "Index [a] in aaa",
			s:     "aaa",
			words: []string{"a"},
			want:  []int{0, 1, 2},
		},
		{
			name:  "Index [a,b] in abcacba",
			s:     "abcacba",
			words: []string{"a", "b"},
			want:  []int{0, 5},
		},
		{
			name:  "Index [,] in abcacba",
			s:     "abcacba",
			words: []string{"", ""},
			want:  []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, FindSubstring(tt.s, tt.words))
		})
	}
}

func Test_NextPermutation(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    []int
	}{
		{
			name:    "123 => 132",
			payload: []int{1, 2, 3},
			want:    []int{1, 3, 2},
		},
		{
			name:    "321 => 123",
			payload: []int{3, 2, 1},
			want:    []int{1, 2, 3},
		},
		{
			name:    "132 => 213",
			payload: []int{1, 3, 2},
			want:    []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NextPermutation(tt.payload)
			assert.Equal(t, tt.want, tt.payload)
		})
	}
}

func Test_SearchTarget(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "nums = [4,5,6,7,0,1,2], target = 0",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "nums = [10,12,0,2,4,6,8], target = 6",
			nums:   []int{10, 12, 0, 2, 4, 6, 8},
			target: 6,
			want:   5,
		},
		{
			name:   "nums = [10,12,14,16,0,2,4,6,8], target 14",
			nums:   []int{10, 12, 14, 16, 0, 2, 4, 6, 8},
			target: 14,
			want:   2,
		},
		{
			name:   "nums = [3,1], target 1",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SearchTarget(tt.nums, tt.target))
		})
	}
}

func Test_LongestValidParentheses(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    int
	}{
		{
			"(())() )) => valued is 6",
			"(())()))",
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LongestValidParentheses(tt.payload))
		})
	}
}
func Test_SearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			"nums = [5,7,7,8,8,10] target=8; want: [3,4]",
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			"nums = [5,7,7,8,8,10] target=8; want: [3,4]",
			[]int{5, 7, 7, 8, 8, 10},
			6,
			[]int{-1, -1},
		},
		{
			"nums = [1] target=1; want: [1]",
			[]int{1},
			1,
			[]int{0, 0},
		},
		{
			"nums = [1,1] target=1; want: [0,1]",
			[]int{1, 1},
			1,
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SearchRange(tt.nums, tt.target))
		})
	}
}

func Test_MyPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{
			name: "2^3 =8",
			x:    2.0,
			n:    3,
			want: 8.0,
		},
		{
			name: "-2^3 =-8",
			x:    -2.0,
			n:    3,
			want: -8.0,
		},
		{
			name: "2^-3 =1/8",
			x:    2,
			n:    -3,
			want: 1.0 / 8.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MyPow(tt.x, tt.n))
			assert.Equal(t, tt.want, MyPow(tt.x, tt.n))
		})
	}
}

func Test_SingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[1,2,1,2,3] => 3",
			nums: []int{1, 2, 1, 2, 3},
			want: 3,
		},
		{
			name: "[1,2,1,2,3] => 3",
			nums: []int{2, 1, 1, 2, 3},
			want: 3,
		},
		{
			name: "[1,2,1,2,3] => 3",
			nums: []int{2, 1, 3, 1, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SingleNumber(tt.nums))
		})
	}
}

func Test_subarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "find 2 from [1,1,1]",
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			name: "find 1 from [1,1,1]",
			nums: []int{1, 1, 1},
			k:    1,
			want: 3,
		},
		{
			name: "find 3 from [1,1,1]",
			nums: []int{1, 1, 1},
			k:    3,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, subarraySum(tt.nums, tt.k))
			assert.Equal(t, tt.want, subarraySumWithPreSum(tt.nums, tt.k))
		})
	}
}
