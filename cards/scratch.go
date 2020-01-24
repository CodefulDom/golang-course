package main

import "fmt"

func otherFunc()  {
  getValue := functionTrial()

  fmt.Println(getValue)
}

func functionTrial() string {
  return "Hey, did this shit work"
}
