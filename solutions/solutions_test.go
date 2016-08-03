package solutions
import (
    "testing"
)

func TestGetSum(t *testing.T) {
    r := getSum(3, 5)
    if r != 8 {
        t.Errorf("getSum(3, 5) failed. Got %d, expected 8.", r)
    }
}

func TestTwoSum(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    target := 9
    r := twoSum(nums, target)
    if r[0] != 0 || r[1] != 1 {
        t.Errorf("twoSum({2, 7, 11, 15}, 9) failed. Got %d, %d, expected 0,1.", r[0], r[1])
    }
}