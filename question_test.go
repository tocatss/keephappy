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
			"5 + 5 => 0 -> 1",
			&ListNode{
				5,
				nil,
			},
			&ListNode{
				5,
				nil,
			},
			&ListNode{
				0,
				&ListNode{
					1,
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
			assert.Equal(t, tt.want, ReverseKGroupAgain(tt.payload, tt.k))
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

func Test_findOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			name:       "ok: 0 -> 1,2 -> 3 => 0,1,2,3",
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0}, {2, 0}, {3, 1}, {3, 2},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name:       "ok: 0 -> 1,2 -> 3 ->4 => 0,1,2,3,4",
			numCourses: 5,
			prerequisites: [][]int{
				{1, 0}, {2, 0}, {3, 1}, {3, 2}, {4, 3},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name:       "not ok: 0 -> 1,2 -> 3 -> 2 => nil",
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0}, {2, 0}, {3, 1}, {3, 2}, {2, 3},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findOrder(tt.numCourses, tt.prerequisites))
		})
	}
}

func Test_ComputeSuffix(t *testing.T) {
	tests := []struct {
		name   string
		suffix []string
		want   int
	}{
		{
			name:   "1+2+3 => 12+3+ => 6",
			suffix: []string{"1", "2", "+", "3", "+"},
			want:   6,
		},
		{
			name:   "1+2x3 => 123x+ => 7",
			suffix: []string{"1", "2", "3", "*", "+"},
			want:   7,
		},
		{
			name:   "1+2x3-10 => 123x+10- => 6",
			suffix: []string{"1", "2", "3", "*", "+", "10", "-"},
			want:   -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ComputeSuffix(tt.suffix))
		})
	}
}

func Test_Convert2Suffix(t *testing.T) {
	tests := []struct {
		name  string
		infix []string
		want  []string
	}{
		{
			name:  "1+2+3 => 12+3+",
			infix: []string{"1", "+", "2", "+", "3"},
			want:  []string{"1", "2", "+", "3", "+"},
		},
		{
			name:  "1+2x3 => 123x+",
			infix: []string{"1", "+", "2", "*", "3"},
			want:  []string{"1", "2", "3", "*", "+"},
		},
		{
			name:  "1+2x3-10 => 123x+10-",
			infix: []string{"1", "+", "2", "*", "3", "-", "10"},
			want:  []string{"1", "2", "3", "*", "+", "10", "-"},
		},
		{
			name:  "1+2x(3-10) => 12310-x+",
			infix: []string{"1", "+", "2", "*", "(", "3", "-", "10", ")"},
			want:  []string{"1", "2", "3", "10", "-", "*", "+"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Convert2Suffix(tt.infix))
		})
	}
}

func Test_IPConvert(t *testing.T) {
	tests := []struct {
		name    string
		address string
		number  int
	}{
		// 下面的网站可以做转换验证。
		// http://www.ab173.com/net/ip2int.php
		{
			name:    "0.2.3.1 => 131841 => 0.2.3.1",
			address: "0.2.3.1",
			number:  131841,
		},
		{
			name:    "",
			address: "255.255.255.1",
			number:  4294967041,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.number, Ip2Number(tt.address))
			assert.Equal(t, tt.address, Number2Ip(tt.number))
		})
	}
}

func Test_Hanoi(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int // 步数
	}{
		{
			name: "3 个盘子",
			n:    3,
			want: 7,
		},
		{
			name: "4 个盘子",
			n:    4,
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Hanoi(tt.n))
		})
	}
}

func Test_MaximumMatching(t *testing.T) {
	tests := []struct {
		name  string
		m     map[string]interface{}
		input string
		want  []string
	}{
		{
			name: "MaximumMatching with rune",
			m: map[string]interface{}{
				"今天":   nil,
				"天气真好": nil,
				"天气":   nil,
				"真好":   nil,
				"我们":   nil,
				"出去玩耍": nil,
				"出去":   nil,
				"玩耍":   nil,
			},
			input: "今天的天气真好我们出去玩耍吧",
			want:  []string{"今天", "的", "天气真好", "我们", "出去玩耍", "吧"},
		},
		{
			name: "MaximumMatching with byte",
			m: map[string]interface{}{
				"bc":     nil,
				"bcde":   nil,
				"bcdek":  nil,
				"deskld": nil,
				"kld":    nil,
				"kl":     nil,
				"kldzf":  nil,
			},
			input: "abcdeskld",
			want:  []string{"a", "bcde", "s", "kld"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaximumMatching(tt.input, tt.m))
		})
	}
}

