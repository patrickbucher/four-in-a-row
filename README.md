# four-in-a-row

Four in a Row implementation in Go, with bots and interactive gameplay.

## Simulation

Run the game once:

    $ go run simulation/simulation.go

Run the game multiple times (12 times):

    $ go run simulation/simulation.go -n 12

## League

Run a tournament (with match and rematch):

    $ go run league/league.go

        Rank  Player              Points     Games       Won      Lost      Tied
    --------  ----------------  --------  --------  --------  --------  --------
           1  Randy Random I.          3         2         1         1         0
           2  Randy Random II.         3         2         1         1         0

Run a tournament with multiple rounds for each match/rematch pairing:

    $ go run league/league.go -n 10

        Rank  Player              Points     Games       Won      Lost      Tied
    --------  ----------------  --------  --------  --------  --------  --------
           1  Randy Random I.         36        20        12         8         0
           2  Randy Random II.        24        20         8        12         0
