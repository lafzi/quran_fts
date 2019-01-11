package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/harunalfat/quran-clean-fts/specialcase"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath     = "./data/lafzi.sqlite"
	fathah     = "َ"
	kasrah     = "ِ"
	dhamah     = "ُ"
	fathahtain = "ً"
	kasrahtain = "ٍ"
	dhamahtain = "ٌ"
	syaddah    = "ّ"
	sukun      = "ْ"
)

func writeSQLiteFile() {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		d1 := []byte("")
		err := ioutil.WriteFile(dbPath, d1, 0644)
		checkErr(err)
		println("DB file created")
	}
}

func setupDB() *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	checkErr(err)
	return db
}

func prepareIndex(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS quran_text")
	checkErr(err)

	query := "CREATE VIRTUAL TABLE quran_text USING fts4(text_reverse)"
	_, err = db.Exec(query)
	checkErr(err)
	println("table 'quran_text' created")
}

func removeDiacritic(in string) string {
	replacer := strings.NewReplacer(syaddah, "", sukun, "", fathah, "", fathahtain, "",
		kasrah, "", kasrahtain, "", dhamah, "", dhamahtain, "")
	return replacer.Replace(in)
}

func insertQuranVerse(db *sql.DB) {
	file, err := os.Open("./data/quran-simple-clean.txt")
	checkErr(err)

	file2, err := os.Open("./data/quran-uthmani.txt")
	checkErr(err)

	defer file.Close()
	defer file2.Close()

	scanner := bufio.NewScanner(file)
	scanner2 := bufio.NewScanner(file2)
	i := 1

	replacer := strings.NewReplacer("أ", "ا", "إ", "ا", "ة", "ه", "آ", "ا")
	for scanner.Scan() {
		scanner2.Scan()
		line := scanner.Text()
		line2 := scanner2.Text()

		split := strings.Split(line, "|")
		surahNum := split[0]
		ayatNum := split[1]
		ayat := split[2]

		split2 := strings.Split(line2, "|")
		ayatFull := split2[2]

		if ayatNum == "1" && surahNum != "1" {
			ayat = ayat[42:len(ayat)]
		}

		ayat = replacer.Replace(ayat)
		sn, err := strconv.Atoi(surahNum)
		checkErr(err)
		an, err := strconv.Atoi(ayatNum)
		checkErr(err)

		if specialcase.IsSpecialCase(sn, an) {
			ayat = specialcase.ReplaceSpecialCase(ayat, an, sn)
		}
		ayatReverse := reverse(ayat)

		query := fmt.Sprintf(`
			INSERT INTO quran_text VALUES('%s')
		`, ayatReverse)

		queryUpdate := fmt.Sprintf(`
			UPDATE ayat_quran SET arabic_text_length = %d, short_arabic_text_length = %d WHERE _id = %d
		`, len(ayatFull), len(ayat), i)

		_, err = db.Exec(query)
		checkErr(err)

		_, err = db.Exec(queryUpdate)
		checkErr(err)

		i++
		if i > 6236 {
			break
		}
	}
	println("Done indexing quran.")

	/*
		_, err = db.Exec(`
			UPDATE
				ayat_quran
			SET short_arabic_text_length = (
				SELECT LENGTH(text_reverse) FROM quran_text WHERE docid = _id
			)
		`)
		checkErr(err)
	*/
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	writeSQLiteFile()
	db := setupDB()
	prepareIndex(db)
	insertQuranVerse(db)
}
