package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func Log(a string) {
	log.Println("[ValorantObliterator]", a)
}

func main() {
	Log("This program is meant to remove all Riot Games/Valorant Data")
	Log("Please exit if you do not want this to happen")
	Log("Starting in 5 seconds...")
	time.Sleep(time.Second * 5)
	var PATHS []string = make([]string, 0)
	var LocalAppData string = os.Getenv("localappdata")
	var ProgramData string = os.Getenv("programdata")
	PATHS = append(PATHS,
		"C:\\Riot Games",
		path.Join(LocalAppData, "Riot Games"),
		path.Join(ProgramData, "Riot Games"),
	)

	for _, path := range PATHS {
		Log(fmt.Sprintf("Removing \"%v\"", path))
		err := os.RemoveAll(path)
		if err != nil {
			Log(fmt.Sprintf("Error removing \"%v\"", path))
			Log(err.Error())
		}
	}
	Log("Done!")
	Log("Closing in 15 seconds...")
	time.Sleep(time.Second * 15)
}
