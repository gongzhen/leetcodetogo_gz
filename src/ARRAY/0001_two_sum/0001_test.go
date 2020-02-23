package _001_two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{7, 2, 11, 15}
	funcs := []func([]int, int) []int{twoSum}

	for _, testFunc := range funcs {
		if res := testFunc(nums, 9); !reflect.DeepEqual(res, []int{1, 0}) {
			t.Error("Failed, two sum")
		}
		if res := testFunc(nums, 6); !reflect.DeepEqual(res, []int{}) {
			t.Error("Failed, two sum")
		}
	}
}

