package main

import (
	"fmt"
	"log"
)

func socialNetwork() {
	defer fmt.Println("FaceBook") //defer keyword makes this line Won't run untill all the remaining code in socialNetwork function runs
	fmt.Println("Youtube")
	fmt.Println("Instagram")
}

func blockchainTech() error {
	defer fmt.Println("Use it to validate the products") // This line will be called when blockchainTech() returns
	fmt.Println("Scam by making useless Coin")
	return fmt.Errorf("Suu! don't say useless!")
	fmt.Println("NFTs yeah Baby!!!")
	return nil
}

func main() {
	socialNetwork()
	err := blockchainTech()
	if err != nil {
		log.Fatal(err)
	}
}