func Test_Gcd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "2 and 3 => 1",
			a:    2,
			b:    3,
			want: 1,
		},
		{
			name: "24 and 36 => 12",
			a:    24,
			b:    36,
			want: 12,
		},
		{
			name: "24 and 8 => 8",
			a:    8,
			b:    24,
			want: 8,
		},
		{
			name: "invalid: 24 and 0 => -1",
			a:    0,
			b:    24,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Gcd(tt.a, tt.b))
		})
	}
}

func Test_RectangleMNM(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		want int
	}{
		{
			name: "3x3 => 9 + 6 + 3 + 6 + 4 + 2 + 3 + 2 + 1",
			n:    3,
			m:    3,
			want: 36,
		},
		{
			name: "3x2 => 6 + 4 + 2 + 3 + 2 + 1",
			n:    3,
			m:    2,
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RectangleMN(tt.n, tt.m))
		})
	}
}

func Test_SquareMN(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		want int
	}{
		{
			name: "3x3 =>6 + 4 + 2 + 4 + 2 + 1 + 2 + 1 + 0",
			n:    3,
			m:    3,
			want: 22,
		},
		{
			name: "3x2 => 4 + 2 + 1 + 2 + 1 + 0",
			n:    3,
			m:    2,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SquareMN(tt.n, tt.m))
		})
	}
}

func Test_MaxProfix(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "[7,1,5,3,6,4] => 1 买入，5 卖出，3买入，6卖出",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			name:   "[1,2,3,4,5] => 1 买入，5 卖出",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "[7,6,5,4,3] => 0",
			prices: []int{7, 6, 5, 4, 3},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaxProfit(tt.prices))
		})
	}
}

