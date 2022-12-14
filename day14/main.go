package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	caveSizeX = 1000
	caveSizeY = 169
)

const (
	coordSeparator string = " -> "
	xySeparator    string = ","
)

type space string

const (
	empty space = "empty"
	rock  space = "rock"
	sand  space = "sand"
)

var renderMap = map[space]string{
	empty: " ",
	rock:  "#",
	sand:  "o",
}

type grid [][]space

type coord struct {
	x int
	y int
}

func (g *grid) build(lineStrs []string) {
	rocks := make([][]coord, 0)
	for _, lineStr := range lineStrs {
		coords := parseCoords(lineStr)
		rocks = append(rocks, coords)
	}

	for _, coords := range rocks {
		for i, co := range coords {
			if (i + 1) < len(coords) {
				startX := co.x
				startY := co.y
				finishX := coords[i+1].x
				finishY := coords[i+1].y

				xElems := getRange(startX, finishX)
				yElems := getRange(startY, finishY)

				for _, xe := range xElems {
					for _, ye := range yElems {
						(*g)[ye][xe] = rock
					}
				}
			}
		}
	}
}

func (g *grid) pourAllSand(until func(coord) bool) {
	for true {
		location := coord{500, 0}

		if isBlocked(*g, location) {
			break
		}

		for true {
			next := getNextLocation(*g, location)
			if next == location {
				break
			}

			location = next
		}

		if until(location) {
			break
		}

		(*g)[location.y][location.x] = sand
	}
}

func (g grid) countSand() int {
	count := 0
	for _, y := range g {
		for _, x := range y {
			if x == sand {
				count++
			}
		}
	}

	return count
}

func (g *grid) render() {
	for _, y := range *g {
		for _, x := range y {
			fmt.Print(renderMap[x])
		}
		fmt.Print("\n")
	}
}

func untilAbyss(location coord) bool {
	if location.x >= caveSizeX-1 || location.y >= caveSizeY-1 {
		return true
	}

	return false
}

func untilBlocked(_ coord) bool {
	return false
}

func getNextLocation(g grid, location coord) coord {
	if location.x+1 >= caveSizeX || location.y+1 >= caveSizeY {
		return location
	}
	if g[location.y+1][location.x] == empty {
		return coord{location.x, location.y + 1}
	}
	if g[location.y+1][location.x-1] == empty {
		return coord{location.x - 1, location.y + 1}
	}
	if g[location.y+1][location.x+1] == empty {
		return coord{location.x + 1, location.y + 1}
	}

	return location
}

func isBlocked(g grid, location coord) bool {
	if g[location.y][location.x] == sand {
		return true
	}

	return false
}

func makeGrid(sizeX int, sizeY int) grid {
	g := make(grid, sizeY)
	for xi, _ := range g {
		g[xi] = make([]space, sizeX)
		for yi, _ := range g[xi] {
			g[xi][yi] = empty
		}
	}

	return g
}

func parseCoords(coordsStr string) []coord {
	coordStrs := strings.Split(coordsStr, coordSeparator)

	coords := make([]coord, 0)
	for _, coordStr := range coordStrs {
		coords = append(coords, parseCoord(coordStr))
	}

	return coords
}

func parseCoord(coordStr string) coord {
	xyStrs := strings.Split(coordStr, xySeparator)
	if len(xyStrs) != 2 {
		panic(fmt.Sprintf("invalid coordinate: %v", coordStr))
	}

	x, errX := strconv.Atoi(xyStrs[0])
	y, errY := strconv.Atoi(xyStrs[1])
	if errX != nil || errY != nil {
		panic(fmt.Sprintf("invalid coordinate: %v", coordStr))
	}

	return coord{x, y}
}

func getRange(x, y int) []int {
	max, min := x, y
	if x < y {
		max = y
		min = x
	}

	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineStrs := make([]string, 0)
	for scanner.Scan() {
		lineStrs = append(lineStrs, scanner.Text())
	}

	caveP1 := makeGrid(caveSizeX, caveSizeY)
	caveP1.build(lineStrs)
	caveP1.pourAllSand(untilAbyss)
	//caveP1.render()
	countP1 := caveP1.countSand()

	caveP2 := makeGrid(caveSizeX, caveSizeY)
	caveP2.build(lineStrs)
	caveP2.pourAllSand(untilBlocked)
	//caveP2.render()
	countP2 := caveP2.countSand()

	fmt.Printf("Part 1: %d\n", countP1)
	fmt.Printf("Part 2: %d\n", countP2)
}
