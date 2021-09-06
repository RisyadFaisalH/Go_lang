package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 10

type all struct{
	dice int
	dicepicked bool
}
var array [N]all

func main() {
	
	var dealerdice, diceindex, score, player, dealer int
	var name, play string
	//var a bool

		fmt.Println("")
		fmt.Print("Enter your name : ")
		fmt.Scanln(&name)
 		fmt.Println("")
 		fmt.Println("WELCOME TO THE GAME", name)
 		fmt.Println("")
 		fmt.Printf("press %q to roll the dice and type %q to end the game\n", "Enter", "QUIT")
		fmt.Scanln(&play)
	
	for play != "QUIT" {


		rand.Seed(time.Now().UnixNano())
	    fmt.Print("Player dice : ")
	    fmt.Println()
	
		for l := 1; l < 9; l++ {
			fmt.Print(" ")
			array[l].dice = rolldice()
			fmt.Printf("Dice (%d) = %d\n", l, array[l].dice)
		}
		fmt.Println()	
	
        //a = true		
 
		for score < 21 {
			dealerdice = rolldice()
			dealer++
			fmt.Println("Dealer dice : ", dealerdice)
			score = score + dealerdice
			fmt.Println("Total score =", score)

			if score < 21 {	
				fmt.Print("Choose your dice : ")
				fmt.Scan(&diceindex)
				for array[diceindex].dicepicked {						
					fmt.Scan(&diceindex)
					array[diceindex].dicepicked = false

				}			  						
			    player++
				score = score + (array[diceindex].dice)
				fmt.Println("Total score =", score)
			}else{
				fmt.Scanln(&play)
 				play = "QUIT" 
 				
		    }
	    }	
		play = "QUIT"
		fmt.Printf("Congrats you get %d point(s)\n", countpoint(player, score, dealer))
	}
  }

func rolldice() int {
	
	return rand.Intn(6) + 1
}


func countpoint(player, score, dealer int) int {
	
	var points int
	
	if player == 8 && score == 21 {
		
		points = 150
	
	} else if player < 8  && score == 21 {
		
		points = 100
	
	} else if player < dealer && score > 21 {
		
		points = 70
	
	} else if player == 8 && score > 21 {
		
		points = 30
	
	} else if dealer > player && score > 21 {
		
		points = 10
	
	} else {
		
		points = 0
	}
	
	return points
}