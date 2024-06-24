package coordinates

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

type Config struct {
	ArrX  []int
	ArrY  []int
	Delay int
}

func SetCoordinates(count int) {
	var arrX, arrY []int

	fmt.Println("-------------!Setting of coordinates is strting now!-------------")
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	for i := 0; i < count; i++ {
		fmt.Printf("Setting of ACTION - %v, please locate your mouse in correct place\n", i+1)

		for j := 5; j >= 0; j-- {
			fmt.Println(j)
			time.Sleep(1 * time.Second)
		}

		x, y := robotgo.Location()
		arrX = append(arrX, x)
		arrY = append(arrY, y)

		fmt.Printf("ACTION - %v DONE with coordinates x=%v; y=%v\n", i+1, x, y)
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Please specify the delay between actions (in seconds)")
	reader := bufio.NewReader(os.Stdin)
	delayStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Invalid argument: %v", err)
	}

	delayStr = strings.TrimSpace(delayStr)
	delay, err := strconv.Atoi(delayStr)
	if err != nil {
		log.Fatalf("Invalid argument: %v", err)
	}

	config := Config{
		ArrX:  arrX,
		ArrY:  arrY,
		Delay: delay,
	}

	err = saveConfig(config, "config.json")
	if err != nil {
		fmt.Printf("Failed to save config file: %v\n", err)
		return
	}
}

func saveConfig(config Config, fileName string) error {
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func LoadConfig(fileName string) (Config, error) {
	var config Config
	data, err := os.ReadFile(fileName)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	return config, err
}
