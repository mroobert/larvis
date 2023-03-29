package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mroobert/larvis"
)

type config struct {
	csv   bool
	hand1 string
	hand2 string
}

func main() {
	var cfg config

	flag.BoolVar(&cfg.csv, "csv", false, "run with hands from internal CSV file")
	flag.StringVar(&cfg.hand1, "hand1", "", "first hand")
	flag.StringVar(&cfg.hand2, "hand2", "", "second hand")
	flag.Parse()

	if cfg.csv {
		file, err := os.Open("./cmd/games.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan() // skip header
		for scanner.Scan() {
			line := scanner.Text()
			hands := strings.Split(line, ",")

			hand1, hand2, err := parseHands(hands[0], hands[1])
			if err != nil {
				log.Fatal(err)
			}
			game := larvis.NewGame(hand1, hand2)
			log.Printf("Hand1: %q; Hand2: %q => %s\n", hand1, hand2, game.Play())
		}
	} else {

		hand1, hand2, err := parseHands(cfg.hand1, cfg.hand2)
		if err != nil {
			log.Fatal(err)
		}
		game := larvis.NewGame(hand1, hand2)
		log.Printf("Hand1: %q; Hand2: %q => %s\n", hand1, hand2, game.Play())
	}

}

func parseHands(hand1, hand2 string) (larvis.Hand, larvis.Hand, error) {
	if hand1 == "" || hand2 == "" {
		return nil, nil, errors.New("invalid hands")
	}

	h1, err := larvis.CreateHand([]rune(hand1))
	if err != nil {
		return larvis.Hand{}, larvis.Hand{}, fmt.Errorf("failed to create hand1: %w", err)
	}
	h2, err := larvis.CreateHand([]rune(hand2))
	if err != nil {
		return larvis.Hand{}, larvis.Hand{}, fmt.Errorf("failed to create hand2: %w", err)
	}

	return h1, h2, nil
}
