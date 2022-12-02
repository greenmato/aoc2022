package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	elfIndex := 0
	elfMap := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			if _, ok := elfMap[elfIndex]; ok {
				elfMap[elfIndex] = elfMap[elfIndex] + calories
			} else {
				elfMap[elfIndex] = calories
			}
		} else {
			elfIndex++
		}
	}

	calories := make([]int, len(elfMap))
	for _, value := range elfMap {
		calories = append(calories, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	fmt.Printf("1: %d\n", calories[0])
	fmt.Printf("2: %d\n", calories[1])
	fmt.Printf("3: %d\n", calories[2])
	fmt.Printf("Total: %d\n", calories[0]+calories[1]+calories[2])
}
