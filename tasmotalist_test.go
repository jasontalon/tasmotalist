package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckForTasmotaDevices(t *testing.T) {
	filename := "nmap.xml"

	wd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	filePath := filepath.Join(wd, filename)

	nmap, err := GetNmapOutput(filePath)

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	devices := FindTasmotaDevices(FindPotentialHosts(nmap.Hosts))

	ResolveTasmotaDeviceName(devices)

	for _, d := range devices {
		fmt.Println(d.Hostnames.Hostname.Name + " " + d.Address.Addr)
	}

}
func TestFindDevices(t *testing.T) {
	filename := "nmap.xml"

	wd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	filePath := filepath.Join(wd, filename)

	nmap, err := GetNmapOutput(filePath)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(nmap)

	validHosts := FindPotentialHosts(nmap.Hosts)

	fmt.Println(validHosts)
}

func TestCheckNmap(t *testing.T) {
	err := CheckNmap()

	fmt.Println(err)
}
func TestGetNmapOutput(t *testing.T) {
	filename := "nmap.xml"
	_, err := ExecuteCmd("nmap", "-oX", filename, "-p", "80", "192.168.192.0-255")

	if err != nil {
		t.Error(err)
	}

	wd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	filePath := filepath.Join(wd, filename)

	nmap, err := GetNmapOutput(filePath)

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	potentialHosts := FindPotentialHosts(nmap.Hosts)
	fmt.Println(potentialHosts)

}

func TestGetActiveInterfaces(t *testing.T) {
	interfaces, err := GetActiveInterfaces()

	if err != nil {
		t.Error(err)
	}

	addrs := FindPrimaryAddresses(interfaces)

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}

func TestExecuteCmd(t *testing.T) {
	output, err := ExecuteCmd("nmap", "-p", "80", "192.168.1.0-255")

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	fmt.Println(output)
}

func TestFlow(t *testing.T) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	nmapFile := "nmap.xml"

	nmapFullPath := path.Join(wd, nmapFile)

	currentIp := GetOutboundIP().String()

	ipRange := currentIp[0:strings.LastIndex(currentIp, ".")] + ".0-255"

	if !CheckNmap() {
		panic("nmap not installed")
	}

	fmt.Println("searching ip range " + ipRange)

	_, err = ExecuteCmd("nmap", "--script-timeout", "2", "-d", "-oX", nmapFile, "-p", "80", ipRange)

	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(nmapFullPath); err != nil {
		panic(err)
	}

	nmapOutput, err := GetNmapOutput(nmapFullPath)

	if err != nil {
		panic(err)
	}

	devices := FindTasmotaDevices(FindPotentialHosts(nmapOutput.Hosts))

	ResolveTasmotaDeviceName(devices)

	t.Log(nmapOutput)
}
