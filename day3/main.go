package main

import (
	"bufio"
	"fmt"
	"os"
)

type groupStatus struct {
	group1HasBadge bool
	group2HasBadge bool
	group3HasBadge bool
}

func halve(str string) (string, string) {
	if len(str)%2 != 0 {
		panic("length of string is odd")
	}

	compartment1 := str[0:(len(str) / 2)]
	compartment2 := str[(len(str) / 2):(len(str))]

	return compartment1, compartment2
}

func getScore(item rune) int {
	if int(item) >= 65 && int(item) <= 90 {
		return int(item) - 38
	}
	if int(item) >= 97 && int(item) <= 122 {
		return int(item) - 96
	}

	panic("invalid item")
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	incorrect := make([]rune, 0)
	priorityTotal := 0

	for scanner.Scan() {
		line := scanner.Text()

		compartment1, compartment2 := halve(line)
		intersect := make(map[rune]int, 0)

		for _, item := range compartment1 {
			intersect[item] = 1
		}
		for _, item := range compartment2 {
			if _, ok := intersect[item]; ok {
				incorrect = append(incorrect, item)
				priorityTotal = priorityTotal + getScore(item)
				break
			}
		}
	}

	fmt.Printf("Total priorities: %d\n", priorityTotal)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	groups := make(map[int][]string, 0)
	badges := make([]rune, 0)
	priorityTotal := 0

	groupNo := 0
	for scanner.Scan() {
		line := scanner.Text()

		if _, ok := groups[groupNo]; ok {
			groups[groupNo] = append(groups[groupNo], line)
		} else {
			groups[groupNo] = []string{line}
		}

		if len(groups[groupNo]) == 3 {
			intersect := make(map[rune]groupStatus, 0)
			for _, item := range groups[groupNo][0] {
				intersect[item] = groupStatus{group1HasBadge: true}
			}
			for _, item := range groups[groupNo][1] {
				if status, ok := intersect[item]; ok {
					status.group2HasBadge = true
					intersect[item] = status
				} else {
					intersect[item] = groupStatus{group2HasBadge: true}
				}
			}
			for _, item := range groups[groupNo][2] {
				if status, ok := intersect[item]; ok {
					if status.group1HasBadge && status.group2HasBadge {
						badges = append(badges, item)
						priorityTotal = priorityTotal + getScore(item)
						break
					}
				}
			}

			groupNo++
		}
	}

	fmt.Printf("Badges priority total: %d\n", priorityTotal)
}

func main() {
	part1()
	part2()
}
