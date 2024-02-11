package kanatrans

import (
	"strings"
)

// MoraeKanaConverter converts morae to Kana characters.
type MoraeKanaConverter struct {
	vowels    string
	moraMap   map[string]string
	doubled   string
	geminate  string
}

// NewMoraeKanaConverter creates a new instance of MoraeKanaConverter.
func NewMoraeKanaConverter() *MoraeKanaConverter {
	return &MoraeKanaConverter{
		vowels:  "aeiou",
		moraMap: map[string]string{
			"a":   "ア", "i": "イ", "u": "ウ", "e": "エ", "o": "オ",
			"ka":  "カ", "ki": "キ", "ku": "ク", "ke": "ケ", "ko": "コ",
			"ga":  "ガ", "gi": "ギ", "gu": "グ", "ge": "ゲ", "go": "ゴ",
			"kya": "キャ", "kyu": "キュ", "kyo": "キョ",
			"kwa": "クァ", "kwi": "クィ", "kwu": "クゥ", "kwe": "クェ", "kwo": "クォ",
			"gya": "ギャ", "gyu": "ギュ", "gyo": "ギョ",
			"gwa": "グァ", "gwi": "グィ", "gwu": "グゥ", "gwe": "グェ", "gwo": "グォ",
			"sa":  "サ", "si": "シ", "su": "ス", "se": "セ", "so": "ソ",
			"sha": "シャ", "shi": "シ", "shu": "シュ", "she": "シェ", "sho": "ショ",
			"ja":  "ジャ", "ji": "ジ", "ju": "ジュ", "je": "ジェ", "jo": "ジョ",
			"jya": "ジャ", "jyu": "ジュ", "jyo": "ジョ",
			"za":  "ザ", "zi": "ジ", "zu": "ズ", "ze": "ゼ", "zo": "ゾ",
			"zya": "ジャ", "zyu": "ジュ", "zyo": "ジョ",
			"ta":  "タ", "ti": "ティ", "tu": "ツ", "te": "テ", "to": "ト",
			"cha": "チャ", "chi": "チ", "chu": "チュ", "che": "チェ", "cho": "チョ",
			"da":  "ダ", "di": "ディ", "du": "ドゥ", "de": "デ", "do": "ド",
			"na":  "ナ", "ni": "ニ", "nu": "ヌ", "ne": "ネ", "no": "ノ",
			"nya": "ニャ", "nyi": "ニィ", "nyu": "ニュ", "nye": "ニェ", "nyo": "ニョ",
			"ha":  "ハ", "hi": "ヒ", "hu": "フ", "he": "ヘ", "ho": "ホ",
			"hya": "ヒャ", "hyu": "ヒュ", "hyo": "ヒョ",
			"fa":  "ファ", "fi": "フィ", "fu": "フ", "fe": "フェ", "fo": "フォ",
			"fya": "フャ", "fyi": "フィ", "fyu": "フュ", "fye": "フェ", "fyo": "フョ",
			"ba":  "バ", "bi": "ビ", "bu": "ブ", "be": "ベ", "bo": "ボ",
			"bya": "ビャ", "byi": "ビィ", "byu": "ビュ", "bye": "ビェ", "byo": "ビョ",
			"pa":  "パ", "pi": "ピ", "pu": "プ", "pe": "ペ", "po": "ポ",
			"pya": "ピャ", "pyu": "ピュ", "pyo": "ピョ",
			"ma":  "マ", "mi": "ミ", "mu": "ム", "me": "メ", "mo": "モ",
			"mya": "ミャ", "myu": "ミュ", "myo": "ミョ",
			"ya":  "ヤ", "yi": "イ", "yu": "ユ", "ye": "イェ", "yo": "ヨ",
			"ra":  "ラ", "ri": "リ", "ru": "ル", "re": "レ", "ro": "ロ",
			"rya": "リャ", "ryu": "リュ", "ryo": "リョ",
			"wa":  "ワ", "wi": "ウィ", "wu": "ウ", "we": "ウェ", "wo": "ウォ",
			"N":   "ン",
		},
		doubled:  "ー",
		geminate: "ッ",
	}
}

// ConvertMorae converts morae to Kana characters.
func (mkc *MoraeKanaConverter) ConvertMorae(morae string) string {
	sounds := strings.Split(morae, ".")
	result := mkc.moraMap[sounds[0]]
	for idx := 1; idx < len(sounds); idx++ {
		if strings.Contains(mkc.vowels, string(sounds[idx][0])) && sounds[idx-1][len(sounds[idx-1])-1:] == sounds[idx][:1] {
			result += mkc.doubled
		} else if !strings.Contains(mkc.vowels, string(sounds[idx][0])) && idx+1 < len(sounds) && sounds[idx] == sounds[idx+1][:1] {
			result += mkc.geminate
		} else {
			result += mkc.moraMap[sounds[idx]]
		}
	}
	return result
}