package voicevoxcorego_test

import (
	"os"
	"strings"
	"testing"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	status := m.Run()
	os.Exit(status)
}

func TestLoadModelAndIsModelLoaded(t *testing.T) {
	t.Log("initialize")
	core := setupCore()
	defer core.Finalize()
	t.Log("initialize done")

	t.Log("assert return false when model is not loaded")
	assert.Equal(t, core.IsModelLoaded(1), false)
	if err := core.LoadModel(1); err != nil {
		t.Fatal(err)
		return
	}
	assert.Equal(t, core.IsModelLoaded(1), true)

	t.Log("assert error when model id is invalid")
	assert.Equal(t, core.IsModelLoaded(9999), false)
	err := core.LoadModel(9999)
	if assert.Error(t, err) {
		assert.NotEqual(t, strings.Contains(err.Error(), "無効なspeaker_idです"), false)
		return
	}
	t.Fatal("error is not occurred")
}

// Ttsの実行を確認するテスト
// nolint:errcheck
func TestTts(t *testing.T) {
	t.Log("initialize")
	core := setupCore()
	defer core.Finalize()
	t.Log("initialize done")

	t.Log("load model")
	core.LoadModel(1)
	t.Log("load model done")

	ttsOptions := core.MakeDefaultTtsOotions()

	t.Log("Test Tts()")

	result, err := core.Tts("テストなのだね", 1, ttsOptions)
	if err != nil {
		t.Fatal(err)
	}

	isWavFile(t, result)
}

// オーディオクエリを発行し、音声合成を行うテスト
// nolint:errcheck
func TestSynthesis(t *testing.T) {
	t.Log("initialize")
	core := setupCore()
	defer core.Finalize()
	t.Log("initialize done")

	t.Log("load model")
	core.LoadModel(1)
	t.Log("load model done")

	// AudioQueryを生成
	aqOptions := core.MakeDefaultAudioQueryOotions()
	query, _ := core.AudioQuery("テストなのだね", 1, aqOptions)

	// 音声合成する
	synOptions := core.MakeDefaultSynthesisOotions()
	result, err := core.Synthesis(query, 1, synOptions)
	if err != nil {
		t.Fatal(err)
	}
	isWavFile(t, result)
}

// nolint:errcheck
func TestPredictDuration(t *testing.T) {
	t.Log("initialize")
	core := setupCore()
	defer core.Finalize()
	t.Log("initialize done")

	t.Log("load model")
	core.LoadModel(1)
	t.Log("load model done")

	// AudioQueryを生成
	aqOptions := core.MakeDefaultAudioQueryOotions()
	query, _ := core.AudioQuery("テストなのだね", 1, aqOptions)
	accentPhrases := query.AccentPharases
	var phonemes []int64
	for _, ap := range accentPhrases {
		moras := ap.Moras
		for _, m := range moras {
			if m.Consonant != "" {
				phonemes = append(phonemes, int64(phonemeIndexOf(m.Consonant)))
			}
			phonemes = append(phonemes, int64(phonemeIndexOf(m.Vowel)))
		}
	}

	duration, _ := core.PredictDuration(1, phonemes)
	assert.Equal(t, len(phonemes), len(duration))

}

func setupCore() voicevoxcorego.VoicevoxCore {
	core := voicevoxcorego.NewVoicevoxCore()
	initOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	core.Initialize(initOptions) //nolint:errcheck

	return core
}

func isWavFile(t *testing.T, bin []byte) {
	t.Helper()

	WAV_MAGICNUMBER_FIRST := []byte{0x52, 0x49, 0x46, 0x46}
	WAV_MAGICNUMBER_SECOND := []byte{0x57, 0x41, 0x56, 0x45}

	t.Log("assert MagicNumber")
	assert.Equal(t, WAV_MAGICNUMBER_FIRST, bin[:4])
	assert.Equal(t, WAV_MAGICNUMBER_SECOND, bin[8:12])
}

func phonemeIndexOf(phoneme string) int {
	for i, v := range PHONEME {
		if phoneme == v {
			return i
		}
	}
	return -1
}

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
