package main

import "fmt"

type sport_team struct {
	name    string
	sport   string
	ranking int
}

func main() {
	m := make(map[string]sport_team)

	sport_team1 := sport_team{
		name:    "Real Madrid",
		sport:   "Football",
		ranking: 1,
	}
	m[sport_team1.name] = sport_team1

	sport_team2 := sport_team{
		name:    "F.C. Barcelona",
		sport:   "Football",
		ranking: 3,
	}

	m[sport_team2.name] = sport_team2

	sport_team3 := sport_team{
		name:    "Atlético de Madrid",
		sport:   "Football",
		ranking: 2,
	}

	m[sport_team3.name] = sport_team3

	for s, v := range m {
		fmt.Printf("Key:%s\nValue:%v\n", s, v)
	}

}
