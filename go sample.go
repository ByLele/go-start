package main

const g string = "to be stronger"

/*//状态协程
type  readOp struct{
	key int
	resp chan int
}
type writeOp struct{
	key int
	value int
	resp chan bool
}*/
/*//线程池
func worker (it int, jobs <-chan int,results  chan<- int) {
	for j:= range jobs {
		fmt.Println("worker",it,"processing id",j)
		time.Sleep(time.Second*2)
		results <- it * 2
	}
}*/

func main() {

/*	//40.排序
	strs := []string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println(strs)
	ints := []int{2,33,3,4,1}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)*/

/*	//39.go状态协程
	var readOps uint64 = 0
	var writeOps uint64 = 0

	reads := make(chan *eadOp)
	writes := make(chan *writeOp)
	go func(){
		state := make(map[int]int)
		for {
			select {
				case read := <-reads:
					read.resp <- state[read.key]
					case write := <-writes:
						state[write.key] = write.key
						write.resp <-true
			}
		}
	}()*/


/*	//互斥锁
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0

	for r :=0;r <100;r++ {
		go func(){
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops,1)
				runtime.Gosched()
			}
		}()
	}
	for w := 0;w < 10;w++ {
		go func(){
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddInt64(&ops,1)
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops",opsFinal)

	mutex.Lock()
	fmt.Println("state",state)
	mutex.Unlock()*/

/*	//37原子计数器
	var ops uint64 = 0;
	for i := 0;i <50;i++ {
		go func() {
			for {
				atomic.AddUint64(&ops,1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops",opsFinal)*/

/*	//34线程池
	jobs := make(chan int,100)
	results := make(chan int,100)

	for w:=0;w <=3;w++ {
		go worker(w,jobs,results)
	}
	for j := 1;j <= 9;j++ {
		jobs<-j
	}
	close(jobs)

	for a:=1;a<=9;a++ {
		<-results
	}*/
	/*	//34打点器
	ticker := time.NewTicker(time.Millisecond*500)
	go func() {
		for em := range  ticker.C {
			fmt.Println("ticker tick",em)
		}
	}()
	time.Sleep(time.Millisecond*1600)
	ticker.Stop()
	fmt.Println("ticker stop")*/
/*	//定时器
	timer1 := time.NewTimer(time.Second*2)
	<-timer1.C
	fmt.Println("timer1 timeout")
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("time2 timeout")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stop")
	}
*/
	//fmt.Println("Hello World!")
	//2.
	//fmt.Println("cat"+"dog")
	//fmt.Println("1+1",1+1)
	//fmt.Println("3.0/2.0",3.0/2.0)
	//fmt.Println(true && false)
	//fmt.Println(true || false)
	//fmt.Println(!false)

	////3.
	//var a = 1
	//var b,c =2,3
	//fmt.Println(a)
	//fmt.Println(b,c)
	//
	//var d = "just go"
	//fmt.Println(d)
	//var e = true
	//fmt.Println(e)
	//var  f int
	//fmt.Println(f)

	////4.
	//fmt.Println(g)
	//const h = 50000
	//fmt.Println(h)
	//const i = 1e8/h
	//fmt.Println(i)
	//fmt.Println(int64(i))
	//fmt.Println(math.Sin(i))
	//5 for+
	//i := 1
	//for i <= 3 {
	//	fmt.Println(i)
	//	i++
	//}
	//for j := 3;j <7;j++ {
	//	fmt.Println(j)
	//}
	//for {
	//	fmt.Println("loop")
	//	break
	//}
	//for k := 0; k < 4;k++ {
	//	if(k % 2 == 0) {
	//		continue;
	//	}
	//	fmt.Println(k)
	//}
	////if
	//if 7 % 2== 0 {
	//	fmt.Println("this not odd")
	//}
	//if 8 % 4 == 0 {
	//	fmt.Println("this is ")
	//}
	//
	//if num:=9;num<0 {
	//	fmt.Println(num,"is negative")
	//}else if num <10 {
	//	fmt.Println(num,"has 1 digit")
	//}else {
	//	fmt.Println(num,"has multiple digits")
	//}
	// 7 switch/case
	//i := 2
	//fmt.Print("where",i,"as")
	//switch i {
	//case 1:
	//	fmt.Println("one")
	//case 2:
	//	fmt.Println("two")
	//case 3:
	//	fmt.Println("three")
	//}
	//switch time.Now().Weekday(){
	//case time.Saturday,time.Sunday:
	//	fmt.Println("weekend")
	//default:
	//	fmt.Println("work day")
	//}
	//t := time.Now()
	//switch  {//此处注意
	//case t.Hour() < 12:
	//	fmt.Println("before noon")
	//default:
	//	fmt.Println("after noon")
	//
	//}
	//whatAmI := func(i interface{}) {
	//	switch t := i.(type) {
	//	case bool:
	//		fmt.Println("bool")
	//	case int :
	//		fmt.Println("int")
	//	case string:
	//		fmt.Println("string")
	//	default:
	//		fmt.Println("other type",t)
	//	}
	//}
	//whatAmI(true)
	//whatAmI(22)
	//whatAmI("safaf")

	////8array
	//var arr[5] int
	//fmt.Println("emp",arr)
	//arr[4] = 100
	//fmt.Println("arr",arr)
	//fmt.Println("arr[4]",arr[4])
	//
	//fmt.Println("len",len(arr))
	//
	//b := [5]int{1,2,3,4,5}
	//fmt.Println("dcl:",b)
	//
	//var twoarr [2][3]int
	//for i := 0;i < 2;i++ {
	//	for j :=0; j <3;j++ {
	//		twoarr[i][j] = i + j
	//	}
	//}
	//fmt.Println("two arr",twoarr)

	////9 slices
	///* 类型仅由 包含元素决定*/
	//s := make([]string,3) //注意make 语法
	//fmt.Println("emp0",s)
	//s[0] = "a"
	//s[1] = "b"
	//s[2] = "c"
	//fmt.Println("set:",s)
	//fmt.Println("get:",s[2])
	//
	//fmt.Println("len",len(s))
	//fmt.Println("cap",cap(s))
	//s = append(s,"d")
	//s = append(s,"e")
	//fmt.Println("appended:",s)
	//c := make([]string,len(s)) //s 复制给c
	//copy(c,s)
	//fmt.Println("copy c by s",c)
	///*切片操作*/
	//l := s[2:5] //slice[low:high]切片操作
	//fmt.Println("sl1",l)
	//l = s[:5]
	//fmt.Println("sl3",l)
	//
	//l = s[2:]
	//fmt.Println("sl3",l)
	//
	//t := []string{"g","h","i"}
	//fmt.Println("dcl",t)
	//
	////多维slice  长度可以不同
	//twoD := make([][]int,3)
	//for i:=0;i < 3;i++ {
	//	innerLen := i+1
	//	twoD[i] = make([]int,innerLen)
	//	for j:=0;j<innerLen;j++ {
	//		twoD[i][j] = i+j
	//	}
	//}
	//
	//fmt.Println("2d:",twoD)

	////map 关联数组
	//// 要创建一个空 map，需要使用内建的 `make`:
	//// `make(map[key-type]val-type)`.
	//m := make(map[string]int)
	//m["k1"] = 7
	//m["k2"] = 13
	//fmt.Println("map",m)
	//
	//v1 := m["k1"]
	//fmt.Println("v1:",v1)
	//
	//fmt.Println("len",len(m))
	//
	////删除
	//delete(m,"k2")
	//fmt.Println("map",m)
	//_,prs := m["k2"]
	//fmt.Println("value",prs)
	//
	//n := map[string]int{"a":33,"b":333}
	//fmt.Println(n)

	//11 range遍历

}