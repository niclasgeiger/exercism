package stringset

import (
	"fmt"
)

const testVersion = 4

type Set []string

func New() Set {
	return []string{}
}

func NewFromSlice(slice []string) Set {
	set := New()
	duplicate := map[string]int{}
	for _, s := range slice {
		if _, ok := duplicate[s]; !ok {
			set = append(set, s)
		}
		duplicate[s] = 1
	}
	return set
}

func (set Set) String() string {
	if set.IsEmpty() {
		return "{}"
	}
	joined := ""
	for _, item := range set {
		joined += fmt.Sprintf("\"%s\", ", item)
	}
	return fmt.Sprintf("{%s}", joined[:len(joined)-2])
}

func (set Set) IsEmpty() bool {
	return set.Length() == 0
}

func (set *Set) Add(s string) {
	if !set.Has(s) {
		*set = append(*set, s)
	}
}

func (set Set) Has(s string) bool {
	for _, item := range set {
		if item == s {
			return true
		}
	}
	return false
}

func (set Set) Length() int {
	return len(set)
}

func Equal(a, b Set) bool {
	if a.Length() != b.Length() {
		return false
	}
	return Union(Difference(a, b), Difference(b, a)).Length() == 0
}

func Subset(a, b Set) bool {
	if a.IsEmpty() {
		return true
	}
	if b.IsEmpty() {
		return false
	}
	for _, item := range a {
		if !b.Has(item) {
			return false
		}
	}
	return true
}
func Union(a, b Set) Set {
	union := New()
	match := map[string]int{}
	for _, item := range a {
		if _, found := match[item]; !found {
			union = append(union, item)
			match[item] = 1
		}
	}
	for _, item := range b {
		if _, found := match[item]; !found {
			union = append(union, item)
			match[item] = 1
		}
	}
	return union
}

func Intersection(a, b Set) Set {
	union := Union(a, b)
	difference := Union(Difference(a, b), Difference(b, a))
	intersection := New()
	for _, item := range union {
		if !difference.Has(item) {
			intersection = append(intersection, item)
		}
	}
	return intersection
}

func Disjoint(a, b Set) bool {
	return Intersection(a, b).Length() == 0
}

func Difference(a, b Set) Set {
	difference := New()
	for _, item := range a {
		if !b.Has(item) {
			difference = append(difference, item)
		}
	}
	return difference
}
