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
		fmt.Println("usage:\n\tgo run tts.go [ text ]")
		os.Exit(127)
	}
	text := os.Args[1]

	core := voicevoxcorego.NewVoicevoxCore()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	err := core.Initialize(initializeOptions)
	if err != nil {
		fmt.Println(err)
	}
	err = core.LoadModel(1)
	if err != nil {
		fmt.Println(err)
	}
	ttsOptions := voicevoxcorego.NewVoicevoxTtsOptions(false, true)
	result, err := core.Tts(text, 1, ttsOptions)
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create("out.wav")
	_, err = f.Write(result)
	if err != nil {
		fmt.Println(err)
	}
}
