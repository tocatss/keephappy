package main

import (
	"bytes"
	"container/heap"
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
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

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// 链接：https://leetcode-cn.com/problems/add-two-numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	next := AddTwoNumbers(l1.Next, l2.Next)

	if sum >= 10 {
		return &ListNode{
			Val:  sum - 10,
			Next: AddTwoNumbers(next, &ListNode{1, nil}),
		}
	}
	return &ListNode{
		Val:  sum,
		Next: AddTwoNumbers(next, &ListNode{1, nil}),
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

func binarySearch(s []int, target int) bool {
	if len(s) == 0 {
		return false
	}

	start := 0
	end := len(s) - 1

	for mid := (start + end) / 2; start <= end; mid = (start + end) / 2 {
		if target == s[mid] {
			return true
		} else if target > s[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

// 盛最多水的容器
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
func maxArea(height []int) int {
	if len(height) < 2 {
		return -1
	}

	left := 0
	right := len(height) - 1
	max := 0
	for left != right {
		if height[left] < height[right] {
			area := (right - left) * height[left]
			if area > max {
				max = area
			}
			left++
			continue
		}

		area := (right - left) * height[right]
		if area > max {
			max = area
		}
		right--
	}
	return max
}

// https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	result := 0
	for left := 0; left < len(height); {
		if height[left] == 0 || left+1 >= len(height) {
			left++
			continue
		}

		// so,find first heigher than left and return value,index.
		right := -1
		for i, v := range height[left+1:] {
			if v >= height[left] {
				right = left + 1 + i
				break
			}
		}
		// 515 or 516
		if right >= 0 {
			blockArea := height[left]
			for i := left; i < right; i++ {
				blockArea += 1 * height[i]
			}
			result += (right-left+1)*height[left] - blockArea
			left = right
			continue
		}

		// maybe 532 or 523, so find max of [2,3].
		maxValue := 0
		for i, v := range height[left+1:] {
			if v >= maxValue {
				maxValue = v
				right = left + 1 + i
			}
		}

		blockArea := maxValue
		for i := left + 1; i <= right; i++ {
			blockArea += 1 * height[i]
		}
		result += (right-left+1)*maxValue - blockArea
		left = right
	}

	return result
}

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j]
	})

	var res [][]int
	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && v == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			rv := nums[right]
			lv := nums[left]
			sum := v + rv + lv
			if sum == 0 {
				res = append(res, []int{v, lv, rv})
				for right = right - 1; right > left; right-- {
					if rv != nums[right] {
						break
					}
				}
				for left = left + 1; left < right; left++ {
					if lv != nums[left] {
						break
					}
				}
			} else if sum > 0 {
				for right = right - 1; right > left; right-- {
					if rv != nums[right] {
						break
					}
				}
			} else {
				for left = left + 1; left < right; left++ {
					if lv != nums[left] {
						break
					}
				}
			}
		}
	}

	return res
}

func LetterCombinations(digits string) []string {
	var res []string
	m := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	for i := 0; i < len(digits); i++ {
		vs, ok := m[digits[i]]
		if !ok {
			continue
		}
		if len(res) == 0 {
			res = vs
			continue
		}
		t := make([]string, 0, 4*len(res))
		for j := 0; j < len(res); j++ {
			for _, s := range vs {
				t = append(t, res[j]+s)
			}
		}
		res = t

	}
	return res
}

func IntToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romas := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var sb strings.Builder
	for i, n := range nums {
		for v := num / n; v > 0; v-- {
			sb.WriteString(romas[i])
		}
		num = num % n
	}
	return sb.String()
}

