package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

func main() {

	if !CheckNmap() {
		panic("nmap not installed")
	}

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	nmapFile := "nmap.xml"

	nmapFullPath := path.Join(wd, nmapFile)

	currentIp := GetOutboundIP().String()

	ipRange := currentIp[0:strings.LastIndex(currentIp, ".")] + ".0-255"

	spin := spinner.New(spinner.CharSets[2], 100*time.Millisecond)

	spin.Prefix = "Scanning network " + ipRange + " "

	spin.Start()

	_, err = ExecuteCmd("nmap", "--script-timeout", "2", "-d", "-oX", nmapFile, "-p", "80", ipRange)
	spin.Stop()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(nmapFullPath); err != nil {
		panic(err)
	}

	spin.Prefix = "Looking for tasmota devices... "

	spin.Start()

	nmapOutput, err := GetNmapOutput(nmapFullPath)

	if err != nil {
		panic(err)
	}

	devices := FindTasmotaDevices(FindPotentialHosts(nmapOutput.Hosts))

	ResolveTasmotaDeviceName(devices)

	spin.Stop()

	if len(devices) == 0 {
		fmt.Println("nothing found.")
		return
	}

	results := ""

	for _, device := range devices {
		results += fmt.Sprintf("%-15s %s\n", device.Address.Addr, device.Hostnames.Hostname.Name)
	}

	fmt.Println(results)
}
