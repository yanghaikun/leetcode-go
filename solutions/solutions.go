package solutions

import (
	"sort"
	"strings"
)

//1. Two Sum
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution.
//Example:
//Given nums = [2, 7, 11, 15], target = 9,
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
//You may assume that each input would have exactly one solution.
func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	//copy a array of nums
	var nums_copy []int = make([]int, len(nums))
	copy(nums_copy, nums)
	//sort
	sort.Ints(nums_copy)
	//find the exactly numbers
	var start, end int
	start = 0
	end = len(nums_copy) - 1
	for start < end {
		if (nums_copy[start] + nums_copy[end]) < target {
			start++
		} else if (nums_copy[start] + nums_copy[end]) > target {
			end--
		} else {
			break
		}
	}

	//find the original loc of numbers in nums
	var loc int = 0
	for i := 0; i < len(nums); i++ {
		if (nums[i] == nums_copy[start]) || (nums[i] == nums_copy[end]) {
			result[loc] = i
			loc++
			if loc > 1 {
				break
			}
		}

	}

	return result
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//2. Add Two Numbers
//You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var i, j, n int
	var first, now *ListNode
	for {
		if l1 == nil && l2 == nil && n == 0 {
			break
		}

		if l1 != nil {
			i = l1.Val
			l1 = l1.Next
		} else {
			i = 0
		}
		if l2 != nil {
			j = l2.Val
			l2 = l2.Next
		} else {
			j = 0
		}

		temp := new(ListNode)
		temp.Val = (i + j + n) % 10
		n = (i + j + n) / 10

		if first == nil {
			first = temp
			now = temp
		} else {
			now.Next = temp
			now = temp
		}
	}

	return first
}

// 3. Longest Substring Without Repeating Characters
// Given a string, find the length of the longest substring without repeating characters.
// Examples:
// Given "abcabcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func lengthOfLongestSubstring(s string) int {
	var loc, offset, max int
	var sub string = ""
	for loc+offset < len(s) {
		if strings.Contains(sub, s[(loc+offset):(loc+offset+1)]) {
			if len(sub) > max {
				max = len(sub)
			}
			sub = sub[1:len(sub)]
			loc++
			offset--
		} else {
			offset++
			sub = s[loc:(loc + offset)]
		}
	}
	if max < len(sub) {
		max = len(sub)
	}

	return max
}

// 4. Median of Two Sorted Arrays
// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// Example 1:
// nums1 = [1, 3]
// nums2 = [2]
// The median is 2.0
// Example 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// The median is (2 + 3)/2 = 2.5
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2, l int
	var result float64
	var nums []int = make([]int, len(nums1)+len(nums2))

	for l < len(nums1)+len(nums2) {
		if l1 < len(nums1) && l2 < len(nums2) {
			if nums1[l1] < nums2[l2] {
				nums[l] = nums1[l1]
				l1++
			} else {
				nums[l] = nums2[l2]
				l2++
			}
		} else {
			if l1 < len(nums1) {
				nums[l] = nums1[l1]
				l1++
			} else {
				nums[l] = nums2[l2]
				l2++
			}
		}
		l++
	}

	if l%2 == 1 {
		result = float64(nums[l/2])
	} else {
		result = float64(nums[l/2-1]+nums[l/2]) / 2
	}

	return result
}

//371. Sum of Two Integers
//Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
//Example:
//Given a = 1 and b = 2, return 3.
func getSum(a int, b int) int {
	if b == 0 {
		return a
	}

	xor := a ^ b
	and := (a & b) << 1

	return getSum(xor, and)
}
