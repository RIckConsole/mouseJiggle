package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/integrii/flaggy"
)

var version = "1.0"
var random = false
var interval = 1

func init() {
	flaggy.SetName("Mouse Jiggle")
	flaggy.SetDescription("Hands getting tired from waking up your work PC?")
	flaggy.SetVersion(version)
	flaggy.DefaultParser.ShowHelpOnUnexpected = true

	flaggy.Bool(&random, "r", "random", "Enable random mouse movements (default is slight jitter)")
	flaggy.Int(&interval, "i", "interval", "Set the time (in seconds) between mouse movements")

	flaggy.Parse()
}

func main() {
	fmt.Println("[*] Engaging Mouse Movements. Press CTRL+C to stop.")
	if random {
		i := 1
		for i == 1 {
			//Get Screen Size
			sx, sy := robotgo.GetScreenSize()
			num := rand.Intn(400)
			robotgo.MoveSmoothRelative(rand.Intn(num-5), rand.Intn(num-5), 2)
			time.Sleep(time.Duration(interval) * time.Second / 2)
			robotgo.MoveSmoothRelative(num-num*2, num-num*2, 2)
			x, y := robotgo.GetMousePos()
			if x == 0 || y == 0 {
				robotgo.MoveSmoothRelative(sx/2, sy/2)
			} else if x == sx || y == sy {
				robotgo.MoveSmoothRelative(sx/2, sy/2)
			}
		}

	}
	i := 1
	for i == 1 {
		robotgo.MoveSmoothRelative(5, 0)
		time.Sleep(time.Duration(interval) * time.Second / 2)
		robotgo.MoveSmoothRelative(-5, 0)
	}
}
