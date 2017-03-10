package main

import (
    "fmt" // for printing
    "net/http" // for making requests
    "io/ioutil" // for reading response
    "strings" // for string manipulation
)

func main() {
    recurring_call() // keep tryin till we get data from ifconfig.me
}

func recurring_call(){
    resp, err := http.Get("http://ifconfig.me/")
    if (err != nil) { // if error in connection
            fmt.Println(err)
            fmt.Println("retrying request to http://ifconfig.me/")
            recurring_call()
    } else {
        defer resp.Body.Close();
        body, err := ioutil.ReadAll(resp.Body) // read the body the output will be an arry of bytes
        if err != nil { // check for any errors if present in Body
            fmt.Println(err)
            return
        }
        s := string(body[:len(body)]) // convert byte array to string
        indexOfIPRead := strings.Index(s, "\"ip_addr\":\"") // get index of "ip_addr": 
        afterIndexOfIPReadString := string(s[indexOfIPRead:len(s)]) // truncate string
        indexOfComma := strings.Index(afterIndexOfIPReadString, ",") // find position of comma IP comes in between ip_addr and ,
        currentIP := string(afterIndexOfIPReadString[10:indexOfComma]) // get current IP
        fmt.Println("Your Current IP on the Internet is:",currentIP) // print IP
    }

}

