package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
	_____                           
	|  |  |___ ___ ___ _____ ___ ___ 
	|     | .'|   | . |     | .'|   |
	|__|__|__,|_|_|_  |_|_|_|__,|_|_|
				  |___|              
	`)
}

func Draw(g *Game, guess string) {
	//Afficher l'état courant de l'echaffaud
	drawTurns(g.TurnsLeft)

	//Etat de la partie
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		+---+
		|   |
		O   |
	   /|\  |
	   / \  |
			|
		========='''
				`

	case 1:
		draw = `
		+---+
		|   |
		O   |
	   /|\  |
	        |
			|
		========='''
				`
	case 2:
		draw = `
		+---+
		|   |
		O   |
	    |   |
			|
			|
		=========''' 
				`
	case 3:
		draw = `
		+---+
		|   |
		O   |
	        |
			|
			|
		========='''  
			`
	case 4:
		draw = `
		+---+
		|   |
		    |
		    |
			|
			|
		=========''' 
				`
	case 5:
		draw = `
		+---+
		    |
		    |
			|
			|
			|
		=========''' 
			`

	case 6:
		draw = `
		    |
			|
			|
			|
			|
		========='''`
	case 7:
		draw = `
		========='''
			`
	case 8:
		draw = `
			
			`
	}

	fmt.Println(draw)

}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed : ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used : ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess !\n")
	case "alreadyGuessed":
		fmt.Printf("letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, letter '%s' is not in the word\n", guess)
	case "lost":
		fmt.Print("You lost ! The word was : ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You won ! The word was : ")
		drawLetters(g.Letters)
	}
}
