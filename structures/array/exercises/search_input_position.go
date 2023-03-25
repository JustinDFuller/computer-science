package exercises

func SearchInsert(nums []int, target int) int {
	high := len(nums) - 1
	low := 0

	for high >= low {
		middle := (low + high) / 2

		if nums[middle] == target {
			return middle
		}

		if middle+1 <= len(nums)-1 && nums[middle] < target && nums[middle+1] > target {
			return middle + 1
		}

		if middle-1 >= 0 && nums[middle] > target && nums[middle-1] < target {
			return middle
		}

		if middle+1 > len(nums)-1 && target > nums[middle] {
			return middle + 1
		}

		if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return 0
}
