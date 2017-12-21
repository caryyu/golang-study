package main
import "fmt"

/**
* 第一种方法
**/
func method1()  {
  var p PhoneAction;

  var p1 = new(NokiaPhone);
  p1.number = 1;
  p1.desc = "I'm Nokias'";

  var p2 = new(IPhone)
  p2.number = 2;
  p2.desc = "I'm Apples'";

  p = p1;
  p.call();

  p = p2
  p.call();
}

/**
* 第二种方法
**/
func method2()  {
  var p1 = new(NokiaPhone);
  p1.number = 1;
  p1.desc = "I'm Nokias'";

  var p2 = new(IPhone)
  p2.number = 2;
  p2.desc = "I'm Apples'";

  p1.call();
  p2.call();
}

// func main()  {
//   method1();
//   method2();
// }

// OOP
type PhoneAction interface {
  call()
}

type Phone struct {
  number int32;
  desc string;
}

type NokiaPhone struct {
  Phone
}

type IPhone struct {
  Phone
}

func (v NokiaPhone) call() {
  fmt.Printf("number=%d,%s,calling by NokiaPhone\n",v.number,v.desc)
}

func (v IPhone) call() {
  fmt.Printf("number=%d,%s,calling by IPhone\n",v.number,v.desc)
}
