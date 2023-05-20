package voicevoxcorego_test

import (
	"os"
	"testing"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	status := m.Run()
	os.Exit(status)
}

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

	WAV_MAGICNUMBER_FIRST := []byte{0x52, 0x49, 0x46, 0x46}
	WAV_MAGICNUMBER_SECOND := []byte{0x57, 0x41, 0x56, 0x45}

	t.Log("assert magicnumber")
	assert.Equal(t, WAV_MAGICNUMBER_FIRST, result[:4])
	assert.Equal(t, WAV_MAGICNUMBER_SECOND, result[8:12])
}

// nolint:errcheck
// func TestSynthesis(t *testing.T) {
// 	t.Log("Run AudioQuery()")

// 	t.Log("initialize")
// 	core := setupCore()
// 	defer core.Finalize()
// 	t.Log("initialize done")

// 	t.Log("load model")
// 	core.LoadModel(1)
// 	t.Log("load model done")

// 	aqOptions := core.MakeDefaultAudioQueryOotions()

// 	query := core.AudioQuery("テストなのだね", 1, aqOptions)
// }

func setupCore() voicevoxcorego.VoicevoxCore {
	core := voicevoxcorego.NewVoicevoxCore()
	initOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	core.Initialize(initOptions) //nolint:errcheck

	return core
}
