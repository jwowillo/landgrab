// Package game implements the landgrab game in a functional fashion.
//
// The initial State of a game is created with the Rules for the game and the
// Players that will be involved in the game. From that point, nothing is
// mutated. A function that advances to the next State clones the State to
// insure it isn't modified and returns a new State representing the game after
// the current Player has moved.
//
// For performance reasons, some internal functions mutate but only for
// performance reasons. These mutations are never visible outside the package.
package game
