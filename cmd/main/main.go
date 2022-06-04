package main

import (
	"dataPipeline/internal/config"
	"dataPipeline/internal/pipe"
	"dataPipeline/internal/ringBuf"
	"dataPipeline/pkg/logging"
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
	go pipe.Read(chInput, done, logger)

	chNegativeChannel := make(chan int)
	go pipe.Negative(chInput, chNegativeChannel, done, logger)

	chNotAMultipleOfThree := make(chan int)
	go pipe.NotAMultipleOfThree(chNegativeChannel, chNotAMultipleOfThree, done, logger)

	chBufferInt := make(chan int)
	go pipe.BufferStage(chNotAMultipleOfThree, chBufferInt, done, buffer, interval, logger)

	for  {
		select {
		case data := <- chBufferInt:
			logger.Infof("Полученные данные ...%d", data)
		case <- done:
			return
		}
	}
}
