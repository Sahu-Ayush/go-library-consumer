package main

import (
	"fmt"

	"github.com/sahu-ayush/go-custom-libraries/systeminfo"
)

func main() {
	hostname, err := systeminfo.GetHostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	fmt.Println("System Information")
	fmt.Println("------------------")
	fmt.Println("Hostname:", hostname)
	fmt.Println("Operating System:", systeminfo.GetOS())
	fmt.Println("System Summary:", systeminfo.GetSystemSummary())
	fmt.Println("Architecture:", systeminfo.GetArchitecture())
	fmt.Println("CPU Count:", systeminfo.GetCPUCount())
	fmt.Println("Go Version:", systeminfo.GetGoVersion())
}
