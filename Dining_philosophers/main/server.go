package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}
var sublogs []string

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

var dining sync.WaitGroup

type Input struct {
	Name                   string `json:"Name"`
	TimeToEat              string `json:"TimeToEat"`
	HowManyDishesToBeEaten string `json:"HowManyDishesToBeEaten"`
}

type Inputs []Input

type PhilosopherOutput struct {
	Name     string `json:"Name"`
	Status string `json:"Status"`
	TimeStamp string `json:"TimeStamp"`
}

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func diningProblem(phName string, dominantHand, otherHand *sync.Mutex, w http.ResponseWriter) {
	
	var philosophersLog []PhilosopherOutput

	philosophersLog = append(philosophersLog, PhilosopherOutput{Name: phName, Status: "Seated", TimeStamp: time.Now().Format("15:04:05.99999999")})
	
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		philosophersLog = append(philosophersLog, PhilosopherOutput{Name:phName, Status:"Hungry", TimeStamp: time.Now().Format("15:04:05.99999999")})
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		philosophersLog = append(philosophersLog, PhilosopherOutput{Name:phName, Status:"Eating", TimeStamp: time.Now().Format("15:04:05.99999999")})
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		philosophersLog = append(philosophersLog, PhilosopherOutput{Name:phName, Status:"Thinking", TimeStamp: time.Now().Format("15:04:05.99999999")})
		rSleep(think)
	}
	philosophersLog = append(philosophersLog, PhilosopherOutput{Name:phName, Status:"Satisfied", TimeStamp: time.Now().Format("15:04:05.99999999")})
	dining.Done()
	philosophersLog = append(philosophersLog, PhilosopherOutput{Name:phName, Status:"Left the table", TimeStamp: time.Now().Format("15:04:05.99999999")})
	
	philosophersLogOut, err := json.Marshal(philosophersLog)
	
	if err != nil {
		panic(err)
	}
	
	w.Write(philosophersLogOut)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var inputs []Input

	err := json.NewDecoder(r.Body).Decode(&inputs)
	
	if err != nil {
		
		// if error is not nil
		// print error
		fmt.Println(err)
	}
	
	inputsOut, err := json.Marshal(inputs)
	
	if err != nil {
		panic(err)
	}
	
	w.Write(inputsOut)

	// printing decoded array
	// values one by one
	for i := range inputs {
		fmt.Println(inputs[i].Name + " - " + inputs[i].TimeToEat +
		" - " + inputs[i].HowManyDishesToBeEaten)
	}
	
	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		
		go diningProblem(ph[i], forkLeft, forkRight, w)
		forkLeft = forkRight
	}
	
	go diningProblem(ph[0], fork0, forkLeft, w)
	
	dining.Wait() // wait for philosphers to finish
	//philosopherLogs.Statuses = append(philosopherLogs.Statuses, sublogs[1]...)
	
	w.WriteHeader(http.StatusOK)
}
