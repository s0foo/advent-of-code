package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func day_01_1() {
	var larger_measurements int
	var lines []string

	file, err := os.Open("data/day_01.txt")
	if err != nil {
		log.Fatalf("Failed to open data/day_01.txt")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	for i := 1; i < len(lines); i++ {
		previous_depth, _ := strconv.Atoi(lines[i-1])
		current_depth, _ := strconv.Atoi(lines[i])
		if current_depth > previous_depth {
			larger_measurements += 1
		}
	}
	log.Println(larger_measurements)
}

func day_01_2() {
	var larger_sums int
	var lines []string

	file, err := os.Open("data/day_01.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	previous_sum_0, _ := strconv.Atoi(lines[0])
	previous_sum_1, _ := strconv.Atoi(lines[1])
	previous_sum_2, _ := strconv.Atoi(lines[2])
	previous_sum := previous_sum_0 + previous_sum_1 + previous_sum_2

	for i := 3; i < len(lines); i++ {
		expired_depth, _ := strconv.Atoi(lines[i-3])
		current_depth, _ := strconv.Atoi(lines[i])
		current_sum := previous_sum - expired_depth + current_depth
		if current_sum > previous_sum {
			larger_sums += 1
		}
		current_sum = previous_sum
	}
	log.Println(larger_sums)
}

func day_02_1() {
	var x_pos int
	var depth int
	var lines []string

	file, err := os.Open("data/day_02.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	for i := 0; i < len(lines); i++ {
		command := strings.Split(lines[i], " ")
		order := command[0]
		value, _ := strconv.Atoi(command[1])
		if order == "forward" {
			x_pos += value
		} else if order == "down" {
			depth += value
		} else if order == "up" {
			depth -= value
		}
	}

	log.Println(x_pos * depth)
}

func day_02_2() {
	var x_pos int
	var depth int
	var aim int
	var lines []string

	file, err := os.Open("data/day_02.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	for i := 0; i < len(lines); i++ {
		command := strings.Split(lines[i], " ")
		order := command[0]
		value, _ := strconv.Atoi(command[1])
		if order == "forward" {
			x_pos += value
			depth += aim * value
		} else if order == "down" {
			aim += value
		} else if order == "up" {
			aim -= value
		}
	}

	log.Println(x_pos * depth)
}

func day_03_1() {
	var cum [12]int
	var val int
	var lines []string
	var gamma_rate string
	var epsilon_rate string

	file, err := os.Open("data/day_03.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	ln_num := len(lines)

	for i := 0; i < ln_num; i++ {
		command := strings.Split(lines[i], "")
		for j := 0; j < 12; j++ {
			val, _ = strconv.Atoi(command[j])
			if val == 1 {
				cum[j] += 1
			}
		}
	}

	for i := 0; i < 12; i++ {
		if cum[i] > ln_num/2 {
			gamma_rate += "1"
			epsilon_rate += "0"
		} else {
			gamma_rate += "0"
			epsilon_rate += "1"
		}
	}

	int_gamma, _ := strconv.ParseInt(gamma_rate, 2, 64)
	int_epsilon, _ := strconv.ParseInt(epsilon_rate, 2, 64)

	log.Println(int_gamma * int_epsilon)
}

func main() {
	day_01_1()
	day_01_2()
	day_02_1()
	day_02_2()
	day_03_1()
}
