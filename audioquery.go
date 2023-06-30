package voicevoxcorego

import "encoding/json"

type (

	/*
		オーディオクエリを表す構造体

		`accent_phrases`にはアクセント句の配列を格納する。

		`spped_scale`は発話速度の倍率を表す。

		`pitch_scale`は音高の倍率を表す。

		`intonation_scale`はイントネーションの倍率を表す。

		`volume_scale`は音量の倍率を表す。

		`pre_phoneme_length`は発声開始前の無音の長さを表す。

		`post_phoneme_length`は発声終了後の無音の長さを表す。

		`output_sampling_rate`は出力音声のサンプリングレートを表す。

		`output_stereo`は出力音声がステレオかどうかを表す。

		`kana`は読み仮名を表す。
	*/
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

	/*
		アクセント句を表す構造体

		`moras`にはモーラの配列を格納する。

		`accent`はアクセント位置を表す。

		`pause_mora`はポーズのモーラを表す。

		`is_interrogative`は疑問文かどうかを表す。
	*/
	AccentPharase struct {
		Moras           []Mora `json:"moras"`
		Accent          uint32 `json:"accent"`
		PauseMora       Mora   `json:"pause_mora,omitempty"`
		IsInterrogative bool   `json:"is_interrogative"`
	}

	/*
		モーラを表す構造体

		`text`はモーラの文字列を表す。

		`consonant`は子音を表す。

		`consonant_length`は子音の長さを表す。

		`vowel`は母音を表す。

		`vowel_length`は母音の長さを表す。

		`pitch`は音高を表す。
	*/
	Mora struct {
		Text            string  `json:"text"`
		Consonant       string  `json:"consonant,omitempty"`
		ConsonantLength float32 `json:"consonant_length,omitempty"`
		Vowel           string  `json:"vowel"`
		VowelLength     float32 `json:"vowel_length"`
		Pitch           float32 `json:"pitch"`
	}
)

/*
AudioQuery構造体をJsonのバイト列に変換する
*/
func (q *AudioQuery) ToJson() ([]byte, error) {
	return json.Marshal(q)
}

/*
AudioQuery構造体をJson文字列に変換する
*/
func (q *AudioQuery) ToJsonString() (string, error) {
	jsonBytes, err := q.ToJson()
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

/*
Jsonのバイト列からAudioQuery構造体を生成する
*/
func NewAudioQueryFromJson(queryJson []byte) (*AudioQuery, error) {
	var query AudioQuery
	err := json.Unmarshal(queryJson, &query)
	if err != nil {
		return &AudioQuery{}, err
	}
	return &query, nil
}
