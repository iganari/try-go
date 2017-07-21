package main

import (
       "fmt"
       "time"
       )

func setSleepTime(sleepTime int){
    fmt.Println("Sleep start  :", sleepTime, "s")
    time.Sleep(time.Duration(sleepTime) * time.Second)
    fmt.Println("Sleep finish :", sleepTime, "s")
}


func main(){
    // fmt.Println("This is sleep script")
    // var num int = 15
    // fmt.Println("num", num,  "sec sleep start")
    // time.Sleep( 1 * time.Second) 
    // time.Sleep( 1 * time.Minute)
    // time.Sleep( 1 * time.Hour)
    // fmt.Println("num", num,  "sec sleep end")
    setSleepTime(5)
}
