package main

import (
	"fmt"
	"strconv"
	"strings"
)

var split_main = strings.Repeat("-", 10)

func main() {
	fmt.Println(split_main + "Codewars Start" + split_main)

	//add function here
	//Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})
	//fmt.Println(EncodeCd(128))
	//TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq")
	//SpinWords("Hey fellow warriors")
	//TwiceAsOld(36, 7)
	//fmt.Println(dots_on_domino_bones(2))
	//CreateBox(10, 10)
	//fmt.Println(Is_valid_ip("12.255.56.1"))
	//Rot13("EBG13 rknzcyr.")

	fmt.Println(AlphabetWar("twjbzwtdmjqwqtzbdjpzdj"))
	fmt.Println(split_main + "Codewars End" + split_main)
}

//---------------------------------8kyu

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

func EvenOrOdd(number int) string {
	//8kyu - Even or Odd
	result := "Odd"
	is_odd_or_even := number % 2

	if is_odd_or_even == 0 {
		result = "Even"
	}

	return result
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	//8kyu - Twice as old
	for i := 0; i >= 0; i++ {
		if (dadYearsOld+i) == 2*(sonYearsOld+i) || (dadYearsOld-i) == 2*(sonYearsOld-i) {
			return i
		}
	}
	return 0

	//smart solution
	//return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))

	//explanation
	// at target, a + x = 2(b+x)
	// : a + x = 2b + 2x
	// : a - 2b = x

}

func countSheep(num int) string {
	//8kyu - If you can't sleep, just count sheep!!
	var result string = ""
	for i := 1; i <= num; i++ {
		result += (strconv.Itoa(i) + " sheep...")
	}

	return result
}

func ExpressionMatter(a int, b int, c int) int {
	//8kyu - Expressions Matter
	largest := a + b + c
	if largest < a+b*c {
		largest = a + b*c
	}
	if largest < a*b+c {
		largest = a*b + c
	}
	if largest < (a+b)*c {
		largest = (a + b) * c
	}
	if largest < a*(b+c) {
		largest = a * (b + c)
	}
	if largest < a*b*c {
		largest = a * b * c
	}

	return largest
}

//---------------------------------7kyu

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

func TwoToOne(s1 string, s2 string) string {
	//7kyu - Two to One
	var all_letters_str string = "abcdefghijklmnopqrstuvwxyz"
	var all_letters = strings.Split(all_letters_str, "") //split the letters because im too lazy to type it out

	var letter_flags [26]int
	for index := range all_letters { //init letter flags cuz im lazy
		letter_flags[index] = 0
	}

	//check for chars

	//old

	// for index := range all_letters {
	// 	for s1_index := range s1 {
	// 		if string(s1[s1_index]) == all_letters[index] {
	// 			//works
	// 			letter_flags[index] = 1
	// 		}
	// 	}

	// 	for s2_index := range s2 {
	// 		if string(s2[s2_index]) == all_letters[index] {
	// 			letter_flags[index] = 1
	// 		}
	// 	}
	// }

	//new - improvement from seeing other solutions -> instead of 2 for loops you can combine both strings and use 1 for loop instead

	var combined_string string = s1 + s2

	for index := range all_letters {
		for string_index := range combined_string {
			if string(combined_string[string_index]) == all_letters[index] {
				letter_flags[index] = 1
			}
		}
	}

	var longest_string string
	//creating the string
	for index := range letter_flags {
		if letter_flags[index] == 1 {
			longest_string += all_letters[index]
		}
	}

	//fmt.Println(longest_string)
	return longest_string
}

func dots_on_domino_bones(n int) int {
	//7kyu - Dots on Domino's Bones
	var total_points int = 0
	for i := 0; i <= n; i++ {
		for j := i; j <= n; j++ {
			total_points += j
			total_points += i
		}
	}
	return total_points
}

func CreateBox(width, length int) [][]int {
	//7kyu - The 'spiraling' box

	//number of lists = length
	//length of list = width

	var box [][]int
	var created_box [][]int
	//var half int = (length + 1) / 2

	for row := 1; row <= (length+1)/2; row++ {
		push_list := []int{}
		for column := 1; column <= (width)/2; column++ { //first half
			if row < column {
				push_list = append(push_list, row)
			} else {
				push_list = append(push_list, column)
			}
		}

		hold_list := push_list
		if width%2 == 1 {
			if row > width/2+1 {
				push_list = append(push_list, width/2+1)
			} else {
				push_list = append(push_list, row)
			}
		}

		for index := (width / 2) - 1; index >= 0; index-- { //second half
			push_list = append(push_list, hold_list[index])
		}
		box = append(box, push_list)
	}
	for i := 0; i < len(box); i++ {
		created_box = append(created_box, box[i])
		//fmt.Println(box[i])
	}
	//if odd number of rows, add the odd row
	var num int = 1
	if length%2 == 1 {
		fmt.Println("Found is odd")
		num = 2
	}
	for i := len(box) - num; i >= 0; i-- {
		created_box = append(created_box, box[i])
		//fmt.Println(box[i])
	}

	for i := 0; i < len(created_box); i++ {
		fmt.Println(created_box[i])
	}
	return created_box
}

//---------------------------------6kyu