func RomanToInt(s string) int {
	m := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	var res int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i+1]] > m[s[i]] {
			res += m[s[i+1]] - m[s[i]]
			i++
			continue
		}
		res += m[s[i]]
	}
	return res
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	res := strs[0]
	for _, v := range strs {
		var sb strings.Builder
		for i := 0; i < len(res) && i < len(v); i++ {
			if res[i] != v[i] {
				break
			}
			sb.WriteByte(res[i])
		}
		res = sb.String()
	}
	return res
}

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := nums[0] + nums[1] + nums[2]

	for start := 0; start < len(nums); start++ {
		left := start + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[start] + nums[left] + nums[right]
			if sum > target {
				res = findNearlyTarget(target, res, sum)
				right--
			} else if sum < target {
				res = findNearlyTarget(target, res, sum)
				left++
			} else {
				return target
			}
		}
	}
	return res
}

func findNearlyTarget(target, a, b int) int {
	aa := target - a
	bb := target - b
	if aa < 0 {
		aa = 0 - aa
	}
	if bb < 0 {
		bb = 0 - bb
	}
	if aa <= bb {
		return a
	}
	return b
}

func FourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	for i := 0; i < len(nums); {
		for j := i + 1; j < len(nums); {
			left := j + 1
			right := len(nums) - 1
			expect := target - nums[i] - nums[j]
			for left < right {
				if expect > nums[left]+nums[right] {
					left = findNextLeft(nums, left, right)
				} else if expect < nums[left]+nums[right] {
					right = findNextRight(nums, left, right)
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left = findNextLeft(nums, left, right)
					right = findNextRight(nums, left, right)
				}
			}
			j = findNextLeft(nums, j, len(nums))
		}
		i = findNextLeft(nums, i, len(nums))

	}
	return res
}

func findNextRight(nums []int, left, right int) int {
	right--
	for left < right {
		if nums[right] != nums[right+1] {
			break
		}
		right--
	}
	return right
}

func findNextLeft(nums []int, left, right int) int {
	left++
	for left < right {
		if nums[left] != nums[left-1] {
			break
		}
		left++
	}
	return left
}

// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 遍历一次，双指针法.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	p := dummyNode
	for i, node := 0, head; node != nil; i, node = i+1, node.Next {
		if i >= n {
			p = p.Next
		}
	}

	willDeleteNode := p.Next
	p.Next = willDeleteNode.Next

	return dummyNode.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummy   *ListNode = &ListNode{}
		crtNode *ListNode = dummy
	)

	for l1 != nil || l2 != nil {
		if l1 == nil {
			crtNode.Next = l2
			l2 = nil
			break
		}
		if l2 == nil {
			crtNode.Next = l1
			l1 = nil
			break
		}

		if l1.Val < l2.Val {
			crtNode.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
			crtNode = crtNode.Next
			continue
		}
		crtNode.Next = &ListNode{Val: l2.Val}
		l2 = l2.Next
		crtNode = crtNode.Next
	}
	return dummy.Next
}

type nsHead []*ListNode

func (h nsHead) Len() int {
	return len(h)
}

