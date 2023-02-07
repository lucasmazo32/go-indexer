package get_index

import "strings"

func GetIndex(path string) string {
	str := path[27:]
	strs := strings.Split(str, "/")[1:3]
	return strings.Join(strs, "/")
}