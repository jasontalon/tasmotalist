package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestTasmoTaList(t *testing.T) {
	addresses := []string{"192.168.1.32", "192.168.1.26", "192.168.1.28"}
	var tasmotalist []TasmotaInfo
	for _, addr := range addresses {
		var tasmota TasmotaInfo

		url := fmt.Sprintf("http://%s/cm?cmnd=status 0", addr)

		res, err := http.Get(url)

		if err != nil {
			continue
		}

		if !(res.StatusCode == 401 || res.StatusCode == 200) {
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			continue
		}

		content := doc.Text()

		err = json.Unmarshal([]byte(content), &tasmota)

		if err != nil {
			continue
		}

		fmt.Println(tasmota)

		tasmotalist = append(tasmotalist, tasmota)
	}

	fmt.Println(tasmotalist)

}
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

	tasmotalist := FindTasmotaDevices(FindPotentialHosts(nmap.Hosts))

	for _, tasmota := range tasmotalist {
		fmt.Println(tasmota)
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
	_, err := CheckNmap()

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
