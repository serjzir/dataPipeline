package main

import (
	"dataPipeline/internal/config"
	"dataPipeline/internal/pipe"
	"dataPipeline/internal/ringBuf"
	"dataPipeline/pkg/logging"
	"fmt"
	"time"
)



func  main()  {
	cfg := config.GetConfig()
	logger := logging.Init()
	interval := time.Duration(cfg.BufferDrainInterval) * time.Second
	buffer := ringBuf.NewIntRingBuff(cfg.BufferSize)
	logger.Infof("Размер буффера: %d, Продолжительность: %d", cfg.BufferSize, cfg.BufferDrainInterval)
	done := make(chan bool)
	chInput := make(chan int)
	go pipe.Read(chInput, done)

	chNegativeChannel := make(chan int)
	go pipe.Negative(chInput, chNegativeChannel, done)

	chNotAMultipleOfThree := make(chan int)
	go pipe.NotAMultipleOfThree(chNegativeChannel, chNotAMultipleOfThree, done)

	chBufferInt := make(chan int)
	go pipe.BufferStage(chNotAMultipleOfThree, chBufferInt, done, buffer, interval)

	for  {
		select {
		case data := <- chBufferInt:
			fmt.Println("Полученные данные ...", data)
		case <- done:
			return
		}
	}
}