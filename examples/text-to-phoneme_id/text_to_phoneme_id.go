package main

import (
	"fmt"
	"os"
	"strings"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
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

func phonemeIndexOf(phoneme string) int {
	for i, v := range PHONEME {
		if phoneme == v {
			return i
		}
	}
	return -1
}

// nolint:errcheck
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage:\n\tgo run phonemeid.go [ text ]")
		os.Exit(127)
	}
	textSlice := args[1:]
	text := strings.Join(textSlice, " ")

	core := voicevoxcorego.New()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	core.Initialize(initializeOptions)
	defer core.Finalize()
	core.LoadModel(1) // ずんだもん

	// audioQueryを発行
	audioQueryOptions := voicevoxcorego.NewVoicevoxAudioQueryOptions(false)
	query, err := core.AudioQuery(text, 1, audioQueryOptions)
	if err != nil {
		fmt.Println(err)
	}
	for _, phrase := range query.AccentPharases {
		for _, mora := range phrase.Moras {
			if mora.Consonant != "" {
				fmt.Printf("%d ", phonemeIndexOf(mora.Consonant))
			}
			fmt.Printf("%d ", phonemeIndexOf(mora.Vowel))
		}
	}
}
