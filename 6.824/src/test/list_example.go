package main

import "fmt"
import "container/list"
import "strings"
import "unicode"

func Map(value string) *list.List {
	list := list.New()

	f := func(c rune) bool {
	    return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	array_str := strings.FieldsFunc(value, f)
	for i := 0; i < len(array_str); i++ {
		fmt.Println("Element", i, "of array is", array_str[i])
		list.PushBack(array_str[i])
	}
	return list
}

func PrintList(value *list.List) {
	// Iterate through list and print its contents.
	for e := value.Front(); e != nil; e = e.Next() {
	    fmt.Println(e.Value)
	}
}

func main() {
	PrintList(Map("   aaaa,bbbb 1111, cccc333.^&*() ddddddd2342"))
}