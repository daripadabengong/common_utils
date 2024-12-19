package main

import (
	"fmt"

	"github.com/daripadabengong/common_utils/domain"
)

func main() {
	test, err := domain.NewPhoneNumber("+628123456789")

	if err != nil {
		fmt.Println("error ", err.Error())
	} else {
		fmt.Println(test.GetValue())
	}

	fmt.Println("Hello world")
}
