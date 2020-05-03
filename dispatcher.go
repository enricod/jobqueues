package main

import "fmt"

// WorkerQueue vedi http://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/
// passiamo una richiesta che contiene un canale in nui il ricevente scriverà
var WorkerQueue chan chan WorkRequest

// StartDispatcher avviamo tutti i worker
func StartDispatcher(nworkers int) {

	WorkerQueue = make(chan chan WorkRequest, nworkers)

	// creiamo n worker queue
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker ", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request")
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
