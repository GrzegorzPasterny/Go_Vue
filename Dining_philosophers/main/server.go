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
	Name     []string `json:"Name"`
	Statuses []string `json:"Status"`
}

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func diningProblem(phName string, dominantHand, otherHand *sync.Mutex, w http.ResponseWriter) []string {

	var subLogs2 []string
	subLogs2 = append(subLogs2, fmt.Sprintf(phName, "Seated"))

	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		subLogs2 = append(subLogs2, fmt.Sprintf(phName, "Hungry"))
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		subLogs2 = append(subLogs2, fmt.Sprintf(phName, "Eating"))
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		subLogs2 = append(subLogs2, fmt.Sprintf(phName, "Thinking"))
		rSleep(think)
	}
	subLogs2 = append(subLogs2, fmt.Sprintf(phName, "Satisfied"))
	dining.Done()
	subLogs2 = append(subLogs2, fmt.Sprintf(phName, "Left the table"))

	return subLogs2
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	var inputs []Input

	Data := []byte(` 
    [ 
        {"Name": "John", "TimeToEat": "3", "HowManyDishesToBeEaten": "3"}, 
        {"Name": "Marta", "TimeToEat": "4", "HowManyDishesToBeEaten": "2"}
    ]`)

	err := json.Unmarshal(Data, &inputs)

	inputsJson, err := json.Marshal(inputs)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(inputsJson)

	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// printing decoded array
	// values one by one
	for i := range inputs {
		fmt.Println(inputs[i].Name + " - " + inputs[i].TimeToEat +
			" - " + inputs[i].HowManyDishesToBeEaten)
	}

	var philosopherLogs PhilosopherOutput
	var sublogs [][]string

	philosopherLogs.Name = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

	philosopherLogs.Statuses = append(philosopherLogs.Statuses, "Table empty")
	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go func(){
			sublogs = append(sublogs, diningProblem(ph[i], forkLeft, forkRight, w))
		}()
		forkLeft = forkRight
	}
	var subLogs3 []string
	go func() {
		subLogs3 = diningProblem(ph[0], fork0, forkLeft, w)
	}()

	dining.Wait() // wait for philosphers to finish
	philosopherLogs.Statuses = append(philosopherLogs.Statuses, sublogs[1]...)
	philosopherLogs.Statuses = append(philosopherLogs.Statuses, "BREAK!!!--------------------------------------")
	philosopherLogs.Statuses = append(philosopherLogs.Statuses, subLogs3...)
	philosopherLogs.Statuses = append(philosopherLogs.Statuses, "Table empty")
	philosopherOutput, err := json.Marshal(philosopherLogs)

	if err != nil {
		panic(err)
	}

	w.Write(philosopherOutput)
}
