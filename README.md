# Variation of a poker game

Cards:

- Each player gets a hand of 5 cards
- Available symbols are `23456789TJQKA`, in order of value (`A` being the more
  valuable)

Combinations in order of value:

- Four of a kind, like `77377`
- Full house, means 3 of a kind, and 2 of a kind, in the same hand, like `KK2K2`
- Triple, like `32666`
- Two pairs, like `77332`
- A pair, like `43K9K`
- High card, when there's none of the above, like `297QJ`

When two hands have the same combination value, the combination with the higher
cards wins. When there are two pairs, the higher pair is compared first, so
`99662` wins against `88776`. When comparing full houses, triples go first, then
pairs, so `88822` defeats `QQ777`. If the combinations are identical, then the
other cards are compared, highest to lowest, so `7T2T6` wins against `TT753`,
because `TT = TT`, `7 = 7`, but `6 > 5`.

Some examples of hands and what the result should be.

| Hand 1  | Hand 2  | Winner |
| :-----: | :-----: | :----: |
| `AAAQQ` | `QQAAA` |  Tie   |
| `53QQ2` | `Q53Q2` |  Tie   |
| `53888` | `88385` |  Tie   |
| `QQAAA` | `AAAQQ` |  Tie   |
| `Q53Q2` | `53QQ2` |  Tie   |
| `88385` | `53888` |  Tie   |
| `AAAQQ` | `QQQAA` | Hand 1 |
| `Q53Q4` | `53QQ2` | Hand 1 |
| `53888` | `88375` | Hand 1 |
| `33337` | `QQAAA` | Hand 1 |
| `22333` | `AAA58` | Hand 1 |
| `33389` | `AAKK4` | Hand 1 |
| `44223` | `AA892` | Hand 1 |
| `22456` | `AKQJT` | Hand 1 |
| `99977` | `77799` | Hand 1 |
| `99922` | `88866` | Hand 1 |
| `9922A` | `9922K` | Hand 1 |
| `99975` | `99965` | Hand 1 |
| `99975` | `99974` | Hand 1 |
| `99752` | `99652` | Hand 1 |
| `99752` | `99742` | Hand 1 |
| `99753` | `99752` | Hand 1 |
| `88822` | `QQ777` | Hand 1 |
| `99662` | `88776` | Hand 1 |
| `QQQAA` | `AAAQQ` | Hand 2 |
| `53QQ2` | `Q53Q4` | Hand 2 |
| `88375` | `53888` | Hand 2 |
| `QQAAA` | `33337` | Hand 2 |
| `AAA58` | `22333` | Hand 2 |
| `AAKK4` | `33389` | Hand 2 |
| `AA892` | `44223` | Hand 2 |
| `AKQJT` | `22456` | Hand 2 |
| `77799` | `99977` | Hand 2 |
| `88866` | `99922` | Hand 2 |
| `9922K` | `9922A` | Hand 2 |
| `99965` | `99975` | Hand 2 |
| `99974` | `99975` | Hand 2 |
| `99652` | `99752` | Hand 2 |
| `99742` | `99752` | Hand 2 |
| `99752` | `99753` | Hand 2 |

# Run

You can run it by using your `local GO installation` or by using the attached `Dockerfile`.

Get repo:

```shell
$ git clone https://github.com/mroobert/larvis.git
```

## Using `GO`

Run a game:

```shell
$ cd cmd

$ go run main.go -hand1 AAAQQ -hand2 QQAAA
```

Run multiple games defined in the `./cmd/games.csv` file:

```shell
$ cd cmd

$ go run main.go -csv true
```

## Using `Docker`

Build a docker image:

```shell
$ docker build -t poker-larvis .
```

Run a game:

```shell
$ docker run poker-larvis -hand1 AAAQQ -hand2 QQAAA
```

Run multiple games defined in the `.cmd/games.csv` file:

```shell
$ docker run poker-larvis -csv true
```
