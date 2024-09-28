package fp_cm_utils

import (
	"errors"
)

// GetMultiLanguageWeekdays returns the names of the weekdays in the specified language.
func GetMultiLanguageWeekdays(lang string) ([]string, error) {
	weekdays := map[string][]string{
		"en": { // English
			"Sunday",
			"Monday",
			"Tuesday",
			"Wednesday",
			"Thursday",
			"Friday",
			"Saturday",
		},
		"bn": { // Bangla
			"রবিবার",
			"সোমবার",
			"মঙ্গলবার",
			"বুধবার",
			"বৃহস্পতিবার",
			"শুক্রবার",
			"শনিবার",
		},
		"hindi": { // Hindi
			"रविवार",
			"सोमवार",
			"मंगलवार",
			"बुधवार",
			"गुरुवार",
			"शुक्रवार",
			"शनिवार",
		},
		"tamil": { // Tamil
			"ஞாயிறு",
			"திங்கடு",
			"செவ்வாய்",
			"புதன்",
			"வியாழன்",
			"வெள்ளி",
			"சனி",
		},
		"telugu": { // Telugu
			"ఆదివారం",
			"సోమవారం",
			"మంగళవారం",
			"బుధవారం",
			"గురువారం",
			"శుక్రవారం",
			"శనివారం",
		},
		"china": {
			"星期日", // Sunday
			"星期一", // Monday
			"星期二", // Tuesday
			"星期三", // Wednesday
			"星期四", // Thursday
			"星期五", // Friday
			"星期六", // Saturday
		},
		"japan": {
			"日曜日", // Sunday
			"月曜日", // Monday
			"火曜日", // Tuesday
			"水曜日", // Wednesday
			"木曜日", // Thursday
			"金曜日", // Friday
			"土曜日", // Saturday
		},
		"german": {
			"Sonntag",    // Sunday
			"Montag",     // Monday
			"Dienstag",   // Tuesday
			"Mittwoch",   // Wednesday
			"Donnerstag", // Thursday
			"Freitag",    // Friday
			"Samstag",    // Saturday
		},
		"france": {
			"dimanche", // Sunday
			"lundi",    // Monday
			"mardi",    // Tuesday
			"mercredi", // Wednesday
			"jeudi",    // Thursday
			"vendredi", // Friday
			"samedi",   // Saturday
		},
	}

	if days, ok := weekdays[lang]; ok {
		return days, nil
	}
	return nil, errors.New("unsupported language code. please check the package supported languages")
}
