package sliceutils

import "strconv"

func Atoi64(strings []string) ([]int64, error) {
	nums := make([]int64, len(strings))

	for i, s := range strings {
		num, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}

	return nums, nil
}

func Atoui64(strings []string) ([]uint64, error) {
	nums := make([]uint64, len(strings))

	for i, s := range strings {
		num, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}

	return nums, nil
}

func Atoi(strings []string) ([]int, error) {
	nums := make([]int, len(strings))

	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}

	return nums, nil
}
