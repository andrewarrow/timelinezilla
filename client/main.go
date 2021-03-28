package main

import (
	"clt/cli"
	"clt/network"
	"clt/persist"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clt ls        # List recent timelines")
	fmt.Println("  clt profile   # List recent timelines")
	fmt.Println("  clt post      # Post new timeline with --text=hi")
	fmt.Println("  clt auth      # Set your username --name=")
	fmt.Println("  clt servers   # List the main list")
	fmt.Println("  clt simulate  # Simulate traffic")
	fmt.Println("  clt toggle    # Toggle follow")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	persist.Init()
	cli.ReadInGlobalVars()

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "profile" {
		username := cli.Username
		if cli.ArgMap["name"] != "" {
			username = cli.ArgMap["name"]
		}
		s := network.DoGet(fmt.Sprintf("timelines/%s", username))
		//fmt.Println(s)
		network.DisplayProfileTimelines(s)
	} else if command == "toggle" {
		cli.EnsureParamPass("name")
		s := network.DoPost(fmt.Sprintf("follow/%s", cli.ArgMap["name"]), []byte{})
		fmt.Println(s)
	} else if command == "ls" {
		s := network.DoGet(fmt.Sprintf("timelines"))
		//fmt.Println(s)
		network.DisplayInboxTimelines(s)
	} else if command == "servers" {
		s := network.DoGet("servers")
		fmt.Println(s)
	} else if command == "auth" {
		persist.SaveToFile("USERNAME", cli.ArgMap["name"])
	} else if command == "simulate" {
		people := []string{"bob", "alice", "candy", "mike", "dave", "chris", "pam"}
		words := []string{"hey", "hello", "what up?", "hi there", "ok then"}

		for _, person := range people {
			for _, word := range words {
				network.PostNewTimeline(word, person)
			}
		}
	} else if command == "post" {
		network.PostNewTimeline(cli.ArgMap["text"], cli.Username)
	}
}
