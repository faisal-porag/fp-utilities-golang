package fp_cm_utils

import (
	"reflect"
	"testing"
)

func TestGetMultiLanguageWeekdays(t *testing.T) {
	tests := []struct {
		lang     string
		expected []string
		hasError bool
	}{
		{"en", []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}, false},
		{"bn", []string{"রবিবার", "সোমবার", "মঙ্গলবার", "বুধবার", "বৃহস্পতিবার", "শুক্রবার", "শনিবার"}, false},
		{"hindi", []string{"रविवार", "सोमवार", "मंगलवार", "बुधवार", "गुरुवार", "शुक्रवार", "शनिवार"}, false},
		{"tamil", []string{"ஞாயிறு", "திங்கடு", "செவ்வாய்", "புதன்", "வியாழன்", "வெள்ளி", "சனி"}, false},
		{"telugu", []string{"ఆదివారం", "సోమవారం", "మంగళవారం", "బుధవారం", "గురువారం", "శుక్రవారం", "శనివారం"}, false},
		{"china", []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}, false},
		{"japan", []string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}, false},
		{"german", []string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"}, false},
		{"france", []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"}, false},
		{"es", nil, true}, // Unsupported language
	}

	for _, test := range tests {
		t.Run(test.lang, func(t *testing.T) {
			result, err := GetMultiLanguageWeekdays(test.lang)

			if test.hasError {
				if err == nil {
					t.Errorf("Expected an error for language %s, got none", test.lang)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error for language %s, got: %v", test.lang, err)
				}
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("For language %s, expected %v, got %v", test.lang, test.expected, result)
				}
			}
		})
	}
}
