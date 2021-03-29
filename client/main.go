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
	fmt.Println("  clt universe  # Display universe_id")
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
	} else if command == "universe" {
		s := network.DoGet(fmt.Sprintf("universe"))
		fmt.Println(s)
	} else if command == "servers" {
		s := network.DoGet("servers")
		fmt.Println(s)
	} else if command == "auth" {
		persist.SaveToFile("USERNAME", cli.ArgMap["name"])
	} else if command == "simulate" {
		people := []string{"bob", "alice", "candy", "mike", "dave", "chris", "pam", "abigail", "emma", "luna",
			"logan", "owen", "liam", "sophia", "santiago", "joe", "dan", "mark", "charles", "kevin",
			"logan2", "owen2", "liam2", "sophia2", "santiago2", "joe2", "dan2", "mark2", "charles2", "kevin2",
			"logan3", "owen3", "liam3", "sophia3", "santiago3", "joe3", "dan3", "mark3", "charles3", "kevin3",
			"logan4", "owen4", "liam4", "sophia4", "santiago4", "joe4", "dan4", "mark4", "charles4", "kevin4",
			"logan5", "owen5", "liam5", "sophia5", "santiago5", "joe5", "dan5", "mark5", "charles5", "kevin5",
			"logan6", "owen6", "liam6", "sophia6"}
		words := []string{"hi there", "ok then"}

		for _, person := range people {
			for _, word := range words {
				network.PostNewTimeline(word, person)
				time.Sleep(time.Millisecond * 20)
			}
		}
	} else if command == "post" {
		network.PostNewTimeline(cli.ArgMap["text"], cli.Username)
	}
}
