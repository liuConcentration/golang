package main

import (
	"fmt"
	"learnGo/retriever/mock"
	real2 "learnGo/retriever/real"
	"time"
)

/**
接口定义的解说, download 和Retriever的角色
download 是使用者,Retriever 是实现者
go语言的接口是由使用者来定义的,传统的面向对象是实现者调用的,

接口变量自带指针
接口变量同样采用值传递,几乎不需要使用接口的指针
指针接收者实现智能以指针方式使用,值接受者都可以

表示任何类型: interface{} 即泛型 这里在上面的目录中我单独写了一个queueinterfa目录可以去看
查看接口变量
Type Assertion
Type  Switch
*/

//实现者
type Retriever interface {
	//在interface中我们这个Get()不用加func这关键字将Get 指定为函数
	Get(url string) string
}

//使用者
func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func inspect(r Retriever) {
	//打印类型
	fmt.Printf("%T %v\n", r, r)
	//可以发现改为&mock.Retriever{...,之后下面的case mock.Retriever:不会被命中了输出不了Contents:
	fmt.Printf("type switch:\n")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Content)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
func main() {
	var r Retriever
	//第一个属性写第一个,第二个写第二个,以此类推
	//这是一个值接受者,但是可以传入指针也没有问题比如&mock.Retriever{...,然后后面输出类型的时候就会发现,里面的类型也变成了指针
	r = mock.Retriever{"this is fake baidu.com", "this is fake title"}
	//打印类型和value
	//输出mock.Retriever {this is fake baidu.com this is fake title} , 可见r的里面还是有类型和值的
	//fmt.Printf("%T %v\n", r, r)
	//封装后打印类型
	inspect(r)

	//改为指针类型, 可以是一个指针类型也可以是一个值类型, 如果是值类型的很显然他回事一个拷贝到r的里面
	//但是结果都是一样的, 因为我们接口一般不会用指针, 通常接口里面类似于含着一个指针
	r = &real2.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	//打印类型和value
	//*real.Retriever &{Mozilla/5.0 1m0s} 可见r的里面还是有类型和值的,如果GET定义为指针类型的话,这里也会输出指针类型
	//fmt.Printf("%T %v\n", r, r)
	//封装后打印类型
	inspect(r)

	//Type assertion 输出类型打印内容
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)
	//错误案例,这里调用r,因为上面r已经被赋值为real.Retriever所以下面的输出会报错
	//可以通过if 来判断一下屏蔽错误
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Content)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}
