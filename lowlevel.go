package voicevoxcorego

// #cgo LDFLAGS: -lvoicevox_core
// #include <voicevox_core.h>
import "C"

type VoicevoxCore struct {
}

func VoicevoxMakeDefaultInitializeOptions() C.VoicevoxInitializeOptions {
	return C.voicevox_make_default_initialize_options()
}

func VoicevoxInitialize(options C.VoicevoxInitializeOptions) C.int {
	return C.voicevox_initialize(options)
}

func VoicevoxGetVersion() *C.char {
	return C.voicevox_get_version()
}

func VoicevoxLoadModel(speakerID C.uint) C.int {
	return C.voicevox_load_model(speakerID)
}

func VoicevoxIsGpuMode() C.bool {
	return C.voicevox_is_gpu_mode()
}

func VoicevoxIsModelLoaded(speakerID C.uint) C.bool {
	return C.voicevox_is_model_loaded(speakerID)
}

func VoicevoxFinalize() C.void {
	return C.voicevox_finalize()
}

func VoicevoxGetMetasJson() *C.char {
	return C.voicevox_get_metas_json()
}

func VoicevoxGetSupportedDevicesJson() *C.char {
	return C.voicevox_get_supported_devices_json()
}

func VoicevoxPredictDuration(
	length C.ulong,
	phonemeVector *C.longlong,
	speakerID C.uint,
	outputPredictDurationDataLength C.ulong,
	outputPredictDurationData **C.float,
) C.int {
	return C.voicevox_predict_duration(
		length,
		phonemeVector,
		speakerID,
		&outputPredictDurationDataLength,
		outputPredictDurationData,
	)
}

func VoicevoxPredictDurationDataFree(predictDurationData *C.float) C.void {
	return C.voicevox_predict_duration_data_free(predictDurationData)
}

func VoicevoxPredictIntonation(
	length C.ulong,
	vowel_phoneme_vector *C.longlong,
	consonantPhonemeVector *C.longlong,
	startAccentVector *C.longlong,
	endAccentVector *C.longlong,
	startAccentPhraseVector *C.longlong,
	endAccentPhraseVector *C.longlong,
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

func VoicevoxPredictIntonationDataFree(predictIntonationData *C.float) C.void {
	return C.voicevox_predict_intonation_data_free(predictIntonationData)
}

func VoicevoxDecode(
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

func VoicevoxDecodeDataFree(decodeData *C.float) C.void {
	return C.voicevox_decode_data_free(decodeData)
}

func VoicevoxMakeDefaultAudioQueryOptions() C.VoicevoxAudioQueryOptions {
	return C.voicevox_make_default_audio_query_options()
}

func VoicevoxAudioQuery(
	text *C.char,
	speakerID C.uint,
	options C.VoicevoxAudioQueryOptions,
	outputAudioQueryJson **C.char,
) C.VoicevoxResultCode {
	return C.voicevox_audio_query(text, speakerID, options, outputAudioQueryJson)
}

func VoicevoxMakeDefaultSynthesisOptions() C.VoicevoxSynthesisOptions {
	return C.voicevox_make_default_synthesis_options()
}

func VoicevoxSynthesis(
	audioQueryJson *C.char,
	speakerID C.uint,
	options C.VoicevoxSynthesisOptions,
	outputWavLength *C.ulong,
	outputWav **C.uchar,
) C.int {
	return C.voicevox_synthesis(audioQueryJson, speakerID, options, outputWavLength, outputWav)
}

func VoicevoxMakeDefaultTtsOptions() C.VoicevoxTtsOptions {
	return C.voicevox_make_default_tts_options()
}

func VoicevoxTts(
	text *C.char,
	speakerID C.uint,
	options C.VoicevoxTtsOptions,
	outputWavLength C.ulong,
	outputWav **C.uchar,
) C.int {
	return C.voicevox_tts(
		text,
		speakerID,
		options,
		&outputWavLength,
		outputWav,
	)
}

func VoicevoxAudioQueryJsonFree(audioQueryJson *C.char) C.void {
	return C.voicevox_audio_query_json_free(audioQueryJson)
}

func VoicevoxWavFree(wav *C.uchar) C.void {
	return C.voicevox_wav_free(wav)
}

func VoicevoxErrorResultToMessage(resultCode C.VoicevoxResultCode) *C.char {
	return C.voicevox_error_result_to_message(resultCode)
}
