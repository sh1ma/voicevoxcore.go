//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage:\n\tgo run audio_predict_duration.go [phoneme_id1, phoneme_id2, phoneme_id3, ...]")
		os.Exit(127)
	}
	rawPhonemeIDs := os.Args[1:]

	var phonemes []int64

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

		phonemes = append(phonemes, int64(phID))
	}

	core := voicevoxcorego.New()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	err := core.Initialize(initializeOptions)
	defer core.Finalize()
	if err != nil {
		fmt.Println(err)
	}
	err = core.LoadModel(1)
	if err != nil {
		fmt.Println(err)
	}

	retValue, _ := core.PredictDuration(1, phonemes)
	for _, v := range retValue {
		println(fmt.Sprintf("%f", v))
	}
}
