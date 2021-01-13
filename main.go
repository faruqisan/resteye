package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
	"github.com/gosuri/uiprogress"
)

func main() {

	printHeader()

	// change the interval, because ticker in second, we times the interval by 60
	sleepInterval := 60 * 20 // remind every 20minutes

	tick := time.NewTicker(1 * time.Second)

	notifyChan := make(chan interface{})
	resumeChan := make(chan interface{})

	ticked := 0
	uiprogress.Start() // start rendering
	bar := uiprogress.AddBar(sleepInterval).AppendCompleted().PrependElapsed()

	go func() {
		for range tick.C {

			if ticked != sleepInterval {
				ticked++
				bar.Incr()
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
			answer, err := dlgs.Question("RestEye", "Do you want to resuming screen timer?", true)
			if err != nil {
				log.Fatal(err)
			}
			if !answer {
				log.Fatal("terminating")
			}
			bar = uiprogress.AddBar(sleepInterval).AppendCompleted().PrependElapsed()
			resumeChan <- nil
		}
	}

}

func notify(msg string) error {
	return beeep.Notify("RestEye!", msg, "")
}

func printHeader() {
	header := `

	____     ___  _____ ______    ___  __ __    ___ 
	|    \   /  _]/ ___/|      |  /  _]|  |  |  /  _]
	|  D  ) /  [_(   \_ |      | /  [_ |  |  | /  [_ 
	|    / |    _]\__  ||_|  |_||    _]|  ~  ||    _]
	|    \ |   [_ /  \ |  |  |  |   [_ |___, ||   [_ 
	|  .  \|     |\    |  |  |  |     ||     ||     |
	|__|\_||_____| \___|  |__|  |_____||____/ |_____|
													 
	
	Starting timer

	`
	fmt.Println(header)
}
