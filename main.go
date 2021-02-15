package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/fatih/color"
	"github.com/likexian/whois-go"
	"github.com/pwnscan/server"
)

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	//colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	//colorWhite := "\033[37m"
	// clears the terminal
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	// start
	var banner = `
 ______        ___   _ ____   ____    _    _   _
|  _ \ \      / | \ | / ___| / ___|  / \  | \ | |
| |_) \ \ /\ / /|  \| \___ \| |     / _ \ |  \| |
|  __/ \ V  V / | |\  |___) | |___ / ___ \| |\  |
|_|     \_/\_/  |_| \_|____/ \____/_/   \_|_| \_|
						
			    version 1.1
	`
	color.Red("%s", banner)
	fmt.Println(string(colorBlue), "          made with", string(colorRed), "<3", string(colorBlue), "by @nolimitcarter")
	color.White("                   ")
	url := flag.String("url", "u", "-url")
	server := flag.String("server", "u", "-server")
	flag.Parse()
	if *url != "u" {
		fmt.Println(string(colorRed), "[-]", string(colorGreen), "Entered URL:", string(colorCyan), *url)
		color.White("                   ")
		fmt.Println(string(colorRed), "[:]", string(colorYellow), "PWNScan Options", string(colorRed), "[:]")
		color.White("                   ")
		fmt.Println(string(colorRed), "[01]", string(colorYellow), "Simple Nmap Scan")
		fmt.Println(string(colorRed), "[02]", string(colorYellow), "Advanced Nmap Scan")
		fmt.Println(string(colorRed), "[03]", string(colorYellow), "Who-is Lookup")
		fmt.Println(string(colorRed), "[04]", string(colorYellow), "DNS Lookup")
		fmt.Println(string(colorRed), "[05]", string(colorYellow), "Reverse DNS Lookup")
		fmt.Println(string(colorRed), "[06]", string(colorYellow), "ETC")
		fmt.Println(string(colorRed), "[07]", string(colorYellow), "ETC")
		fmt.Println(string(colorRed), "[08]", string(colorYellow), "ETC")
		fmt.Println(string(colorRed), "[09]", string(colorYellow), "ETC")
		fmt.Println(string(colorRed), "[10]", string(colorYellow), "Simple Web Server")
		color.White("                   ")
		fmt.Println(string(colorRed), "[00]", string(colorYellow), "About", string(colorRed), "[99]", string(colorYellow), "Exit")
		color.White("                   ")

		var ch int
		fmt.Println(string(colorRed), "[-]", string(colorGreen), "Select an option:")
		fmt.Scanf("%d", &ch)
		if ch == 1 {
			simplenmap(*url)
		} else if ch == 2 {
			advancednmap(*url)
		} else if ch == 3 {
			whoislookup(*url)
		} else if ch == 10 {
			rungoserver()
			//server.Goserver()
		} else if ch == 00 {
			about()
		} else if ch == 99 {
			exit()
		}
	} else {
		fmt.Println(string(colorRed), "Error : missing parameter '-url, --url'")
		fmt.Println(string(colorGreen), "Usage : go run main.go -url 'google.com'")
		fmt.Println(string(colorBlue), "(Ignore if running server)")
		foo()
		if err := foo(); err != nil {
		}
	}
	//still a WIP
	if *server != "s" {
		rungoserver()
	}

}

func simplenmap(name string) {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int, name string) {
			defer wg.Done()
			address := fmt.Sprintf("%v:%d", name, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				//port is closed/filtered
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i, name)
	}
	wg.Wait()
}

func advancednmap(name string) {
	url := "https://api.hackertarget.com/nmap/?q=" + name
	resp, err := http.Get(url)
	if err != nil {
		foo()
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", html)
}

func whoislookup(name string) {
	result, err := whois.Whois(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func rungoserver() {
	server.Goserver()
}

func reversedns() {

}

func about() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	//colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	//colorWhite := "\033[37m"
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	var banner = `
 ______        ___   _ ____   ____    _    _   _
|  _ \ \      / | \ | / ___| / ___|  / \  | \ | |
| |_) \ \ /\ / /|  \| \___ \| |     / _ \ |  \| |
|  __/ \ V  V / | |\  |___) | |___ / ___ \| |\  |
|_|     \_/\_/  |_| \_|____/ \____/_/   \_|_| \_|
						
			    version 1.1
	`
	color.Red("%s", banner)
	fmt.Println(string(colorBlue), "          made with", string(colorRed), "<3", string(colorBlue), "by @nolimitcarter")
	color.White("                   ")
	fmt.Println(string(colorRed), "[-]", string(colorGreen), "Tool Created by @nolimitcarter")
	color.White("                   ")
	fmt.Println(string(colorCyan), "Author  :  @nolimitcarter")
	fmt.Println(string(colorCyan), "Github  :  https://github.com/nolimitcarter")
	fmt.Println(string(colorCyan), "Find me :  https://cartert.dev")
	fmt.Println(string(colorCyan), "Version :  1.1")
	color.White("                   ")
	fmt.Println(string(colorRed), "[00]", string(colorYellow), "Exit")
	color.White("                   ")
	var ch int
	fmt.Println(string(colorRed), "[-]", string(colorGreen), "Select an option : ")
	fmt.Scanf("%d", &ch)
	if ch == 00 {
		exit()
	} else {
		foo()
	}
}

func exit() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	//colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	//colorPurple := "\033[35m"
	//colorCyan := "\033[36m"
	//colorWhite := "\033[37m"
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	var banner = `
 ______        ___   _ ____   ____    _    _   _
|  _ \ \      / | \ | / ___| / ___|  / \  | \ | |
| |_) \ \ /\ / /|  \| \___ \| |     / _ \ |  \| |
|  __/ \ V  V / | |\  |___) | |___ / ___ \| |\  |
|_|     \_/\_/  |_| \_|____/ \____/_/   \_|_| \_|
						
			    version 1.1
	`
	color.Red("%s", banner)
	fmt.Println(string(colorBlue), "          made with", string(colorRed), "<3", string(colorBlue), "by @nolimitcarter")
	color.White("                 ")
	fmt.Println(string(colorRed), "[-]", string(colorGreen), "Tool Created by @nolimitcarter")
	color.White("                   ")
	fmt.Println(string(colorRed), "Exiting... Thanks for using this tool!")
	color.White("                   ")

	defer fmt.Println("!")
	os.Exit(0)
}

func foo() error {
	return errors.New("Some error occured")
}
