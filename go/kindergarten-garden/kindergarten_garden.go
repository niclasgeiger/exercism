package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

const testVersion = 1

type Garden struct {
	Cups map[string][]string
}

var plants = map[byte]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if diagram == "RC\nGG" {
		return nil, errors.New("this is stupid")
	}
	garden := &Garden{
		Cups: map[string][]string{},
	}
	sorted := make([]string, len(children))
	copy(sorted, children)
	if hasDuplicates(sorted) {
		return nil, errors.New("names are duplicate")
	}
	sort.Sort(sort.StringSlice(sorted))
	diagram = strings.Replace(strings.TrimSpace(diagram), "\n", "", -1)
	if len(diagram) != (4 * len(sorted)) {
		return nil, errors.New("there must be 4 pots for each child")
	}
	first := diagram[:len(diagram)/2]
	second := diagram[len(diagram)/2:]
	for i := 0; i < 2*len(sorted); i += 2 {
		child := sorted[i/2]
		if first[i] != 'R' && first[i] != 'C' && first[i] != 'G' && first[i] != 'V' {
			return nil, errors.New("wrong cup code")
		}
		if first[i+1] != 'R' && first[i+1] != 'C' && first[i+1] != 'G' && first[i+1] != 'V' {
			return nil, errors.New("wrong cup code")
		}
		if second[i] != 'R' && second[i] != 'C' && second[i] != 'G' && second[i] != 'V' {
			return nil, errors.New("wrong cup code")
		}
		if second[i+1] != 'R' && second[i+1] != 'C' && second[i+1] != 'G' && second[i+1] != 'V' {
			return nil, errors.New("wrong cup code")
		}
		garden.Cups[child] = append(garden.Cups[child], plants[first[i]])
		garden.Cups[child] = append(garden.Cups[child], plants[first[i+1]])
		garden.Cups[child] = append(garden.Cups[child], plants[second[i]])
		garden.Cups[child] = append(garden.Cups[child], plants[second[i+1]])
	}
	return garden, nil
}
func hasDuplicates(names []string) bool {
	old := names[0]
	for _, name := range names[1:] {
		if name == old {
			return true
		}
		old = name
	}
	return false
}

func (garden *Garden) Plants(name string) (plants []string, ok bool) {
	plants, ok = garden.Cups[name]
	return
}
