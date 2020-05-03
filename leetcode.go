package main

import (
	"bytes"
	"container/heap"
	"errors"
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
	var (
		hasLenKAfterHead func(head *ListNode) bool = func(head *ListNode) bool {
			crt := head.Next
			for kk := k; kk > 0; kk-- {
				if crt == nil {
					return false
				}
				crt = crt.Next
			}
			return true
		}
		move2Head func(head *ListNode, times int) *ListNode = func(head *ListNode, times int) *ListNode {
			crt := head.Next
			for ; times > 0; times-- {
				next := crt.Next
				crt.Next = next.Next
				next.Next = head.Next
				head.Next = next
			}
			return crt
		}
	)

	// 1->2->3->4
	// k=2 2->1->4->3
	dummy := &ListNode{Next: head}
	p := dummy
	for hasLenKAfterHead(p) {
		p = move2Head(p, k-1)
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
	notFound := []int{-1, -1}
	if len(nums) == 0 {
		return notFound
	}

	for start, end := 0, len(nums)-1; start <= end; {
		switch mid := (start + end) / 2; true {
		case nums[mid] < target:
			start = mid + 1
		case nums[mid] > target:
			end = mid - 1
		case nums[mid] == target:
			var (
				ans     = make([]int, 2)
				hasNext = func(index, target int) bool {
					return index <= len(nums)-1 && nums[index] == target
				}
				hasLast = func(index, target int) bool {
					return index >= 0 && nums[index] == target
				}
			)
			for i, j := mid, mid; hasLast(i, target) || hasNext(j, target); {
				if hasLast(i, target) {
					ans[0] = i
					i--
				}
				if hasNext(j, target) {
					ans[1] = j
					j++
				}
			}
			return ans
		}
	}
	return notFound
}
