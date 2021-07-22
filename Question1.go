package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"

)

func main(){
	var numberOfStrings int//stores the number of strings that we need
	var mutex = &sync.Mutex{}
	fmt.Scanln(&numberOfStrings)
	frequencyMap := make(map[string]int)//this will store the count of each individual rune(string)
	givenSetOfString := make([]string,numberOfStrings)
	for i := 0; i < numberOfStrings ; i++{
		fmt.Scanln(&givenSetOfString[i])
	}
	jobsChannel := make(chan string,numberOfStrings)//a channel to send the strings for the worker to work on
	workerChannel := make(chan bool)//channel for workers
	go func() {
		for{
			currentJob,jobsLeft := <-jobsChannel
			if jobsLeft == true{
				for _, char := range currentJob{
					mutex.Lock()
					frequencyMap[string(char)]++
					mutex.Unlock()
				}
			} else {
				workerChannel <- true
				return
			}
		}
	}()
	//send the strings over the job channel
	for i := 0; i < numberOfStrings ; i++{
		jobsChannel <- givenSetOfString[i]
	}
	close(jobsChannel) //close the channel as no longer required
	fmt.Println("{")
	keys := make([]string, 0, len(frequencyMap))
	for key := range frequencyMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)//sort on the map on the basis of key in frequencyMap
	for _, key := range keys {
		fmt.Println(strconv.Quote(key), ":",frequencyMap[key])
	}
	fmt.Println("}")
	<-workerChannel//await the worker
}
