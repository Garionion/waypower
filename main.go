package main

import (
	"fmt"
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
	var totalPower int
	var power int

	files, err := ioutil.ReadDir(searchPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if re.Match([]byte(f.Name())) {
			powerBytes, err := ioutil.ReadFile(filepath.Join(searchPath, f.Name(), "power_now"))
			if err != nil {
				log.Fatal(err)
			}
			power, err = strconv.Atoi(strings.Trim(string(powerBytes), "\n"))
			if err != nil {
				log.Fatal(err)
			}
			totalPower += power
		}
	}

	//divide trough 1.000.000 to get power in Watts
	totalPower = totalPower / 1000000

	fmt.Println(totalPower, "W")

}
