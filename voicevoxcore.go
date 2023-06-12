package voicevoxcorego

// #include <stdint.h>
import "C"
import (
	"encoding/json"
	"unsafe"

	"golang.org/x/xerrors"
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
func (r *VoicevoxCore) Initialize(options VoicevoxInitializeOptions) error {
	if r.initialized {
		err := xerrors.Errorf("Already initialized")
		return err
	}
	code := r.voicevoxInitialize(*options.Raw)
	if code != 0 {
		err := r.raiseError(code)
		return err
	}
	r.initialized = true
	return nil
}

// 音声合成モデルをロードする関数
func (r *VoicevoxCore) LoadModel(speakerID uint) error {
	id := C.uint(speakerID)
	code := r.voicevoxLoadModel(id)
	if code != 0 {
		err := r.raiseError(code)
		return err
	}
	return nil
}

/*
Text to Speechを実行する関数。実行結果はwavファイルフォーマットのバイト列。

Sample: https://github.com/sh1ma/sample-tts
*/
func (r *VoicevoxCore) Tts(text string, speakerID int, options VoicevoxTtsOptions) ([]byte, error) {
	ctext := C.CString(text)
	cspeakerID := C.uint(speakerID)

	datap, sizep, data, _ := makeDataReceiver[*C.uchar, C.ulong]()

	defer r.voicevoxWavFree(*datap)
	code := r.voicevoxTts(ctext, cspeakerID, *options.Raw, sizep, datap)
	if code != 0 {
		err := r.raiseError(code)
		return nil, err
	}

	slice := unsafe.Slice(data[0], *sizep)
	sliceUnsafe := unsafe.SliceData(slice)
	slicebytes := C.GoBytes(unsafe.Pointer(sliceUnsafe), C.int((len(slice))))
	return slicebytes, nil
}

/*
Audio Queryを基に音声合成を実行する関数。実行結果はwavファイルフォーマットのバイト列。

Sample: https://github.com/sh1ma/sample-synthesis
*/
func (r *VoicevoxCore) Synthesis(
	audioQuery string,
	speakerID int,
	options VoicevoxSynthesisOptions,
) ([]byte, error) {
	ctext := C.CString(audioQuery)
	cspeakerID := C.uint(speakerID)

	datap, sizep, data, _ := makeDataReceiver[*C.uchar, C.ulong]()

	defer r.voicevoxWavFree(*datap)
	code := r.voicevoxSynthesis(ctext, cspeakerID, *options.Raw, sizep, datap)
	if int(code) != 0 {
		err := r.raiseError(code)
		return nil, err
	}

	slice := unsafe.Slice(data[0], *sizep)
	sliceUnsafe := unsafe.SliceData(slice)
	slicebytes := C.GoBytes(unsafe.Pointer(sliceUnsafe), C.int((len(slice))))
	return slicebytes, nil
}

// `Initialize()` のデフォルトオプションを生成する
func (r *VoicevoxCore) MakeDefaultInitializeOptions() VoicevoxInitializeOptions {
	raw := r.voicevoxMakeDefaultInitializeOptions()
	return VoicevoxInitializeOptions{Raw: &raw}
}

// `Tts()` のデフォルトオプションを生成する
func (r *VoicevoxCore) MakeDefaultTtsOotions() VoicevoxTtsOptions {
	raw := r.voicevoxMakeDefaultTtsOptions()
	return VoicevoxTtsOptions{Raw: &raw}
}

// `AudioQuery()` のデフォルトオプションを生成する
func (r *VoicevoxCore) MakeDefaultAudioQueryOotions() VoicevoxAudioQueryOptions {
	raw := r.voicevoxMakeDefaultAudioQueryOptions()
	return VoicevoxAudioQueryOptions{Raw: &raw}
}

// `Synthesis()` のデフォルトオプションを生成する
func (r *VoicevoxCore) MakeDefaultSynthesisOotions() VoicevoxSynthesisOptions {
	raw := r.voicevoxMakeDefaultSynthesisOptions()
	return VoicevoxSynthesisOptions{Raw: &raw}
}

// オーディオクエリを発行する
func (r *VoicevoxCore) AudioQuery(text string, speakerID uint, options VoicevoxAudioQueryOptions) (AudioQuery, error) {
	ctext := C.CString(text)
	cSpeakerID := C.uint(speakerID)

	data := make([]*C.char, 1)
	datap := &data[0]
	defer r.voicevoxAudioQueryJsonFree(*datap)

	code := r.voicevoxAudioQuery(ctext, cSpeakerID, *options.Raw, datap)
	if code != 0 {
		err := r.raiseError(code)
		return AudioQuery{}, err
	}

	audioQueryJsonBytes := []byte(C.GoString(*datap))
	var audioQuery AudioQuery
	if err := json.Unmarshal(audioQueryJsonBytes, &audioQuery); err != nil {
		return AudioQuery{}, err
	}

	return audioQuery, nil
}

// ファイナライズ
func (r *VoicevoxCore) Finalize() {
	r.voicevoxFinalize()
}

// メタ情報のjsonを取得する
func (r *VoicevoxCore) GetMetasJson() string {
	cResult := r.voicevoxGetMetasJson()

	return C.GoString(cResult)
}

// サポートしているデバイス一覧のjsonを取得する
func (r *VoicevoxCore) GetSupportedDevicesJson() string {
	cResult := r.voicevoxGetSupportedDevicesJson()

	return C.GoString(cResult)
}

// Coreのバージョンを取得する
func (r *VoicevoxCore) GetCoreVersion() string {
	cResult := r.voicevoxGetVersion()

	return C.GoString(cResult)
}

// Gpuモードが有効になっているか確認する
func (r *VoicevoxCore) IsGpuMode() bool {
	cResult := r.voicevoxIsGpuMode()

	return bool(cResult)
}

// モデルがロードされているか確認する
func (r *VoicevoxCore) IsModelLoaded(speakerID uint) bool {
	cSpeakerID := C.uint(speakerID)
	cResult := r.voicevoxIsModelLoaded(cSpeakerID)

	return bool(cResult)
}

// 音素長を取得
func (r *VoicevoxCore) PredictDuration(speakerID int, phonemeVector []int64) ([]float32, error) {

	length := len(phonemeVector)

	datap, sizep, data, _ := makeDataReceiver[*C.float, C.ulong]()

	cPhonemeVectoTmp := &phonemeVector[0]
	cPhonemeVectorPtr := (*C.int64_t)(cPhonemeVectoTmp)
	cPhonemeVectorLength := (C.ulong)(length)

	defer r.voicevoxPredictDurationDataFree(*datap)
	code := r.voicevoxPredictDuration(cPhonemeVectorLength, cPhonemeVectorPtr, C.uint(speakerID), sizep, datap)

	if code != 0 {
		err := r.raiseError(code)
		return nil, err
	}

	slice := unsafe.Slice(data[0], *sizep)

	retValue := make([]float32, length)

	for i, v := range slice {
		retValue[i] = float32(v)
	}

	return retValue, nil
}

// 音高を取得
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
			return nil, xerrors.Errorf("%+w", "全てのベクター変数は同じ長さでなければなりません。")
		}
	}

	// Cの関数に渡す引数を用意する
	cLength := C.ulong(length)
	cSpeakerID := C.uint(speakerID)

	// []int64を[]C.int64に変換する関数を用意する
	int64ToCtype := sliceToCtype[int64, C.int64_t]

	// それぞれのVectorに上述の　`int64ToCtype` を適用する
	cVowelPhonemeVector := int64ToCtype(vowelPhonemeVector)
	cConsonantPhonemeVecor := int64ToCtype(consonantPhonemeVector)
	cStartAccentVector := int64ToCtype(startAccentPhraseVector)
	cEndAccentVector := int64ToCtype(endAccentVector)
	cStartAccentPhraseVector := int64ToCtype(startAccentPhraseVector)
	cEndAccentPhraseVector := int64ToCtype(endAccentPhraseVector)

	// 返り値のデータの出力先を用意する
	datap, sizep, data, _ := makeDataReceiver[*C.float, C.ulong]()

	defer r.voicevoxPredictIntonationDataFree(*datap)

	code := r.voicevoxPredictIntonation(
		cLength,
		&cVowelPhonemeVector, &cConsonantPhonemeVecor,
		&cStartAccentVector, &cEndAccentVector,
		&cStartAccentPhraseVector, &cEndAccentPhraseVector,
		cSpeakerID,
		sizep, datap,
	)
	if code != 0 {
		err := r.raiseError(code)
		return nil, err
	}

	slice := unsafe.Slice(data[0], *sizep)

	retValue := make([]float32, length)

	for i, v := range slice {
		retValue[i] = float32(v)
	}

	return retValue, nil
}

