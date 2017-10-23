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

	for i := 0; i < len(websites); i++ {
		cmd := exec.Command("ping", "-c", "10", websites[i])
		output, err := cmd.Output()
		out := string(output[:])
		if err == nil {
			split := strings.Split(out, "\n")

			last := split[len(split)-2]

			stats := strings.Split(strings.Split(last, "=")[1], " ")
			fmt.Println(stats)

		}
	}
}