func Multiple3And5(number int) int {
	//6kyu - Multiples of 3 or 5

	var sum int

	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func SpinWords(str string) string {
	//6kyu - Stop gninnipS My sdroW!

	all_words := strings.Split(str, " ")
	for index := range all_words {
		if len(all_words[index]) >= 5 {
			var new_string string
			var new_string_index int
			for letter_index := range all_words[index] {
				new_string_index = len(all_words[index]) - letter_index - 1
				new_string += string(all_words[index][new_string_index])
			}
			all_words[index] = new_string
		}
	}
	return strings.Join(all_words, " ")
}

func Is_valid_ip(ip string) bool {
	//6kyu - IP Validation

	sections := strings.Split(ip, ".")
	if len(sections) != 4 {
		return false
	}
	for i := 0; i < len(sections); i++ {
		if string(sections[i][0]) == "0" && len(sections[i]) != 1 {
			return false
		}
		bit, err := strconv.Atoi(string(sections[i]))
		if err == nil {
			if bit < 0 || bit > 255 {
				return false
			}
		} else {
			return false
		}
	}
	return true

	//improvments
	//there's a .ParseIP() function lol
}

//--------------------------------5kyu

func Rot13(msg string) string {
	//5kyu - ROT13

	//97-122 for lower
	//65-90 for upper

	new_string := ""
	//fmt.Println("Word is: " + msg)
	for index := 0; index < len(msg); index++ {
		var rot_13_value int
		if int(msg[index]) <= 90 && int(msg[index]) >= 65 {
			//for lower range
			rot_13_value = int(msg[index]) + 13
			if rot_13_value > 90 {
				rot_13_value = rot_13_value - 90 + 64
			}
			new_string += string(rune(rot_13_value))
		} else if int(msg[index]) <= 122 && int(msg[index]) >= 97 {
			//for upper range
			rot_13_value = int(msg[index]) + 13
			if rot_13_value > 122 {
				rot_13_value = rot_13_value - 122 + 96
			}
			new_string += string(rune(rot_13_value))
		} else {
			new_string += string(msg[index])
		}
	}

	//fmt.Println(new_string)
	return new_string
}

func AlphabetWar(fight string) string {
	//5kyu - Alphabet war - Wo lo loooooo priests join the war

	var right_power int = 0
	var left_power int = 0
	var winner string = ""

	//loop through then check t/j? - add power straight away
	//loop, save the character from the previous loop iteration -> if the current changes the previous, deduct power and add accordingly
	var tag byte = '-'
	for letter := 0; letter < len(fight); letter++ {
		fmt.Print("Index ")
		fmt.Print(string(fight[letter]))
		fmt.Println(letter)
		if fight[letter] == 't' || fight[letter] == 'j' {
			tag = fight[letter]
			fmt.Println("Skipped!")
			continue
		}

		if letter < (len(fight) - 1) {
			if (fight[letter+1] == 't' && tag == 'j') || (fight[letter+1] == 'j' && tag == 't') {
				right_power += AlphabetWarRight(fight[letter])
				fmt.Print("0Right: ")
				fmt.Println(right_power)
				left_power += AlphabetWarLeft(fight[letter])
				fmt.Print("0Left: ")
				fmt.Println(left_power)
				continue
			} else if fight[letter+1] == 't' {
				left_power += AlphabetWarLeft(fight[letter])
				fmt.Print("1Left: ")
				fmt.Println(left_power)
				left_power += AlphabetWarRight(fight[letter])
				fmt.Print("Left: ")
				fmt.Println(left_power)
				continue
			} else if fight[letter+1] == 'j' {
				right_power += AlphabetWarRight(fight[letter])
				fmt.Print("1Right: ")
				fmt.Println(right_power)
				right_power += AlphabetWarLeft(fight[letter])
				fmt.Print("Right: ")
				fmt.Println(right_power)
				continue
			}
		}
		if tag == 'j' {
			right_power += AlphabetWarRight(fight[letter])
			fmt.Print("2Right: ")
			fmt.Println(right_power)
			right_power += AlphabetWarLeft(fight[letter])
			fmt.Print("Right: ")
			fmt.Println(right_power)
			tag = '-'
			continue
		} else if tag == 't' {
			left_power += AlphabetWarLeft(fight[letter])
			fmt.Print("2Left: ")
			fmt.Println(left_power)
			left_power += AlphabetWarRight(fight[letter])
			fmt.Print("Left: ")
			fmt.Println(left_power)
			tag = '-'
			continue
		} else {
			right_power += AlphabetWarRight(fight[letter])
			fmt.Print("3Right: ")
			fmt.Println(right_power)
			left_power += AlphabetWarLeft(fight[letter])
			fmt.Print("3Left: ")
			fmt.Println(left_power)
			continue
		}

	}

	if right_power > left_power {
		winner = "Right side wins!"
	} else if left_power > right_power {
		winner = "Left side wins!"
	} else {
		winner = "Let's fight again!"
	}
	return winner
}

func AlphabetWarLeft(letter byte) int {
	left := []byte{'w', 'p', 'b', 's'} //t
	power := []int{4, 3, 2, 1}
	for index := 0; index < len(left); index++ {
		if left[index] == letter {
			return power[index]
		}
	}
	return 0
}

func AlphabetWarRight(letter byte) int {
	right := []byte{'m', 'q', 'd', 'z'} //j
	power := []int{4, 3, 2, 1}
	for index := 0; index < len(right); index++ {
		if right[index] == letter {
			return power[index]
		}
	}
	return 0
}
