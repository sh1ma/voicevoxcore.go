package voicevoxcorego

// #include <stdint.h>
import "C"
import (
	"errors"
	"unsafe"
)

// VoicevoxCore is top-level API Wrapper instance
type VoicevoxCore struct {
	*RawVoicevoxCore
	initialized bool
}

// VoicevoxCore のコンストラクタ関数
func NewVoicevoxCore() (core VoicevoxCore) {
	core = VoicevoxCore{}
	return
}

// C APIを通じてVoicevox_coreを初期化する関数
func (r *VoicevoxCore) Initialize(options VoicevoxInitializeOptions) (err error) {
	if r.initialized {
		err = errors.New("Already initialized")
		return
	}
	r.voicevoxInitialize(*options.Raw)
	r.initialized = true
	return
}

// 音声合成モデルをロードする関数
func (r *VoicevoxCore) LoadModel(speakerID int) (err error) {
	id := C.uint(speakerID)
	code := r.voicevoxLoadModel(id)
	if int(code) != 0 {
		err = errors.New("Model load Failed")
	}
	return
}

// Text to Speechを実行する関数。実行結果はwavファイルフォーマットのバイト列。
// Sample: https://github.com/sh1ma/sample-tts
func (r *VoicevoxCore) Tts(text string, speakerID int, options VoicevoxTtsOptions) (slicebytes []byte, err error) {
	ctext := C.CString(text)
	cspeakerID := C.uint(speakerID)

	// TODO: このあたりを関数にまとめる
	var size int
	data := make([]*C.uchar, 1)
	// NOTE: ここキャストする必要ない
	datap := (**C.uchar)(&data[0])
	defer r.voicevoxWavFree(*datap)

	sizep := (*C.ulong)(unsafe.Pointer(&size))

	code := r.voicevoxTts(ctext, cspeakerID, *options.Raw, sizep, datap)
	if int(code) != 0 {
		err = errors.New("Failed TTS")
		return
	}

	slice := unsafe.Slice(data[0], *sizep)
	sliceUnsafe := unsafe.SliceData(slice)
	slicebytes = C.GoBytes(unsafe.Pointer(sliceUnsafe), C.int((len(slice))))
	return
}

// Audio Queryを基に音声合成を実行する関数。実行結果はwavファイルフォーマットのバイト列。
// Sample: https://github.com/sh1ma/sample-synthesis
func (r *VoicevoxCore) Synthesis(
	audioQuery string,
	speakerID int,
	options VoicevoxSynthesisOptions,
) (slicebytes []byte, err error) {
	ctext := C.CString(audioQuery)
	cspeakerID := C.uint(speakerID)

	// HACK: 煩雑なコード。Tts()のTODOと一緒の関数にまとめる
	var size int
	data := make([]*C.uchar, 1)
	// NOTE: ここキャストする必要ない
	datap := (**C.uchar)(&data[0])
	defer r.voicevoxWavFree(*datap)

	sizep := (*C.ulong)(unsafe.Pointer(&size))

	code := r.voicevoxSynthesis(ctext, cspeakerID, *options.Raw, sizep, datap)
	if int(code) != 0 {
		err = errors.New("Failed TTS")
		return
	}

	slice := unsafe.Slice(data[0], *sizep)
	sliceUnsafe := unsafe.SliceData(slice)
	slicebytes = C.GoBytes(unsafe.Pointer(sliceUnsafe), C.int((len(slice))))
	return
}

func (r *VoicevoxCore) MakeDefaultInitializeOptions() VoicevoxInitializeOptions {
	raw := r.voicevoxMakeDefaultInitializeOptions()
	return VoicevoxInitializeOptions{Raw: &raw}
}

func (r *VoicevoxCore) MakeDefaultTtsOotions() VoicevoxTtsOptions {
	raw := r.voicevoxMakeDefaultTtsOptions()
	return VoicevoxTtsOptions{Raw: &raw}
}

func (r *VoicevoxCore) MakeDefaultAudioQueryOotions() VoicevoxAudioQueryOptions {
	raw := r.voicevoxMakeDefaultAudioQueryOptions()
	return VoicevoxAudioQueryOptions{Raw: &raw}
}

func (r *VoicevoxCore) MakeDefaultSynthesisOotions() VoicevoxSynthesisOptions {
	raw := r.voicevoxMakeDefaultSynthesisOptions()
	return VoicevoxSynthesisOptions{Raw: &raw}
}

func (r *VoicevoxCore) MakeDefaultAccentPhrasesOptions() {

}

