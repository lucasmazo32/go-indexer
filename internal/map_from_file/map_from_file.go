package map_from_file

import (
	"os"
	"regexp"
	"strings"

	"indexer.com/indexer/internal/check_error"
)

func MapFromFile(file string) *map[string]string {
	dat, err := os.ReadFile(file)
	check_error.Check(err)
	data := string(dat)
	reg := regexp.MustCompile("\n[a-zA-Z\\-]+: ")
	subjectReg := regexp.MustCompile("\n[a-zA-Z\\-]+: ")
	dataArr := reg.Split(data[12:], -1)
	mapArr := []string{"\nMessage-ID: "}
	subArr := subjectReg.FindAllString(data, -1)
	mapArr = append(mapArr, subArr...)
	mapArr = append(mapArr, "\nMessage: ")
	transformData := map[string]string{}
	subLen := len(dataArr)
	for i, v := range dataArr {
		if i == subLen-1 {
			subV := strings.SplitN(v, "\n", 2)
			k1 := strings.TrimPrefix(mapArr[i], "\n")
			lenK1 := len(k1)
			k2 := strings.TrimPrefix(mapArr[i+1], "\n")
			lenK2 := len(k2)
			transformData[k1[0:lenK1-2]] = subV[0]
			if len(subV) == 1 {
				transformData[k2[0:lenK2-2]] = ""
			} else {
				transformData[k2[0:lenK2-2]] = subV[1]
			}
		} else {
			k := strings.TrimPrefix(mapArr[i], "\n")
			lenK := len(k)
			transformData[k[0:lenK-2]] = v
		}
	}
	return &transformData
}
