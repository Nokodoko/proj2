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

type name struct{
    field type
}

type FightType struct {
	Movement int  //Attack speed (scale of 1 to 10)
	Weapon   bool //Does this charcter use a weapon: if yes, add kill method will weapon
	Strength int  //Take more LifePoints points per hit
	Finish   FinishMoves
}

type Person struct {
	Name          string
	Aliases       []string
	LifePoints    int
	FightingStyle fightType
}
