package utils

import (
	"fmt"

	col "github.com/sulicat/goboi/colors"
)

func PanicOnErr(e error) {
	if e != nil {
		fmt.Printf(col.BgBrightRed+"ERROR"+col.Reset+" %s\n", e)
		panic(e)
	}
}
