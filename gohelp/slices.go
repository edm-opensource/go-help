package gohelp

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
RemoveDuplicateString finds and removes any duplicate values for a given string slice
*/
func RemoveDuplicateString(slice []string) []string {
	encountered := make(map[string]bool)
	result := []string{}

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

/*
RemoveAllInstancesOfStrings returns a copy of the first slice parameter,
with all instances of the rest of the arguments removed
*/
func RemoveAllInstancesOfStrings(s []string, toBeRemoved ...string) []string {
	length := len(s)
	var res []string

	for i := 0; i < length; i++ {
		if !ContainsString(toBeRemoved, s[i]) {
			res = append(res, s[i])
		}
	}

	return res
}

/*
RemoveAllInstancesOfInts returns a copy of the first slice parameter,
with all instances of the rest of the arguments removed
*/
func RemoveAllInstancesOfInts(s []int, toBeRemoved ...int) []int {
	length := len(s)
	var res []int

	for i := 0; i < length; i++ {
		if !ContainsInt(toBeRemoved, s[i]) {
			res = append(res, s[i])
		}
	}

	return res
}

/*
RemoveAllInstancesOfFloat64s returns a copy of the first slice parameter,
with all instances of the rest of the arguments removed
*/
func RemoveAllInstancesOfFloat64s(s []float64, toBeRemoved ...float64) []float64 {
	length := len(s)
	var res []float64

	for i := 0; i < length; i++ {
		if !ContainsFloat64(toBeRemoved, s[i]) {
			res = append(res, s[i])
		}
	}

	return res
}

/*
ContainsString returns true if the slice contains the string
*/
func ContainsString(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

/*
ContainsInt returns true if the slice contains the int
*/
func ContainsInt(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

/*
ContainsFloat64 returns true if the slice contains the float64
*/
func ContainsFloat64(slice []float64, item float64) bool {
	set := make(map[float64]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
