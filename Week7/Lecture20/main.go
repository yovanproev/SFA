package main

import (
	cocktails "cocktails/bartender"
	"fmt"
)

func main() {
	fmt.Printf("What would you like to drink?\nThere are 635 drinks in the Database!\nProvide at least 3 letters in order to narrow your search or\nif you know the full name of the drink please provide it.\n")

	var newCocktail cocktails.CocktailBartender

	cocktails.DrinksHandler(newCocktail)
	newCocktail = cocktails.Bartender.Start(newCocktail)

}

// Output:
// What would you like to drink?
// There are 635 drinks in the Database!
// Provide at least 3 letters in order to narrow your search or
// if you know the full name of the drink please provide it.
// gin
// -- Gin Fizz
// -- Gin Sour
// -- Pink Gin
// -- Gin Daisy
// -- Gin Sling
// -- Gin Smash
// -- Gin Toddy
// -- Gin Tonic
// -- Gin Lemon
// -- Gin Cooler
// -- Gin Squirt
// -- Gin Rickey
// -- Gin Swizzle
// -- Gin and Soda
// -- Gin And Tonic
// -- Royal Gin Fizz
// -- Ramos Gin Fizz
// -- Gin Basil Smash
// -- Sloe Gin Cocktail
// -- Pineapple Gingerale Smoothie

//                                         <- Now you can choose from the listed drinks! ->

// gin daisy
// 2 oz  Gin
// 1 oz  Lemon juice
// 1/2 tsp superfine  Sugar
// 1/2 tsp  Grenadine
// 1  Maraschino cherry
// 1  Orange
// In a shaker half-filled with ice cubes
//  combine the gin
//  lemon juice
//  sugar
//  and grenadine. Shake well. Pour into an old-fashioned glass and garnish with the cherry and the orange slice.

//                                         <- Now you can choose from the listed drinks! ->

// marg
// -- Margarita
// -- Blue Margarita
// -- Tommy's Margarita
// -- Whitecap Margarita
// -- Strawberry Margarita
// -- Smashed Watermelon Margarita

//                                         <- Now you can choose from the listed drinks! ->

// tommy's margarita
// 4.5 cl Tequila
// 1.5 cl Lime Juice
// 2 spoons Agave syrup
// Shake and strain into a chilled cocktail glass.

//                                         <- Now you can choose from the listed drinks! ->

// nothing
// exit status 1