func (r *VoicevoxCore) AudioQuery(text string, speakerID uint, options VoicevoxAudioQueryOptions) string {
	ctext := C.CString(text)
	cSpeakerID := C.uint(speakerID)

	data := make([]*C.char, 1)
	datap := &data[0]
	defer r.voicevoxAudioQueryJsonFree(*datap)

	r.voicevoxAudioQuery(ctext, cSpeakerID, *options.Raw, datap)

	audioQueryJson := C.GoString(*datap)

	return audioQueryJson
}

func (r *VoicevoxCore) Finalize() {
	r.voicevoxFinalize()
}

func (r *VoicevoxCore) ErrorResultToMessage(resultCode int) string {
	cResultCode := C.int(resultCode)
	retValue := r.voicevoxErrorResultToMessage(cResultCode)

	return C.GoString(retValue)
}

func (r *VoicevoxCore) GetMetasJson() string {
	cResult := r.voicevoxGetMetasJson()

	return C.GoString(cResult)
}

func (r *VoicevoxCore) GetSupportedDevicesJson() string {
	cResult := r.voicevoxGetSupportedDevicesJson()

	return C.GoString(cResult)
}

func (r *VoicevoxCore) GetCoreVersion() string {
	cResult := r.voicevoxGetVersion()

	return C.GoString(cResult)
}

func (r *VoicevoxCore) IsGpuMode() bool {
	cResult := r.voicevoxIsGpuMode()

	return bool(cResult)
}

func (r *VoicevoxCore) IsModelLoaded(speakerID uint) bool {
	cSpeakerID := C.uint(speakerID)
	cResult := r.voicevoxIsModelLoaded(cSpeakerID)

	return bool(cResult)
}

func (r *VoicevoxCore) PredictDuration(speakerID int, phonemeVector []int64) []float32 {

	length := len(phonemeVector)

	var size uint32
	data := make([]*C.float, 1)
	datap := &data[0]
	sizep := (*C.ulong)(unsafe.Pointer(&size))
	cPhonemeVectoTmp := &phonemeVector[0]
	cPhonemeVectorPtr := (*C.int64_t)(cPhonemeVectoTmp)
	phonemeVectorLength := (C.ulong)(length)
	r.voicevoxPredictDuration(phonemeVectorLength, cPhonemeVectorPtr, C.uint(speakerID), sizep, datap)
	defer r.voicevoxPredictDurationDataFree(*datap)

	slice := unsafe.Slice(data[0], *sizep)

	retValue := make([]float32, length)

	for i, v := range slice {
		retValue[i] = float32(v)
	}

	return retValue
}

func (r *VoicevoxCore) PredictIntonation(
	speakerID int,
	vowelPhonemeVector, consonantPhonemeVector []int64,
	startAccentVector, endAccentVector []int64,
	startAccentPhraseVector, endAccentPhraseVector []int64,
) ([]float32, error) {

	// 全てのlengthが同数かのチェック
	length := len(vowelPhonemeVector)
	otherLengths := []int{
		len(consonantPhonemeVector),
		len(startAccentPhraseVector), len(endAccentVector),
		len(startAccentPhraseVector), len(endAccentPhraseVector),
	}
	for _, l := range otherLengths {
		if length != l {
			return nil, errors.New("全てのベクター変数は同じ長さでなければなりません。")
		}
	}

	cLength := C.ulong(length)
	cSpeakerID := C.uint(speakerID)

	// []int64を[]C.int64に変換する関数
	int64ToCtype := sliceToCtype[int64, C.int64_t]

	// それぞれのVectorに `int64ToCtype` を適用する
	cVowelPhonemeVector := int64ToCtype(vowelPhonemeVector)
	cConsonantPhonemeVecor := int64ToCtype(consonantPhonemeVector)
	cStartAccentVector := int64ToCtype(startAccentPhraseVector)
	cEndAccentVector := int64ToCtype(endAccentVector)
	cStartAccentPhraseVector := int64ToCtype(startAccentPhraseVector)
	cEndAccentPhraseVector := int64ToCtype(endAccentPhraseVector)

	// 返り値のデータの出力先を用意する
	datap, sizep, data, _ := makeDataReceiver[*C.float, C.ulong]()

	defer r.voicevoxPredictIntonationDataFree(*datap)

	r.voicevoxPredictIntonation(
		cLength,
		&cVowelPhonemeVector, &cConsonantPhonemeVecor,
		&cStartAccentVector, &cEndAccentVector,
		&cStartAccentPhraseVector, &cEndAccentPhraseVector,
		cSpeakerID,
		sizep, datap,
	)

	slice := unsafe.Slice(data[0], *sizep)

	retValue := make([]float32, length)

	for i, v := range slice {
		retValue[i] = float32(v)
	}

	return retValue, nil
}
