package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

// select {
// case v := <-ch:
//     // use v
// default:
//     // receiving from ch would block
//     // so do something else
// }

// read-only

// func main() {
//     ch := make(chan int)
//     readCh(ch)
// }

// func readCh(ch <-chan int) {
//     // ch can only be read from
//     // in this function
// }

// write-only

// func writeCh(ch chan<- int) {
//     // ch can only be written to
//     // in this function
// }
