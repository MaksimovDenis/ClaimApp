package main

import (
	"ClaimApp/coordinates"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {

	fmt.Println("Please input 1, if you need to set coordinates, 2 - if coordinates settings already installed")

	reader := bufio.NewReader(os.Stdin)
	choiceStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Invalid argument %v", err)
	}

	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		log.Fatalf("Invalid argument %v", err)
	}

	if choice == 1 {
		fmt.Println("Please input the amount of actions for your mouse")

		reader := bufio.NewReader(os.Stdin)
		amoutStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Invalid argument %v", err)
		}

		amoutStr = strings.TrimSpace(amoutStr)
		amount, err := strconv.Atoi(amoutStr)
		if err != nil {
			log.Fatalf("Invalid argument %v", err)
		}

		coordinates.SetCoordinates(amount)

	} else if choice == 2 {
		coord, err := coordinates.LoadConfig("config.json")
		if err != nil {
			log.Fatalf("Failed to load config file, please set the coordinates: %v", err)
		}

		count := len(coord.ArrX)

		fmt.Println(`Please input the count of your telegram accounts: `)
		reader = bufio.NewReader(os.Stdin)
		countOfAppsStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Invalid argument %v", err)
		}
		countOfAppsStr = strings.TrimSpace(countOfAppsStr)

		countOfApps, err := strconv.Atoi(countOfAppsStr)
		if err != nil {
			log.Fatalf("Invalid argument %v", err)
		}

		path := `C:\Users\User\Desktop\Farm`

		for i := 0; i < countOfApps; i++ {
			cmd := exec.Command(path + `\` + strconv.Itoa(i+1) + `\` + strconv.Itoa(i+1))
			err := cmd.Start()
			if err != nil {
				log.Fatalf("Failed to start telegram app №: %v", i+1)
			}

			time.Sleep(10 * time.Second)
			log.Printf("Telegram №%v is running\n", i+1)

			for j := 0; j < count; j++ {
				robotgo.Move(coord.ArrX[j], coord.ArrY[j])
				robotgo.Click()
				time.Sleep(time.Second * 2)
				robotgo.Click()
				time.Sleep(time.Duration(coord.Delay) * time.Second)
			}

			err = cmd.Process.Kill()
			if err != nil {
				log.Printf("Failed to close telegram app №%v\n", i+1)
			}
			log.Printf("Telegram app №%v has been closed\n", i+1)
		}

	} else {
		log.Fatal("Invalid argument")
		os.Exit(0)
	}

	err = shutdown()
	if err != nil {
		fmt.Println("Error shutting down the system:", err)
	}
}

func shutdown() error {
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	return cmd.Run()
}