func Test_MaxProfitWithTwice(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "[1,2,4,2,5,7,2,4,9,0] => buy 1 sell 7, buy 2 sell 9",
			prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			want:   13,
		},
		// {
		// 	name:   "[1,2,4,2,5,7,2,4,9,0] => buy 1 sell 7, buy 2 sell 9",
		// 	prices: []int{397, 6621, 4997, 7506, 8918, 1662, 9187, 3278, 3890, 514, 18, 9305, 93, 5508, 3031, 2692, 6019, 1134, 1691, 4949, 5071, 799, 8953, 7882, 4273, 302, 6753, 4657, 8368, 3942, 1982, 5117, 563, 3332, 2623, 9482, 4994, 8163, 9112, 5236, 5029, 5483, 4542, 1474, 991, 3925, 4166, 3362, 5059, 5857, 4663, 6482, 3008, 3616, 4365, 3634, 270, 1118, 8291, 4990, 1413, 273, 107, 1976, 9957, 9083, 7810, 4952, 7246, 3275, 6540, 2275, 8758, 7434, 3750, 6101, 1359, 4268, 5815, 2771, 126, 478, 9253, 9486, 446, 3618, 3120, 7068, 1089, 1411, 2058, 2502, 8037, 2165, 830, 7994, 1248, 4993, 9298, 4846, 8268, 2191, 3474, 3378, 9625, 7224, 9479, 985, 1492, 1646, 3756, 7970, 8476, 3009, 7457, 8922, 2980, 577, 2342, 4069, 8341, 4400, 2923, 2730, 2917, 105, 724, 518, 5098, 6375, 5364, 3366, 8566, 8838, 3096, 8191, 2414, 2575, 5528, 259, 573, 5636, 4581, 9049, 4998, 2038, 4323, 7978, 8968, 6665, 8399, 7309, 7417, 1322, 6391, 335, 1427, 7115, 853, 2878, 9842, 2569, 2596, 4760, 7760, 5693, 9304, 6526, 8268, 4832, 6785, 5194, 6821, 1367, 4243, 1819, 9757, 4919, 6149, 8725, 7936, 4548, 2386, 5354, 2222, 8777, 2041, 1, 2245, 9246, 2879, 8439, 1815, 5476, 3200, 5927, 7521, 2504, 2454, 5789, 3688, 9239, 7335, 6861, 6958, 7931, 8680, 3068, 2850, 1181, 1793, 7138, 2081, 532, 2492, 4303, 5661, 885, 657, 4258, 131, 9888, 9050, 1947, 1716, 2250, 4226, 9237, 1106, 6680, 1379, 1146, 2272, 8714, 8008, 9230, 6645, 3040, 2298, 5847, 4222, 444, 2986, 2655, 7328, 1830, 6959, 9341, 2716, 3968, 9952, 2847, 3856, 9002, 1146, 5573, 1252, 5373, 1162, 8710, 2053, 2541, 9856, 677, 1256, 4216, 9908, 4253, 3609, 8558, 6453, 4183, 5354, 9439, 6838, 2682, 7621, 149, 8376, 337, 4117, 8328, 9537, 4326, 7330, 683, 9899, 4934, 2408, 7413, 9996, 814, 9955, 9852, 1491, 7563, 421, 7751, 1816, 4030, 2662, 8269, 8213, 8016, 4060, 5051, 7051, 1682, 5201, 5427, 8371, 5670, 3755, 7908, 9996, 7437, 4944, 9895, 2371, 7352, 3661, 2367, 4518, 3616, 8571, 6010, 1179, 5344, 113, 9347, 9374, 2775, 3969, 3939, 792, 4381, 8991, 7843, 2415, 544, 3270, 787, 6214, 3377, 8695, 6211, 814, 9991, 2458, 9537, 7344, 6119, 1904, 8214, 6087, 6827, 4224, 7266, 2172, 690, 2966, 7898, 3465, 3287, 1838, 609, 7668, 829, 8452, 84, 7725, 8074, 871, 3939, 7803, 5918, 6502, 4969, 5910, 5313, 4506, 9606, 1432, 2762, 7820, 3872, 9590, 8397, 1138, 8114, 9087, 456, 6012, 8904, 3743, 7850, 9514, 7764, 5031, 4318, 7848, 9108, 8745, 5071, 9400, 2900, 7341, 5902, 7870, 3251, 7567, 2376, 9209, 9000, 1491, 7030, 2872, 7433, 1779, 362, 5547, 7218, 7171, 7911, 2474, 914, 2114, 8340, 8678, 3497, 2659, 2878, 2606, 7756, 7949, 2006, 656, 5291, 4260, 8526, 4894, 1828, 7255, 456, 7180, 8746, 3838, 6404, 6179, 5617, 3118, 8078, 9187, 289, 5989, 1661, 1204, 8103, 2, 6234, 7953, 9013, 5465, 559, 6769, 9766, 2565, 7425, 1409, 3177, 2304, 6304, 5005, 9559, 6760, 2185, 4657, 598, 8589, 836, 2567, 1708, 5266, 1754, 8349, 1255, 9767, 5905, 5711, 9769, 8492, 3664, 5134, 3957, 575, 1903, 3723, 3140, 5681, 5133, 6317, 4337, 7789, 7675, 3896, 4549, 6212, 8553, 1499, 1154, 5741, 418, 9214, 1007, 2172, 7563, 8614, 8291, 3469, 677, 4413, 1961, 4341, 9547, 5918, 4916, 7803, 9641, 4408, 3484, 1126, 7078, 7821, 8915, 1105, 8069, 9816, 7317, 2974, 1315, 8471, 8715, 1733, 7685, 6074, 257, 5249, 4688, 8549, 5070, 5366, 2962, 7031, 6059, 8861, 9301, 7328, 6664, 5294, 8088, 6500, 6421, 1518, 4321, 5336, 2623, 8742, 1505, 9941, 1716, 2820, 4764, 6783, 906, 2450, 2857, 7515, 4051, 7546, 2416, 9121, 9264, 1730, 6152, 1675, 592, 1805, 9003, 7256, 7099, 3444, 3757, 9872, 4962, 4430, 1561, 7586, 3173, 3066, 3879, 1241, 2238, 8643, 8025, 3144, 7445, 882, 7012, 1496, 4780, 9428, 617, 396, 1159, 3121, 2072, 1751, 4926, 7427, 5359, 8378, 871, 5468, 8250, 5834, 9899, 9811, 9772, 9424, 2877, 3651, 7017, 5116, 8646, 5042, 4612, 6092, 2277, 1624, 7588, 3409, 1053, 8206, 3806, 8564, 7679, 2230, 6667, 8958, 6009, 2026, 7336, 6881, 3847, 5586, 9067, 98, 1750, 8839, 9522, 4627, 8842, 2891, 6095, 7488, 7934, 708, 3580, 6563, 8684, 7521, 9972, 6089, 2079, 130, 4653, 9758, 2360, 1320, 8716, 8370, 9699, 6052, 1603, 3546, 7991, 670, 3644, 6093, 9509, 9518, 7072, 4703, 2409, 3168, 2191, 6695, 228, 2124, 3258, 5264, 9645, 9583, 1354, 1724, 9713, 2359, 1482, 8426, 3680, 6551, 3148, 9731, 8955, 4751, 9629, 6946, 5421, 9625, 9391, 1282, 5495, 6464, 5985, 4256, 5984, 4528, 952, 6212, 6652, 562, 1476, 6297, 145, 9182, 8021, 6211, 1542, 5856, 4637, 1574, 2407, 7785, 1305, 1362, 2536, 934, 4661, 4309, 559, 4052, 1943, 2406, 516, 4280, 6662, 2852, 8808, 7614, 9064, 1813, 4529, 6893, 8110, 4674, 2427, 2484, 7237, 3969, 8340, 1874, 5543, 7099, 6011, 3200, 8461, 8547, 486, 9474, 9208, 7397, 9879, 7503, 9803, 6747, 1783, 6466, 9600, 6944, 432, 8664, 8757, 4961, 1909, 6867, 5988, 4337, 5703, 3225, 4658, 4043, 1452, 6554, 1142, 7463, 9754, 5956, 2363, 241, 1782, 7923, 7638, 1661, 5427, 3794, 8409, 7210, 260, 8009, 4154, 692, 3025, 9263, 2006, 4935, 2483, 7994, 5624, 8186, 7571, 282, 8582, 9023, 6836, 6076, 6487, 6591, 2032, 8850, 3184, 3815, 3125, 7174, 5476, 8552, 968, 3885, 2115, 7580, 8246, 2621, 4625, 1272, 1885, 6631, 6207, 4368, 4625, 8183, 2554, 8548, 8465, 1136, 7572, 1654, 7213, 411, 4597, 5597, 5613, 7781, 5764, 8738, 1307, 7593, 7291, 8628, 7830, 9406, 6208, 6077, 2027, 833, 7349, 3912, 7464, 9908, 4632, 8441, 8091, 7187, 6990, 2908, 4675, 914, 4562, 8240, 1325, 9159, 190, 6938, 3292, 5954, 2028, 4600, 9899, 9319, 3228, 7730, 5077, 9436, 159, 7105, 6622, 7508, 7369, 4086, 3768, 2002, 8880, 8211, 5541, 2222, 1119, 216, 3136, 5682, 4809, 813, 1193, 4999, 4103, 4486, 7305, 6131, 9086, 7205, 5451, 2314, 1287, 528, 8102, 1446, 3985, 4724, 5306, 1355, 5163, 9074, 9709, 4043, 7285, 5250, 2617, 4756, 1818, 2105, 6790, 6627, 2918, 7984, 7978, 7021, 2470, 1636, 3152, 7908, 8841, 4955, 222, 6480, 5484, 4676, 7926, 5821, 9401, 3232, 7176, 916, 8658, 3237, 1311, 5943, 8487, 3928, 7051, 306, 6033, 3842, 3285, 8951, 1826, 7616, 2324, 648, 9252, 5476, 8556, 4445, 6784},
		// 	want:   13,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaxProfitWithTwice(tt.prices))
		})
	}
}

