package main

import (
  "log"
)

func threadYE(x int){
  LABEL:
    log.Print(x)
  goto LABEL
}

func main(){
  go threadYE(1)
  go threadYE(2)    //Go concurrency will use multiple CPU cores! Nice!
  go threadYE(3)    //You can watch them be used in HW Monitor
  go threadYE(4)
  go threadYE(5)
  go threadYE(6)
  go threadYE(7)
  go threadYE(8)

  select{}
}