func (h nsHead) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h nsHead) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *nsHead) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *nsHead) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	// if len(lists) == 0 {
	// 	return nil
	// }
	// var mergeK func(ns []*ListNode, start, end int) *ListNode
	// mergeK = func(ns []*ListNode, start, end int) *ListNode {
	// 	if start+1 >= end {
	// 		return ns[start]
	// 	}
	// 	if start+2 >= end {
	// 		return mergeTwoLists(ns[start], ns[start+1])
	// 	}
	// 	mid := (end - start) / 2
	// 	return mergeTwoLists(mergeK(ns, start, mid), mergeK(ns, mid, end))
	// }
	// return mergeK(lists, 0, len(lists))

	// ANOTHER..
	// if len(lists) == 0 {
	// 	return nil
	// }
	// res := lists[0]
	// for i := 1; i < len(lists); i++ {
	// 	res = mergeTwoLists(res, lists[i])
	// 	printLink(res)
	// }
	// return res
	ns := make(nsHead, 0, len(lists))
	for _, n := range lists {
		if n != nil {
			ns = append(ns, n)
		}
	}

	h := &ns
	heap.Init(h)
	dummy := &ListNode{}
	crt := dummy
	for h.Len() > 0 {
		item := heap.Pop(h)
		node := item.(*ListNode)
		crt.Next = node
		crt = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

func printLink(l *ListNode) {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	log.Print(res)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	p := dummy

	revertK := func() {
		newTail := p.Next
		for i := 0; i < k-1; i++ {
			left, right := newTail, newTail.Next

			left.Next = right.Next
			right.Next = p.Next
			p.Next = right
		}
		p = newTail
	}
	hasK := func() bool {
		n := p.Next
		for i := 0; i < k; i++ {
			if n == nil {
				return false
			}
			n = n.Next
		}
		return true
	}

	for {
		if !hasK() {
			break
		}
		revertK()
	}

	return dummy.Next
}

func RemoveDuplicates(nums []int) int {
	// var (
	// 	ans       int
	// 	move2Tail func(index int)
	// )
	// move2Tail = func(index int) {
	// 	v := nums[index]
	// 	for j := index + 1; j < len(nums); j++ {
	// 		nums[j-1] = nums[j]
	// 	}
	// 	nums[len(nums)-1] = v
	// }

	// for i, j := 0, 0; i < len(nums) && j < len(nums); {
	// 	if i > 0 && nums[i] == nums[i-1] {
	// 		move2Tail(i)
	// 		j++
	// 		continue
	// 	}
	// 	if i > 0 && nums[i] < nums[i-1] {
	// 		break
	// 	}
	// 	ans++
	// 	i++
	// }
	// return ans
	if len(nums) <= 1 {
		return len(nums)
	}

	p, q := 0, 1
	for q < len(nums) {
		if nums[p] == nums[q] {
			q++
			continue
		}
		nums[p+1] = nums[q]
		p++
		q++
	}

	return p + 1
}

func RemoveElement(nums []int, val int) int {
	var p, q int
	for q < len(nums) {
		if nums[q] == val {
			q++
			continue
		}
		nums[p] = nums[q]
		p++
		q++
	}
	return p
}

func MoveZeroes(nums []int) {
	var q, p int
	for q < len(nums) {
		if nums[q] == 0 {
			q++
			continue
		}
		nums[p] = nums[q]
		if q > p {
			nums[q] = 0
		}
		p++
		q++
	}
}

func RemoveDuplicates2(nums []int) int {
	p, q, count := 1, 1, 0
	for p < len(nums) {
		if nums[p] == nums[p-1] {
			count++
			if count > 1 {
				p++
				continue
			}
			nums[q] = nums[p]
			p++
			q++
			continue
		}
		count = 0
		nums[q] = nums[p]
		p++
		q++
	}
	return q
}

func Divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	var (
		divide func(dividend, divisor, ans int) int
		mark   bool
	)
	if dividend < 0 && divisor < 0 {
		dividend = 0 - dividend
		divisor = 0 - divisor
	} else if dividend < 0 {
		dividend = 0 - dividend
		mark = true
	} else if divisor < 0 {
		divisor = 0 - divisor
		mark = true
	}

	divide = func(newDividend, newDivisor, ans int) int {
		if newDividend < newDivisor {
			return 0
		} else if newDividend == newDivisor {
			return 1
		} else if newDividend > newDivisor+newDivisor {
			if ans == 0 {
				ans = 1
			}

			ans += ans
			return divide(newDividend, newDivisor+newDivisor, ans)
		} else if newDividend == newDivisor+newDivisor {
			if ans == 0 {
				ans = 1
			}

			ans += ans
			return ans
		} else { // newDividend > newDivisor
			if newDivisor == divisor {
				return 1
			}
			return ans + divide(newDividend-newDivisor, divisor, 0)
		}
	}

	ans := divide(dividend, divisor, 0)
	if mark {
		ans = 0 - ans
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	if ans < math.MinInt32 {
		return math.MinInt32
	}

	return ans
}

func FindSubstring(s string, words []string) []int {
	// Union words waste too much time.
	// if s == "" {
	// 	return nil
	// }
	// if len(words) == 0 {
	// 	return nil
	// }
	// // Union all word.
	// var (
	// 	path     []string
	// 	visited  []bool = make([]bool, len(words))
	// 	allWords [][]string
	// 	dfs      func(visited []bool, path []string)
	// )

	// dfs = func(visited []bool, path []string) {
	// 	if len(path) == len(words) {
	// 		copied := make([]string, len(words))
	// 		copy(copied, path)
	// 		allWords = append(allWords, copied)
	// 		return
	// 	}

	// 	for i, v := range words {
	// 		if visited[i] {
	// 			continue
	// 		}

	// 		path = append(path, v)
	// 		visited[i] = true

	// 		dfs(visited, path)

	// 		visited[i] = false
	// 		path = path[:len(path)-1]

	// 	}
	// }

	// dfs(visited, path)

	// ansMap := make(map[int]interface{})
	// for _, v := range allWords {
	// 	substr := strings.Join(v, "")
	// 	if len(substr) == 0 {
	// 		ansMap[0] = nil
	// 		continue
	// 	}

	// 	lastIndex := 0
	// 	i := strings.Index(s, substr)
	// 	for i != -1 {
	// 		index := lastIndex + i
	// 		ansMap[index] = nil

	// 		if index+1 >= len(s) {
	// 			break
	// 		}

	// 		lastIndex = index + 1
	// 		i = strings.Index(s[index+1:], substr)
	// 	}

	// }

	// ans := make([]int, 0, len(ansMap))
	// for k := range ansMap {
	// 	ans = append(ans, k)
	// }

	// return ans
	if s == "" || len(words) == 0 {
		return nil
	}
	if len(words[0]) == 0 {
		return []int{0}
	}
	var (
		wl  = len(words[0])
		wsl = wl * len(words)
		wm  = make(map[string]int)
		cwm = make(map[string]int)

		isWordEqual = func(ss string, cwm, wm map[string]int) bool {
			for i := 0; i < len(ss); i += wl {
				k := ss[i : i+wl]
				v, ok := cwm[k]
				if !ok || v == wm[k] {
					return false
				}
				cwm[k]++
			}
			return true
		}
		resetCopyWordMap = func(cwm map[string]int) {
			for k := range cwm {
				cwm[k] = 0
			}
		}
	)

	for _, v := range words {
		if _, ok := wm[v]; !ok {
			wm[v] = 1
			continue
		}
		wm[v]++
	}
	for k := range wm {
		cwm[k] = 0
	}

	ans := make([]int, 0)
	for start := 0; start+wsl <= len(s); start++ {
		ss := s[start : start+wsl]
		if ok := isWordEqual(ss, cwm, wm); ok {
			ans = append(ans, start)
			resetCopyWordMap(cwm)
			continue
		}
		resetCopyWordMap(cwm)
	}
	return ans
}

func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := len(nums) - 1; i > 0; i-- {
		// Find nextPermutation => i-1.
		if nums[i] > nums[i-1] {
			// Find the smallest one which bigger than nums[i-1] in nums[i:]
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					// Swap i-i and smallest one.
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}
			// Reverse nums[i:].
			sublen := len(nums) - 1 - i
			base := i
			for j := 0; ; j++ {
				source := j + base
				target := sublen - j + base
				nums[source], nums[target] = nums[target], nums[source]
				if source == target || source+1 == target {
					return
				}
			}
		}
	}

	// Not found nextPermutation,reverse nums[:]
	for i := 0; ; i++ {
		j := len(nums) - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
		if j == i+1 || j == i {
			return
		}
	}
}

