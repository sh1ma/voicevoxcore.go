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
		fmt.Println("usage:\n\tgo run tts.go [ audioquery_json_path ]")
		os.Exit(127)
	}
	jsonPath := os.Args[1]

	file, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("ファイル読み込みに失敗")
		os.Exit(1)
	}

	audioQuery, err := voicevoxcorego.NewAudioQueryFromJson(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	core := voicevoxcorego.NewVoicevoxCore()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	err = core.Initialize(initializeOptions)
	defer core.Finalize()
	if err != nil {
		fmt.Println(err)
	}
	err = core.LoadModel(1)
	if err != nil {
		fmt.Println(err)
	}
	ttsOptions := voicevoxcorego.NewVoicevoxSynthesisOptions(false)
	result, err := core.Synthesis(audioQuery, 1, ttsOptions)
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create("out.wav")
	defer f.Close()
	_, err = f.Write(result)
	if err != nil {
		fmt.Println(err)
	}
}
