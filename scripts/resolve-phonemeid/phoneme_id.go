package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var PHONEME []string = []string{
	"pau",
	"A",
	"E",
	"I",
	"N",
	"O",
	"U",
	"a",
	"b",
	"by",
	"ch",
	"cl",
	"d",
	"dy",
	"e",
	"f",
	"g",
	"gw",
	"gy",
	"h",
	"hy",
	"i",
	"j",
	"k",
	"kw",
	"ky",
	"m",
	"my",
	"n",
	"ny",
	"o",
	"p",
	"py",
	"r",
	"ry",
	"s",
	"sh",
	"t",
	"ts",
	"ty",
	"u",
	"v",
	"w",
	"y",
	"z",
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage:\n\tgo run phonemeid.go [ phoneme_id1, phoneme_id2, phoneme_id3, ... ]")
		os.Exit(127)
	}

	rawPhonemeIDs := args[1:]

	var resolved []string

	for _, v := range rawPhonemeIDs {
		phID, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("phoneme_id must be integer")
			return
		}

		if phID > 44 || phID < 0 {
			fmt.Println("phoneme_id range must be 0 - 44")
			return
		}

		resolved = append(resolved, PHONEME[phID])
	}

	fmt.Println(strings.Join(resolved, " "))
}
