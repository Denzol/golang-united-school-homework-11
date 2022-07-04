package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	ch1 := make(chan int64, n)
	ch2 := make(chan user, n)
	for i := 0; i < int(pool); i++ {
		go func() {
			for id := range ch1 {
				ch2 <- getOne(id)
			}
		}()
	}
	go func() {
		for user := range ch2 {
			res = append(res, user)
			wg.Done()
		}
	}()
	for j := 0; j < int(n); j++ {
		wg.Add(1)
		ch1 <- int64(j)
	}
	wg.Wait()
	close(ch1)
	close(ch2)
	return res
}
