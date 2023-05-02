//go:build ignore

package main

import (
	"fmt"
	"os"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage:\n\tgo run audio_query.go [ text ]")
		os.Exit(127)
	}
	text := os.Args[1]

	core := voicevoxcorego.NewVoicevoxCore()
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
	audioQueryOptions := voicevoxcorego.NewVoicevoxAudioQueryOptions(false)
	result := core.AudioQuery(text, 1, audioQueryOptions)
	fmt.Println(result)
}
