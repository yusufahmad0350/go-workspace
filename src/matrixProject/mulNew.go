package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type pair struct {
    row, col int
}

const length = 512

var start time.Time
var rez [length][length]int

func main() {
    const threadlength = 1
    pairs := make(chan pair, 1000)
    var wg sync.WaitGroup
    var a [length][length]int
    var b [length][length]int
    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
            a[i][j] = rand.Intn(10)
            b[i][j] = rand.Intn(10)
        }
    }
    wg.Add(threadlength)
    for i := 0; i < threadlength; i++ {
        go Calc(pairs, &a, &b, &rez, &wg)
    }
    start = time.Now()
    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
            pairs <- pair{row: i, col: j}
        }
    }
    close(pairs)
    wg.Wait()
    elapsed := time.Since(start)
    fmt.Println("Binomial took ", elapsed)

    for i := 0; i < length; i++ {
        for j := 0; j < length; j++ {
           // fmt.Print(rez[i][j])
            //fmt.Print(" ")
        }
        //fmt.Println(" ")
    }
}

func Calc(pairs chan pair, a, b, rez *[length][length]int, wg *sync.WaitGroup) {
    for {
        pair, ok := <-pairs
        if !ok {
            break
        }
        rez[pair.row][pair.col] = 0
        for i := 0; i < length; i++ {
            rez[pair.row][pair.col] += a[pair.row][i] * b[i][pair.col]
        }
    }
    wg.Done()
}