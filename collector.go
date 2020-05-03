package main

import (
	"net/http"
	"time"
)

// WorkQueueWorkRequestChan A buffered channel that we can send work requests on.
var WorkQueueWorkRequestChan = make(chan WorkRequest, 100)

// Collector funzione invocata dal web server
// crea una WorkRequest e la invia alla WorkQueue
func Collector(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the delay.
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check to make sure the delay is anywhere from 1 to 10 seconds.
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	// Now, we retrieve the person's name from the request.
	name := r.FormValue("name")

	// Just do a quick bit of sanity checking to make sure the client actually provided us with a name.
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
	workReq := WorkRequest{Name: name, Delay: delay}

	// Push the work onto the queue.
	WorkQueueWorkRequestChan <- workReq
	//fmt.Println("work request queued")

	// And let the user know their work request was created.
	w.WriteHeader(http.StatusCreated)
	return
}
