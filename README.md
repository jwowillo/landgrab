[![Go Report Card](https://goreportcard.com/badge/github.com/jwowillo/trim)](https://goreportcard.com/report/github.com/jwowillo/landgrab)
<!--- TODO: Add test coverage badge. -->

# landgrab

landgrab is a board game and AI platform with multiple interfaces for playing.

Provided interfaces are a CLI app, a web app, and mobile apps. AIs can be
compared through all interfaces.

The rules of the game are:

* Each player has n = 4 pieces on a 2n + 1 by 2n + 1 board represented as a
  grid.
* Each piece has 3 life and does 1 damage initially. Both of these are increased
  by one when a piece levels up.
* Each player can optionally move any of their pieces one space in the 8
  cardinal directions on the grid per turn.
* Collisions between pieces cause units to damage all contacting enemy units
  equal to their damage attribute.
* Pieces are removed from the board, or destroyed, if they have taken damage
  greater than or equal to their life attribute.
* Pieces gain a level when they participate in destroying another piece,
  allowing their damage and life to be boosted by a fixed amount.
* Players with no pieces at the end of a turn lose and it is possible for both
  players to lose.
* Each player has 30 seconds to make a move each turn.

## Installation

Run `make` to build all interfaces. More specifically:

* `make cli`: Builds the CLI app.
* `make web`: Builds the web app.
* `make mobile`: Builds the mobile apps.

## Documentation

* Site: http://www.jwowillo.com/landgrab/
* API Documentation: https://godoc.org/github.com/jwowillo/landgrab
* Requirements: doc/requirements.pdf
* Design: doc/design.pdf

## Tests

Run tests with `sh run_tests.sh`. Pass the `--bench` flag to also run
benchmarks.


## Running

Run the CLI with `landgrab_run_cli` after running `make run_cli`. Accepted flags
are:

* `--wait`: Prompt the user to press enter before the next turn. Defaults to
  true.
* `--player1`: Don't prompt the user for a player one and use this instead.
* `--player2`: Don't prompt the user for a player two and use this instead.

Run the web application with `landgrab_run_web` after running `make run_web`.
Accepted flags are:

* `--host`:
* `--port`:
