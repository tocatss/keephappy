package main

import (
	"bytes"
	"errors"
	"math"
	"strconv"
	"strings"
)

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

func TwoSum(nums []int, target int) []int {
	// result := make([]int,0,len(nums))
	// for i:= 0; i<len(nums); i++ {
	// 	for j:=i+1; j<len(nums); j++ {
	// 		if nums[i] + nums[j] == target {
	// 			result = append(result,i,j)
	// 			return result
	// 		}
	// 	}
	// }
	// return result
	m := make(map[int]int)
	for i, v := range nums {
		expect := target - v
		if _, ok := m[expect]; ok {
			return []int{m[expect], i}
		}
		m[v] = i
	}
	return nil
}

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// // Validate.
	// if l1 == nil || l2 == nil {
	// 	return nil
	// }

	// dummyNode := &ListNode{}
	// headNode := dummyNode
	// up := 0 // 1 || 0
	// for l1 != nil || l2 != nil || up != 0 {
	// 	val := 0
	// 	if l1 != nil {
	// 		val += l1.Val
	// 	}
	// 	if l2 != nil {
	// 		val += l2.Val
	// 	}
	// 	val += up

	// 	up = val / 10
	// 	val = val % 10

	// 	headNode.Next = &ListNode{val, nil}
	// 	headNode = headNode.Next
	// 	if l1 != nil {
	// 		l1 = l1.Next
	// 	}
	// 	if l2 != nil {
	// 		l2 = l2.Next
	// 	}
	// }
	// return dummyNode.Next

	// Validate
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	nextNode := AddTwoNumbers(l1.Next, l2.Next)

	if sum >= 10 {
		upNode := &ListNode{1, nil}
		return &ListNode{
			sum % 10,
			AddTwoNumbers(upNode, nextNode),
		}
	}
	return &ListNode{
		sum,
		nextNode,
	}
}

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 输入: "abcabcbb"
// 输出: 3
func LengthOfLongestSubstring(s string) int {
	// max := 0
	// for i := 0; i < len(s); i++ {
	// 	a := map[byte]interface{}{
	// 		s[i]: nil,
	// 	}
	// 	for j := i + 1; j < len(s); j++ {
	// 		if _, ok := a[s[j]]; ok {
	// 			break
	// 		}
	// 		a[s[j]] = nil
	// 	}
	// 	length := len(a)
	// 	if length > max {
	// 		max = length
	// 	}
	// }
	// return max
	var result, left, right int
	sw := s[left:right]

	for ; right < len(s); right++ {
		if index := strings.IndexByte(sw, s[right]); index != -1 {
			left += index + 1
		}
		sw = s[left : right+1]
		if len(sw) > result {
			result = len(sw)
		}
	}
	return result
}

func LongestPalindrome(s string) string {
	if s == "" {
		return s
	}
	if len(s) == 1 {
		return s
	}

	var tsb strings.Builder
	for i := 0; i < len(s); i++ {
		tsb.WriteString("#")
		tsb.WriteByte(s[i])
	}
	tsb.WriteString("#")

	ts := tsb.String()
	result := ""
	for i := 0; i < len(ts); i++ {
		ss := findPalindromeByMark(ts, i)
		if len(ss) > len(result) {
			result = ss
		}
	}
	return strings.ReplaceAll(result, "#", "")
}

func findPalindromeByMark(s string, mark int) string {
	if len(s) <= mark {
		return ""
	}
	if mark == 0 || mark+1 == len(s) {
		return string(s[mark])
	}

	left := s[:mark]
	mid := s[mark]
	right := s[mark+1:]

	var result, resultLeft, resultRight strings.Builder
	for len(left) > 0 && len(right) > 0 {
		l := left[len(left)-1]
		r := right[0]
		if l != r {
			break
		}
		resultLeft.WriteByte(l)
		resultRight.WriteByte(r)
		left = strPop(left)
		right = strShift(right)
	}
	result.WriteString(strReverse(resultLeft.String()))
	result.WriteByte(mid)
	result.WriteString(resultRight.String())

	return result.String()
}

func strPop(s string) string {
	if s == "" {
		return s
	}
	return s[:len(s)-1]
}

func strShift(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return ""
	}
	return s[1:]
}

