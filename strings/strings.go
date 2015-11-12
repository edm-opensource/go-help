package strings

import "strings"

/*
TitleFirstWord bra skit va
*/
func TitleFirstWord(s string) string {
	res := strings.ToUpper(string([]rune(s)[0])) + s[1:]
	return res
}
