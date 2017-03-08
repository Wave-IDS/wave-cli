package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strings"
)

func main() {
	login()
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := stdin.ReadString('\n')
		text := strings.TrimSpace(line)
		parts := strings.Split(text, " ")
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "new":
			if len(parts) < 2 {
				log.Error("no arguments provided to 'new'")
				continue
			}
			switch parts[1] {
			case "collector":
				if len(parts) < 3 {
					log.Error("no name argument provided to 'new collector'")
					continue
				}
				newCollector(parts[2])
			case "user":
			}
		case "help":
			//printHelp()
		case "exit":
			os.Exit(0)
		default:
			log.Error("command not recognized, try 'help'")
		}
	}
}
