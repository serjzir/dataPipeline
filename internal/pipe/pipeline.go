package pipe

import (
	"bufio"
	"dataPipeline/internal/ringBuf"
	"dataPipeline/pkg/logging"
	"os"
	"strconv"
	"strings"
	"time"
)


func Read(inputPipe chan <- int, done chan bool) {
	logger := logging.Init()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		if strings.EqualFold(data, "exit") {
			logger.Info("Программа завершила работу")
			close(done)
			return
		}
		i, err := strconv.Atoi(data)
		if err != nil {
			logger.Info("Программа работает только с числами")
			continue
		}
		inputPipe <- i
	}
}
func Negative(previousChannel <-chan int, next chan <- int, done <- chan bool) {
	for {
		select {
		case data := <- previousChannel :
			if data > 0 {
				next <- data
			}
		case <- done:
			return
		}
	}
}

func NotAMultipleOfThree(previousChannel <-chan int, next chan <- int, done <- chan bool) {
	for {
		select {
		case data := <- previousChannel :
			if data % 3 == 0 {
				next <- data
			}
		case <- done:
			return
		}
	}
}

func BufferStage(previousChannel <-chan int, next chan <- int, done <- chan bool, buff *ringBuf.IntRingBuff, interval time.Duration) {
	for {
		select {
		case data := <- previousChannel:
			buff.Push(data)
		case <- time.After(interval):
			dataInBuffer := buff.Get()
			if dataInBuffer != nil {
				for _, data := range dataInBuffer {
					next <- data
				}
			}
		case <-done:
			return
		}
	}
}