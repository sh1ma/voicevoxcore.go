package voicevoxcorego

// #cgo LDFLAGS: -lvoicevox_core
// #include <voicevox_core.h>
import "C"

// RawVoicevoxCore is a function group that wraps the C API
type RawVoicevoxCore struct{}

func (r *RawVoicevoxCore) voicevoxMakeDefaultInitializeOptions() C.VoicevoxInitializeOptions {
	return C.voicevox_make_default_initialize_options()
}

func (r *RawVoicevoxCore) voicevoxInitialize(options C.VoicevoxInitializeOptions) C.int {
	return C.voicevox_initialize(options)
}

func (r *RawVoicevoxCore) voicevoxGetVersion() *C.char {
	return C.voicevox_get_version()
}

func (r *RawVoicevoxCore) voicevoxLoadModel(speakerID C.uint) C.int {
	return C.voicevox_load_model(speakerID)
}

func (r *RawVoicevoxCore) voicevoxIsGpuMode() C.bool {
	return C.voicevox_is_gpu_mode()
}

func (r *RawVoicevoxCore) voicevoxIsModelLoaded(speakerID C.uint) C.bool {
	return C.voicevox_is_model_loaded(speakerID)
}

func (r *RawVoicevoxCore) voicevoxFinalize() C.void {
	return C.voicevox_finalize()
}

func (r *RawVoicevoxCore) voicevoxGetMetasJson() *C.char {
	return C.voicevox_get_metas_json()
}

func (r *RawVoicevoxCore) voicevoxGetSupportedDevicesJson() *C.char {
	return C.voicevox_get_supported_devices_json()
}

func (r *RawVoicevoxCore) voicevoxPredictDuration(
	length C.ulong,
	phonemeVector *C.int64_t,
	speakerID C.uint,
	outputPredictDurationDataLength *C.ulong,
	outputPredictDurationData **C.float,
) C.int {
	return C.voicevox_predict_duration(
		length,
		phonemeVector,
		speakerID,
		outputPredictDurationDataLength,
		outputPredictDurationData,
	)
}

func (r *RawVoicevoxCore) voicevoxPredictDurationDataFree(predictDurationData *C.float) C.void {
	return C.voicevox_predict_duration_data_free(predictDurationData)
}

func (r *RawVoicevoxCore) voicevoxPredictIntonation(
	length C.ulong,
	vowel_phoneme_vector *C.int64_t,
	consonantPhonemeVector *C.int64_t,
	startAccentVector *C.int64_t,
	endAccentVector *C.int64_t,
	startAccentPhraseVector *C.int64_t,
	endAccentPhraseVector *C.int64_t,
	speakerID C.uint,
	outputPredictIntonationDataLength *C.ulong,
	outputPredictIntonationData **C.float,
) C.int {
	return C.voicevox_predict_intonation(
		length,
		vowel_phoneme_vector,
		consonantPhonemeVector,
		startAccentVector,
		endAccentVector,
		startAccentPhraseVector,
		endAccentPhraseVector,
		speakerID,
		outputPredictIntonationDataLength,
		outputPredictIntonationData,
	)
}

func (r *RawVoicevoxCore) voicevoxPredictIntonationDataFree(predictIntonationData *C.float) C.void {
	return C.voicevox_predict_intonation_data_free(predictIntonationData)
}

func (r *RawVoicevoxCore) voicevoxDecode(
	length C.ulong,
	phonemeSize C.ulong,
	f0 *C.float,
	phonemeVector *C.float,
	speakerID C.uint,
	outputDecodeDataLength *C.ulong,
	outputDecodeData **C.float,
) C.int {
	return C.voicevox_decode(
		length,
		phonemeSize,
		f0,
		phonemeVector,
		speakerID,
		outputDecodeDataLength,
		outputDecodeData)
}

func (r *RawVoicevoxCore) voicevoxDecodeDataFree(decodeData *C.float) C.void {
	return C.voicevox_decode_data_free(decodeData)
}

func (r *RawVoicevoxCore) voicevoxMakeDefaultAudioQueryOptions() C.VoicevoxAudioQueryOptions {
	return C.voicevox_make_default_audio_query_options()
}

func (r *RawVoicevoxCore) voicevoxAudioQuery(
	text *C.char,
	speakerID C.uint,
	options C.VoicevoxAudioQueryOptions,
	outputAudioQueryJson **C.char,
) C.VoicevoxResultCode {
	return C.voicevox_audio_query(text, speakerID, options, outputAudioQueryJson)
}

func (r *RawVoicevoxCore) voicevoxMakeDefaultSynthesisOptions() C.VoicevoxSynthesisOptions {
	return C.voicevox_make_default_synthesis_options()
}

func (r *RawVoicevoxCore) voicevoxSynthesis(
	audioQueryJson *C.char,
	speakerID C.uint,
	options C.VoicevoxSynthesisOptions,
	outputWavLength *C.ulong,
	outputWav **C.uchar,
) C.int {
	return C.voicevox_synthesis(audioQueryJson, speakerID, options, outputWavLength, outputWav)
}

func (r *RawVoicevoxCore) voicevoxMakeDefaultTtsOptions() C.VoicevoxTtsOptions {
	return C.voicevox_make_default_tts_options()
}

func (r *RawVoicevoxCore) voicevoxTts(
	text *C.char,
	speakerID C.uint,
	options C.VoicevoxTtsOptions,
	outputWavLength *C.ulong,
	outputWav **C.uchar,
) C.int {
	return C.voicevox_tts(
		text,
		speakerID,
		options,
		outputWavLength,
		outputWav,
	)
}

func (r *RawVoicevoxCore) voicevoxAudioQueryJsonFree(audioQueryJson *C.char) C.void {
	return C.voicevox_audio_query_json_free(audioQueryJson)
}

func (r *RawVoicevoxCore) voicevoxWavFree(wav *C.uchar) C.void {
	return C.voicevox_wav_free(wav)
}

func (r *RawVoicevoxCore) voicevoxErrorResultToMessage(resultCode C.VoicevoxResultCode) *C.char {
	return C.voicevox_error_result_to_message(resultCode)
}
