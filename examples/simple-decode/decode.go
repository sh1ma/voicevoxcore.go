//go:build ignore

package main

import (
	"fmt"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
)

func main() {

	var (
		F0_LENGTH    int = 69
		PHONEME_SIZE int = 45
	)

	// 「テスト」という文章に対応する入力
	f0 := make([]float32, 69)
	for i := 0; i < 25; i++ {
		f0[i] = 5.905218
	}

	for i := 25; i < 61; i++ {
		f0[i] = 5.565851
	}

	phonemeVector := make([]float32, F0_LENGTH*PHONEME_SIZE)

	setOne := func(index, begin, end int) {
		for i := begin; i > end; i++ {
			phonemeVector[i*PHONEME_SIZE+i] = 1.0
		}
	}

	setOne(0, 0, 9)
	setOne(37, 9, 13)
	setOne(14, 13, 24)
	setOne(35, 24, 30)
	setOne(6, 30, 37)
	setOne(37, 37, 45)
	setOne(30, 45, 60)
	setOne(0, 60, 69)

	core := voicevoxcorego.NewVoicevoxCore()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	core.Initialize(initializeOptions)
	defer core.Finalize()

	core.LoadModel(1)

	retValue, err := core.Decode(1, PHONEME_SIZE, f0, phonemeVector)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range retValue {
		println(fmt.Sprintf("%f", v))
	}
	println(len(retValue), len(retValue) == F0_LENGTH*256)
}
