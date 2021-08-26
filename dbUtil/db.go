package dbUtil

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	at "github.com/sidhaler/attg/attconf"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Lifetime            = time.Minute * 3
	MaxConns            = 5
	MaxIdleConns        = 5
	OK           string = "----------------------------------- \n"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var s at.Atcfg

func OpenCon() (db *sql.DB) {
	cfg := s.Getconf()
	s.FatalWarns()
	port := strconv.Itoa(cfg.Port)
	fmt.Println(cfg.Usr + ":" + cfg.Passwd + "@tcp(" + cfg.Host + ":" + port + ")/" + cfg.Database)
	db, err := sql.Open("mysql", cfg.Usr+":"+cfg.Passwd+"@tcp("+cfg.Host+":"+port+")/"+cfg.Database)
	check(err)
	db.SetConnMaxLifetime(Lifetime)
	db.SetMaxOpenConns(MaxConns)
	db.SetMaxIdleConns(MaxIdleConns)
	return db
}

func Fetchall() {
	con := OpenCon()
	defer con.Close()
	rows, err := con.Query("SELECT * FROM " + s.Getconf().Table)
	check(err)

	columns, err := rows.Columns()
	check(err)

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		check(err)
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	check(err)
}

func writer(path string, str []string, perms os.FileMode) {

	file, err := os.OpenFile(path, os.O_RDWR, perms)
	check(err)
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range str {
		fmt.Fprintln(w, line)
	}
	fmt.Println("written...")
}

func ImportAll(path string) {
	_, err := os.Create(path)
	check(err)
	p, err := os.Stat(path)
	check(err)
	perms := p.Mode()
	con := OpenCon()
	defer con.Close()
	rows, err := con.Query("SELECT * FROM " + s.Getconf().Table)
	check(err)

	columns, err := rows.Columns()
	check(err)

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var bomba []string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		check(err)
		var value, elo string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			elo = columns[i] + ":" + value + "\n"
			bomba = append(bomba, elo)
		}
		bomba = append(bomba, OK)
		//writer(path, OK, perms)
	}
	writer(path, bomba, perms)
	fmt.Println(bomba)
	check(err)
}

// work but its insane bugged
func FetchwithID(key string) {
	cn := OpenCon()
	defer cn.Close()
	rows, err := cn.Query("SELECT * FROM " + s.Getconf().Table + " WHERE id = " + key)
	check(err)

	columns, err := rows.Columns()

	val := make([]uint8, len(columns))

	scanArgs := make([]interface{}, len(val))

	for i := range val {
		scanArgs[i] = &val
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		check(err)
		var value string
		for i, col := range val {
			if col == 0 {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	check(err)
}

func TestOpen() {
	db := OpenCon()
	fmt.Println("success connected to db !")
	db.Close()
}
