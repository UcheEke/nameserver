package main

import (
	"bufio"
	"os"
	"log"
	"strings"
	"math/rand"
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
}

// dataStrip(): Data Processing function. Processes string data in file <filename> using a specific process function
func dataStrip(filename, sep string, processFunc func(str, sep string) []string) []string {
	fd, err := os.Open(filename)
	errorCheck(err)

	var (
		line []string
		data []string
	)

	defer fd.Close()
	{
		s := bufio.NewScanner(fd)
		for (s.Scan()) {
			str := s.Text()
			line = processFunc(str,sep)
			for _, datapoint := range line {
				data = append(data, string(datapoint))
			}
		}
	}
	return data
}

// Process Functions: Functions used by dataStrip() to handle the specific requirements of webscraped data
// getFirstnames processes a list of 400+ popular firstnames scraped from
// http://nameberry.com/popular_names/Nameberry
func getFirstnames(str string, sep string) []string {
	result := []string{}
	line := strings.Split(str, sep)
	if len(line) == 3 {
		result = append(result, strings.TrimSpace(line[1]))
		result = append(result, strings.TrimSpace(line[2]))
	}
	//fmt.Println(result)
	return result
}

// getLastnames processes a list of popular 300 european lastnames scraped from
// http://en.geneanet.org/genealogy/1/Surname.php
func getLastnames(str string, sep string) []string {
	result := []string{}
	line := strings.Split(str, sep)
	var lastname string
	if len(line) == 2 {
		lastname = strings.TrimSpace(line[0])
		lastname = strings.ToLower(lastname)
		lastname = strings.Title(lastname)
		result = append(result, lastname)
	}
	return result
}


type Person struct {
	ID int			`json:"id"`
	Firstname string 	`json:"first_name"`
	Lastname string		`json:"last_name"`
	Age int			`json:"age"`
}

type PeopleServer struct {
	firstnames []string
	lastnames []string
}

func (ps *PeopleServer) ServeHTTP(w http.ResponseWriter, req *http.Request){
	var people []Person

	for i := 0; i < 10;  i++ {
		people = append(people, createPerson(ps.firstnames, ps.lastnames, i+1))
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	// For testing purposes - REMOVE
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if data, err := json.Marshal(people); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error marshaling JSON: %s\n", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func createPerson(forenames, surnames []string, id int) Person {
	person := Person{}

	var baseAge, ageRange int
	baseAge = 18
	ageRange = 47

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	person.Firstname = forenames[rand.Intn(len(forenames))]
	person.Lastname = surnames[rand.Intn(len(surnames))]
	person.Age = baseAge + rand.Intn(ageRange)
	person.ID = id

	return person
}

func main(){
	mux := http.NewServeMux()
	ps := PeopleServer{}

	ps.firstnames = dataStrip("firstnames.txt","\t", getFirstnames)
	ps.lastnames = dataStrip("lastnames.txt","\t", getLastnames)

	mux.Handle("/", &ps)

	port := 8001
	fmt.Printf("Listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}


