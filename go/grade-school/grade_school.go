package school

import (
	"sort"
	"strings"
)

const testVersion = 1

type Grade struct {
	Level    int
	Students []string
}
type Grades []Grade
type School struct {
	Grades Grades
}

func New() *School {
	return &School{}
}

func (s *School) Add(student string, level int) {
	for i, g := range s.Grades {
		if g.Level == level {
			s.Grades[i].Students = append(s.Grades[i].Students, student)
			sort.Sort(s.Grades[i])
			return
		}
	}
	grade := Grade{
		Level:    level,
		Students: []string{student},
	}
	s.Grades = append(s.Grades, grade)
}

func (s *School) Grade(level int) []string {
	for _, g := range s.Grades {
		if g.Level == level {
			return g.Students
		}
	}
	return []string{}
}
func (s *School) Enrollment() []Grade {
	sort.Sort(s.Grades)
	return s.Grades
}

// -------Sorting Implmentation for Grade-----
func (grade Grade) Len() int {
	return len(grade.Students)
}

func (grade Grade) Less(i, j int) bool {
	return strings.Compare(grade.Students[i], grade.Students[j]) < 0
}

func (grade Grade) Swap(i, j int) {
	grade.Students[i], grade.Students[j] = grade.Students[j], grade.Students[i]
}

// -------Sorting Implmentation for Grades-----
func (grades Grades) Len() int {
	return len(grades)
}

func (grades Grades) Less(i, j int) bool {
	return grades[i].Level < grades[j].Level
}

func (grades Grades) Swap(i, j int) {
	grades[i], grades[j] = grades[j], grades[i]
}
