package executor

import (
	"log"
	"reflect"
	"testing"
)

func TestStringIsEmpty(t *testing.T) {
	i1 := ""
	i2 := "."
	i3 := "abcd"
	i4 := "\\"
	e1 := false
	e2 := true
	e3 := true
	e4 := true
	r1 := stringIsNotEmpty(i1)
	r2 := stringIsNotEmpty(i2)
	r3 := stringIsNotEmpty(i3)
	r4 := stringIsNotEmpty(i4)
	if r1 != e1 {
		log.Fatalf("error r1 %v, expected %v", r1, e1)
	} else if r2 != e2 {
		log.Fatalf("error r2 %v, expected %v", r2, e2)
	} else if r3 != e3 {
		log.Fatalf("error r3 %v, expected %v", r3, e3)
	} else if r4 != e4 {
		log.Fatalf("error r4 %v, expected %v", r4, e4)
	}
}

func TestFilterMap(t *testing.T) {
	i1 := map[string]string{"key1": "", "key2": "abc", "key3": "\\"}
	e1 := map[string]string{"key2": "abc", "key3": "\\"}
	filterMap(i1, stringIsNotEmpty)
	if !reflect.DeepEqual(i1, e1) {
		log.Fatalf("error")
	}
}

func TestUnescapeQuotes(t *testing.T) {
	i1 := `SELECT 1 \\"coucou\\"`
	i2 := "SELECT 1"
	i3 := ""
	i4 := "\\"
	e1 := `SELECT 1 "coucou"`
	e2 := "SELECT 1"
	e3 := ""
	e4 := ""

	r1 := unescapeQuotes(i1)
	r2 := unescapeQuotes(i2)
	r3 := unescapeQuotes(i3)
	r4 := unescapeQuotes(i4)

	if r1 != e1 {
		log.Fatalf("r1 is '%v', expected '%v'", r1, e1)
	} else if r2 != e2 {
		log.Fatalf("r2 is '%v', expected '%v'", r2, e2)
	} else if r3 != e3 {
		log.Fatalf("r3 is '%v', expected '%v'", r3, e3)
	} else if r4 != e4 {
		log.Fatalf("r4 is '%v', expected '%v'", r4, e4)
	}

}

func TestDuplicateQueryParameterInArray(t *testing.T) {
	log.Println("TestDuplicateQueryParameterInArray")
	i1 := map[string]string{"mydataset1": "SELECT 1"}
	j1 := constructorQueryParameter(i1)
	i2 := map[string]string{"mydataset1": "SELECT 1"}
	j2 := constructorQueryParameter(i2)
	i3 := map[string]string{"mydataset2": "SELECT 1"}
	j3 := constructorQueryParameter(i3)
	l := []QueryParameter{j1, j2, j3}

	e := []QueryParameter{j1, j3}

	r := removeDuplicateQueryParameterInArray(l)

	if len(r) != 2 {
		log.Fatalf("Filtered QueryParameters are %v long, expected 2", len(r))
	}
	if r[0] != e[0] {
		log.Fatalf("First query is %v, expected %v", r[0], e[0])
	}
	if r[1] != e[1] {
		log.Fatalf("First query is %v, expected %v", r[1], e[1])
	}
}

func TestDuplicaterStrInArray(t *testing.T) {
	log.Println("TestDuplicaterStrInArray")
	i1 := []string{"SELECT 1", "SELECT 1", "SELECT 2"}
	i2 := []string{}

	e1 := []string{"SELECT 1", "SELECT 2"}

	r1 := removeDuplicateStrInArray(i1)
	r2 := removeDuplicateStrInArray(i2)

	if len(r1) != 2 {
		log.Fatalf("Filtered QueryParameters are %v long, expected 2", len(r1))
	}
	if len(r2) != 0 {
		log.Fatalf("Filtered QueryParameters are %v long, expected 0", len(r2))
	}
	if r1[0] != e1[0] {
		log.Fatalf("First query is %v, expected %v", r1[0], e1[0])
	}
	if r1[1] != e1[1] {
		log.Fatalf("First query is %v, expected %v", r1[1], e1[1])
	}
}
