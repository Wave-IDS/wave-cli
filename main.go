package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/joho/godotenv/autoload"
	"io"
	"os"
	"strings"
)

func main() {
	login()
	defer logout()
	fmt.Println("Wave CLI")
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(fmt.Sprintf("%s@wave-cli> ", username))
		line, err := stdin.ReadString('\n')
		if err == io.EOF {
			fmt.Println("exit")
			return
		}
		text := strings.TrimSpace(line)
		parts := strings.Split(text, " ")
		if len(parts) == 0 || parts[0] == "" {
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
				if len(parts) < 3 {
					log.Error("no name argument provided to 'new user'")
					continue
				}
				newUser(parts[2])
			default:
				fmt.Println("unknown argument to 'new'")
			}
		case "get":
			if len(parts) < 2 {
				log.Error("no arguments provided to 'get'")
				continue
			}
			switch parts[1] {
			case "collectors":
				getCollectors()
			case "users":
				getUsers()
			case "tls":
				getTLS()
			default:
				fmt.Println("unknown argument to 'get'")
			}
		case "update":
			if len(parts) < 2 {
				log.Error("no arguments provided to 'update'")
				continue
			}
			switch parts[1] {
			case "password":
				if len(parts) < 4 {
					log.Error("not enough arguments provided to 'update password'")
					continue
				}
				updateUserPassword(parts[2], parts[3])
			case "assignpassword":
				if len(parts) < 4 {
					log.Error("not enough arguments provided to 'update assignpassword'")
					continue
				}
				assignUserPassword(parts[2], parts[3])
			case "username":
				if len(parts) < 3 {
					log.Error("not enough arguments provided to 'update username'")
					continue
				}
				updateUserName(parts[2])
			case "tls":
				if len(parts) < 4 {
					log.Error("not enough arguments provided to 'update tls'")
					continue
				}
				setTLS(parts[2], parts[3])
			default:
				fmt.Println("unknown argument to 'update'")
			}
		case "delete":
			if len(parts) < 2 {
				log.Error("no arguments provided to 'delete'")
				continue
			}
			switch parts[1] {
			case "collector":
				if len(parts) < 3 {
					log.Error("no name argument provided to 'delete collector'")
					continue
				}
				deleteCollector(parts[2])
			case "user":
				if len(parts) < 3 {
					log.Error("no name argument provided to 'delete user'")
					continue
				}
				deleteUser(parts[2])
			default:
				fmt.Println("unknown argument to 'delete'")
			}
		case "stream":
			if len(parts) < 2 {
				log.Error("no arguments provided to 'delete'")
				continue
			}
			switch parts[1] {
			case "visualizer":
				streamVisualEvents()
			default:
				fmt.Println("unknown argument to 'visualize'")
			}
		case "help":
			printHelp()
		case "exit":
			return
		default:
			log.Error("command not recognized, try 'help'")
		}
	}
}

func printHelp() {
	fmt.Println("\nAvailable commands:\n")
	fmt.Println("\thelp")
	fmt.Println("\texit\n")
	fmt.Println("\tnew")
	fmt.Println("\t\tcollector <name>")
	fmt.Println("\t\tuser <username> <password>")
	fmt.Println("\tget")
	fmt.Println("\t\tcollectors")
	fmt.Println("\t\tusers")
	fmt.Println("\t\ttls")
	fmt.Println("\tupdate")
	fmt.Println("\t\tpassword <old> <new>")
	fmt.Println("\t\tassignpassword <username> <password>")
	fmt.Println("\t\tusername <name>")
	fmt.Println("\t\ttls <cert file> <key file>")
	fmt.Println("\tdelete")
	fmt.Println("\t\tcollector <name>")
	fmt.Println("\t\tuser <username>")
	fmt.Println("\tstream")
	fmt.Println("\t\tvisualizer\n")
}