func SearchTarget(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target > nums[len(nums)-1] && target < nums[0] {
		return -1
	}

	bSearch := func(start, end int) int {
		for start <= end {
			m := (start + end) / 2
			if nums[m] < target {
				start = m + 1
			} else if nums[m] > target {
				end = m - 1
			} else {
				return m
			}
		}
		return -1
	}

	if target <= nums[len(nums)-1] {
		for start, end := 0, len(nums)-1; start <= end; {
			mid := (start + end) / 2
			if nums[mid] > nums[end] {
				start = mid + 1
				continue
			}
			if nums[mid] > target {
				end = mid - 1
				continue
			}

			return bSearch(mid, end)
		}
		return -1
	}

	for start, end := 0, len(nums)-1; start <= end; {
		mid := (start + end) / 2
		if nums[mid] < nums[0] {
			end = mid - 1
			continue
		}
		if nums[mid] < target {
			start = mid + 1
			continue
		}

		return bSearch(start, mid)
	}
	return -1
}

func LongestValidParentheses(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		match, path, valued := 0, 0, 0
		for j := i; j < len(s); j++ {
			switch s[j] {
			case '(':
				path++
			case ')':
				path--
			}
			if path >= 0 {
				match++
				if path == 0 {
					valued = match
				}
				continue
			}
			break
		}
		if valued > ans {
			ans = valued
		}
	}
	return ans
}

