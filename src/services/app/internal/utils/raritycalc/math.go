package raritycalc

func Min[T int | int32 | int64 | uint | uint32 | uint64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T int | int32 | int64 | uint | uint32 | uint64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
