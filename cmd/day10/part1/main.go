package main

import (
	"advent-of-code-25/internal/utils"
	"errors"
	"log"
	"slices"
	"strings"
	"unicode"
)

type Button []int

type Buttons []Button
type Lights []rune

type LightB = int
type ButtonB = int

const On = '#'
const Off = '.'

func parseLightsToBinary(lights Lights) LightB {
	result := 0
	for i := len(lights) - 1; i >= 0; i-- {
		result <<= 1
		if lights[i] == '#' {
			result |= 1
		}
	}
	return result
}

func parseButtonToBinary(button Button) ButtonB {
	result := 0
	for _, b := range button {
		result |= (1 << b)
	}
	return result
}

func parseBinaryToButton(buttonB ButtonB) Button {
	var buttons []int
	for i := 0; buttonB > 0; i++ {
		if buttonB&1 == 1 {
			buttons = append(buttons, i)
		}
		buttonB >>= 1
	}
	return buttons
}

func buttonsBtoButtons(buttonsB []ButtonB) []Button {
	return utils.Map(buttonsB, func(_ int, buttonB ButtonB) Button {
		return parseBinaryToButton(buttonB)
	})
}

func printPermutationAsButtons(buttonsB []ButtonB) {
	utils.Debug("BUTTONS", buttonsBtoButtons(buttonsB))
}

func parseButtonsToBinary(buttons Buttons) []ButtonB {
	var result []int
	for _, button := range buttons {
		result = append(result, parseButtonToBinary(button))
	}
	return result
}

func pressButton(lightsBinary LightB, buttonBinary ButtonB) LightB {
	return lightsBinary ^ buttonBinary
}

func pressAllButtonsInOrder(buttonsBinary []ButtonB) LightB {
	lightState := 0
	for _, button := range buttonsBinary {
		lightState = pressButton(lightState, button)
	}
	return lightState
}

func countForRow(lights Lights, buttons Buttons) (int, error) {
	desiredLights := parseLightsToBinary(lights)
	binaryButtons := parseButtonsToBinary(buttons)
	maxDepth := len(buttons)
	//fmt.Println(binaryButtons)
	currentPermutation := make([]int, len(binaryButtons))
	copy(currentPermutation, binaryButtons)

	var findMaxLengthPermutation func(int, int) bool
	findMaxLengthPermutation = func(currentDepth int, targetDepth int) bool {
		if currentDepth == targetDepth {
			if currentDepth == 3 {
				printPermutationAsButtons(currentPermutation[0:currentDepth])
			}
			// click all the buttons
			lightResult := pressAllButtonsInOrder(currentPermutation[0:currentDepth])
			return lightResult == desiredLights
		}
		for depth := currentDepth; depth < maxDepth; depth++ {
			currentPermutation[depth], currentPermutation[currentDepth] = currentPermutation[currentDepth], currentPermutation[depth]

			found := findMaxLengthPermutation(currentDepth+1, targetDepth)

			currentPermutation[depth], currentPermutation[currentDepth] = currentPermutation[currentDepth], currentPermutation[depth]
			if found {
				return true
			}
		}
		return false
	}
	for currentDepth := 1; currentDepth <= maxDepth; currentDepth++ {
		// go through all permutations, shortest first
		//fmt.Println("Current Depth:", currentDepth)
		if findMaxLengthPermutation(0, currentDepth) {
			return currentDepth, nil
		}
	}
	return 0, errors.New("no solution")
}

func countAllRows(inputPath string) int {
	presses := 0
	buttonBoundaries := []rune{'[', ']'}
	err := utils.ProcessInput(inputPath, func(line string) error {
		//[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
		parts := strings.Split(line, " ")
		lights := utils.Filter([]rune(parts[0]), func(_ int, r rune) bool {
			return !slices.Contains(buttonBoundaries, r)
		})
		buttonStrs := parts[1 : len(parts)-1]
		buttons := utils.Map(buttonStrs, func(_ int, str string) Button {
			return utils.Reduce[rune, Button]([]rune(str), func(_ int, btn Button, r rune) Button {
				if unicode.IsDigit(r) {
					return append(btn, int(r-'0'))
				}
				return btn
			}, Button{})
		})
		result, err := countForRow(lights, buttons)
		if err != nil {
			return errors.New("no solution for row: " + line)
		}
		presses += result
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return presses
}

func main() {
	println(countAllRows("../input.txt"))
}