func SearchRange(nums []int, target int) []int {
	// notFound := []int{-1, -1}
	// if len(nums) == 0 {
	// 	return notFound
	// }

	// for start, end := 0, len(nums)-1; start <= end; {
	// 	switch mid := (start + end) / 2; true {
	// 	case nums[mid] < target:
	// 		start = mid + 1
	// 	case nums[mid] > target:
	// 		end = mid - 1
	// 	case nums[mid] == target:
	// 		var (
	// 			ans     = make([]int, 2)
	// 			hasNext = func(index, target int) bool {
	// 				return index <= len(nums)-1 && nums[index] == target
	// 			}
	// 			hasLast = func(index, target int) bool {
	// 				return index >= 0 && nums[index] == target
	// 			}
	// 		)
	// 		for i, j := mid, mid; hasLast(i, target) || hasNext(j, target); {
	// 			if hasLast(i, target) {
	// 				ans[0] = i
	// 				i--
	// 			}
	// 			if hasNext(j, target) {
	// 				ans[1] = j
	// 				j++
	// 			}
	// 		}
	// 		return ans
	// 	}
	// }
	var (
		leftSearch = func(nums []int, target int) int {
			left, right := 0, len(nums) // 左闭右开区间
			for left < right {
				switch mid := left + (right-left)/2; {
				case nums[mid] == target:
					right = mid // 相等,接着找左边界.
				case nums[mid] < target:
					left = mid + 1
				case nums[mid] > target:
					right = mid // 右开区间,right = mid
				}
			}

			if left == len(nums) || nums[left] != target {
				return -1
			}
			return left
		}
		rightSearch = func(nums []int, target int) int {
			left, right := 0, len(nums)
			for left < right {
				switch mid := left + (right-left)/2; {
				case nums[mid] == target:
					left = mid + 1 // 找右边界.
				case nums[mid] < target:
					left = mid + 1
				case nums[mid] > target:
					right = mid
				}
			}

			if left == 0 || nums[left-1] != target {
				return -1
			}
			return left - 1
		}
	)
	return []int{
		leftSearch(nums, target),
		rightSearch(nums, target),
	}
}

// 即计算 x 的 n 次幂函数。
// 2分法
func MyPow(x float64, n int) float64 {
	// 1.将n进行二进制表示：
	// n=9 => 1001 => 8*1+4*0+2*0+1*1
	// x^n => x^8*x^0*x^0*x^1
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}

	calc := func(x float64, n int) float64 {
		var (
			ans     float64 = 1
			iterate float64 = x
		)
		for i := n; i > 0; {
			if i&1 == 1 {
				ans *= iterate
			}
			iterate *= iterate
			i = i >> 1
		}
		return ans
	}

	if n < 0 {
		return 1 / calc(x, -n)
	}
	return calc(x, n)
}

