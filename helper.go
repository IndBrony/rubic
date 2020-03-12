package rubic

import "math"

func ReverseSlice(input []string) (out []string) {
	out = make([]string, len(input), len(input))
	for index, element := range input {
		reversedIndex := int(math.Abs(float64(index - (len(input) - 1))))
		out[reversedIndex] = element
	}
	return
}
