package main

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, 0, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	return NumArray{
		sum: sum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
