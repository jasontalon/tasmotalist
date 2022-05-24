package main

import (
	"bufio"
	"encoding/xml"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func ResolveTasmotaDeviceName(hosts []Host) {
	for i := range hosts {
		res, err := http.Get("http://" + hosts[i].Address.Addr)

		if err != nil {
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			continue
		}

		title := doc.Find("title").Text()

		title = strings.ReplaceAll(title, " - Main Menu", "")

		hosts[i].Hostnames.Hostname.Name = title
	}
}

func FindTasmotaDevices(hosts []Host) []Host {
	return Filter[Host](hosts, func(host Host, _ int) bool {
		res, err := http.Get("http://" + host.Address.Addr)

		if err != nil {
			return false
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			return false
		}

		found := false

		doc.Find("a").Each(func(_ int, selection *goquery.Selection) {
			if found {
				return
			}

			content := selection.Text()

			found = strings.HasPrefix(content, "Tasmota") && strings.HasSuffix(content, "by Theo Arends")
		})

		return found
	})

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

func CheckNmap() bool {
	output, err := ExecuteCmd("which", "nmap")

	if err != nil {
		return false
	}

	if strings.Contains(output, "not found") {
		return false
	}

	return true
}
func GetNmapOutput(filePath string) (output NmapOutput, err error) {
	xmlFile, err := os.Open(filePath)
	defer xmlFile.Close()
	if err != nil {
		return output, err
	}

	byteValue, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		return output, err
	}

	xml.Unmarshal(byteValue, &output)

	return output, nil
}

func ExecuteCmd(cmd string, args ...string) (string, error) {
	proc := exec.Command(cmd, args...)

	reader, err := proc.StdoutPipe()

	if err != nil {
		return "", nil
	}

	scanner := bufio.NewScanner(reader)

	var output string

	go func() {
		for scanner.Scan() {
			output += scanner.Text() + "\n"
		}
	}()

	err = proc.Start()

	if err != nil {
		return "", err
	}

	err = proc.Wait()

	if err != nil {
		return "", err
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

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
