package executor

import (
	"log"
	"strings"
)

// unescapeQuotes Removes \ from a string if precedes a double quote
func unescapeQuotes(vs string) string {
	r := strings.ReplaceAll(vs, "\\", "")
	return r
}

func stringIsNotEmpty(s string) bool {
	return s != ""
}

func removeDuplicateStrInArray(s []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range s {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		} else {
			log.Println("WARNING: DUPLICATE QUERY IN OPTIMIZATION PLAN")
			// TODO: Send a notification to developper if this happens
		}
	}
	return list
}

func removeDuplicateQueryParameterInArray(qs []QueryParameter) []QueryParameter {
	// TODO: Duplicate query should only be dataset and statement, not mmvName
	rqs := make([]QueryParameter, 0)
	allKeys := make(map[QueryParameter]bool)
	for _, item := range qs {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			rqs = append(rqs, item)
		} else {
			log.Println("WARNING: DUPLICATE QUERY IN OPTIMIZATION PLAN")
			// TODO: Send a notification to developper if this happens
		}
	}
	return rqs
}

// filterMap delete entries from a map that return false when passed to function f
func filterMap(ms map[string]string, f func(string) bool) {
	for k, v := range ms {
		if !f(v) {
			delete(ms, k)
		}
	}
}
