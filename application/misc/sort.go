package misc

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// SortQuery for
func SortQuery(keys url.Values) string {
	var (
		queryKey   []string
		queryData  string
		queryValue string
		data       string
	)

	for i, _ := range keys {
		queryKey = append(queryKey, i)
	}

	for _, v := range queryKey {
		splitted := strings.Split(v, ",")
		for _, item := range splitted {
			if "sort" == item[0:4] {
				data = item[5 : len(item)-1]
				queryValue = keys.Get(item)
			}
		}
	}

	if data != "" {
		if queryValue == "DESC" {
			queryData += fmt.Sprintf("%s %s ", data, "DESC")
		} else {
			queryData += fmt.Sprintf("%s ", data)
		}
	} else {
		queryData = "updated_at DESC"
	}

	return queryData
}

func DeleteMultipleArray(data []int) []int {
	sort.Ints(data)
	j := 0
	for i := 1; i < len(data); i++ {
		if data[j] == data[i] {
			continue
		}
		j++

		data[j] = data[i]
	}
	result := data[:j+1]
	return result
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
