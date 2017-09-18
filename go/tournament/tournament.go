package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const testVersion = 4

const (
	WIN  = "win"
	LOSS = "loss"
	DRAW = "draw"
)

type Tournament struct {
	Table
}

// sort.Interface
type Table map[Team]Score

type Score struct {
	Matches int
	Win     int
	Draw    int
	Lost    int
	Points  int
}

type Team string

func Tally(reader io.Reader, writer io.Writer) error {
	tournament := &Tournament{
		Table: make(map[Team]Score),
	}
	bufReader := bufio.NewReader(reader)
	line, _, err := bufReader.ReadLine()
	for ; line != nil; line, _, err = bufReader.ReadLine() {
		if err != nil {
			return err
		}
		info := strings.Split(string(line), ";")
		if len(info) > 2 {
			err := tournament.AddScore(Team(info[0]), Team(info[1]), info[2])
			if err != nil {
				return err
			}
		} else if len(line) > 0 && line[0] != '#' {
			return errors.New("Wrong Format")
		}
	}
	tournament.PrintScore(writer)
	return nil
}

func (t *Tournament) PrintScore(writer io.Writer) {
	sorted := SortedTournament{}
	for team, score := range t.Table {
		entry := SortedEntry{
			Team:  team,
			Score: score,
		}
		sorted = append(sorted, entry)
	}
	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	sort.Sort(sorted)
	for _, entry := range sorted {
		fill := ""
		for i := 0; i < 31-len(entry.Team); i++ {
			fill += " "
		}
		line := fmt.Sprintf("%s%s|  %d |  %d |  %d |  %d |  %d\n", entry.Team, fill, entry.Matches, entry.Win, entry.Draw, entry.Lost, entry.Points)
		writer.Write([]byte(line))
	}
}

func (t *Tournament) AddScore(team1, team2 Team, outcome string) error {
	switch outcome {
	case WIN:
		{
			t.Table[team1] = t.Table[team1].AddWin()
			t.Table[team2] = t.Table[team2].AddLost()
		}
	case DRAW:
		{
			t.Table[team1] = t.Table[team1].AddDraw()
			t.Table[team2] = t.Table[team2].AddDraw()
		}
	case LOSS:
		{
			t.Table[team1] = t.Table[team1].AddLost()
			t.Table[team2] = t.Table[team2].AddWin()
		}
	default:
		{
			return errors.New("Unknown outcome!")
		}
	}
	return nil
}

func (s Score) AddWin() Score {
	s.Win++
	s.Points += 3
	s.Matches++
	return s
}

func (s Score) AddLost() Score {
	s.Lost++
	s.Matches++
	return s
}

func (s Score) AddDraw() Score {
	s.Draw++
	s.Points++
	s.Matches++
	return s
}

type SortedTournament []SortedEntry

type SortedEntry struct {
	Team
	Score
}

func (s SortedTournament) Len() int {
	return len(s)
}

func (s SortedTournament) Less(i, j int) bool {
	if s[i].Points > s[j].Points {
		return true
	}
	if s[i].Points < s[j].Points {
		return false
	}
	return strings.Compare(string(s[i].Team), string(s[j].Team)) < 0
}

func (s SortedTournament) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