// func Test_MaxProfitWithTwiceWithoutCache(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		prices []int
// 		want   int
// 	}{
// 		// {
// 		// 	name:   "[1,2,4,2,5,7,2,4,9,0] => buy 1 sell 7, buy 2 sell 9",
// 		// 	prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
// 		// 	want:   13,
// 		// },
// 		{
// 			name:   "[1,2,4,2,5,7,2,4,9,0] => buy 1 sell 7, buy 2 sell 9",
// 			prices: []int{397, 6621, 4997, 7506, 8918, 1662, 9187, 3278, 3890, 514, 18, 9305, 93, 5508, 3031, 2692, 6019, 1134, 1691, 4949, 5071, 799, 8953, 7882, 4273, 302, 6753, 4657, 8368, 3942, 1982, 5117, 563, 3332, 2623, 9482, 4994, 8163, 9112, 5236, 5029, 5483, 4542, 1474, 991, 3925, 4166, 3362, 5059, 5857, 4663, 6482, 3008, 3616, 4365, 3634, 270, 1118, 8291, 4990, 1413, 273, 107, 1976, 9957, 9083, 7810, 4952, 7246, 3275, 6540, 2275, 8758, 7434, 3750, 6101, 1359, 4268, 5815, 2771, 126, 478, 9253, 9486, 446, 3618, 3120, 7068, 1089, 1411, 2058, 2502, 8037, 2165, 830, 7994, 1248, 4993, 9298, 4846, 8268, 2191, 3474, 3378, 9625, 7224, 9479, 985, 1492, 1646, 3756, 7970, 8476, 3009, 7457, 8922, 2980, 577, 2342, 4069, 8341, 4400, 2923, 2730, 2917, 105, 724, 518, 5098, 6375, 5364, 3366, 8566, 8838, 3096, 8191, 2414, 2575, 5528, 259, 573, 5636, 4581, 9049, 4998, 2038, 4323, 7978, 8968, 6665, 8399, 7309, 7417, 1322, 6391, 335, 1427, 7115, 853, 2878, 9842, 2569, 2596, 4760, 7760, 5693, 9304, 6526, 8268, 4832, 6785, 5194, 6821, 1367, 4243, 1819, 9757, 4919, 6149, 8725, 7936, 4548, 2386, 5354, 2222, 8777, 2041, 1, 2245, 9246, 2879, 8439, 1815, 5476, 3200, 5927, 7521, 2504, 2454, 5789, 3688, 9239, 7335, 6861, 6958, 7931, 8680, 3068, 2850, 1181, 1793, 7138, 2081, 532, 2492, 4303, 5661, 885, 657, 4258, 131, 9888, 9050, 1947, 1716, 2250, 4226, 9237, 1106, 6680, 1379, 1146, 2272, 8714, 8008, 9230, 6645, 3040, 2298, 5847, 4222, 444, 2986, 2655, 7328, 1830, 6959, 9341, 2716, 3968, 9952, 2847, 3856, 9002, 1146, 5573, 1252, 5373, 1162, 8710, 2053, 2541, 9856, 677, 1256, 4216, 9908, 4253, 3609, 8558, 6453, 4183, 5354, 9439, 6838, 2682, 7621, 149, 8376, 337, 4117, 8328, 9537, 4326, 7330, 683, 9899, 4934, 2408, 7413, 9996, 814, 9955, 9852, 1491, 7563, 421, 7751, 1816, 4030, 2662, 8269, 8213, 8016, 4060, 5051, 7051, 1682, 5201, 5427, 8371, 5670, 3755, 7908, 9996, 7437, 4944, 9895, 2371, 7352, 3661, 2367, 4518, 3616, 8571, 6010, 1179, 5344, 113, 9347, 9374, 2775, 3969, 3939, 792, 4381, 8991, 7843, 2415, 544, 3270, 787, 6214, 3377, 8695, 6211, 814, 9991, 2458, 9537, 7344, 6119, 1904, 8214, 6087, 6827, 4224, 7266, 2172, 690, 2966, 7898, 3465, 3287, 1838, 609, 7668, 829, 8452, 84, 7725, 8074, 871, 3939, 7803, 5918, 6502, 4969, 5910, 5313, 4506, 9606, 1432, 2762, 7820, 3872, 9590, 8397, 1138, 8114, 9087, 456, 6012, 8904, 3743, 7850, 9514, 7764, 5031, 4318, 7848, 9108, 8745, 5071, 9400, 2900, 7341, 5902, 7870, 3251, 7567, 2376, 9209, 9000, 1491, 7030, 2872, 7433, 1779, 362, 5547, 7218, 7171, 7911, 2474, 914, 2114, 8340, 8678, 3497, 2659, 2878, 2606, 7756, 7949, 2006, 656, 5291, 4260, 8526, 4894, 1828, 7255, 456, 7180, 8746, 3838, 6404, 6179, 5617, 3118, 8078, 9187, 289, 5989, 1661, 1204, 8103, 2, 6234, 7953, 9013, 5465, 559, 6769, 9766, 2565, 7425, 1409, 3177, 2304, 6304, 5005, 9559, 6760, 2185, 4657, 598, 8589, 836, 2567, 1708, 5266, 1754, 8349, 1255, 9767, 5905, 5711, 9769, 8492, 3664, 5134, 3957, 575, 1903, 3723, 3140, 5681, 5133, 6317, 4337, 7789, 7675, 3896, 4549, 6212, 8553, 1499, 1154, 5741, 418, 9214, 1007, 2172, 7563, 8614, 8291, 3469, 677, 4413, 1961, 4341, 9547, 5918, 4916, 7803, 9641, 4408, 3484, 1126, 7078, 7821, 8915, 1105, 8069, 9816, 7317, 2974, 1315, 8471, 8715, 1733, 7685, 6074, 257, 5249, 4688, 8549, 5070, 5366, 2962, 7031, 6059, 8861, 9301, 7328, 6664, 5294, 8088, 6500, 6421, 1518, 4321, 5336, 2623, 8742, 1505, 9941, 1716, 2820, 4764, 6783, 906, 2450, 2857, 7515, 4051, 7546, 2416, 9121, 9264, 1730, 6152, 1675, 592, 1805, 9003, 7256, 7099, 3444, 3757, 9872, 4962, 4430, 1561, 7586, 3173, 3066, 3879, 1241, 2238, 8643, 8025, 3144, 7445, 882, 7012, 1496, 4780, 9428, 617, 396, 1159, 3121, 2072, 1751, 4926, 7427, 5359, 8378, 871, 5468, 8250, 5834, 9899, 9811, 9772, 9424, 2877, 3651, 7017, 5116, 8646, 5042, 4612, 6092, 2277, 1624, 7588, 3409, 1053, 8206, 3806, 8564, 7679, 2230, 6667, 8958, 6009, 2026, 7336, 6881, 3847, 5586, 9067, 98, 1750, 8839, 9522, 4627, 8842, 2891, 6095, 7488, 7934, 708, 3580, 6563, 8684, 7521, 9972, 6089, 2079, 130, 4653, 9758, 2360, 1320, 8716, 8370, 9699, 6052, 1603, 3546, 7991, 670, 3644, 6093, 9509, 9518, 7072, 4703, 2409, 3168, 2191, 6695, 228, 2124, 3258, 5264, 9645, 9583, 1354, 1724, 9713, 2359, 1482, 8426, 3680, 6551, 3148, 9731, 8955, 4751, 9629, 6946, 5421, 9625, 9391, 1282, 5495, 6464, 5985, 4256, 5984, 4528, 952, 6212, 6652, 562, 1476, 6297, 145, 9182, 8021, 6211, 1542, 5856, 4637, 1574, 2407, 7785, 1305, 1362, 2536, 934, 4661, 4309, 559, 4052, 1943, 2406, 516, 4280, 6662, 2852, 8808, 7614, 9064, 1813, 4529, 6893, 8110, 4674, 2427, 2484, 7237, 3969, 8340, 1874, 5543, 7099, 6011, 3200, 8461, 8547, 486, 9474, 9208, 7397, 9879, 7503, 9803, 6747, 1783, 6466, 9600, 6944, 432, 8664, 8757, 4961, 1909, 6867, 5988, 4337, 5703, 3225, 4658, 4043, 1452, 6554, 1142, 7463, 9754, 5956, 2363, 241, 1782, 7923, 7638, 1661, 5427, 3794, 8409, 7210, 260, 8009, 4154, 692, 3025, 9263, 2006, 4935, 2483, 7994, 5624, 8186, 7571, 282, 8582, 9023, 6836, 6076, 6487, 6591, 2032, 8850, 3184, 3815, 3125, 7174, 5476, 8552, 968, 3885, 2115, 7580, 8246, 2621, 4625, 1272, 1885, 6631, 6207, 4368, 4625, 8183, 2554, 8548, 8465, 1136, 7572, 1654, 7213, 411, 4597, 5597, 5613, 7781, 5764, 8738, 1307, 7593, 7291, 8628, 7830, 9406, 6208, 6077, 2027, 833, 7349, 3912, 7464, 9908, 4632, 8441, 8091, 7187, 6990, 2908, 4675, 914, 4562, 8240, 1325, 9159, 190, 6938, 3292, 5954, 2028, 4600, 9899, 9319, 3228, 7730, 5077, 9436, 159, 7105, 6622, 7508, 7369, 4086, 3768, 2002, 8880, 8211, 5541, 2222, 1119, 216, 3136, 5682, 4809, 813, 1193, 4999, 4103, 4486, 7305, 6131, 9086, 7205, 5451, 2314, 1287, 528, 8102, 1446, 3985, 4724, 5306, 1355, 5163, 9074, 9709, 4043, 7285, 5250, 2617, 4756, 1818, 2105, 6790, 6627, 2918, 7984, 7978, 7021, 2470, 1636, 3152, 7908, 8841, 4955, 222, 6480, 5484, 4676, 7926, 5821, 9401, 3232, 7176, 916, 8658, 3237, 1311, 5943, 8487, 3928, 7051, 306, 6033, 3842, 3285, 8951, 1826, 7616, 2324, 648, 9252, 5476, 8556, 4445, 6784},
// 			want:   13,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			assert.Equal(t, tt.want, MaxProfitWithTwiceWithoutCache(tt.prices))
// 		})
// 	}
// }
