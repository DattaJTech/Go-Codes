package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
    "strconv"
    "time"
)

func main() {
    readIntervalPeriod := os.Args[1]
    repeatDelay, err := strconv.ParseUint(readIntervalPeriod,10,0) 
    if (err != nil) { // if error in parsing
        fmt.Println(err)
        fmt.Println("please restart with valid read interval")
        return
    }
    for {
        out, err := exec.Command("wmic", "cpu", "get" , "loadpercentage").Output()
        if err != nil {
            log.Fatal(err)
        }

        outputString :=  string(out[:len(out)])     
        indexOfNewline1 := strings.Index(outputString,"\n")
        loadPercentageString := string(out[indexOfNewline1+1:len(out)]) 
        indexOfNewline2 := strings.Index(loadPercentageString,"\n")
        loadpercentage := string(loadPercentageString[0:indexOfNewline2])
        fmt.Println("Current load percentage is ", loadpercentage)
        time.Sleep(time.Duration(repeatDelay) * time.Second)
    }
}