//
// Converts duration in format like 1.5h or 30m to hours value.
//
// For example:
//    1.5h -> 1.5
//    30m -> 0.5
//    100m -> 1.6666666666666667
//

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("0")
		return
	}

	durationString := os.Args[1]
	var isHours bool
	var isMinutes bool

	if strings.HasSuffix(durationString, "h") {
		isHours = true
	} else if strings.HasSuffix(durationString, "m") {
		isMinutes = true
	}

	var durationCut string = durationString
	if isHours || isMinutes {
		durationCut = durationString[:len(durationString)-1]
	}

	durationFloat, err := strconv.ParseFloat(durationCut, 10)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Println("0")
		return
	}

	if isMinutes {
		durationFloat /= 60.0
	}

	fmt.Printf("%v", durationFloat)
}
