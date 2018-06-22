package gocommon

import (
	"os"
	"os/signal"

	log "github.com/lotusdeng/log4go"
)

var GlobalQuitChannel chan os.Signal
var SignalQuitChannel chan os.Signal

//main start call this
func InitAppQuit() {
	GlobalQuitChannel = make(chan os.Signal)
	SignalQuitChannel = make(chan os.Signal, 1)
	signal.Notify(SignalQuitChannel, os.Interrupt)

	//goroutine receive quit signal, to set AppDataSingleton.AppExit

	go func() {
		<-SignalQuitChannel
		if GlobalQuitChannel != nil {
			log.Info("receive Signal Quit， close GlobalQuitChannel")
			close(GlobalQuitChannel)
			GlobalQuitChannel = nil
		}
	}()
}

func SignalAppQuit() {
	SignalQuitChannel <- os.Interrupt
}

func IsAppQuit() bool {
	if GlobalQuitChannel == nil {
		return true
	} else {
		return false
	}
}

func WaitAppQuit() {
	<-GlobalQuitChannel
}

//main end call this
func UinitAppQuit() {
	if GlobalQuitChannel != nil {
		close(GlobalQuitChannel)
		GlobalQuitChannel = nil
	}
	if SignalQuitChannel != nil {
		close(SignalQuitChannel)
		SignalQuitChannel = nil
	}
}
