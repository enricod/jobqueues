package main

import "fmt"

// WorkerQueueChanChan vedi http://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/
// passiamo una richiesta che contiene un canale in nui il ricevente scriverà
var WorkerQueueChanChan chan chan WorkRequest

// StartDispatcher avviamo tutti i worker
func StartDispatcher(nworkers int) {

	WorkerQueueChanChan = make(chan chan WorkRequest, nworkers)

	// creiamo n worker queue
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker ", i+1)
		worker := NewWorker(i+1, WorkerQueueChanChan)
		worker.Start()
	}

	go func() {
		for {
			select {
			// WorkQueueWorkRequestChan è creato dal Collector
			case workReq := <-WorkQueueWorkRequestChan:
				// attende work request passate dal Collector
				fmt.Println("Received work request ", workReq.Name)
				go func() {
					worker := <-WorkerQueueChanChan

					fmt.Println("Dispatching work request ", workReq.Name)
					worker <- workReq
				}()
			}
		}
	}()
}
