package main

import (  
    "fmt"
    "net"
    "os"
    "strconv"
    "bytes"
	"flag"
	"strings"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

const (  
    CONN_HOST = ""
    CONN_TYPE = "tcp"
	SQL_DB = "go"
	SQL_USER = "go"
	SQL_PASS = "go"
)

func main() {
	
	numbPort := flag.Int("port", 5555, "port, integer")
	flag.Parse()
	strPort := strconv.Itoa(*numbPort)
	
	l, err := net.Listen(CONN_TYPE, ":" + strPort)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()
	
	fmt.Println("Listening on " + CONN_HOST + ":" + strPort)
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	_ = reqLen
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
  
	n := bytes.Index(buf, []byte{0})  
	instring := string(buf[:n-1])
	instring = strings.ToLower(instring)
	words := strings.Fields(instring)

	if len(words) > 0 {
		word := words[0] // get only first word from phrase
		conn.Write([]byte("We receive your word. This is: " + word + "\n"));
		// Push the word in the DB.
		insertWord(word)
	}

	conn.Close()
}

func insertWord(word string) {  
	db, err := sql.Open("mysql", SQL_USER + ":" + SQL_PASS + "@/" + SQL_DB)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("insert into words (word) values (?)", word)
	if err != nil {
		panic(err.Error())
	}
}
