package common

func SlicePlace(find byte, slice []byte) int {
	if len(slice) == 0 {
		return -1
	}
	for i := 0; i < len(slice); i++ {
		if find == slice[i] {
			return i
		}
	}
	return 0
}
