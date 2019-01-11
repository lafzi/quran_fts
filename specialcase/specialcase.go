package specialcase

import "strings"

type AyatPos struct {
	SuratNo   int
	AyatNo    int
	AyatLong  string
	AyatShort string
}

func ReplaceSpecialCase(in string, ayatNo, suratNo int) string {
	var ap AyatPos
	for _, c := range getSpecialCase() {
		if ayatNo == c.AyatNo && suratNo == c.SuratNo {
			ap = c
			break
		}
	}

	return strings.Replace(in, ap.AyatShort, ap.AyatLong, -1)
}

func getSpecialCase() []AyatPos {
	return []AyatPos{
		AyatPos{SuratNo: 2, AyatNo: 1, AyatLong: "الف لام ميم", AyatShort: "الم"},
		AyatPos{SuratNo: 3, AyatNo: 1, AyatLong: "الف لام ميم", AyatShort: "الم"},
		AyatPos{SuratNo: 7, AyatNo: 1, AyatLong: "الف لام ميم صاد", AyatShort: "المص"},
		AyatPos{SuratNo: 10, AyatNo: 1, AyatLong: "الف لام راء", AyatShort: "الر"},
		AyatPos{SuratNo: 11, AyatNo: 1, AyatLong: "الف لام راء", AyatShort: "الر"},
		AyatPos{SuratNo: 12, AyatNo: 1, AyatLong: "الف لام راء", AyatShort: "الر"},
		AyatPos{SuratNo: 13, AyatNo: 1, AyatLong: "الف لام ميم راء", AyatShort: "المر"},
		AyatPos{SuratNo: 14, AyatNo: 1, AyatLong: "الف لام راء", AyatShort: "الر"},
		AyatPos{SuratNo: 15, AyatNo: 1, AyatLong: "الف لام راء", AyatShort: "الر"},
		AyatPos{SuratNo: 19, AyatNo: 1, AyatLong: "كاف ها يا عين صاد", AyatShort: "كهيعص"},
		AyatPos{SuratNo: 20, AyatNo: 1, AyatLong: "طه", AyatShort: "طه"},
		AyatPos{SuratNo: 26, AyatNo: 1, AyatLong: "طا سين ميم", AyatShort: "طسم"},
		AyatPos{SuratNo: 27, AyatNo: 1, AyatLong: "طا سين", AyatShort: "طس"},
		AyatPos{SuratNo: 28, AyatNo: 1, AyatLong: "طا سين ميم", AyatShort: "طسم"},
		AyatPos{SuratNo: 36, AyatNo: 1, AyatLong: "ياسين", AyatShort: "يس"},
		AyatPos{SuratNo: 38, AyatNo: 1, AyatLong: "صاد", AyatShort: "ص"},
		AyatPos{SuratNo: 40, AyatNo: 1, AyatLong: "حاميم", AyatShort: "حم"},
		AyatPos{SuratNo: 41, AyatNo: 1, AyatLong: "حاميم", AyatShort: "حم"},
		AyatPos{SuratNo: 42, AyatNo: 1, AyatLong: "حاميم", AyatShort: "حم"},
		AyatPos{SuratNo: 43, AyatNo: 1, AyatLong: "حاميم", AyatShort: "حم"},
		AyatPos{SuratNo: 44, AyatNo: 1, AyatLong: "حاميم", AyatShort: "حم"},
		AyatPos{SuratNo: 45, AyatNo: 1, AyatLong: "حاميم", AyatShort: "حم"},
		AyatPos{SuratNo: 46, AyatNo: 1, AyatLong: "حاميم", AyatShort: "حم"},
		AyatPos{SuratNo: 42, AyatNo: 2, AyatLong: "عين سين قاف", AyatShort: "عسق"},
		AyatPos{SuratNo: 50, AyatNo: 1, AyatLong: "قاف", AyatShort: "ق"},
		AyatPos{SuratNo: 68, AyatNo: 1, AyatLong: "نون", AyatShort: "ن"},
	}
}

func IsSpecialCase(suratNo, ayatNo int) bool {
	for _, c := range getSpecialCase() {
		if ayatNo == c.AyatNo && suratNo == c.SuratNo {
			return true
		}
	}
	return false
}
