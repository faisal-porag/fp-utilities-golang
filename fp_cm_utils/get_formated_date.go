package fp_cm_utils

import (
	"errors"
	"time"
)

// GetMultiLanguageFormatDate returns a formatted date string in the specified language with short month names.
func GetMultiLanguageFormatDate(date time.Time, lang string) (string, error) {
	dateFormats := map[string]string{
		"en":     "Mon, Jan 2, 2006",      // Short date format for English
		"bn":     "সোম. ২ জানুয়ারী ২০০৬", // Short date format for Bengali
		"france": "lun. 2 janv. 2006",     // Short date format for French
		"german": "Mo. 2. Jan. 2006",      // Short date format for German
		"china":  "2006年1月2日周一",           // Short date format for Chinese
		"japan":  "2006年1月2日月",            // Short date format for Japanese
		"hindi":  "सोम. 2 जन. 2006",       // Short date format for Hindi
	}

	if format, ok := dateFormats[lang]; ok {
		return date.Format(format), nil
	}
	return "", errors.New("unsupported language code. please check the package supported languages")
}
