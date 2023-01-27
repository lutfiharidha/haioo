package question_1_5

import "fmt"

func main() {
	input := []string{
		"telkom",
		"telbom",
	}
	if len(input[0]) != len(input[1]) {
		fmt.Println(false)
		return
	}
	var changes int
	for k, v := range input[0] {
		if string(input[1][k]) != string(v) {
			changes++
		}
	}

	if changes > 1 {
		fmt.Println(false)
		return
	}

	fmt.Println(true)
}
