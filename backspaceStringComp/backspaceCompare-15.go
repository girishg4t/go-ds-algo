package backspacestringcomp

import "fmt"

//  s = "ab#c", t = "ad#c"
func BackspaceCompare(s string, t string) bool {
	i := len(s) - 1
	count := 0
	for i > -1 {
		if string(s[i]) == "#" {
			count++
			i--
			continue
		}
		j := 0
		newIndex := i
		for j != count {
			s = s[:newIndex-j] + "#" + s[newIndex+1:]
			j++
		}
		count = 0
		i--
	}
	fmt.Println(s)
	return true
}
