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
	number := int(n)
	parallel := int(pool)
	var wg sync.WaitGroup
	sem := make(chan struct{}, parallel)
	for i := 0; i < number; i++ {
		wg.Add(1)
		sem <- struct{}{}
		time.Sleep(time.Millisecond * 2)
		go func(j int) {
			user := getOne(int64(j))
			res = append(res, user)
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}
