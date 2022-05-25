package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

var spin = spinner.New(spinner.CharSets[2], 100*time.Millisecond)

func standby() {
	fmt.Println("Press Enter Key to exit.")
	fmt.Scanln()
}

func main() {
	spin.Prefix = "Performing nmap check "

	spin.Start()
	nmapV, err := CheckNmap()
	spin.Stop()
	if err != nil {
		fmt.Println("Nmap is required. Please install by visiting the link below.")
		fmt.Println()
		if runtime.GOOS == "windows" {
			fmt.Println("Official Website: https://nmap.org/download#windows")
			fmt.Println("Chocolatey: https://community.chocolatey.org/packages/nmap#install")
		}
		if runtime.GOOS == "darwin" {
			fmt.Println("Official Website: https://nmap.org/download.html#macosx")
			fmt.Println("Homebrew: https://formulae.brew.sh/formula/nmap")
		}
		fmt.Println()
		standby()
		return
	}

	fmt.Println(nmapV)

	nmapFile := "nmap.xml"

	nmapFile = filepath.Join(os.TempDir(), nmapFile)

	currentIp, err := GetOutboundIP()

	if err != nil {
		fmt.Println(err)
		standby()
		return
	}

	currentIpStr := currentIp.String()

	ipRange := currentIpStr[0:strings.LastIndex(currentIpStr, ".")] + ".0-255"

	spin.Prefix = "Scanning network " + ipRange + " "

	spin.Start()

	b, err := ExecuteCmd("nmap", "--unprivileged", "--script-timeout", "2", "-d", "-oX", nmapFile, "-p", "80", ipRange)
	spin.Stop()
	if err != nil {
		fmt.Println(b, err)
		standby()
		return
	}

	if _, err := os.Stat(nmapFile); err != nil {
		fmt.Println(err)
		standby()
		return
	}

	spin.Prefix = "Looking for tasmota devices... "

	spin.Start()

	nmapOutput, err := GetNmapOutput(nmapFile)

	if err != nil {
		fmt.Println(err)
		standby()
		return
	}

	tasmotalist := FindTasmotaDevices(FindPotentialHosts(nmapOutput.Hosts))

	spin.Stop()

	if len(tasmotalist) == 0 {
		fmt.Println("nothing found.")
		standby()
		return
	}

	sort.SliceStable(tasmotalist, func(i, j int) bool {
		return tasmotalist[i].Status.DeviceName > tasmotalist[j].Status.DeviceName
	})

	//find maximum length of device name
	var (
		namelens []int
		verlens  []int
	)

	for _, tasmota := range tasmotalist {
		namelens = append(namelens, len(tasmota.Status.DeviceName))
		verlens = append(verlens, len(tasmota.StatusFWR.Version))
	}

	sort.Ints(namelens)
	sort.Sort(sort.Reverse(sort.IntSlice(namelens)))
	sort.Ints(verlens)
	sort.Sort(sort.Reverse(sort.IntSlice(verlens)))

	padding := []int{len("255.255.255.255"), namelens[0], verlens[0]}

	results := fmt.Sprintf("%-*s %-*s %-*s\n", padding[0], "IP Address", padding[1], "Device Name", padding[2], "Version")

	for _, tasmota := range tasmotalist {
		results += fmt.Sprintf("%-*s %-*s %-*s\n", padding[0], tasmota.StatusNET.IPAddress, padding[1], tasmota.Status.DeviceName, padding[2], tasmota.StatusFWR.Version)
	}

	fmt.Println(results)

	fmt.Println("done.")

	standby()
}
