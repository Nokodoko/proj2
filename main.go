package main

import "fmt"

//I am going to make characters and then have them perform an action against eachother. The goal of this is to:
//1. struct creation
//2. interfaces (polymorphism)
//3. switch statements
func main() {
	fmt.Println("vim-go")
}

//TODO:
//1. Make character (struct)
//2. Make actions (methods)
//3. Create scenarios (switches)

type FinishMoves struct {
	Painless string
	Painful  string
	Brutal   string
}

type FightType struct {
	Movement   int  //Attack speed (scale of 1 to 10)
	Weapon     bool //Does this charcter use a weapon: if yes, add kill method will weapon
	Strength   int  //Take more LifePoints points per hit
	ComboChain string
	Finish     FinishMoves
}

type Person struct {
	name         string
	relationship Fighter
	LifePoints   int
}

func (p Person) loseLife() int {
	return 0
}

func (p Person) attack(p2 Fighter) int {
	return 0
}

type Fighter struct {
	Name          string
	Aliases       []string
	LifePoints    int
	FightingStyle fightType
}

func (p Fighter) attack(p2 Fighter) int {
	return 0
}

func (p Fighter) loseLife() int {
	return 0
}

type Player interface {
	attack()
	loseLife()
}
