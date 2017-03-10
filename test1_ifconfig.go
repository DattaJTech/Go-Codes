package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
)

func main() {
    recurring_call()
}

func recurring_call(){
    resp, err := http.Get("http://ifconfig.me/")
    if (err != nil) {
            fmt.Println(err)
            fmt.Println("retrying request to http://ifconfig.me/")
            recurring_call()
    } else {
        defer resp.Body.Close();
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println(err)
            return
        }

        s := string(body[:len(body)])
        indexOfIPRead := strings.Index(s, "\"ip_addr\":\"")
        afterIndexOfIPReadString := string(s[indexOfIPRead:len(s)])
        indexOfComma := strings.Index(afterIndexOfIPReadString, ",")
        currentIP := string(afterIndexOfIPReadString[10:indexOfComma])
        fmt.Println("Your Current IP on the Internet is:",currentIP)
    }
}

