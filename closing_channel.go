package main

func countReports(numSentCh chan int) int {
	sum := 0

	for {
		v, ok := <-numSentCh
		if ok {
			sum += v
		} else {
			break
		}
	}
	return sum
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
