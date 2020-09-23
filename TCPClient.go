package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Print("Error in Dial", err.Error())
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Now please speaker name: ")
	readString, _ := reader.ReadString('\n')
	name := strings.Trim(readString, "\r\n")
	for {
		fmt.Printf("Please type words to server! type Q to quit!")
		message, _ := reader.ReadString('\n')
		msg := strings.Trim(message, "\r\n")
		if msg == "Q" {
			return
		}
		_, _ = dial.Write([]byte(name + " says: " + msg))
	}

}
