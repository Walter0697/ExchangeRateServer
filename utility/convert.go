package utility

import "strconv"

func StringToUint(input string) (uint, error) {
	u64, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(u64), nil
}
