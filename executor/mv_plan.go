package executor

import (
	"sort"
)

const (
	DELETE = "delete"
	CREATE = "create"
)

type plan map[MvDefinition]string

func PlanEqual(a, b plan) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func MvsEqual(a, b []MvDefinition) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].IdInt < b[j].IdInt
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i].IdInt < b[j].IdInt
	})
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func MvIsInMvArray(mv MvDefinition, mvs []MvDefinition) bool {
	for _, v := range mvs {
		if mv == v {
			return true
		}
	}
	return false
}

func BuildPlan(newMvs []MvDefinition, oldMvs []MvDefinition) plan {
	p := make(plan)
	for _, v := range newMvs {
		if !MvIsInMvArray(v, oldMvs) {
			p[v] = CREATE
		}
	}
	for _, v := range oldMvs {
		if !MvIsInMvArray(v, newMvs) {
			p[v] = DELETE
		}
	}
	return p
}
