package day7

import (
	"strings"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

type Equation struct {
	Result int
	Values []int
}

// Parse each line into the equations struct
func ParseEquationData(line string) *Equation {
	lineSplit := strings.Split(line, ":")

	//Convert Values to ints and add them to array
	numbersSplit := strings.Split(lineSplit[1], " ")
	var values []int
	for _, number := range numbersSplit {
		if number == "" {
			continue
		}
		values = append(values, util.MustAtoI(number))
	}
	return &Equation{util.MustAtoI(lineSplit[0]), values}
}

// Evaluates a equation in a written form
func EvaluateEquation(values []int, operations []rune) int {
	if len(values)-1 != len(operations) || len(values) == 0 || len(operations) == 0 {
		// There should be a operation between each values
		return 0
	}

	currentSum := values[0]

	// Get the next values and perform the operation the currentSum
	for i, operation := range operations {
		switch operation {
		case '*':
			currentSum *= values[i+1]
		default:
			currentSum += values[i+1]
		}
	}
	return currentSum
}

// Generates a permuation of the operations variable at a specific length
func GenerateOperationPermutations(operations []rune, length int) [][]rune {
	var results [][]rune
	var backtrack func(current []rune)

	backtrack = func(current []rune) {
		if len(current) == length {
			// Once we have the deciered length, copy it accross and exit the function
			temp := make([]rune, len(current))
			copy(temp, current)
			results = append(results, temp)
			return
		}
		for _, op := range operations {
			backtrack(append(current, op))
		}
	}
	backtrack([]rune{})
	return results
}

// Generates a permuation of the operation to check the value of the equation can be made
func HasOperationsInEquation(equation *Equation) bool {
	operationsLength := len(equation.Values) - 1
	operationPermutations := GenerateOperationPermutations([]rune{'*', '+'}, operationsLength)

	for _, operationPermutation := range operationPermutations {
		evaluatedEqation := EvaluateEquation(equation.Values, operationPermutation)

		if evaluatedEqation == equation.Result {
			return true
		}
	}
	return false
}

func Day7Challenge1(input []string) int {
	equations := util.ArrayMap(input, ParseEquationData)
	validEquations := util.ArrayMap(equations, HasOperationsInEquation)

	//Sum up where the the equations are valid
	sum := 0
	for i := range equations {
		if validEquations[i] {
			sum += equations[i].Result
		}
	}
	return sum
}
