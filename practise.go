package main

import "fmt"
import "time"

const s string = "constant"
func main()  {
        //value
        fmt.Println("hello world")
        fmt.Println("go" + "lang")
        fmt.Println("1+1 =", 1+1)
        fmt.Println("7.0/3.0 =", 7.0/3.0)
        fmt.Println(true && false)
        fmt.Println(true || false)
        fmt.Println(!true)
        //variables, The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short" in this case
        var aa = "initial"
        fmt.Println(aa)

        var bb, c int = 1, 2
        fmt.Println(bb, c)

        var d = true
        fmt.Println(d)

        var e string
        fmt.Println(e)

        f := "short"
        fmt.Println(f)
        //Constants, supports constants of character, string, boolean, and numeric values
        const s string = "another"
        fmt.Println(s)

        // for is the only looping construct
        i := 1
        for i<=3 {
                fmt.Println(i)
                i++
        }
        for {
                fmt.Println("test for usage")
                break;
        }
        //go don't need parentheses around condition, but nned braces are required
        // 1 < 2? fmt.Println(1):fmt.Println(2), there is no ternary if in Go

        //switch
        whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
        whatAmI("hey")

        //Arrays,an array is a numbered sequence of elements of a specific length.
        var a [5]int
    fmt.Println("emp:", a)
        a[4] = 100
    fmt.Println("set:", a)
        fmt.Println("get:", a[4])
        fmt.Println("len:", len(a))

        b := [5]int{1, 2, 3, 4, 5}
        fmt.Println("dcl:", b)

        var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
        fmt.Println("2d: ", twoD)

        //slices,more power than array
        ss := make([]string, 3)
        fmt.Println("emp:", s)
        ss[0] = "a"
    ss[1] = "b"
    ss[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
        cc := make([]string, len(s))
    copy(cc, ss)
        fmt.Println("cpy:", cc)

        l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
        fmt.Println("sl2:", l)

    l = s[2:]
        fmt.Println("sl3:", l)

        t := []string{"g", "h", "i"}
        fmt.Println("dcl:", t)

        twoDD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoDD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoDD[i][j] = i + j
        }
    }
        fmt.Println("2d: ", twoDD)

        //map
        m := make(map[string]int)
        m["k1"] = 7
        m["k2"] = 13
        fmt.Println("map:", m)

        //range, use for iterator on, arrays, slice, map
        for ii, ci := range "go" {
                fmt.Println(ii, ci)
        }


        res := plusPlus(1, 2, 3)
        fmt.Println("1+2+3 =", res)

        sums := []int{1, 2, 3, 4}
        sum(sums...)

        nextInt := intSeq()

        fmt.Println(nextInt())
        fmt.Println(nextInt())
        fmt.Println(nextInt())
        fmt.Println(nextInt())

        fmt.Println(fact(5))

        fmt.Println(person{"Bob", 20})

        r := rect{width: 10, height: 5}
        fmt.Println("area: ", r.area())
        fmt.Println("perim:", r.perim())

        //the interface of Go
        // the releation of struct and Go
        //error
        //Go goroutines
        testf("direct")
        go testf("goroutine")
        go func(msg string) {
                fmt.Println(msg)
        }("going")
        fmt.Scanln()
        fmt.Println("done")

        //go Channels
        //By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
        task := make(chan string)
        go func(){ task <- "ping"}()
        tt := <-task
        fmt.Println(tt)
        //channel buffeer
        task1 := make(chan string, 2)
        task1 <- "ren"
        task1 <- "jin"
        fmt.Println(<- task1)
        fmt.Println(<- task1)
        //use channels to synchronize execution across goroutines
        done := make(chan bool, 1)
        go worker(done)
        <- done
        //When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.
        // Channel Directions
        pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
        fmt.Println(<-pongs)
        //select usage,Go’s select lets you wait on multiple channel operations
        c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
        }()
        select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
        }
        c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
        }()
        select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
        }
        //Non-Blocking Channel Operations, use select default case to release
    messages := make(chan string)
        signals := make(chan bool)
        select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
        }
        //Closing Channels
        close(messages)
        close(signals)
        //Range over Channels
        //This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
        queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
        close(queue)
        for elem := range queue {
        fmt.Println(elem)
        }
        // Timers and Tickers
        // Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals
        ticker := time.NewTicker(500 * time.Millisecond)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
        }()
        time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
        fmt.Println("Ticker stopped")
        //Workers Pool
        jobs := make(chan int, 100)
        results := make(chan int, 100)
        for w := 1; w <= 3; w++ {
        go workers(w, jobs, results)
        }
        for j := 1; j <= 5; j++ {
        jobs <- j
        }
        close(jobs)
        for a := 1; a <= 5; a++ {
        <-results
        }
        //Rate Limiting
        // Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.
        // Atomic Counters
        // sync/atomic package for atomic counters accessed by multiple goroutines
        // sort function
}
func workers(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}
func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
        fmt.Println("done")
        done <- true
}
func testf(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
//functions, have multiple consecutive parameters of the same type

func plusPlus(a, b, c int) int {
        return a + b + c
}

//One is multiple return values
func vals() (int, int) {
    return 3, 7
}
// Accepting a variable number of arguments
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
//form closures
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
//recursion
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

//value and reference
func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}
//struct
type person struct {
    name string
    age  int
}

type rect struct {
    width, height int
}
// use point, can change the instance of struct's value
func (r *rect) area() int {
    return r.width * r.height
}
// use value is also well, just calculate and get value
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
