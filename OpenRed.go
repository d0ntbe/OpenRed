package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println(" ")
		fmt.Println("You forget the args. See Usage! Type: h/-h/help")
		fmt.Println("With Love from ZL...")
		return
	}

	if os.Args[1] == "-h" || os.Args[1] == "-help" || os.Args[1] == "help" || os.Args[1] == "h" {

		fmt.Println("USAGE: ")
		fmt.Println("Example: go run OpenRed.go www.fsaexample.com")
		fmt.Println("--<-@  With Love from ZL...")
		return
	}

	conn, err := net.Dial("tcp", os.Args[1]+":80")
	if err != nil {
		fmt.Println("Error connection!", err)
		fmt.Println("Example for target: ddddd.com")
		fmt.Println("If proto is closed for http, this method does not work!")
		return
	}
	defer conn.Close()

	openreditect := "evil.cococo"

	reqString := fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\n", "/", openreditect)

	reqString += "\r\n"

	_, err = conn.Write([]byte(reqString))
	if err != nil {
		fmt.Println("Error request:", err)
		return
	}

	// читаем ответ
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, nil)
	if err != nil {
		fmt.Println("Error read answer", err)
		return
	}
	defer response.Body.Close()

	location := response.Header.Get("Location")
	if strings.Contains(location, openreditect) == false {

		fmt.Println("OPEN REDIRECT NOT FOUND!!!")

	}
	if strings.Contains(location, openreditect) == true {
		fmt.Println("Header Location: in response", location)
		fmt.Println("OPEN REDIRECT FOUND!!!")

	}

}
