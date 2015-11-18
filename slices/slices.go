package slices

/*
RemoveDuplicateInts finds and removes any duplicate values for a given int slice
*/
func RemoveDuplicateInts(slice []int) []int {
	lenSlice := len(slice)

	if lenSlice < 2 {
		return slice
	}

	for i := 0; i < lenSlice; i++ {
		for j := i + 1; j < lenSlice; j++ {
			if slice[i] == slice[j] {
				slice = append(slice[:j], slice[j+1:]...)
				lenSlice = len(slice)
				j--
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

/*
ReverseInts reverses the order of items in a slice of ints
*/
func ReverseInts(s []int) []int {
	length := len(s) - 1
	if length < 1 {
		return s
	}

	var res []int

	for i := length; i >= 0; i-- {
		res = append(res, s[i])
	}

	return res
}

/*
ReverseFloat64s reverses the order of items in a slice of floats64s
*/
func ReverseFloat64s(s []float64) []float64 {
	length := len(s) - 1
	if length < 1 {
		return s
	}

	var res []float64

	for i := length; i >= 0; i-- {
		res = append(res, s[i])
	}

	return res
}
