package main

import (
	"fmt"
	"io/ioutil"
	"github.com/getlantern/systray"
	"time"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	country:="indonesia"
	systray.SetIcon(getIcon("assets/corona.ico"))
	coronaID := systray.AddMenuItem("Corona Indonesia", "Corona Indonesia")
	coronaItaly := systray.AddMenuItem("Corona Italy", "Corona Italy")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			systray.SetTitle(corona(country))
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-coronaID.ClickedCh:
				country= "indonesia"
			case <-coronaItaly.ClickedCh:
				country= "italy"
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleaning stuff here.
}


func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}