func strReverse(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

type direct string

var (
	directUp   direct = "up"
	directDown direct = "down"
	directStop direct = "stop"
)

func Zconvert(s string, numRows int) string {
	// Validate
	if numRows <= 0 || s == "" {
		return s
	}

	bs := make([]bytes.Buffer, numRows)
	cur := 0
	d := directDown
	for i := 0; i < len(s); i++ {
		bs[cur].WriteByte(s[i])

		d = nextDirection(d, cur, numRows)
		if d == directStop {
			continue
		} else if d == directDown {
			cur++
		} else {
			cur--
		}
	}

	var result strings.Builder
	for _, b := range bs {
		result.WriteString(b.String())
	}
	return result.String()
}

func nextDirection(d direct, cur, numRows int) direct {
	if d == directDown && cur == numRows-1 {
		if cur == 0 {
			return directStop
		}
		return directUp
	}
	if d == directUp && cur == 0 {
		return directDown
	}
	return d
}

type linkNode struct {
	data  string
	prior *linkNode
	next  *linkNode
}

func newLinkList(n int) *linkNode {
	dummyNode := &linkNode{}
	prior := dummyNode
	for i := 0; i < n; i++ {
		node := &linkNode{data: strconv.Itoa(i)}
		node.prior = prior
		prior.next = node
		prior = node
	}
	prior.next = dummyNode
	dummyNode.prior = prior
	return dummyNode
}

func IntReverse(x int) int {
	// if origin > math.MaxInt32 || origin < math.MinInt32 {
	// 	return 0
	// }

	// var sb strings.Builder
	// s := strconv.Itoa(origin)
	// if s[0] == '-' {
	// 	sb.WriteByte('-')
	// }

	// for i := len(s) - 1; i >= 0; i-- {
	// 	if s[i] == '-' {
	// 		break
	// 	}
	// 	sb.WriteByte(s[i])
	// }

	// result, err := strconv.Atoi(sb.String())
	// if result > math.MaxInt32 || result < math.MinInt32 {
	// 	return 0
	// }
	// if err != nil {
	// 	log.Print(err)
	// 	return 0
	// }
	// return result
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	r := 0
	t := x
	if x < 0 {
		t = 0 - x
	}
	for t > 0 {
		v := t % 10
		t = t / 10

		if 10*r+v > math.MaxInt32 || -10*r-v < math.MinInt32 {
			return 0
		}
		r = r*10 + v
	}

	if x < 0 {
		return 0 - r
	}
	return r
}

type shareStack struct {
	data []string
	top1 int
	top2 int
}

func newShareStack(i int) *shareStack {
	if i <= 0 {
		return nil
	}

	data := make([]string, i)
	return &shareStack{
		data: data,
		top1: -1,
		top2: i,
	}
}

func (s *shareStack) dump() []string {
	return s.data
}

func (s *shareStack) push(v, direct string) error {
	if direct != "left" && direct != "right" {
		return errors.New("direct is not right")
	}
	if s.top1+1 == s.top2 {
		return errors.New("full")
	}

	if direct == "left" {
		s.top1++
		s.data[s.top1] = v
		return nil
	}
	s.top2--
	s.data[s.top2] = v
	return nil
}

func (s *shareStack) pop(direct string) (string, error) {
	if direct != "left" && direct != "right" {
		return "", errors.New("direct is not right")
	}
	if direct == "left" && s.top1 == -1 {
		return "", errors.New("left is null")
	}
	if direct == "left" && s.top2 == len(s.data) {
		return "", errors.New("right is null")
	}

	if direct == "left" {
		v := s.data[s.top1]
		s.data[s.top1] = ""
		s.top1--

		return v, nil
	}

	v := s.data[s.top2]
	s.data[s.top2] = ""
	s.top2++
	return v, nil
}

func FindNthSmallestOf2Slice(s1, s2 []int, k int) int {
	if len(s1) == 0 {
		return s2[k-1]
	}
	if len(s2) == 0 {
		return s1[k-1]
	}
	// let len(s1) is always shortest.
	if len(s1) > len(s2) {
		return FindNthSmallestOf2Slice(s2, s1, k)
	}
	if k == 1 {
		return intMin(s1[0], s2[0])
	}

	a := s1[len(s1)-1]
	if isLengthLarger(s1, k/2) {
		a = s1[k/2-1]
	}
	b := s2[k/2-1]

	if a <= b {
		if isLengthLarger(s1, k/2) {
			s1 = s1[k/2:]
			k -= k / 2
			return FindNthSmallestOf2Slice(s1, s2, k)
		}
		// clear s1
		k -= len(s1)
		return FindNthSmallestOf2Slice(nil, s2, k)
	}
	s2 = s2[k/2:]
	k -= k / 2
	return FindNthSmallestOf2Slice(s1, s2, k)
}

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	if (l1+l2)%2 == 0 {
		mid1 := FindNthSmallestOf2Slice(nums1, nums2, (l1+l2)/2)
		mid2 := FindNthSmallestOf2Slice(nums1, nums2, (l1+l2)/2+1)
		return (float64(mid1) + float64(mid2)) / 2
	}
	mid := FindNthSmallestOf2Slice(nums1, nums2, (l1+l2)/2+1)
	return float64(mid)
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isLengthLarger(s []int, i int) bool {
	return len(s) > i
}

func isIntPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	v := x
	var r int
	for v >= r {
		r = 10*r + v%10
		if v == r || v/10 == r {
			return true
		}
		v = v / 10
	}
	return false
}
