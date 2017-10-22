package main

import (
	"os/exec"
	"fmt"
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
		cmd := exec.Command("echo", "hi");
		output, err := cmd.Output();
		fmt.Println(output);
	}
}