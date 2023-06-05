//go:build ignore

package main

import (
	"fmt"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
)

func main() {

	core := voicevoxcorego.NewVoicevoxCore()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")

	core.Initialize(initializeOptions)
	defer core.Finalize()
	core.LoadModel(1)

	// 子音(consonant)[0] -> 母音(vowel)[0] -> 子音(consonant)[1] -> 母音(vowel)[1] ... という順で読み解く。
	// 「テスト」という文章に対応する入力
	// ref: https://github.com/VOICEVOX/voicevox_core/blob/f32cafd1c18337abd6467de61944281eda54b73b/crates/voicevox_core/src/publish.rs#L898
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
