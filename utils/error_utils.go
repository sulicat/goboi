package utils

import (
	"fmt"
	"os"
	"strings"

	col "github.com/sulicat/goboi/colors"
)

func PanicOnErr(e error) {
	if e != nil {
		fmt.Printf(col.BgBrightRed+"ERROR"+col.Reset+" %s\n", e)
		panic(e)
	}
}

func ExecutableDir() string {
	path, _ := os.Executable()
	li := strings.LastIndex(path, "/")
	path = path[:li]
	return path + "/"
}
