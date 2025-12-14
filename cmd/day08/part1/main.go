package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	a int
	b int
	d float64
}
type Pairs []Pair

type Box struct {
	x float64
	y float64
	z float64
	c int
}
type Boxes []Box

const BoxUnassignedC int = -1

func newBox(x, y, z float64) Box {
	return Box{x, y, z, BoxUnassignedC}
}

type BoxId int

func calcDistance(ax, ay, az, bx, by, bz float64) float64 {
	dx := ax - bx
	dy := ay - by
	dz := az - bz

	dSquaredSums := dx*dx + dy*dy + dz*dz

	return math.Sqrt(float64(dSquaredSums))
}

func junctionBoxes(filePath string, numPairs int) int {
	boxes := Boxes{}
	pairs := Pairs{}

	utils.ProcessInput(filePath, func(line string) error {
		numStrs := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(numStrs[0], 64)
		y, _ := strconv.ParseFloat(numStrs[1], 64)
		z, _ := strconv.ParseFloat(numStrs[2], 64)
		boxes = append(boxes, newBox(x, y, z))
		return nil
	})

	utils.Debug("Boxes", boxes)

	for i := 0; i+1 < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			ax, ay, az := boxes[i].x, boxes[i].y, boxes[i].z
			bx, by, bz := boxes[j].x, boxes[j].y, boxes[j].z
			d := calcDistance(ax, ay, az, bx, by, bz)
			pairs = append(pairs, Pair{i, j, d})
		}
	}

	utils.Debug("Pairs", pairs)

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].d < pairs[j].d
	})

	utils.Debug("Sorted Pairs", pairs)
	pairs = pairs[:numPairs]

	utils.Debug("First "+strconv.Itoa(numPairs)+" Pairs", pairs)

	updateAll := func(oldC, newC int) {
		for i := range boxes {
			if boxes[i].c == oldC {
				boxes[i].c = newC
			}
		}
	}

	for i, pair := range pairs {
		a, b := pair.a, pair.b
		if boxes[a].c == BoxUnassignedC && boxes[b].c == BoxUnassignedC {
			// if both are unassigned, add them to the current circuit
			boxes[a].c = i
			boxes[b].c = i
		} else if boxes[a].c != BoxUnassignedC && boxes[b].c == BoxUnassignedC {
			// if boxA is assigned and boxB is not, then add boxB to boxA's circuit
			boxes[b].c = boxes[a].c
		} else if boxes[a].c == BoxUnassignedC && boxes[b].c != BoxUnassignedC {
			// if boxB is assigned and boxA is not, then add boxA to boxB's circuit
			boxes[a].c = boxes[b].c
		} else if boxes[a].c != boxes[b].c {
			// if both are assigned and different, choose the smaller, and update all in the larger to the smaller
			if boxes[a].c < boxes[b].c {
				updateAll(boxes[b].c, boxes[a].c)
				boxes[b].c = boxes[a].c
			} else {
				updateAll(boxes[a].c, boxes[b].c)
				boxes[a].c = boxes[b].c
			}
		}
		utils.Debug("Updated Boxes for circuit: "+strconv.Itoa(i), boxes)
		// if both are the same, do nothing
	}

	circSizeById := make(map[int]int)
	for i, box := range boxes {
		if box.c == BoxUnassignedC {
			circSizeById[i*BoxUnassignedC]++
		} else {
			circSizeById[box.c]++
		}
		fmt.Println(circSizeById)
	}

	utils.Debug("allCircuitSizes", circSizeById)

	allSizes := []int{}
	for _, cSize := range circSizeById {
		println(cSize)
		allSizes = append(allSizes, cSize)
	}
	utils.Debug("allSizes", allSizes)

	sort.Slice(allSizes, func(i, j int) bool {
		return allSizes[i] > allSizes[j]
	})

	utils.Debug("sorted allSizes", allSizes)

	return allSizes[0] * allSizes[1] * allSizes[2]
}

func main() {
	result := junctionBoxes("../input.txt", 1000)
	fmt.Print("RESULT ---\n")
	fmt.Printf("RESULT: %d\n", result)
	fmt.Print("RESULT ---\n")
}
