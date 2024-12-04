package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"types/exercise"
)

func main() {
	Lecture()

	//Accept Interfaces, Return Structs
	//using StringProcessor
	stringResult := ProcessData(StringProcessor{}, "Hello world!")
	fmt.Printf("String Processor Result: %+v\n", stringResult)

	//Using JSONProcessor
	jsonResult := ProcessData(JSONProcessor{}, "Hello world!")
	fmt.Printf("JSON Processor Result: %+v\n", jsonResult)
}

func Lecture() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	// ? -> Exercise -----------------------------
	teams := []exercise.Team{
		{Teamname: "Manchester United", Playername: []string{"Rashford, Fernandes, Amad"}},
		{Teamname: "Arsenal", Playername: []string{"Saka, Odegard, Martinelli"}},
	}

	league := exercise.League{
		Teams: teams,
		Wins:  make(map[string]int),
	}

	league.MatchResult("Manchester United", 2, "Arsenal", 1)
	league.MatchResult("Manchester United", 2, "Arsenal", 2)
	league.MatchResult("Manchester United", 4, "Arsenal", 2)
	league.MatchResult("Manchester United", 5, "Arsenal", 2)

	fmt.Println(league.Wins["Manchester United"])
	exercise.RankPrinter(league, os.Stdout)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total,
		c.lastUpdated)
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s,  age %d", p.FirstName, p.LastName, p.Age)
}

// ------------------
// DataProcessor interface defines the contract for processing data
type DataProcessor interface {
	Process(input string) string
}

// StringProcessor implements the DataProcessor interface
type StringProcessor struct{}

func (sp StringProcessor) Process(input string) string {
	return strings.ToUpper(input)
}

// JSONProcessor implements the DataProcessor interface
type JSONProcessor struct{}

func (jp JSONProcessor) Process(input string) string {
	return fmt.Sprintf(`{"processed": "%s"}`, input)
}

// Result is a concrete struct that can evolve over time
type Result struct {
	ProcessedData string
	Length        int
	// We can add more fields later without breaking existing code
	Timestamp string
}

// ProcessData accepts any type that implements DataProcessor interface
// But returns a concrete Result struct
func ProcessData(processor DataProcessor, input string) Result {
	processedData := processor.Process(input)

	return Result{
		ProcessedData: processedData,
		Length:        len(processedData),
		// Initially, we might not even include this field
		// Timestamp: time.Now().Format(time.RFC3339),  // Can be added later
	}
}

// ? ----------------------------------
func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

//factory function to create an instance of a simpleDataStore:

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Matthew",
			"2": "John",
			"3": "Phillipians",
		},
	}
}

//business logic to hookup a usera and say hello or goodbye

type DataStore interface {
	UserNameForId(userId string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LogAdapter func(message string)

func (lg LogAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) sayHello(userId string) (string, error) {
	sl.l.Log("in SayHello for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unkown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodBye(userId string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unkown user")
	}
	return "GoodBye, " + name, nil
}
