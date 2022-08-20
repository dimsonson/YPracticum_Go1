package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	StartTime time.Time
	Split     []time.Duration
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
	//fmt.Println(sw)
}

// запустить/сбросить секундомер
func (stw *Stopwatch) Start() {
	stw.StartTime = time.Now()
	stw.Split = nil
}

// сохранить промежуточное время
func (stw *Stopwatch) SaveSplit() {
	stw.Split = append(stw.Split, time.Since(stw.StartTime))
}

// вернуть текущие результаты
func (stw *Stopwatch) GetResults() []time.Duration {
	return stw.Split
}

/* package main

import (
    "fmt"
    "time"
)

type Stopwatch struct {
    startTime time.Time
    splits    []time.Time
}

func (sw *Stopwatch) Start() {
    sw.startTime = time.Now()
    sw.splits = nil
}

func (sw *Stopwatch) SaveSplit() {
    sw.splits = append(sw.splits, time.Now())
}

func (sw Stopwatch) GetResults() (retResults []time.Duration) {
    for _, splitTime := range sw.splits {
        retResults = append(retResults, splitTime.Sub(sw.startTime))
    }

    return
}

func main() {
    sw := Stopwatch{}
    sw.Start()

    time.Sleep(1 * time.Second)
    sw.SaveSplit()

    time.Sleep(500 * time.Millisecond)
    sw.SaveSplit()

    time.Sleep(300 * time.Millisecond)
    sw.SaveSplit()

    fmt.Println(sw.GetResults())
}  */
