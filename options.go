package voicevoxcorego

// #include <voicevox_core.h>
import "C"

type (
	// `VoicevoxCore`を初期化する際のオプションを表す構造体
	VoicevoxInitializeOptions struct {
		// C構造体
		raw *C.VoicevoxInitializeOptions
	}
	// `AudioQuery()`を実行する際のオプションを表す構造体
	VoicevoxAudioQueryOptions struct {
		// C構造体
		raw *C.VoicevoxAudioQueryOptions
	}
	// `Synthesis()`を実行する際のオプションを表す構造体
	VoicevoxSynthesisOptions struct {
		// C構造体
		raw *C.VoicevoxSynthesisOptions
	}
	// `Tts()`を実行する際のオプションを表す構造体
	VoicevoxTtsOptions struct {
		// C構造体
		raw *C.VoicevoxTtsOptions
	}
)

/*
`VoiceVoxCore`の初期化オプションを生成する関数
*/
func NewVoicevoxInitializeOptions(accelerationMode int, cpuNumThreads int, loadAllModels bool, openJtalkDictDir string) (options VoicevoxInitializeOptions) {
	raw := C.VoicevoxInitializeOptions{
		acceleration_mode:   C.int(accelerationMode),
		cpu_num_threads:     C.ushort(cpuNumThreads),
		load_all_models:     C.bool(loadAllModels),
		open_jtalk_dict_dir: C.CString(openJtalkDictDir),
	}
	options = VoicevoxInitializeOptions{raw: &raw}
	return
}

/*
初期化オプションの`accelerationMode`をアップデートする関数
*/
func (o *VoicevoxInitializeOptions) UpdateAccelerationMode(accelerationMode int) {
	o.raw.acceleration_mode = C.int(accelerationMode)
}

/*
初期化オプションの`cpuNumThreads`をアップデートする関数
*/
func (o *VoicevoxInitializeOptions) UpdateCpuNumThreads(cpuNumThreads int) {
	o.raw.cpu_num_threads = C.ushort(cpuNumThreads)
}

/*
初期化オプションの`loadAllModels`をアップデートする関数
*/
func (o *VoicevoxInitializeOptions) UpdateLoadAllModels(loadAllModels bool) {
	o.raw.load_all_models = C.bool(loadAllModels)
}

/*
初期化オプションの`openJtalkDictDir`をアップデートする関数
*/
func (o *VoicevoxInitializeOptions) UpdateOpenJtalkDictDir(openJtalkDictDir string) {
	o.raw.open_jtalk_dict_dir = C.CString(openJtalkDictDir)
}

/*
`AudioQuery()`の初期化オプションを生成する関数
*/
func NewVoicevoxAudioQueryOptions(kana bool) (options VoicevoxAudioQueryOptions) {
	raw := C.VoicevoxAudioQueryOptions{
		kana: C.bool(kana),
	}

	options = VoicevoxAudioQueryOptions{raw: &raw}
	return
}

/*
`AudioQuery()`のオプションの`kana`をアップデートする関数
*/
func (o *VoicevoxAudioQueryOptions) UpdateKana(kana bool) {
	o.raw.kana = C.bool(kana)
}

/*
`Synthesis()`の初期化オプションを生成する関数
*/
func NewVoicevoxSynthesisOptions(enableInterrogativeUpspeak bool) *VoicevoxSynthesisOptions {
	raw := C.VoicevoxSynthesisOptions{
		enable_interrogative_upspeak: C.bool(enableInterrogativeUpspeak),
	}

	options := VoicevoxSynthesisOptions{raw: &raw}
	return &options
}

/*
`AudioQuery()`のオプションの`kana`をアップデートする関数
*/
func (o *VoicevoxSynthesisOptions) UpdateInterrogativeUpspeak(kana bool) {
	o.raw.enable_interrogative_upspeak = C.bool(kana)
}

/*
`Tts()`の初期化オプションを生成する関数
*/
func NewVoicevoxTtsOptions(kana bool, enableInterrogativeUpspeak bool) *VoicevoxTtsOptions {
	raw := C.VoicevoxTtsOptions{
		kana:                         C.bool(kana),
		enable_interrogative_upspeak: C.bool(enableInterrogativeUpspeak),
	}

	options := &VoicevoxTtsOptions{raw: &raw}
	return options
}

/*
`Tts()`のオプションの`kana`をアップデートする関数
*/
func (o *VoicevoxTtsOptions) UpdateKana(kana bool) {
	o.raw.kana = C.bool(kana)
}

/*
`Tts()`のオプションの`kana`をアップデートする関数
*/
func (o *VoicevoxTtsOptions) UpdateInterrogativeUpspeak(kana bool) {
	o.raw.enable_interrogative_upspeak = C.bool(kana)
}
