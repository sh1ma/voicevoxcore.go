package voicevoxcorego

import "encoding/json"

type (
	AudioQuery struct {
		AccentPharases     []AccentPharase `json:"accent_phrases"`
		SpeedScale         float32         `json:"speed_scale"`
		PitchScale         float32         `json:"pitch_scale"`
		IntonationScale    float32         `json:"intonation_scale"`
		VolumeScale        float32         `json:"volume_scale"`
		PrePhonemeLength   float32         `json:"pre_phoneme_length"`
		PostPhonemeLength  float32         `json:"post_phoneme_length"`
		OutputSamplingRate float32         `json:"output_sampling_rate"`
		OutputStereo       bool            `json:"output_stereo"`
		Kana               string          `json:"kana"`
	}

	AccentPharase struct {
		Moras           []Mora `json:"moras"`
		Accent          uint32 `json:"accent"`
		PauseMora       Mora   `json:"pause_mora,omitempty"`
		IsInterrogative bool   `json:"is_interrogative"`
	}

	Mora struct {
		Text            string  `json:"text"`
		Consonant       string  `json:"consonant,omitempty"`
		ConsonantLength float32 `json:"consonant_length,omitempty"`
		Vowel           string  `json:"vowel"`
		VowelLength     float32 `json:"vowel_length"`
		Pitch           float32 `json:"pitch"`
	}
)

func (q *AudioQuery) ToJson() ([]byte, error) {
	return json.Marshal(q)
}

func (q *AudioQuery) ToJsonString() (string, error) {
	jsonBytes, err := q.ToJson()
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func NewAudioQueryFromJson(queryJson []byte) (AudioQuery, error) {
	var query AudioQuery
	err := json.Unmarshal(queryJson, &query)
	if err != nil {
		return AudioQuery{}, err
	}
	return query, nil
}
