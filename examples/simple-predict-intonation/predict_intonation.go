//go:build ignore

package main

import (
	"fmt"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
)

func main() {
	// args := os.Args
	// if len(args) < 2 {
	// 	fmt.Println("usage:\n\tgo run audio_query.go")
	// 	os.Exit(127)
	// }
	// text := os.Args[1]

	core := voicevoxcorego.NewVoicevoxCore()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")

	core.Initialize(initializeOptions)
	defer core.Finalize()
	core.LoadModel(1)

	vowelPhonemeVector := []int64{0, 14, 6, 30, 0}
	consonantPhonemeVector := []int64{-1, 37, 35, 37, -1}
	startAccentVector := []int64{0, 1, 0, 0, 0}
	endAccentVector := []int64{0, 1, 0, 0, 0}
	startAccentPhraseVector := []int64{0, 1, 0, 0, 0}
	endAccentPhraseVector := []int64{0, 0, 0, 1, 0}

	retValue, err := core.PredictIntonation(
		1,
		vowelPhonemeVector, consonantPhonemeVector,
		startAccentVector, endAccentVector,
		startAccentPhraseVector, endAccentPhraseVector,
	)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range retValue {
		println(fmt.Sprintf("%f", v))
	}
}
