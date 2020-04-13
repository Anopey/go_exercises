package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	handleURL(readURL(conn), conn)
}

func readURL(conn net.Conn) string {
	scanner := bufio.NewScanner(conn)
	i := 0
	var url string = ""
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if i == 0 {
			//request
			fields := strings.Fields(line)
			if len(line) < 3 {
				break
			}
			m := fields[0]
			url = fields[1]
			fmt.Println("***METHOD: ", m)
		}
		if line == "" {
			//reached end of request
			break
		}
		i++
	}
	return url
}

func handleURL(url string, conn net.Conn) {
	if url == "" {
		//was invalid.
		fmt.Fprint(conn, "400 Bad Request")
	} else {
		//was valid.
		fmt.Println("URL Was:", url)
		body := `<!DOCTYPE html>
		<html>
		<head>
			<title>Welcome to the website!</title>
		</head>
		<body>Your requested URL was: ` + url + `</body>
		</html>`

		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}
}
