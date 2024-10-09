package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
)

func makeResponse() string {
	response := "Hello World!\n\n"
	trueWords := []string{"true", "yes", "da"}
	debugFlag := os.Getenv("MY_SPECIAL_DEBUG_VARIABLE")
	version := os.Getenv("MY_SPECIAL_VERSION")
	if slices.Contains(trueWords, strings.ToLower(debugFlag)) {
		hostname := getHostname()
		ipAddresses, err := getLocalIPAddresses()
		if err != nil {
			return response + "Error getting debug info\n" + err.Error()
		}
		return response + fmt.Sprintf("DEBUG INFO:\nRunning on host: %s\nVersion: %s\nLocal IP addresses:\n%s", hostname, version, strings.Join(ipAddresses, "\n"))
	}
	return response

}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func handler(w http.ResponseWriter, _ *http.Request) {
	response := makeResponse()
	_, err := fmt.Fprintf(w, response)
	if err != nil {
		log.Fatal(err)
	}
}

func getLocalIPAddresses() ([]string, error) {
	var LocalIPAddresses []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			LocalIPAddresses = append(LocalIPAddresses, ip.String())
		}
	}
	return LocalIPAddresses, nil
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
