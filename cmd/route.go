package cmd

import (
	"fmt"
	"github.com/boggydigital/clo"
	"github.com/boggydigital/glo/internal"
	"strconv"
)

func Route(req *clo.Request, defs *clo.Definitions) error {

	if req == nil {
		return clo.Route(req, defs)
	}

	switch req.Command {
	case "convert":
		for _, sg := range req.ArgValues("grams") {
			fg, err := strconv.ParseFloat(sg, 64)
			if err != nil {
				return fmt.Errorf("invalid grams value: %s", sg)
			}
			pounds, ounces := internal.Convert(fg)

			fmt.Printf("%vg is about %vlb", sg, pounds)
			if ounces > 0.01 {
				fmt.Printf(" %.2foz", ounces)
			}
			fmt.Println()
		}
		return nil
	default:
		return clo.Route(req, defs)
	}
}
