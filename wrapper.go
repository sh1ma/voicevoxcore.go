package voicevoxcorego

import "C"
import (
	"errors"
	"unsafe"
)

// VoicevoxCore is top-level API Wrapper instance
type VoicevoxCore struct {
	rawCore     *RawVoicevoxCore
	initialized bool
}

// VoicevoxCore のコンストラクタ関数
func NewVoicevoxCore() (core VoicevoxCore) {
	core = VoicevoxCore{rawCore: &RawVoicevoxCore{}}
	return
}

// C APIを通じてVoicevox_coreを初期化する関数
func (r *VoicevoxCore) Initialize(options VoicevoxInitializeOptions) (err error) {
	if r.initialized {
		err = errors.New("Already initialized")
		return
	}
	r.rawCore.VoicevoxInitialize(*options.Raw)
	r.initialized = true
	return
}

// 音声合成モデルをロードする関数
func (r *VoicevoxCore) LoadModel(speakerID int) (err error) {
	id := C.uint(speakerID)
	code := r.rawCore.VoicevoxLoadModel(id)
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
	datap := (**C.uchar)(&data[0])
	defer r.rawCore.VoicevoxWavFree(*datap)

	sizep := (*C.ulong)(unsafe.Pointer(&size))

	code := r.rawCore.VoicevoxTts(ctext, cspeakerID, *options.Raw, sizep, datap)
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
	datap := (**C.uchar)(&data[0])
	defer r.rawCore.VoicevoxWavFree(*datap)

	sizep := (*C.ulong)(unsafe.Pointer(&size))

	code := r.rawCore.VoicevoxSynthesis(ctext, cspeakerID, *options.Raw, sizep, datap)
	if int(code) != 0 {
		err = errors.New("Failed TTS")
		return
	}

	slice := unsafe.Slice(data[0], *sizep)
	sliceUnsafe := unsafe.SliceData(slice)
	slicebytes = C.GoBytes(unsafe.Pointer(sliceUnsafe), C.int((len(slice))))
	return
}
