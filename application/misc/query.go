package misc

import (
	"strconv"
	"strings"
)

func QueryConvert(query string) []uint64 {
	var dataArray []uint64

	splitted := strings.Split(query, ",")
	for _, v := range splitted {
		dataUint, _ := strconv.ParseUint(v, 10, 32)
		dataArray = append(dataArray, dataUint)
	}

	return dataArray
}
