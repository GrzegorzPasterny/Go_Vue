package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

// const hunger = 3                // Number of times each philosopher eats
// const think = time.Second / 100 // Mean think time
// const eat = time.Second / 100   // Mean eat time

var dining sync.WaitGroup

type Input struct {
	Name                   string `json:"Name"`
	TimeToEat              string `json:"TimeToEat"`
	TimeToThink			   string `json:"TimeToThink"`	
	HowManyDishesToBeEaten string `json:"HowManyDishesToBeEaten"`
}

type Inputs []Input

var id = 0

type PhilosopherOutput struct {
	Id 			int `json:"id"`
	Name     string `json:"Name"`
	Status string `json:"Status"`
	TimeStamp string `json:"TimeStamp"`
}

func main() {

	http.HandleFunc("/DiningPhilosophers", homePageHandler)
	
	fs := http.FileServer(http.Dir("../philosophers_frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func SendAsJson(pl PhilosopherOutput, w http.ResponseWriter) {
	philosophersLogOut, err := json.Marshal(pl)
	
	if err != nil {
		panic(err)
	}
	
	w.Write(philosophersLogOut)
	comma := []byte(",")
	w.Write(comma)
}

func diningProblem(ph Input, dominantHand, otherHand *sync.Mutex, w http.ResponseWriter) {
	
	var philosophersLog PhilosopherOutput

	id++
	philosophersLog = PhilosopherOutput{Id: id, Name: ph.Name, Status: "Seated", TimeStamp: time.Now().Format("15:04:05.99999999")}
	SendAsJson(philosophersLog, w)

	timeToEat, err := strconv.Atoi(ph.TimeToEat)
	if err != nil {
        fmt.Println(err)
	}

	timeToThink, err := strconv.Atoi(ph.TimeToThink)
	if err != nil {
        fmt.Println(err)
	}

	howManyDishesToBeEaten, err := strconv.Atoi(ph.HowManyDishesToBeEaten)
	if err != nil {
        fmt.Println(err)
	}
	

	think := time.Duration(timeToThink) * time.Second / 100 
	eat := time.Duration(timeToEat) * time.Second / 100

	h := fnv.New64a()
	h.Write([]byte(ph.Name))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := howManyDishesToBeEaten; h > 0; h-- {
		id++
		philosophersLog = PhilosopherOutput{Id: id, Name:ph.Name, Status:"Hungry", TimeStamp: time.Now().Format("15:04:05.99999999")}
		SendAsJson(philosophersLog, w)
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		id++
		philosophersLog = PhilosopherOutput{Id: id, Name:ph.Name, Status:"Eating", TimeStamp: time.Now().Format("15:04:05.99999999")}
		SendAsJson(philosophersLog, w)
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		id++
		philosophersLog = PhilosopherOutput{Id: id, Name:ph.Name, Status:"Thinking", TimeStamp: time.Now().Format("15:04:05.99999999")}
		SendAsJson(philosophersLog, w)
		rSleep(think)
	}
	id++
	philosophersLog = PhilosopherOutput{Id: id, Name:ph.Name, Status:"Satisfied", TimeStamp: time.Now().Format("15:04:05.99999999")}
	SendAsJson(philosophersLog, w)
	dining.Done()
	id++
	philosophersLog = PhilosopherOutput{Id: id, Name:ph.Name, Status:"Left the table", TimeStamp: time.Now().Format("15:04:05.99999999")}
	SendAsJson(philosophersLog, w)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var inputs []Input

	err := json.NewDecoder(r.Body).Decode(&inputs)
	
	if err != nil {
		
		// if error is not nil
		// print error
		fmt.Println(err)
	}
	
	a := []byte("[")
	w.Write(a)

	dining.Add(len(inputs))
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(inputs); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(inputs[i], forkLeft, forkRight, w)
		forkLeft = forkRight
	}
	
	go diningProblem(inputs[0], fork0, forkLeft, w)
	
	dining.Wait() // wait for philosphers to finish
	
	id++
	philosophersLog := PhilosopherOutput{Id: id, Name:"Feast ended", Status:"Everyone's satisfied", TimeStamp: time.Now().Format("15:04:05.99999999")}
	philosophersLogOut, err := json.Marshal(philosophersLog)
	
	if err != nil {
		panic(err)
	}
	
	w.Write(philosophersLogOut)

	a = []byte("]")
	w.Write(a)
	id=0
	//w.WriteHeader(http.StatusOK)

}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}