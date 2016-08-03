package solutions
import (
    "sort"
)

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

//1.Two Sum
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution.
//Example:
//Given nums = [2, 7, 11, 15], target = 9,
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].
//You may assume that each input would have exactly one solution.
func twoSum(nums []int, target int) []int {
    result := []int{0,0}
    //copy a array of nums
    var nums_copy []int = make([]int, len(nums))
    copy(nums_copy, nums)
    //sort
    sort.Ints(nums_copy)
    //find the exactly numbers
    var start, end int
    start = 0;
    end = len(nums_copy) - 1
    for (start < end) {
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

