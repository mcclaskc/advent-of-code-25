package main

import (
	"advent-of-code-25/cmd/day12"
	"advent-of-code-25/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

//type State [][]int
//
//type StateCache map[]

func calculateAreaNeeded(packageAmounts []int, packages day12.Packages) int {
	return utils.Reduce(packageAmounts, func(i, acc, curr int) int {
		return acc + (curr * packages[i].Area)
	}, 0)
}

func countRow(hght, wdth int, packageAmounts []int, packages day12.Packages) int {
	totalArea := hght * wdth
	areaNeeded := calculateAreaNeeded(packageAmounts, packages)
	if totalArea < areaNeeded {
		return 0
	}
	//stateCache
	return 1
}

func countRegions(inputPath string) int {
	packages := day12.GetPackages()
	sum := 0
	utils.ProcessInput(inputPath, func(line string) error {
		//4x4: 0 0 0 0 2 0
		parts := strings.Split(line, ": ")
		dimensions := strings.Split(parts[0], "x")
		hght, _ := strconv.Atoi(dimensions[0])
		wdth, _ := strconv.Atoi(dimensions[1])
		packageAmounts := utils.Map(strings.Split(parts[1], " "), func(_ int, s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		sum += countRow(hght, wdth, packageAmounts, packages)
		return nil
	})
	return sum
}

func main() {
	fmt.Println(countRegions("../input.txt"))
}
