package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name   string
	win    int
	draw   int
	loss   int
	points int
	played int
}

type Teams map[string]Team

var teams Teams

var validResults = map[string]int{
	"win":  3,
	"loss": 0,
	"draw": 1,
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams = make(Teams)
	bufScannerIo := bufio.NewScanner(reader)
	bufScannerIo.Split(bufio.ScanLines)

	for bufScannerIo.Scan() {
		if err := processLine(bufScannerIo.Text()); err != nil {
			return err
		}
	}

	writeOutput(writer, sortTeamsByPoints(teams))

	return nil
}

func processLine(line string) error {
	if line == "" || strings.HasPrefix(line, "#") {
		return nil
	}

	values := strings.Split(line, ";")

	if len(values) != 3 {
		return errors.New("No valid line")
	}

	result := values[2]
	if _, ok := validResults[result]; !ok {
		return errors.New("No valid result")
	}

	team1 := values[0]
	teamResult1 := getTeam(team1)
	teamResult1.played += 1
	team2 := values[1]
	teamResult2 := getTeam(team2)
	teamResult2.played += 1

	if result == "win" {
		teamResult1.win += 1
		teamResult1.points += validResults["win"]
		teamResult2.loss += 1
		teamResult2.points += validResults["loss"]
	}

	if result == "draw" {
		teamResult1.draw += 1
		teamResult1.points += validResults["draw"]
		teamResult2.draw += 1
		teamResult2.points += validResults["draw"]
	}

	if result == "loss" {
		teamResult1.loss += 1
		teamResult1.points += validResults["loss"]
		teamResult2.win += 1
		teamResult2.points += validResults["win"]
	}

	teams[team1] = teamResult1
	teams[team2] = teamResult2

	return nil
}

func getTeam(name string) Team {
	if team, ok := teams[name]; ok {
		return team
	}

	return Team{
		name: name,
	}
}

func sortTeamsByPoints(teams Teams) []Team {
	sortedTeams := make([]Team, 0, len(teams))

	for _, value := range teams {
		sortedTeams = append(sortedTeams, value)
	}

	sort.Slice(sortedTeams, func(a, b int) bool {
		if sortedTeams[a].points > sortedTeams[b].points {
			return true
		}

		if sortedTeams[a].points < sortedTeams[b].points {
			return false
		}

		return strings.Compare(sortedTeams[a].name, sortedTeams[b].name) <= 0
	})

	return sortedTeams
}

func writeOutput(writer io.Writer, teams []Team) {
	writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))

	for _, team := range teams {
		writer.Write([]byte(fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n",
			team.name, team.played, team.win, team.draw, team.loss, team.points)))
	}
}
