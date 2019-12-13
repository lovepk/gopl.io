package main

import (
	"fmt"
)

type interface1 interface {
	echo(s string)
	echo1(s string)
}

type type1 struct {
	s string
}

func (t type1)echo(s string)  {
	t.s = s;
	fmt.Println(t.s);
}

func (p *type1)echo1(s string)  {
	p.s = s;
	fmt.Println(p.s);
}

func (p *type1)echo2(s string)  {
	p.s = s;
	fmt.Println(p.s);
}


func main() {
	var iValue interface1;

	var tPointer *type1;
	var tValue type1;
	tPointer = &type1{};
	tValue = type1{};

	iValue = tValue;  
	//  此处不能赋值的原因是tValue是type1类型的，不是*type1类型的。又因为T类型的值不能拥有所有*T指针的方法，所以type1类型只实现了echo，而没有实现echo1，故type1类型没有实现接口interface1，故不能赋值。
	fmt.Printf("%T\n", iValue);
	tValue.echo("taaaaa");
	tValue.echo1("taaaaa");
	iValue.echo("aaaaa");
	iValue.echo1("aaaaa");
	iValue.echo2("aaaaa");
	// go语言是静态类型语言，接口变量是静态类型的，接口的静态类型决定了什么方法可以通过接口型变量调用
	// 比如interface{}型的接口变量，可以接受任何值得赋值，但是不能直接对它持有的值做操作，因为interface{}没有任何方法

	iValue = tPointer;
	fmt.Printf("%T\n", iValue);
	tPointer.echo1("pbbbb");
	tPointer.echo("pbbbb");
	iValue.echo1("bbbb");
	iValue.echo("bbbb");
	iValue.echo2("aaaaa");
}