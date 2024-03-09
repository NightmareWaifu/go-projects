package main

import (
	"fmt"
	"strconv"
	"strings"
)

var split_main = strings.Repeat("-", 10)

func main() {
	fmt.Println(split_main + "Codewars Start" + split_main)
	Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})
	fmt.Println(split_main + "Codewars End" + split_main)
}

func Quadratic(x1, x2 int) [3]int {
	//8kyu - Quadratic Coefficients Solver
	//ax^2 + bx + c = 0 where a = 1
	var a int = 1
	var b int = -(x1) + -(x2)
	var c int = -(x1) * -(x2)
	return [3]int{a, b, c}
}

func QuarterOf(month int) int {
	//8kyu - Quarter of the year
	// 1-12, 1 = Jan, 12 = Dec
	//1-3, 4-6, 7-9, 10-12

	// if 1 <= month && month <= 3 {
	// 	return 1
	// } else if 4 <= month && month <= 6 {
	// 	return 2
	// } else if 7 <= month && month <= 9 {
	// 	return 3
	// } else if 10 <= month && month <= 12 {
	// 	return 4
	// } else {
	// 	return 0
	// }

	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	case 10, 11, 12:
		return 4
	default:
		return 0
	}

}

func Points(games []string) int {
	//8kyu - Total amount of points
	//10 matches played
	var total_points int
	var win int = 3
	var draw int = 1
	var lose int = 0

	for index := range games {
		score := strings.Split(games[index], ":")

		team_points, err_team := strconv.Atoi(score[0])
		opponent_points, err_opponent := strconv.Atoi(score[1])

		if err_team != nil || err_opponent != nil {
			return 1
		}

		if team_points > opponent_points {
			total_points += win
		} else if team_points == opponent_points {
			total_points += draw
		} else {
			total_points += lose
		}
	}
	return total_points
}
