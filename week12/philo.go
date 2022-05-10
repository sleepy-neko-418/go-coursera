package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	number int
	sync.Mutex
}

type Philosopher struct {
	number         int
	eatNumber      int
	leftChopStick  *ChopStick
	rightChopStick *ChopStick
}

type eatingTicket int

func main() {
	const PHILO_NUM = 5
	const CONCURRENT_EATING = 2

	wg := new(sync.WaitGroup)

	// Create the philosophers and chopsticks
	chops := make([]*ChopStick, PHILO_NUM)
	philos := make([]*Philosopher, PHILO_NUM)

	for i := 0; i < PHILO_NUM; i++ {
		chops[i] = &ChopStick{number: i + 1}
	}

	for i := 0; i < PHILO_NUM; i++ {
		philos[i] = &Philosopher{
			number:         i + 1,
			eatNumber:      0,
			leftChopStick:  chops[i],
			rightChopStick: chops[(i+1)%PHILO_NUM],
		}
	}

	// Channel to control eating permission
	eatingPermission := make(chan eatingTicket, CONCURRENT_EATING)
	// Insert two eating tickets into the permission channel
	for i := 1; i <= CONCURRENT_EATING; i++ {
		eatingPermission <- eatingTicket(i)
	}

	// Feast
	for i := 0; i < PHILO_NUM; i++ {
		wg.Add(1)
		go philos[i].eat(wg, eatingPermission)
	}

	wg.Wait()
}

func (p *Philosopher) eat(wg *sync.WaitGroup, eatingPermission chan eatingTicket) {
	for {
		// Acquire eating ticket
		ticket := <-eatingPermission
		fmt.Println("Philosopher no", p.number, "acquired eating ticket no", ticket)
		if p.leftChopStick.TryLock() {
			if p.rightChopStick.TryLock() {
				fmt.Println("Starting to eat", p.number)
				fmt.Println("Finishing eating", p.number)
				p.eatNumber += 1
				p.rightChopStick.Unlock()
			}
			p.leftChopStick.Unlock()
		}

		// Return the ticket for other philosopher
		eatingPermission <- ticket
		fmt.Println("Philosopher no", p.number, "returned eating ticket no", ticket)

		if p.eatNumber == 3 {
			fmt.Println(p.number, "is full. No more eating")
			wg.Done()
			return
		}
	}
}
