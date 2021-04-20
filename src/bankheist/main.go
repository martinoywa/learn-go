package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludeGuards := rand.Intn(100)

	if eludeGuards >= 50 {
		fmt.Println("Awesome! You made it past the guards. Continue. 😎")
	} else {
		isHeistOn = !isHeistOn
		fmt.Println("Darn! You got caught by guards, better luck next time. 😪")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Vault open! Grab and Go Go Go!!! 🤑🤑🤑")
	} else if isHeistOn {
		isHeistOn = !isHeistOn
		fmt.Println("Sorry, you couldn't open the vault. 😪")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0 :
			isHeistOn = !isHeistOn
			fmt.Println("Complete heist was unsuccessful! Caught on camera! 📹")
		case 1 :
			isHeistOn = !isHeistOn
			fmt.Println("Complete heist was unsuccessful! There were more guards in the vault location! 💂‍♂️")
		case 3 :
			isHeistOn = !isHeistOn
			fmt.Println("Complete heist was unsuccessful! Turns out the vault doors close automatically if someone's inside! 😨🤯")
		default :
			fmt.Println("Heist complete! Start the gateway car! 😎🤑🥳")
		}
	}

	if isHeistOn {
		amtStole := 10000 + rand.Intn(1000000)
		fmt.Printf("We scored $%d! 😎🤑🤑🤑😎 \n", amtStole)
	}

	fmt.Printf("Current heist status: %v", isHeistOn)
}