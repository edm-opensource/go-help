package strings

import "testing"

func TestTitleFirstWord_1(t *testing.T) {
	if TitleFirstWord("hello world") != "Hello world" {
		t.Error("hello world was not parsed sucessfully")
	}
}
func TestTitleFirstWord_2(t *testing.T) {
	if TitleFirstWord("Hello world") != "Hello world" {
		t.Error("Hello world was not parsed sucessfully")
	}
}
