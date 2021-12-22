package main

import "fmt"

//lv1 使用两个goroutine轮流打印100以内的数
//var ch chan int
//var exit chan string
//
//func print1() {
//	for i := 0; i <= 100; i++ {
//		ch <- 1
//		if i%2 == 1 {
//			fmt.Println(i)
//		}
//	}
//}
//func print2() {
//	for i := 0; i <= 100; i++ {
//		<-ch
//		if i%2 == 0 {
//			fmt.Println(i)
//		}
//		if i == 100 {
//			exit <- "x"
//		}
//	}
//}
//
//func main() {
//	ch = make(chan int)
//	exit = make(chan string)
//	go print1()
//	go print2()
//	<-exit
//}

//或者？？
//var ch chan int
//var exit chan string
//var wg sync.WaitGroup
//
//func print1() {
//	for i := 0; i <= 100; i++ {
//		ch <- 1
//		if i%2 == 1 {
//			fmt.Println(i)
//		}
//	}
//	wg.Done()
//}
//func print2() {
//	for i := 0; i <= 100; i++ {
//		<-ch
//		if i%2 == 0 {
//			fmt.Println(i)
//		}
//	wg.Done()
//}

//func main() {
//	ch = make(chan int)
//	wg.Add(2)
//	go print1()
//	go print2()
//	wg.Wait()
//}

//lv2
//a.使用os.Create函数，在你的项目目录下创建一个"plan.txt"文件，
//b.使用file.Write将"I’m not afraid of difficulties and insist on learning programming",写入"plan.txt"中。
//c.使用file.Read函数读取"plan.txt"的内容，并打印到控制台

//func main() {
//	//写入函数相关内容（写入文本里）
//	path := "E:\\GoProjects\\homework\\homework3\\plan.txt"
//	File, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)//括号：路径，操作模式，权限
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer File.Close()
//	strr := "I’m not afraid of difficulties and insist on learning programming\r\n"
//	writer := bufio.NewWriter(File)
//	for i := 0; i < 2; i++ {
//		writer.WriteString(strr)
//
//		writer.Flush()  //缓存写入内容
//	}
//	//读取函数文本里的相关内容
//	Filee, err := os.Open(path) //将os.Open改成 ioutil.ReadFile()可以一次性将文件读取到位，此时返回值是File，  []byte类型的数组，底下打印时可写：fmt.Println(string(File))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//fmt.Println(File)
//	//defer File.Close()
//	reader := bufio.NewReader(Filee)
//	for {
//		str, err := reader.ReadString('\n')
//		fmt.Println(str)   //逐行读取文件
//		if err == io.EOF { //文件结束
//			break
//		}
//	}
//}

// lv3
//下面这段代码是有错误的，要求更改部分代码，实现0~9的打印（最好在错误的地方，注释并标明错误原因）
//func main() {
//	over := make(chan bool)
//	for i := 0; i < 10; i++ {
//		go func() {
//			fmt.Println(i)
//		}()
//		if i == 9 {
//			over <- true
//		}
//	}
//	<-over
//	fmt.Println("over!!!")
//}
//修改后：
func main() {
	over := make(chan bool, 1)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 9 {
			over <- true
		}
	}
	<-over
	fmt.Println("over!!!")
}
