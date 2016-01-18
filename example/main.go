package main

import (
	"fmt"
	"github.com/alyu/configparser"
	"github.com/xyproto/jpath"
	"github.com/xyproto/mattersend"
	"log"
)

func main() {
	// Read the configuration from hook.conf
	config, err := configparser.Read("hook.conf")
	if err != nil {
		log.Fatalln(err)
	}
	section, err := config.Section("host")
	if err != nil {
		log.Fatalln(err)
	}
	urlStr := section.ValueOf("url")

	// Create a JSON node tree
	rootNode := jpath.NewNode()
	rootNode.Set("text", `---
##### Build Break - Project X - December 12, 2015 - 15:32 GMT +0
| Component  | Tests Run   | Tests Failed                                   |
|:-----------|:------------|:-----------------------------------------------|
| Server     | 948         | :ghost: 0                           |
| Web Client | 123         | :warning: [2 (see details)](http://linktologs) |
| iOS Client | 78          | :poop: [3 (see details)](http://linktologs) |
---`)
	rootNode.Set("username", "HookBot9000")

	// Send the message to the Mattermost host
	if status, err := mattersend.Send(urlStr, rootNode); err != nil || status != "200 OK" {
		log.Fatalln("FAIL", err, status)
	}
	fmt.Println("SUCCESS")
}