func MyPowRecursion(x float64, n int) float64 {
	// x^8 = x^4*x^4 = x^2*x^2*x^2*x^2 ...
	if n == 0 {
		return 1
	}

	var recursion func(x float64, n int) float64
	recursion = func(x float64, n int) float64 {
		if n == 1 {
			return x
		}
		y := recursion(x, n/2)
		if n%2 == 0 {
			return y * y
		}
		return y * y * x
	}
	if n > 0 {
		return recursion(x, n)
	}
	return 1.0 / recursion(x, -n)
}

// 添加辅助栈用于记录最小值,空间换时间.
type MinStack struct {
	nums    []int
	minNums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		nums:    []int{},
		minNums: []int{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, x)
		this.minNums = append(this.minNums, x)
		return
	}

	min := this.minNums[len(this.minNums)-1]
	if x < min {
		this.nums = append(this.nums, x)
		this.minNums = append(this.minNums, x)
		return
	}
	this.nums = append(this.nums, x)
	this.minNums = append(this.minNums, min)
}

func (this *MinStack) Pop() {
	if len(this.nums) == 0 {
		return
	}

	this.nums = this.nums[0 : len(this.nums)-1]
	this.minNums = this.minNums[0 : len(this.minNums)-1]
}

func (this *MinStack) Top() int {
	if len(this.nums) == 0 {
		return -1
	}
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.nums) == 0 {
		return -1
	}
	return this.minNums[len(this.minNums)-1]
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}

	parentNodes := []*TreeNode{root}
	for len(parentNodes) > 0 {
		childNodes := make([]*TreeNode, 0, len(parentNodes)*2)
		vals := make([]int, 0, len(parentNodes))
		for _, v := range parentNodes {
			vals = append(vals, v.Val)
			if v.Left != nil && v.Right != nil {
				childNodes = append(childNodes, v.Left, v.Right)
			} else if v.Left != nil {
				childNodes = append(childNodes, v.Left)
			} else if v.Right != nil {
				childNodes = append(childNodes, v.Right)
			}
		}
		ans = append(ans, vals)
		parentNodes = childNodes
	}
	return ans
}

func LevelOrderDFS(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level > len(ans)-1 {
			ans = append(ans, []int{node.Val})
		} else {
			ans[level] = append(ans[level], node.Val)
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)

	return ans
}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// 0和任何数做异或运算都是其本身； 0 XOR a = a
	// 任何数和其自身做异或都是 0；a XOR a = 0
	// XOR 满足交换律和结合律
	if len(nums) == 0 {
		return -1
	}

	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

// 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
// 示例 1 :
// 输入:nums = [1,1,1], k = 2
// 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
func subarraySum(nums []int, k int) int {
	ans := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

func subarraySumWithPreSum(nums []int, k int) int {
	preSum := make(map[int]int, 0)
	preSum[0] = 1
	ans := 0

	// sum(i,j) = sum(0,j)-sum(0,i-1) = k
	// sum(0,j) -k = sum(0,i-1)
	sum := 0
	for _, v := range nums {
		sum += v
		want := sum - k
		ans += preSum[want]
		preSum[sum]++
	}
	return ans
}

func ReverseKGroupAgain(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	var (
		reverse2Group func(last, left, right *ListNode)
		isLenghEnough func(last *ListNode, k int) bool
	)
	reverse2Group = func(last, left, right *ListNode) {
		left.Next = right.Next
		right.Next = last.Next
		last.Next = right
	}
	isLenghEnough = func(last *ListNode, k int) bool {
		for k > 0 {
			if last.Next != nil {
				last = last.Next
				k--
				continue
			}
			return false
		}
		return true
	}

	dummyNode := &ListNode{Next: head}
	last := dummyNode
	for isLenghEnough(last, k) {
		left := last.Next
		for i := k; i > 1; i-- {
			right := left.Next
			reverse2Group(last, left, right)
		}
		last = left
	}
	return dummyNode.Next
}

// 现在你总共有 n 门课需要选，记为 0 到 n-1。
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
// 给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
// 可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

// 排序： 把有向无环图 变成 线性的排序。
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	graph := make(map[int][]int)
	for _, pre := range prerequisites {
		k, v := pre[1], pre[0]
		if _, ok := graph[k]; !ok {
			graph[k] = make([]int, 0)
		}
		graph[k] = append(graph[k], v)

		inDegree[v]++
	}

	ans := make([]int, 0)

	findZero := func(nums []int) []int {
		ans := make([]int, 0, len(nums))
		for i, v := range nums {
			if v == 0 {
				ans = append(ans, i)
			}
		}
		return ans
	}
	// 每次找对性能不好呢～
	// for zerolist := findZero(inDegree); len(zerolist) > 0; zerolist = findZero(inDegree) {
	// 	for _, from := range zerolist {
	// 		inDegree[from]--
	// 		for _, to := range graph[from] {
	// 			inDegree[to]--
	// 		}
	// 		ans = append(ans, from)
	// 	}
	// }

	zerolist := findZero(inDegree)
	for i := 0; i < len(zerolist); i++ {
		from := zerolist[i]
		ans = append(ans, from)
		for _, to := range graph[from] {
			inDegree[to]--
			if inDegree[to] == 0 {
				zerolist = append(zerolist, to)
			}
		}
	}

	if len(ans) == numCourses {
		return ans
	}
	return nil
}

