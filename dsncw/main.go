package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	websites := [...]string{
		"google.com",
		"google.co.uk",
		"google.com.au",
		"youtube.com",
		"instagram.com",
		"yandex.ru",
		"wikipedia.org",
		"stackoverflow.com",
		"linkedin.com",
		"yahoo.com",
		"yahoo.co.uk",
		"yahoo.com.au",
		"facebook.com",
		"netflix.com",
		"ebay.co.uk",
		"bbc.co.uk",
		"www.bbc.co.uk",
		"reddit.com",
		"github.com",
		"taobao.com",
		"mail.ru",
		"wikia.com",
		"loopsofzen.uk",
		"iinet.net.au"}

	fmt.Println("Host, IPv4, Avg RTT, Min RTT, Max RTT, IPv6, Avg RTT, Min RTT, Max RTT")

	for i := 0; i < len(websites); i++ {
		cmd := exec.Command("ping", "-4", "-c", "10", websites[i])
		output, err := cmd.Output()
		out := string(output[:])

		cmd6 := exec.Command("ping", "-6", "-c", "10", websites[i])
		output6, err6 := cmd6.Output()
		out6 := string(output6[:])
		if err == nil && err6 == nil {
			split := strings.Split(out, "\n")

			first := split[0]
			ip := strings.Split(strings.Split(first, "(")[1], ")")[0]
			last := split[len(split)-2]

			stats := strings.Split(strings.Split(strings.Split(last, "=")[1], " ")[1], "/")

			split6 := strings.Split(out6, "\n")

			first6 := split6[1]
			ip6 := strings.Split(strings.Split(first6, "(")[1], ")")[0]
			last6 := split6[len(split6)-2]

			stats6 := strings.Split(strings.Split(strings.Split(last6, "=")[1], " ")[1], "/")

			fmt.Println(websites[i] + "," + ip + "," + stats[1] + "," + stats[0] + "," + stats[2] + "," + ip6 + "," + stats6[1] + "," + stats6[0] + "," + stats6[2])

		}
	}
}
