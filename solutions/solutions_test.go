package solutions

import (
	"fmt"
	"testing"
)

//1
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	r := twoSum(nums, target)
	if r[0] != 0 || r[1] != 1 {
		t.Errorf("twoSum({2, 7, 11, 15}, 9) failed. Got %d, %d, expected 0,1.", r[0], r[1])
	}
}

//2
func TestAddTwoNumbers(t *testing.T) {
	var toString = func(l *ListNode) string {
		var s string = "["
		for l != nil {
			s += fmt.Sprintf("%d,", l.Val)
			l = l.Next
		}

		s += "]"
		return s
	}

	var a1 *ListNode = new(ListNode)
	var a2 *ListNode = new(ListNode)
	var a3 *ListNode = new(ListNode)

	var b1 *ListNode = new(ListNode)
	var b2 *ListNode = new(ListNode)
	var b3 *ListNode = new(ListNode)

	a1.Val = 5
	b1.Val = 5
	r := addTwoNumbers(a1, b1)
	if r.Val != 0 || r.Next.Val != 1 {
		t.Errorf("addTwoNumber(5, 5) failed. Got %v, Expected [0,1]", toString(r))
	}

	a1.Val = 2
	a2.Val = 4
	a3.Val = 3
	a1.Next = a2
	a2.Next = a3

	b1.Val = 5
	b2.Val = 6
	b3.Val = 4
	b1.Next = b2
	b2.Next = b3

	r = addTwoNumbers(a1, b1)
	if r.Val != 7 || r.Next.Val != 0 || r.Next.Next.Val != 8 {
		t.Errorf("addTwoNumber(2-->4-->3, 5-->6-->4) failed. Got %v, Expected [7,0,8]", toString(r))
	}
}

//3
func TestLengthOfLongestSubstring(t *testing.T) {
	var s string
	var r int
	s = "abcabcbb"
	r = lengthOfLongestSubstring(s)
	if r != 3 {
		t.Errorf("lengthOfLongestSubstring(ahcahcbb) failed. Got %d, Expected 3.", r)
	}

	s = "bbbbb"
	r = lengthOfLongestSubstring(s)
	if r != 1 {
		t.Errorf("lengthOfLongestSubstring(ahcahcbb) failed. Got %d, Expected 1.", r)
	}

	s = "aab"
	r = lengthOfLongestSubstring(s)
	if r != 2 {
		t.Errorf("lengthOfLongestSubstring(ahcahcbb) failed. Got %d, Expected 2.", r)
	}
}

//4
func TestFindMedianSortedArrays(t *testing.T) {
	var nums1, nums2 []int
	var r float64
	nums1 = []int{1, 3}
	nums2 = []int{2}
	r = findMedianSortedArrays(nums1, nums2)
	if r != 2.0 {
		t.Errorf("findMedianSortedArrays({1,3}, {2}) failed. Got %f, expected 2.0.", r)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	r = findMedianSortedArrays(nums1, nums2)
	if r != 2.5 {
		t.Errorf("findMedianSortedArrays({1, 2}, {3, 4}) failed. Got %f, expected 2.5.", r)
	}

	nums1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	nums2 = []int{0, 6}
	r = findMedianSortedArrays(nums1, nums2)
	if r != 10.5 {
		t.Errorf("findMedianSortedArrays({1, 2}, {3, 4}) failed. Got %f, expected 10.5.", r)
	}
}

//371
func TestGetSum(t *testing.T) {
	r := getSum(3, 5)
	if r != 8 {
		t.Errorf("getSum(3, 5) failed. Got %d, expected 8.", r)
	}
}
