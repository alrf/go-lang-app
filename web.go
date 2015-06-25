package main

import (  
	"fmt"
	"net/http"
	"log"
	"strconv"
	"flag"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"strings"
	"encoding/json"
	"sort"
)

const (  
	SQL_DB = "go"
	SQL_USER = "go"
	SQL_PASS = "go"
)

type Counter struct {
	Count			int			`json:"count"`
	Top_5_words		[]string	`json:"top_5_words"`
	Top_5_letters	[]string	`json:"top_5_letters"`
}

type Counters []Counter


func main() {

	numbPort := flag.Int("port", 8080, "port, integer")
	flag.Parse()
	strPort := strconv.Itoa(*numbPort)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/stats", StatsIndex)

    log.Fatal(http.ListenAndServe(":" + strPort, router))
	
}

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Main page!")
}

func StatsIndex(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", SQL_USER + ":" + SQL_PASS + "@/" + SQL_DB)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Count the words
	var countWord = 0
	err = db.QueryRow("select count(word) from words").Scan(&countWord)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("The square number of 13 is: %d", countWord)

	// Top 5 words
	rows, err := db.Query("select word from words group by word order by count(1) desc limit 5")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	topWord := make([]string,0)
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		topWord = append(topWord, name)
		//fmt.Printf("%s\n", name)
    }
	
	// Top 5 letters
	rows, err = db.Query("select distinct word from words")
    if err != nil {
        panic(err.Error())
    }
	defer rows.Close()
	var resString = ""
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			panic(err.Error())
		}
		resString += name
		//fmt.Printf("%s\n", name)
	}
	//fmt.Printf("\n" + resString + "\n")
	
	visited := make(map[string]bool)
	resDict := make(map[string]int)
	for _, element := range strings.Split(resString, "") {
		if visited[element] {
			resDict[element] = strings.Count(resString, element)
			//fmt.Printf(element+" count: %d\n", strings.Count(resString, element))
		}
		visited[element] = true
	}
	var i = 0
	topLetter := make([]string,0)
	for _, letter := range sortedKeys(resDict) {
		if(i == 5){ break }
		//fmt.Println(letter, resDict[letter])
		topLetter = append(topLetter, letter)
		i++
	}

	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
		
	counters := Counters{
		Counter{Count: countWord, Top_5_words: topWord, Top_5_letters: topLetter},
    }
	
	if err := json.NewEncoder(w).Encode(counters); err != nil {
        panic(err)
    }
}
