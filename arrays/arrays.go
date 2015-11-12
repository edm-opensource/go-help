package arrays

func RemoveDuplicates(slice []int) []int {
	lenSlice := len(slice)

	for i := 0; i < lenSlice; i++ {
		for j := i + 1; j < lenSlice; j++ {
			if slice[i] == slice[j] {
				slice = append(slice[:j], slice[j+1:]...)
				lenSlice = len(slice)
			}
		}
	}

	return slice
}
