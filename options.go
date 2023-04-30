package voicevoxcorego

// #cgo LDFLAGS: -lvoicevox_core
// #include <voicevox_core.h>
import "C"

type (
	VoicevoxInitializeOptions struct {
		Raw *C.VoicevoxInitializeOptions
	}

	VoicevoxAudioQueryOptions struct {
		Raw *C.VoicevoxAudioQueryOptions
	}
	VoicevoxSynthesisOptions struct {
		Raw *C.VoicevoxSynthesisOptions
	}
	VoicevoxTtsOptions struct {
		Raw *C.VoicevoxTtsOptions
	}
)

func NewVoicevoxInitializeOptions(accelerationMode int, cpuNumThreads int, loadAllModels bool, openJtalkDictDir string) (options VoicevoxInitializeOptions) {
	raw := C.VoicevoxInitializeOptions{
		acceleration_mode:   C.int(accelerationMode),
		cpu_num_threads:     C.ushort(cpuNumThreads),
		load_all_models:     C.bool(loadAllModels),
		open_jtalk_dict_dir: C.CString(openJtalkDictDir),
	}
	options = VoicevoxInitializeOptions{Raw: &raw}
	return
}

func NewVoicevoxAudioQueryOptions(kana bool) (options VoicevoxAudioQueryOptions) {
	raw := C.VoicevoxAudioQueryOptions{
		kana: C.bool(kana),
	}

	options = VoicevoxAudioQueryOptions{Raw: &raw}
	return
}

func NewVoicevoxSynthesisOptions(enableInterrogativeUpspeak bool) (options VoicevoxSynthesisOptions) {
	raw := C.VoicevoxSynthesisOptions{
		enable_interrogative_upspeak: C.bool(enableInterrogativeUpspeak),
	}

	options = VoicevoxSynthesisOptions{Raw: &raw}
	return
}

func NewVoicevoxTtsOptions(kana bool, enableInterrogativeUpspeak bool) (options VoicevoxTtsOptions) {
	raw := C.VoicevoxTtsOptions{
		kana:                         C.bool(kana),
		enable_interrogative_upspeak: C.bool(enableInterrogativeUpspeak),
	}

	options = VoicevoxTtsOptions{Raw: &raw}
	return
}
