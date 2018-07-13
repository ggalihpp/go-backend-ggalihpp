package example

import (
	"fmt"
	"strconv"
	"strings"
)

// This layer wil handle the business proccess/any logic that happen in your app

func logicExample(input string) (int, error) {
	toNum, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return toNum + 42, nil
}

func Solution(N int) int {
	// write your code in Go 1.4

	binary := strconv.FormatInt(int64(N), 2)

	fmt.Println("BINER: ", binary)
	cek := CheckGap(binary)

	return cek

}

func CheckGap(binary string) (output int) {

	max := 0

	bs := strings.Split(binary, "")

	for _, v := range bs {

		if v == "1" {
			if max > output {
				output = max
			}
			max = 0
		} else {
			max++
		}

	}

	return
}
