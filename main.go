package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	const searchPath = "/sys/class/power_supply"
	var re = regexp.MustCompile(`(?ms)BAT[0-9]+`)
	var batteries []string
	var totalPower int
	var power int

	files, err := ioutil.ReadDir(searchPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if re.Match([]byte(f.Name())) {
			log.Println("Found Battery:", f.Name())
			batteries = append(batteries, f.Name())
		}
	}

	for _, bat := range batteries {
		powerBytes, err := ioutil.ReadFile(filepath.Join(searchPath, bat, "power_now"))
		if err != nil {
			log.Fatal(err)
		}
		power, err = strconv.Atoi(strings.Trim(string(powerBytes), "\n"))
		if err != nil {
			log.Fatal(err)
		}
		totalPower += power
	}
	log.Println("total power consumption:", totalPower)

}