// 后缀表达式求值
// 遇到数字进栈，操作符弹出栈顶两个元素计算后再入栈。
func ComputeSuffix(suffix []string) int {
	stack := make([]int, 0, len(suffix))

	for i := 0; i < len(suffix); i++ {
		s := strings.TrimSpace(suffix[i])
		v, err := strconv.Atoi(s)
		if err == nil {
			stack = append(stack, v)
			continue
		}

		ss := make([]int, len(stack)-2, len(stack)-1)
		copy(ss, stack[:len(stack)-2])

		if s == "+" {
			ss = append(ss, stack[len(stack)-2]+stack[len(stack)-1])
		} else if s == "-" {
			ss = append(ss, stack[len(stack)-2]-stack[len(stack)-1])
		} else if s == "*" {
			ss = append(ss, stack[len(stack)-2]*stack[len(stack)-1])
		} else if s == "/" {
			// FIXME: 假设可以整除
			ss = append(ss, stack[len(stack)-2]/stack[len(stack)-1])
		} else {
			panic("incorrect value")
		}

		stack = ss

	}

	return stack[0]
}

// 中缀转后缀
// 准备两个栈(s1,s2),一个用于存数字一个用来存操作符
func Convert2Suffix(infix []string) []string {
	s1, s2 := make([]string, 0, len(infix)), make([]string, 0, len(infix))

	for _, v := range infix {
		// number?
		if _, err := strconv.Atoi(v); err == nil {
			s1 = append(s1, v)
			continue
		}

		if len(s2) == 0 || s2[len(s2)-1] == "(" {
			s2 = append(s2, v)
			continue
		}

		switch v {
		case "(":
			s2 = append(s2, v)
		case "+", "-":
			i := len(s2) - 1
			for i >= 0 && s2[i] != "(" {
				s1 = append(s1, s2[i])
				i--
			}
			s2 = s2[:i+1]
			s2 = append(s2, v)
		case "*", "/":
			if s2[len(s2)-1] == "+" || s2[len(s2)-1] == "-" {
				s2 = append(s2, v)
				continue
			}
			s1 = append(s1, s2[len(s2)-1])
			s2[len(s2)-1] = v
		case ")":
			i := len(s2) - 1
			for {
				if i < 0 {
					panic("unexpect error")
				}
				if s2[i] == "(" {
					s2 = s2[:i]
					break
				}

				s1 = append(s1, s2[i])
				i--
			}
		}
	}

	for i := len(s2) - 1; i >= 0; i-- {
		s1 = append(s1, s2[i])
	}

	return s1
}

