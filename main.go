package main

import (
	//Notice all packages are all in alphabetical order. It needs to be that way.
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//first function. "/" should display my biography.
func bio(w http.ResponseWriter, r *http.Request) { //handler funtion 1.
	w.Write([]byte("Hi there, my name is Reneil Abner Coleman. I was born and raised in the beautiful district of Toledo until I moved to the sweet city of Belmopan for about two years now. \n"))
	w.Write([]byte("I have been a student at University of Belize since 2019 persuing my Associates Degree in Computer Science. \n"))
	w.Write([]byte("I was born on April 18th, 2002 which makes me 20 years old today. \n"))
	w.Write([]byte("I've worked two jobs so far in my life, both of which are at a call center and I can genuinely say that I've enjoyed them. \n"))
	w.Write([]byte("I really like living in belmopan, I feel like there are way more oppertunities here than my birth place in Toledo. \n"))
	w.Write([]byte("I really like playing basketball and being a social media content creator, those are my main hobbies and it has become a passion of mine. \n"))
	w.Write([]byte("I don't really like being around negative company that has to real goals and aren't passionate about something, it makes me feel that I am being misled from my purpose. \n"))
	w.Write([]byte("Another thing that bothers me would be having a hard time balancing my lifestyle, it can get stressful sometimes, but I believe as long as we can have a little fun and be productive we should be alright. \n"))
}

var ListofQuotes = []string{ // slice/ array of different quotes which will be selected at random by funtion (random)
	"Do not let yesterday take up too much of today",
	"If you’re going through hell, keep going.",
	"Every man dies. Not every man lives.",
	"We need much less than we think we need.",
	"If things go wrong, don’t go with them.",
	"Everything has beauty, but not everyone sees it.",
}

//Second funtion "/random" should display a random quote chosen from a map of quotes.
func random(w http.ResponseWriter, r *http.Request) { //handler funtion #2
	quotes := ListofQuotes[rand.Intn(len(ListofQuotes))]
	w.Write([]byte(quotes))
}

// third function "/greeting" should display time and date along with a message.
func greeting(w http.ResponseWriter, r *http.Request) { // handler function #3
	Time := time.Now().Format("3:04 pm")
	Day := time.Now().Weekday().String()
	Month := time.Now().Month().String()
	Date := time.Now()
	w.Write([]byte("The time right now is currently "))
	w.Write([]byte(Time))
	w.Write([]byte(" "))
	w.Write([]byte(Day))
	w.Write([]byte(" "))
	w.Write([]byte(Month))
	w.Write([]byte(" "))
	fmt.Println(Date.Day())
	w.Write([]byte("\nGood Day! Hope you're having a wonferul time, Enjoy they rest of your day!"))

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", bio)
	mux.HandleFunc("/random", random)
	mux.HandleFunc("/greeting", greeting)

	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
