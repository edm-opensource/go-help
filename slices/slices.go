package slices

/*
RemoveDuplicateInts finds and removes any duplicate values for a given int slice
*/
func RemoveDuplicateInts(slice []int) []int {
	duplicate := make(map[int]bool)
	result := []int{}

	for v := range slice {
		if !duplicate[slice[v]] {
			duplicate[slice[v]] = true
			result = append(result, slice[v])
		}
	}
	return result
}

/*
RemoveDuplicateFloat64 finds and removes any duplicate values for a given float64 slice
*/
func RemoveDuplicateFloat64(slice []float64) []float64 {
	encountered := make(map[float64]bool)
	result := []float64{}

	for v := range slice {
		if !encountered[slice[v]] {
			encountered[slice[v]] = true
			result = append(result, slice[v])
		}
	}
	return result
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