// IP地址一般是一个32位的二进制数
// 如果将IP地址转换成二进制表示应该有32为那么长但是它通常被分割为4个“8位二进制数”（也就是4个字节每，每个代表的就是小于2的8 次方）。
// IP地址通常用“点分十进制”表示成（a.b.c.d）的形式，其中，a,b,c,d都是0~255之间的十进制整数。
// 例：点分十进IP地址（100.4.5.6），实际上是32位二进制数（01100100.00000100.00000101.00000110）
func Ip2Number(s string) int {
	ss := strings.Split(s, ".")
	if len(ss) != 4 {
		return -1
	}
	nums := make([]int, 4)
	for i := 0; i < 4; i++ {
		v, err := strconv.Atoi(ss[i])
		if err != nil {
			return -1
		}
		nums[i] = v
	}

	// nums[0] * 2^24 + nums[1] * 2^16 + nums[2] * 2^8 + nums[1]
	// return nums[0]*256*256*256 + nums[1]*256*256 + nums[2]*256 + nums[3]
	return nums[0]<<24 + nums[1]<<16 + nums[2]<<8 + nums[3]
}

func Number2Ip(num int) string {
	if num < 0 || num > math.MaxUint32 {
		return ""
	}

	ans := make([]string, 4)
	for i := 0; i < 4; i++ {
		n := num
		switch i {
		case 0:
			ans[i] = strconv.Itoa(n >> 24)
		case 1:
			ans[i] = strconv.Itoa(n >> 16 & 255)
		case 2:
			ans[i] = strconv.Itoa(n >> 8 & 255)
		case 3:
			ans[i] = strconv.Itoa(n & 255)
		}
	}

	return strings.Join(ans, ".")
}

// 汉诺塔
// 一股脑地考虑每一步如何移动很困难，我们可以换个思路。先假设一共10个盘子除最下面的盘子之外，我们已经成功地将上面的9个盘子移到了b柱，此时只要将最下面的盘子由a移动到c即可。
// 接下来将b柱作为起始，将8个成功移到a把第九个移动到c即可。
// 规律：即每次都是先将其他圆盘移动到辅助柱子上，并将最底下的圆盘移到c柱子上，然后再把原先的柱子作为辅助柱子，并重复此过程。
func Hanoi(n int) int {
	var (
		res    int
		move   func(index int, from, to string)
		hannoi func(nums int, from, by, to string)
	)

	move = func(index int, from, to string) {
		res++
		fmt.Printf("MOVE %d FROM %s TO %s\n", index, from, to)
	}
	hannoi = func(nums int, from, by, to string) {
		if nums == 1 {
			move(1, from, to)
			return
		}

		hannoi(nums-1, from, to, by) //将前n-1块经由c挪到b
		move(nums, from, to)         // 将第n块挪到c
		hannoi(nums-1, by, from, to) // 将剩下的n-1块经由a挪到c
	}

	hannoi(n, "A", "B", "C")

	fmt.Println("OK")
	return res
}

// 最大正向匹配分词算法
// 输入： 今天的天气真好我们出去玩耍吧
// 字典：今天， 天气真好，天气，真好， 我们，出去玩耍，出去，玩耍
// 输出： 今天 的 天气真好 我们 出去玩耍 吧
// TODO: Trie树 ?
func MaximumMatching(input string, m map[string]interface{}) []string {

	if len(m) == 0 {
		return []string{input}
	}
	if len(input) <= 1 {
		return []string{input}
	}

	rs := []rune(input)
	ans := make([]string, 0, len(rs))
	for i := 0; i < len(rs); {
		isFound := false
		for j := len(rs); j > i; j-- {
			k := string(rs[i:j])
			if _, ok := m[k]; ok {
				ans = append(ans, k)
				i = j
				isFound = true
				break
			}
		}

		if isFound {
			continue
		}

		ans = append(ans, string(rs[i:i+1]))
		i++
	}

	return ans
}
