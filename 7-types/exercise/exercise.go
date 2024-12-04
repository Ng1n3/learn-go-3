/*
1. You have been asked to manage a basketball league and are going to write a
program to help you. Define two types. The first one, called Team, has a field
for the name of the team and a field for the player names. The second type is
called League and has a field called Teams for the teams in the league and a
field called Wins that maps a team’s name to its number of wins.
2. Add two methods to League. The first method is called MatchResult. It takes
four parameters: the name of the first team, its score in the game, the name of
the second team, and its score in the game. This method should update the
Wins field in League. Add a second method to League called Ranking that
returns a slice of the team names in order of wins. Build your data structures
and call these methods from the main function in your program using some
sample data.
3. Define an interface called Ranker that has a single method called Ranking
that returns a slice of strings. Write a function called RankPrinter with two
parameters, the first of type Ranker and the second of type io.Writer. Use
the io.WriteString function to write the values returned by Ranker to the
io.Writer, with a newline separating each result. Call this function from
main.

*/

package exercise

import (
	"io"
	"sort"
)

type Team struct {
	Teamname   string
	Playername []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League )MatchResult(firstTeamName string, firstTeamScore int, secondTeamName string, secondTeamScore int) {
 if firstTeamScore > secondTeamScore {
  l.Wins[firstTeamName]++
 } else if secondTeamScore > firstTeamScore {
  l.Wins[secondTeamName]++
 }
}

func (l League) Ranking() []string {
  teamNames := make([]string, 0, len(l.Wins))
  for teamName := range l.Wins {
    teamNames = append(teamNames, teamName)
  }

  sort.Slice(teamNames, func(i, j int) bool {
    return l.Wins[teamNames[i]] > l.Wins[teamNames[j]]
  })

  return teamNames
}

type Ranker interface {
  Ranking()[]string
}

func RankPrinter (ranker Ranker, writer io.Writer) {
  rankings := ranker.Ranking()
  for _, rank := range rankings {
    io.WriteString(writer, rank + "\n")
  } 
}
