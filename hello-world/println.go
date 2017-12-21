package main
import (
  "fmt"
  "log"
  "os"
)

func main() {
  log.Println(os.Args[0]);
  fmt.Println("Hello World")
  fmt.Printf("%s=%d\n","1+1",1+1);
}
