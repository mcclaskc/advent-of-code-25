package main

import (
	"advent-of-code-25/internal/utils"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Button []int
type Buttons []Button
type Joltages []int
type Matrix [][]float64

const epsilon = 1e-9

func solveGaussian(matrix Matrix) ([]float64, []int) {
	rows := len(matrix)
	cols := len(matrix[0]) - 1

	pivotCols := make([]int, 0)
	pivotRow := 0

	for col := 0; col < cols && pivotRow < rows; col++ {
		sel := -1
		for i := pivotRow; i < rows; i++ {
			// FIX 1: Use Epsilon for pivot selection to avoid noise
			if math.Abs(matrix[i][col]) > epsilon {
				sel = i
				break
			}
		}

		if sel == -1 {
			continue
		}

		matrix[pivotRow], matrix[sel] = matrix[sel], matrix[pivotRow]
		pivotCols = append(pivotCols, col)

		divisor := matrix[pivotRow][col]
		for j := col; j <= cols; j++ {
			matrix[pivotRow][j] /= divisor
		}

		for i := 0; i < rows; i++ {
			if i != pivotRow {
				factor := matrix[i][col]
				for j := col; j <= cols; j++ {
					matrix[i][j] -= factor * matrix[pivotRow][j]
				}
			}
		}

		pivotRow++
	}
	return nil, pivotCols
}

func countForRow(joltages Joltages, buttons Buttons) int {
	matrix := make([][]float64, len(joltages))
	for i, joltage := range joltages {
		row := make([]float64, len(buttons)+1)
		row[len(row)-1] = float64(joltage)
		for b, button := range buttons {
			if slices.Contains(button, i) {
				row[b] = 1
			} else {
				row[b] = 0
			}
		}
		matrix[i] = row
	}

	_, pivotCols := solveGaussian(matrix)

	// Consistency Check
	for i := len(pivotCols); i < len(matrix); i++ {
		if math.Abs(matrix[i][len(buttons)]) > epsilon {
			return 0
		}
	}

	lastColumn := []float64{}
	answerCol := len(buttons)
	for _, row := range matrix {
		lastColumn = append(lastColumn, row[answerCol])
	}

	if len(pivotCols) == len(buttons) {
		// EASY PATH
		total := 0
		for _, val := range lastColumn {
			rounded := math.Round(val)
			if math.Abs(val-rounded) > epsilon {
				return 0
			}
			if rounded < 0 {
				return 0
			}
			total += int(rounded)
		}
		return total
	} else {
		// HARD PATH
		isFree := make([]bool, len(buttons))
		for i := 0; i < len(buttons); i++ {
			isFree[i] = true
		}
		for _, p := range pivotCols {
			isFree[p] = false
		}

		var freeVarIndices []int
		for i, free := range isFree {
			if free {
				freeVarIndices = append(freeVarIndices, i)
			}
		}

		// Dynamic limit calculation
		maxPresses := 0
		for _, j := range joltages {
			if j > maxPresses {
				maxPresses = j
			}
		}
		maxPresses += 1

		minTotalPresses := math.MaxInt
		foundSolution := false

		var solve func(idx int, currentPresses []int)
		solve = func(idx int, currentPresses []int) {
			if idx == len(freeVarIndices) {
				valid := true

				for i, pivotCol := range pivotCols {
					row := matrix[i]
					constant := row[len(buttons)]
					val := constant
					for _, fIdx := range freeVarIndices {
						val -= row[fIdx] * float64(currentPresses[fIdx])
					}

					rounded := math.Round(val)

					// FIX 2: Check rounded < 0, NOT val < 0
					// This prevents -0.00000001 from failing the check
					if rounded < 0 || math.Abs(val-rounded) > epsilon {
						valid = false
						break
					}
					currentPresses[pivotCol] = int(rounded)
				}

				if valid {
					sum := 0
					for _, p := range currentPresses {
						sum += p
					}
					if sum < minTotalPresses {
						minTotalPresses = sum
						foundSolution = true
					}
				}
				return
			}

			fIdx := freeVarIndices[idx]
			for val := 0; val <= maxPresses; val++ {
				currentPresses[fIdx] = val
				solve(idx+1, currentPresses)
			}
		}

		solve(0, make([]int, len(buttons)))

		if foundSolution {
			return minTotalPresses
		}
		return 0
	}
}

// ... rest of your main/countAllRows code ...
func countAllRows(inputPath string) int {
	presses := 0
	err := utils.ProcessInput(inputPath, func(line string) error {
		parts := strings.Split(line, " ")
		joltagePart := parts[len(parts)-1]
		joltageInnards := joltagePart[1 : len(joltagePart)-1]
		joltageStrs := strings.Split(joltageInnards, ",")
		joltages := utils.Map(joltageStrs, func(_ int, s string) int {
			n, _ := strconv.Atoi(s)
			return n
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
		presses += countForRow(joltages, buttons)
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
