package main
import "fmt"

/**
* 第一种方法
**/
func method1()  {
  var p Phone;

  var p1 = NokiaPhone{number: 1, desc:"I'm Nokias'"};

  var p2 = IPhone{number: 2, desc:"I'm Apples'"};

  p = p1;
  p.call();

  p = p2
  p.call();
}

/**
* 第二种方法
**/
func method2()  {
  var p1 Phone = NokiaPhone{number: 1, desc:"I'm Nokias'"};

  var p2 Phone = IPhone{number: 2, desc:"I'm Apples'"};

  p1.call();

  p2.call();
}

/**
* 第三种方法
**/
func method3()  {
  var p1 = NokiaPhone{number: 1, desc:"I'm Nokias'"};
  p1.call();
}

// func main()  {
//   method1();
//   method2();
//   method3();
// }

// OOP
type Phone interface {
  call()
}

type NokiaPhone struct {
  number int32;
  desc string;
}

type IPhone struct {
  number int32;
  desc string;
}

func (v NokiaPhone) call() {
  fmt.Printf("number=%d,%s,calling by NokiaPhone\n",v.number,v.desc)
}

func (v IPhone) call() {
  fmt.Printf("number=%d,%s,calling by IPhone\n",v.number,v.desc)
}
