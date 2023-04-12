package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mroobert/larvis"
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
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Game failed: %v", r)
		}
	}()

	var cfg config
	flag.BoolVar(&cfg.csv, "csv", false, "Run with hands provided from an internal CSV file")
	flag.StringVar(&cfg.hand1, "hand1", "", "Provides first hand")
	flag.StringVar(&cfg.hand2, "hand2", "", "Provides second hand")
	flag.Parse()

	parseFlags(cfg)

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
				log.Printf("Hand1: %q; Hand2: %q %v", hands[0], hands[1], errs)
				fmt.Println()
				continue
			}

			ranker := rank.NewRanker()
			decider := tiebreak.NewDecider()
			g := game.NewGame(hand1, hand2, ranker, decider)
			res, err := g.Play()
			if err != nil {
				log.Printf("Hand1: %q; Hand2: %q\nfailed to play: %v", hand1, hand2, err)
				continue
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
		p := game.NewGame(hand1, hand2, ranker, decider)
		res, err := p.Play()
		if err != nil {
			log.Fatalf("failed to play: %v", err)
		}
		log.Printf("Hand1: %q; Hand2: %q => %s\n", hand1, hand2, res)
	}

}

func parseFlags(cfg config) {
	if cfg.hand1 == "" && cfg.hand2 == "" && !cfg.csv {
		flag.Usage()
		os.Exit(1)
	}
}

func parseHands(hand1, hand2 string) (larvis.Hand, larvis.Hand, *inputerrs.InputErrors) {
	inErrs := inputerrs.NewInputErrors()

	h1, err := larvis.CreateHand([]rune(hand1))
	if err != nil {
		inErrs.AddError(inputerrs.Hand1Key, fmt.Sprintf("failed to create hand1: %v", err))
	}

	h2, err := larvis.CreateHand([]rune(hand2))
	if err != nil {
		inErrs.AddError(inputerrs.Hand2Key, fmt.Sprintf("failed to create hand2: %v", err))
	}

	return h1, h2, inErrs
}
