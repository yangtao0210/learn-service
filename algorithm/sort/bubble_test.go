package sort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	nums1 := []int{1, 0, 3, 2, 4, 7, 9, 2, 4, 5}
	nums2 := []int{0, 1, 2, 3, 4}
	fmt.Println(nums1, nums2)
	BubbleSort(nums1)
	fmt.Println(nums1)
}
