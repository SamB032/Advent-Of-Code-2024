package day7

import (
	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

// Evaluates a equation in a written form
func EvaluateEquationExtra(values []int, operations []rune) int {
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
		case '|':
			joinedString := util.IntToString(currentSum) + util.IntToString(values[i+1])
			currentSum = util.MustAtoI(joinedString)
		default:
			currentSum += values[i+1]
		}
	}
	return currentSum
}

// Generates a permuation of the operation to check the value of the equation can be made
func HasOperationsInEquationExtra(equation *Equation) bool {
	operationsLength := len(equation.Values) - 1
	operationPermutations := GenerateOperationPermutations([]rune{'*', '+', '|'}, operationsLength)

	for _, operationPermutation := range operationPermutations {
		evaluatedEqation := EvaluateEquationExtra(equation.Values, operationPermutation)

		if evaluatedEqation == equation.Result {
			return true
		}
	}
	return false
}

func Day7Challenge2(input []string) int {
	equations := util.ArrayMap(input, ParseEquationData)
	validEquations := util.ArrayMap(equations, HasOperationsInEquationExtra)

	//Sum up where the the equations are valid
	sum := 0
	for i := range equations {
		if validEquations[i] {
			sum += equations[i].Result
		}
	}
	return sum
}
