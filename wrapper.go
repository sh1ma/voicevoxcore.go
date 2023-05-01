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

func NewVoicevoxCore() (core VoicevoxCore) {
	core = VoicevoxCore{rawCore: &RawVoicevoxCore{}}
	return
}

func (r *VoicevoxCore) Initialize(options VoicevoxInitializeOptions) (err error) {
	if r.initialized {
		err = errors.New("Already initialized")
		return
	}
	r.rawCore.VoicevoxInitialize(*options.Raw)
	r.initialized = true
	return
}

func (r *VoicevoxCore) LoadModel(speakerID int) (err error) {
	id := C.uint(speakerID)
	code := r.rawCore.VoicevoxLoadModel(id)
	if int(code) != 0 {
		err = errors.New("Model load Failed")
	}
	return
}

func (r *VoicevoxCore) Tts(text string, speakerID int, options VoicevoxTtsOptions) (slicebytes []byte, err error) {
	ctext := C.CString(text)
	cspeakerID := C.uint(speakerID)

	var size int
	data := make([]*C.uchar, 1)
	// cbyte := C.getCByte()
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