// phnemeVectorを元にデコードする
func (r *VoicevoxCore) Decode(speakerID uint, phonemeSize int, f0 []float32, phonemeVector []float32) ([]float32, error) {
	length := len(f0)

	float32ToCtype := sliceToCtype[float32, C.float]

	cSpeakerID := (C.uint)(speakerID)
	cLength := (C.uintptr_t)(length)
	cPhonemeSize := (C.uintptr_t)(phonemeSize)
	cF0 := float32ToCtype(f0)
	cPhonemeVector := float32ToCtype(phonemeVector)

	datap, sizep, data, _ := makeDataReceiver[*C.float, C.uintptr_t]()

	defer r.voicevoxDecodeDataFree(*datap)

	code := r.voicevoxDecode(cLength, cPhonemeSize, &cF0, &cPhonemeVector, cSpeakerID, sizep, datap)
	if code != 0 {
		err := r.raiseError(code)
		return nil, err
	}

	slice := unsafe.Slice(data[0], *sizep)

	var retValue []float32

	for _, v := range slice {
		retValue = append(retValue, float32(v))
	}

	return retValue, nil
}

// ErrorResultCode をメッセージに変換する
func (r *VoicevoxCore) ErrorResultToMessage(resultCode int) string {
	cResultCode := C.int(resultCode)
	message := r.errorResultToMessageInternal(cResultCode)

	return message
}

func (r *VoicevoxCore) raiseError(resultCode C.int) error {
	message := r.errorResultToMessageInternal(resultCode)
	err := xerrors.Errorf("%+w", xerrors.New(message))

	return err
}

func (r *VoicevoxCore) errorResultToMessageInternal(resultCode C.int) string {
	message := C.GoString(r.voicevoxErrorResultToMessage(resultCode))

	return message
}
