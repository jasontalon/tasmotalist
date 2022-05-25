package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func FindTasmotaDevices(hosts []Host) (tasmotalist []TasmotaInfo) {
	for _, host := range hosts {
		var tasmota TasmotaInfo

		url := fmt.Sprintf("http://%s/cm?cmnd=status%s", host.Address.Addr, "%200")

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

		if tasmota.Warning != "" {
			tasmota.StatusNET.IPAddress = host.Address.Addr
			tasmota.Status.DeviceName = "Secured. (Requires password)"
			tasmota.StatusFWR.Version = "?"
		}

		tasmotalist = append(tasmotalist, tasmota)
	}

	return tasmotalist
}

func FindPotentialHosts(hosts []Host) []Host {
	return Filter[Host](hosts, func(host Host, _ int) bool {
		if host.Status.State != "up" && strings.Contains(host.Status.Reason, "refused") {
			return false
		}

		_, portFound := Find[Port](host.Ports.Port, func(port Port) bool {
			return port.State.State == "open" && port.Portid == "80"
		})

		return portFound
	})
}

func CheckNmap() (string, error) {
	cmd := "which"

	if runtime.GOOS == "windows" {
		cmd = "where"
	}

	output, err := ExecuteCmd(cmd, "nmap")

	if err != nil && output == "" {
		return "", err
	}

	if strings.Contains(output, "not found") || strings.Contains(output, "Could not find") {
		return "", errors.New("nmap not found")
	}

	output, err = ExecuteCmd("nmap", "-V")
	return output, nil
}
func GetNmapOutput(filePath string) (output NmapOutput, err error) {
	xmlFile, err := os.Open(filePath)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(xmlFile)

	if err != nil {
		return output, err
	}

	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		return output, err
	}

	err = xml.Unmarshal(byteValue, &output)

	if err != nil {
		return output, err
	}

	return output, nil
}

func ExecuteCmd(cmd string, args ...string) (output string, err error) {
	proc := exec.Command(cmd, args...)

	reader, err := proc.StdoutPipe()
	errreader, err := proc.StderrPipe()
	if err != nil {
		return "", nil
	}

	scanner := bufio.NewScanner(reader)
	errscanner := bufio.NewScanner(errreader)

	go func() {
		for errscanner.Scan() {
			output += errscanner.Text() + "\n"
		}
	}()

	go func() {
		for scanner.Scan() {
			output += scanner.Text() + "\n"
		}
	}()

	err = proc.Start()

	if err != nil {
		return output, err
	}

	err = proc.Wait()

	time.Sleep(time.Second * 1)

	if err != nil {
		return output, err
	}

	return output, nil
}

func GetActiveInterfaces() ([]net.Interface, error) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}
	var validinterfaces []net.Interface
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()

		if err != nil {
			return nil, err
		}

		if !(iface.Flags&net.FlagUp > 0) {
			continue
		}

		ipv4 := FindPrimaryAddress(addrs)

		if ipv4 != nil {
			validinterfaces = append(validinterfaces, iface)
		}
	}
	return validinterfaces, nil
}

func FindPrimaryAddresses(ifaces []net.Interface) (addrs []net.IP) {
	for _, i := range ifaces {
		a, _ := i.Addrs()
		addrs = append(addrs, FindPrimaryAddress(a))
	}
	return addrs
}

func FindPrimaryAddress(addrs []net.Addr) net.IP {
	for _, addr := range addrs {
		ipStr := addr.String()
		i := strings.LastIndex(ipStr, "/")
		if i < 0 {
			continue
		}
		ipStr = ipStr[0:i]
		ip := net.ParseIP(ipStr)
		if ip == nil {
			continue
		}
		ipv4 := ip.To4()
		if ipv4 == nil {
			continue
		}
		return ipv4
	}
	return nil
}

func Filter[V any](collection []V, predicate func(V, int) bool) []V {
	result := []V{}

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var result T
	return result, false
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
