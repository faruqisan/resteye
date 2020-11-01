package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {

	sleepInterval := 60 * 20 // remind every 20minutes

	tick := time.NewTicker(1 * time.Second)

	notifyChan := make(chan interface{})
	resumeChan := make(chan interface{})

	reader := bufio.NewReader(os.Stdin)

	ticked := 0
	go func() {
		for range tick.C {

			if ticked != sleepInterval {
				ticked++
				continue
			}

			err := notify("Get some rest, look away from the screen")
			if err != nil {
				log.Fatal(err)
			}
			notifyChan <- nil
			<-resumeChan
			ticked = 0
		}
	}()

	for {
		select {
		case <-notifyChan:
			fmt.Print("Press anything for resuming: ")
			reader.ReadString('\n')
			resumeChan <- nil
		}
	}

}

func notify(msg string) error {
	return beeep.Notify("RestEye!", msg, "")
}
