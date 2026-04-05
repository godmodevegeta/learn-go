package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println("Hello, there!")
//	coin := []int{1,2,3}
//	target := 5
//	fmt.Println("Minimum coins needed:", min_coin(coin, target))
//
//	coin = []int{1,2,3}
//	target = 32
//	fmt.Println("Minimum coins needed:", min_coin(coin, target))

	Alice := person{
		FirstName: "Alice",
		LastName: "Jones",
		Age: 23,
	}
	fmt.Println(Alice.String())
	fmt.Println("=========")
	
	var c Counter
	fmt.Println(c.String())
	c.Increment()	
	fmt.Println(c.String())

	fmt.Println("======Methods are functions too=====")

	adder := Adder{
		start: 5,
	}
	fmt.Println(adder.AddTo(10))
	fmt.Println("Demonstration of method value")
	f1 := adder.AddTo
	fmt.Println(f1(11))
	fmt.Println("Demonstration of method expression")
	f2 := Adder.AddTo
	fmt.Println(f2(adder, 12))

	fmt.Println("======== iota as enumerations ==========")

	fmt.Println(Uncategorized, Personal, Spam, Social, Advertisments)
	fmt.Println(Field1, Field2, Field3, Field4, Field5, Field6)

	fmt.Println("==== Use Embedding for Composition =====")
	manager := Manager {
		Employee: Employee{
			Name: "Bob",
			ID: "246",
		},
	}
	fmt.Printf("I'm the useless Manager. My name is %s, and my ID is %s.\n", manager.Name, manager.ID)
	
	fmt.Println("===== Embedding is not Inheritance =====")

  o := Outer{}
	o.CallHello()

	fmt.Println("====== Interfaces are Type-Safe Duck Typing ======")
	
	

	fmt.Println("_____ Chapter 7: Ex 1 ______")

	teamA := Team{
		TeamName:   "Lakers",
		PlayerNames: []string{"LeBron", "AD", "Reaves"},
	}

	teamB := Team{
		TeamName:   "Bulls",
		PlayerNames: []string{"LaVine", "DeRozan", "Vucevic"},
	}

	teamC := Team{
		TeamName:   "Warriors",
		PlayerNames: []string{"Curry", "Klay", "Draymond"},
	}

// Create league
	league := League{
		Teams: []Team{teamA, teamB, teamC},
		Wins:  make(map[string]int),
	}

	// Simulate matches
	league.MatchResult("Lakers", 102, "Bulls", 98)
	league.MatchResult("Warriors", 120, "Lakers", 115)
	league.MatchResult("Bulls", 110, "Warriors", 105)
	league.MatchResult("Lakers", 130, "Warriors", 136)
	league.MatchResult("Lakers", 130, "Bulls", 136) 
	league.MatchResult("Bulls", 130, "Warriors", 136)

	fmt.Println("League Final Championship")
	for _,team := range league.Teams {
		fmt.Println("Team", team.TeamName, ":", league.Wins[team.TeamName], "wins")
	}
	ranking := league.Ranking()
	for i, team := range ranking{
		switch i{ 
			case 0:
				fmt.Println("Winner:", team)
			case 1:
				fmt.Println("Runner Up:", team)
			default:
				fmt.Println("Numeber", i, ":", team)
			}
	}

}

type Team struct {
	TeamName string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins map[string]int
}

func (l *League)MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if score1 > score2 {
		l.Wins[team1] += 1
	} else if score2 > score1 {
		l.Wins[team2] += 1
	}
}

func (l *League)Ranking() []string {
	var response []string

	return response
}

type Incrementer interface {
	Increment()
}

type Stringer interface {
	String() string
}
type Inner struct {}
func (i Inner) SayHello() {
	fmt.Println("Hello from Inner")
}
func (i Inner) CallHello() {
	i.SayHello()
}
type Outer struct {
	Inner
}
func (o Outer) SayHello() {
	fmt.Println("Hello from Outer")
}

type MailCategory int
const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisments
)

const (
	Field1 = 0
	Field2 = 1 + iota	
	Field3 = 20
	Field4
	Field5 = iota
	Field6
)

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}
type IntTree struct {
	val int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{
			val: val,
		}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

type Counter struct {
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}
func (c Counter) String() string{
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type person struct {
	FirstName string
	LastName string
	Age int
}

type Score int
type Converter func(string)Score
type TeamScores map[string]Score

type HighScore Score

type Employee struct {
	Name string
	ID string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (p person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
func min_coin(coins []int, target int) int {
	if target == 0 {
		return 0
	}
	min_coins := 99999
	for _, coin := range(coins) {
		if coin <= target {
			min_coins = min(min_coins, min_coin(coins, target - coin) + 1)
		}
	}
	return min_coins
}



