package main

import (
	"fmt"
	"math"
	"time"
)


var Pi float64
func init() {
	Pi = math.Atan(1) * 4
}
var a = "G"
var a1 string
/*//协程 watiGroup
func worker(id int,wg* sync.WaitGroup) {
	fmt.Println("worker %d starting",id)
	time.Sleep(time.Second*1)
	fmt.Println("worker %d done",id)
	wg.Done()
}*/

/*//线程池
func worker(id int,jobs <-chan int,results  chan<- int,) {
	for j := range jobs {
		fmt.Println("jobs",id,"prosess id",j)
		time.Sleep(time.Second*2)
		results <- j*2
	}
}*/

func main(){
	//36速率限制
	request := make(chan int,5)
	for i := 1;i <= 5;i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(time.Millisecond*100)
	for req := range request {
		<- limiter
		fmt.Println("request",req,time.Now())
	}
	burstyLimiter := make(chan time.Time)
/*	//	//35协程 sync.watiGroup
	var wg sync.WaitGroup
	for i :=1;i <= 5;i++ {
		wg.Add(1)
		worker(i,&wg)
	}
	wg.Wait()*/



	/*	//34线程池                   需要想想。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。需要想想
	jobs := make(chan int,100)
	results := make(chan int,100)
	for w:=1;w <=3;w++ {
		go worker(w,jobs,results)
	}
	for j := 1;j <=9;j++ {
		jobs <-j
	}
	close(jobs)
	for a:=1;a <=9;a++ {
		<-results
	}
*/

	//33打点器


/*	//32 定时器
	time1 := time.NewTimer(time.Second*2)
	<-time1.C
	fmt.Println("timer1 expired")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stop")
	}*/
/*	//31通道的遍历
	queue := make(chan string,2)
	queue <- "send 1"
	queue <- "send 2"

	close(queue)
	for ele := range queue {
		fmt.Println(ele)
	}
*/
/*	//30 通道关闭
	jobs := make(chan int,5)
	done := make(chan bool)
	go func() {
		j,more := <- jobs
		for {
			if more {
				fmt.Println("jobs is", j)
			} else {
				fmt.Println("jobs done")
				done <- true
				return
			}
		}
	}()
	for j:=0;j < 3;j++ {
		jobs <- j
		fmt.Println("send jobs",j)
	}
	close(jobs)
	<-done*/
/*	//29非阻塞通道
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg1 := <-message:
		fmt.Println(msg1)
	default:
		fmt.Println("default 1")
	}
	msg := "hi"
	select {
		case  message  <-msg:
			fmt.Println(message)
	default:
		fmt.Println("default 2")
	}

	select {
		case msg := <- message:
			fmt.Println("received message",msg)
			case sig := <-signals:
				fmt.Println("received signals",sig)
	default:
		fmt.Println("no active")
	}*/
/*//超时处理
	c1 := make(chan string,1)
	go func() {
		time.Sleep(time.Second*2)
		c1 <- "result 1"
	}()
	select {
		case res := <-c1:
			fmt.Println(res)
			case <-time.After(time.Second*1):
			fmt.Println("timeout 1")
	}
	c2 := make(chan string,2)
	go func() {
		time.Sleep(time.Second*2)
		c2<-"result 2"
	}()
	select {
	case res2 := <-c2:
		fmt.Println(res2)
		case <-time.After(time.Second*3):
		fmt.Println("timeout 2")
}*/

/*//通道方向
//发送
func ping(pings chan<-string,msg string) {
	pings <- msg
}
//接受
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
*/
/*//通道通信
func worker(done chan bool) {
	fmt.Printf("doing .....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done<- true
}*/

/*//errror处理
func f1(arg int)(int, error) {
	if(arg == 42) {
		return -1,errors.New("arg is 42")
	}
	return arg+3,nil
}
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return	fmt.Sprintf("%d-%s",e.arg,e.prob)
}
func f2(arg int) (int ,error) {
	if(arg == 42) {
		return -1,errors.New("can`t work")
	}
	return arg+3,nil
}*/
/*//协程
func f(from string){
	for i:=0;i<3;i++ {
		fmt.Println(from,": ",i)
	}
}*/
/*type geometry interface {
	area() float64
	preim() float64
}
type rect struct {
	weight,height float64
}
type circle struct {
	radis float64
}

func (r  rect) area() float64 {
	return r.height*r.weight
}
func (r rect) preim() float64 {
	return 2*r.height + 2*r.weight
}
func (c circle) area() float64{
	return math.Pi*c.radis*c.radis
}
func (c circle) preim() float64 {
	return math.Pi*c.radis*2
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.preim())
}*/
/*type person struct {
	name string
	age int
}*/
/*func zeroval (val int) {
	val = 0
}
func zeroptr (ptr *int) {
	*ptr = 0
}*/
/*func fact(a int) int {
	if(a == 0) {
		return 1
	}
	return a * fact(a-1)
}*/
/*//闭包
func intSeq() func() int {
	i := 0
	return func() int{
		i++
		return i
	}
}*/
/*func mul_par(nums ...int) {
	fmt.Print(nums," ")
	total := 0
	for _,num := range nums{
		total += num
	}
	fmt.Println(total)
}
*//*func mul_ret() (int,int) {
	return 3,4
}
func plus(a,b int) int {
	return a+b
}
func plusplus(a,b,c int) int {
	return a+b+c
}*/
/*func main() {*/
	//var isToF bool +
/*	//28.通道选择器 select
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second*1)
		//fmt.Println("one")
		c1<-"One"
	}()
	go func() {
		time.Sleep(time.Second*2)
		//fmt.Println("two")
		c2 <-"Two"
	}()

	for i := 0;i <2;i++ {
		select {
		case  msg1 := <-c1:
			fmt.Println("Received",msg1)

			case msg2 := <-c2:
				fmt.Println("Received",msg2)
		}
	}*/
/*	//27.通道方向  指定
	pings := make(chan string,1)
	pongs := make(chan string,2)

	ping(pings,"this is a string")
	pong(pings,pongs)
	fmt.Println(<-pongs)*/
/*	//26.通道同步  通道来同步 Go 协程间的执行状态
	done := make(chan bool)
	go worker(done)
	<-done*/

/*	//25 通道缓冲
	message := make(chan string,2)

	message <- "channel"
	message <- "buffer"

	fmt.Println(<-message)
	fmt.Println(<-message)*/
/*	//24 通道  协程间通信
	message := make(chan string)
	go func() { message <- "ping"}()
	msg := <-message
	fmt.Println(msg)*/


/*	//23 协程
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")*/
/*	//22 错误处理
	for _,i := range []int{7,42} {
		if r,e := f1(i); e != nil {
			fmt.Println("f1 failed:",e)
		}else {
			fmt.Println("f1 worked",r)
		}
	}

	for _,i := range []int{7,42} {
		if r,e := f2(i); e != nil {
			fmt.Println("f1 failed:",e)
		}else {
			fmt.Println("f1 worked",r)
		}
	}*/
/*	//21 接口
	r := rect{3,4}
	p := circle{10}
	measure(r)
	measure(p)*/
/*	//20 方法
	r1 := rect{2,3}
	fmt.Println(r1.area())*/

/*	//19 结构体
	fmt.Println(person{"Bob",230})
	fmt.Println(person{name:"alice",age:23})
	fmt.Println(person{name:"adward"})
	fmt.Println(&person{"a",3})

	s := person{"cc",25}
	sp := &s
	fmt.Println(sp.name)
	sp.age = 40
	fmt.Println(sp)*/
/*	//18 指针
	val := 10
	fmt.Println(val)
	zeroval(val)
	fmt.Println(val)
	p := &val
	fmt.Println(*p)
	zeroptr(p)
	fmt.Println(*p)*/

/*	//17 递归
	fmt.Println(fact(7))
	//16 匿名函数  匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
*/
/*	//15 变参函数
	nums := []int{1,2,3,4,5,6,7}
	mul_par(nums...)
	mul_par(1,2,3,4)*/
/*	//14 多返回值
	fmt.Println(mul_ret())*/

/*	//13 函数
	fmt.Println(plus(1,2))
	fmt.Println(plusplus(1,2,3))
*//*	//12 range 遍历
	nums := [3]int {3,4,5}
	sum := 0
	for _,num := range nums {
		sum += num
	}
	fmt.Println(sum)
	for i,num1 := range nums {
		if(3 == num1) {
			fmt.Println(i)
		}
	}
	kvs := map[string]string {"a":"apple","b":"banan"}
	for k,v := range kvs {
		fmt.Println(k,v)
	}
	for i,c := range "go" {
		fmt.Println(i,c)
	}*/
	//关联数组


/*	//切片
	s := make([] string,3)
	fmt.Println("s:",s)
	s[0] = "1"
	fmt.Println(s)
	s = append(s,"4")
	fmt.Println(s)*/


	//// 返回一个“返回int的函数”
	//func fibonacci() func() int {
	//	sum := 0
	//	return func() int {
	//		sum +=
	//	}
	//}
	//
	//
	//f := fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(f(i))}
	/*
	//函数闭包
	func adder() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}
	pos,nes = adder(),adder();

	 */
	/*映射


	type Vertex struct {
		a,b float64
	}
	var m map[string] Vertex
	m = make(map[string] Vertex)
	m["aa"] = Vertex{1.33,2.22}
	fmt.Println(m["aa"])
	fmt.Println(m)*/
	/*
	//切片

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[1:3]
	fmt.Println(a)
	a[1] = "XXXX"
	fmt.Println(names)
	 */
	/*
	//切片

	arr := [8]int{1,2,3,4,5,6,7,8}
	var sli []int = arr[1:5]
	fmt.Println(sli)
	*/
	/*
	//数组
	var prim [4] int = [4]int{1,2,3,4}
	fmt.Println(prim)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	 */
	/* 数字字符串转换
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
	*/

	/*分割拼接
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str3) */

	/* 字符串分割
	strings.fiel
	strings.Fields("a b c")
	strings.Split("aabcc",'b')

	*/
	/*
	4.5.字符串替换

	 */
	/* 字符串 3. 包含定位
	var indexNum int
	indexNum = strings.Index("abc","c")
	fmt.Print(indexNum)
	indexNum = strings.LastIndex("asdfghjkl","fgh")
	fmt.Print(indexNum)

	/* 字符串 2.包含关系
	isToF = strings.Contains("asdfghjkl","sdfghjkl")
	fmt.Print(isToF)
	/*
	字符串  1.前后字符比较

	var isSS bool

	isSS = strings.HasPrefix("aaabbbccc","abc")
	isSS = strings.HasSuffix("aaccdd","dd")
	fmt.Print(isSS)/*
	/* 随机数
	for i := 0;i <10;i++  {
		a := rand.Int()
		fmt.Printf("%d /",a)
	}
	for i := 0; i < 5; i++ {
		a := rand.Intn(10)
		fmt.Printf("%d /",a)
	}

	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0 ; i < 10 ; i++ {
		fmt.Printf("%2.2f /",100* rand.Float32())
	}*/


	/*a1 = "G" 4.4例
	print(a1)
	f1()

	n() //例1
	m()
	n()

	 var twoPi = 2 * Pi
	fmt.Print("next:\n\n")
	fmt.Print(twoPi)

	*/
	/*var b int
	_ , b = 5 , 7
	fmt.Print(b)
	var osSysName string = runtime.GOOS
	fmt.Print("this is a:%s\n",osSysName)
	path := os.Getenv("PATH")

	fmt.Print(path)
	fmt.Printf("\n")*/
}
/*func f1() {
	a1 = "O"
	print(a1) //O
	f2()  //G
}
func f2() {
	print(a1)
}
func n() {
	print(a)
}
func m() {
	a = "O"
	// a := "O"
	print(a)
}*/