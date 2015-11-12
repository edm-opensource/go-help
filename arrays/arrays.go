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

/*
ReverseStrings reverses the order of items in a slice of strings
*/
func ReverseStrings(s []string) []string {
	length := len(s) - 1
	if length < 1 {
		return s
	}

	var res []string

	for i := length; i >= 0; i-- {
		res = append(res, s[i])
	}

	return res
}
