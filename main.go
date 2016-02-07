package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"sync"
	"os"
	"os/exec"
)

func main() {
	logger := make(chan string)
	go func() {
		f, err := os.OpenFile("/data/gs.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
		    panic(err)
		}

		defer f.Close()

		for {
			text := time.Now().Format("2006/01/02 15:04:05") + " > " + <-logger + "\n"
			if _, err = f.WriteString(text); err != nil {
			    panic(err)
			}
		}
	}()

	fmt.Println("Started")

	go hourly(logger)
	go daily(logger)

	time.Sleep(time.Second*10)

	select{}
}

func hourly(logger chan string) {
	for {
		hourly := readHourly()
		execScripts(hourly, logger)

		then := time.Now()
		var t = then.Minute()*60 + then.Second()
		sleepTime := time.Hour - time.Duration(t)*time.Second
		logger <- "Sleep for " + sleepTime.String()
		time.Sleep(sleepTime)
	}
}

func daily(logger chan string) {
	for {
		daily := readDaily()
		execScripts(daily, logger)

		then := time.Now()
		var t = then.Hour()*3600 + then.Minute()*60 + then.Second()
		sleepTime := time.Hour*24 - time.Duration(t)*time.Second
		logger <- "Sleep for " + sleepTime.String()
		time.Sleep(sleepTime)	
	}
}

func readHourly() (hourlySlice []string){
	hourly, _ := ioutil.ReadDir("/data/hourly/")
	for _, f := range hourly {
		hourlySlice = append(hourlySlice, "/data/hourly/" + f.Name())
	}
	return
}

func readDaily() (dailySlice []string){
    daily, _ := ioutil.ReadDir("/data/daily/")
    for _, f := range daily {
    	dailySlice = append(dailySlice, "/data/daily/" + f.Name())
    }
    return
}

func execScripts(scripts []string, logger chan string) {
	var wg sync.WaitGroup
	wg.Add(len(scripts))
	for _, f := range scripts {
		var t = f
		go func() {

			out, err := exec.Command("/bin/bash", t).Output()
			if err != nil {
				logger <- t + " - " + err.Error()
			}

			logger <- ("output>> " + string(out) + " <<\n")
			logger <- "executed " + t

			defer wg.Done()
		}()
	}
	wg.Wait()
}
