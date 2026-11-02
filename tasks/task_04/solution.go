package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {

	if len(nums) < 2 {
		return Stats{}
	}

	var cnt int = len(nums) - 1
	var delta int64 = nums[1] - nums[0]
	var sum, mini, maxi int64 = delta, delta, delta

	for i := 2; i <= cnt; i++ {
		delta = nums[i] - nums[i-1]
		sum += delta
		mini = min(mini, delta)
		maxi = max(maxi, delta)
	}

	return Stats{cnt, sum, mini, maxi}
}
