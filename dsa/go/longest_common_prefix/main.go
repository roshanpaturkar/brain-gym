package main

import (
	"fmt"
	"sort"
	"strings"
)

func longest_common_prefix(strArray []string) string {
	sort.Slice(strArray, func(i, j int) bool {
		return len(strArray[i]) < len(strArray[j])
	})

	pre := ""

	small_word := strings.Split(strArray[0], "")

	for i :=0; i < len(small_word); i++ {
		found := true
		for j := 1; j <len(strArray); j++ {
			if small_word[i] != string(strArray[j][i]) {
				found = false
			}
		}

		if found {
			pre += small_word[i]
		}
	}

	return pre
}

func main() {

	name := "roshan"
	name1 := "rohan"

	abc := strings.HasPrefix(name, string(name1[:2]))
	fmt.Println(abc)

	prf := longest_common_prefix([]string{"flower","flow","flight"})
	fmt.Println(prf)

	prf = longest_common_prefix([]string{"dog","racecar","car"})
	fmt.Println(prf)
}