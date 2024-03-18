package main

import (
	"fmt"
	"strconv"
	"strings"
)

var split_main = strings.Repeat("-", 10)

func main() {
	fmt.Println(split_main + "Codewars Start" + split_main)
	//Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})
	//add function here
	fmt.Println(EncodeCd(128))
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

func EncodeCd(n uint8) string {
	//7kyu - Encode data on CD (Compact Disc) surface
	//8 bit = 128-64-32-16-8-4-2-1 -> where n < 256 [0..255]
	//convert n to a binary sequence
	var one int = 1
	var bit_8_notation = [8]int{0, 0, 0, 0, 0, 0, 0, 0} //where index = 0 is 8 and index = 1 is 0
	var bit_8, bit_7, bit_6, bit_5, bit_4, bit_3, bit_2, bit_1 uint8 = 128, 64, 32, 16, 8, 4, 2, 1
	current_number := n

	if current_number >= bit_8 {
		current_number -= bit_8
		bit_8_notation[0] = one
	}
	if current_number >= bit_7 {
		current_number -= bit_7
		bit_8_notation[1] = one
	}
	if current_number >= bit_6 {
		current_number -= bit_6
		bit_8_notation[2] = one
	}
	if current_number >= bit_5 {
		current_number -= bit_5
		bit_8_notation[3] = one
	}
	if current_number >= bit_4 {
		current_number -= bit_4
		bit_8_notation[4] = one
	}
	if current_number >= bit_3 {
		current_number -= bit_3
		bit_8_notation[5] = one
	}
	if current_number >= bit_2 {
		current_number -= bit_2
		bit_8_notation[6] = one
	}
	if current_number >= bit_1 {
		current_number -= bit_1
		bit_8_notation[7] = one
	}

	var cd = [9]string{"P", "P", "P", "P", "P", "P", "P", "P", "P"}

	//0 - 0:1, 1 - 1:2, 2 - 2:3 ... n - n:n+1
	//end[1 1 0 1 1 1 1 0]start
	for notation := range bit_8_notation {
		notation++            //notation start from 1
		index := 8 - notation //index is used for binary, notation is used for cd
		if bit_8_notation[index] == 1 {
			if cd[notation-1] == "P" {
				cd[notation] = "L"
			} else {
				cd[notation] = "P"
			}
		} else {
			if cd[notation-1] == "P" {
				cd[notation] = "P"
			} else {
				cd[notation] = "L"
			}
		}
	}

	return strings.Join(cd[:], "") //convert cd list to slice before joining
}
