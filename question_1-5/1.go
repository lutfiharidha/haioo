package question_1_5

import (
	"encoding/json"
	"fmt"
)

var money []int = []int{
	100000,
	50000,
	20000,
	10000,
	5000,
	2000,
	1000,
	500,
	200,
	100,
}

func main() {
	input := 145000
	basket := make(map[int]int)
	for _, v := range money {
		if input < v {
			continue
		}
		b := input / v
		input = input - (v * b)
		basket[v] = b
		if input < 100 && input != 0 {
			basket[100] = basket[100] + 1
		}
	}

	bRes, _ := json.Marshal(basket)
	fmt.Println(string(bRes))
}
