package gotcha

func toHex(code Chars) int {
	if code < Zero {
		return -1
	} else if code <= Nine {
		return int(code) - int(Zero)
	} else if code < UpperA {
		return -1
	} else if code <= UpperF {
		return int(code) - int(UpperA) + 10
	} else if code < LowerA {
		return -1
	} else if code <= LowerF {
		return int(code) - int(LowerA) + 10
	} else {
		return -1
	}
}

func hasMask(mask, flags int) bool {
	return (mask & flags) == flags
}