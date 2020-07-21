package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	usage := "glo: grams (g) to pounds (lb) and ounces (oz) converter\n" +
		"usage: glo [grams]"

	if len(os.Args) < 2 {
		fmt.Println(usage)
	} else {
		grams, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Printf("\"%s\" is not a valid value in grams\n", os.Args[1])
			os.Exit(1)
		}

		pounds, ounces := Convert(grams)

		fmt.Printf("%vg is about %vlb", grams, pounds)
		if ounces > 0.01 {
			fmt.Printf(" %.2foz", ounces)
		}
		fmt.Println()
	}
}
