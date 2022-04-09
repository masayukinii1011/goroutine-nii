package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Printf("WaitGroup\n")
	useWaitGroup()

	fmt.Printf("--------------------\n")

	fmt.Printf("Channel\n")
	useChannel()
}

func useWaitGroup() {
	var wg sync.WaitGroup // 並行処理の完了を待つために使用

	for i := 0; i < 3; i++ { 	// 3つのゴルーチンを生成
		wg.Add(1) // ゴルーチン生成前にカウンターをインクリメント
		go exectueInWaitGroup(i+1, &wg)
	}

	wg.Wait() // カウンターが0になるまで終了せずに待つ
}

func exectueInWaitGroup(num int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("ゴルーチン%d：%d回目\n", num, i+1)
	}

	wg.Done() // 処理が完了したらカウンターをデクリメント
}


func useChannel() {
	ch1 := make(chan string) // stringを受け取るChannel
	ch2 := make(chan string)
	ch3 := make(chan string)

	go sendToChannel(1, ch1)
	go sendToChannel(2, ch2)
	go sendToChannel(3, ch3)

	receiveFromChannel(ch1, ch2, ch3)
}

func sendToChannel(num int, ch chan<- string) {
	i := 0
	for {
		time.Sleep(1000 * time.Millisecond)

		i++
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(10)

		if x == 0 { // 1割の確率でChannelをcloseしてループを抜ける
			close(ch)
			break
		}

		ch <- fmt.Sprintf("ゴルーチン%d：%d回目\n", num, i) // Channelに送信
	}
}

func receiveFromChannel(ch1 <-chan string, ch2 <-chan string, ch3 <-chan string) {
	//Channelがcloseされたら終了する
	for c := range ch1 {
		fmt.Println(c)
	}
	for c := range ch2 {
		fmt.Println(c)
	}
	for c := range ch3 {
		fmt.Println(c)
	}
}
