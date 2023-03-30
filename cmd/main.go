package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	poker "github.com/mroobert/larvis"
	"github.com/mroobert/larvis/game"
	"github.com/mroobert/larvis/inputerrs"
	"github.com/mroobert/larvis/rank"
	"github.com/mroobert/larvis/tiebreak"
)

type config struct {
	csv   bool
	hand1 string
	hand2 string
}

func main() {
	var cfg config

	flag.BoolVar(&cfg.csv, "csv", false, "Run with hands provided from an internal CSV file")
	flag.StringVar(&cfg.hand1, "hand1", "", "Provides first hand")
	flag.StringVar(&cfg.hand2, "hand2", "", "Provides second hand")
	flag.Parse()

	if cfg.csv {
		file, err := os.Open("./games.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan() // skip header
		for scanner.Scan() {
			line := scanner.Text()
			hands := strings.Split(line, ",")

			hand1, hand2, errs := parseHands(hands[0], hands[1])
			if len(errs.Errors) > 0 {
				log.Fatal(errs)
			}

			ranker := rank.NewRanker()
			decider := tiebreak.NewDecider()
			p := poker.NewPoker(hand1, hand2, ranker, decider)
			res, err := p.Play()
			if err != nil {
				log.Fatalf("failed to play: %v", err)
			}
			log.Printf("Hand1: %q; Hand2: %q => %s\n", hand1, hand2, res)
		}
	} else {
		hand1, hand2, errs := parseHands(cfg.hand1, cfg.hand2)
		if len(errs.Errors) > 0 {
			log.Fatal(errs)
		}

		ranker := rank.NewRanker()
		decider := tiebreak.NewDecider()
		p := poker.NewPoker(hand1, hand2, ranker, decider)
		res, err := p.Play()
		if err != nil {
			log.Fatalf("failed to play: %v", err)
		}
		log.Printf("Hand1: %q; Hand2: %q => %s\n", hand1, hand2, res)
	}

}

func parseHands(hand1, hand2 string) (game.Hand, game.Hand, *inputerrs.InputErrors) {
	inputErrs := inputerrs.NewInputErrors()

	if hand1 == "" && hand2 == "" {
		inputErrs.AddError(inputerrs.Hand1Key, "empty hand is not permitted")
		inputErrs.AddError(inputerrs.Hand2Key, "empty hand is not permitted")
	}

	h1, err := game.CreateHand([]rune(hand1))
	if err != nil {
		inputErrs.AddError(inputerrs.Hand1Key, fmt.Sprintf("failed to create hand1: %v", err))
	}

	h2, err := game.CreateHand([]rune(hand2))
	if err != nil {
		inputErrs.AddError(inputerrs.Hand2Key, fmt.Sprintf("failed to create hand2: %v", err))
	}

	return h1, h2, inputErrs
}
