package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Bartender interface {
	Start() CocktailBartender
}

type CocktailBartender struct {
	UserInput string
	FirstPick []string
	Drinks    []struct {
		StrDrink        string      `json:"strDrink"`
		StrInstructions string      `json:"strInstructions"`
		StrIngredient1  string      `json:"strIngredient1"`
		StrIngredient2  string      `json:"strIngredient2"`
		StrIngredient3  string      `json:"strIngredient3"`
		StrIngredient4  interface{} `json:"strIngredient4"`
		StrIngredient5  interface{} `json:"strIngredient5"`
		StrIngredient6  interface{} `json:"strIngredient6"`
		StrIngredient7  interface{} `json:"strIngredient7"`
		StrIngredient8  interface{} `json:"strIngredient8"`
		StrIngredient9  interface{} `json:"strIngredient9"`
		StrIngredient10 interface{} `json:"strIngredient10"`
		StrMeasure1     string      `json:"strMeasure1"`
		StrMeasure2     string      `json:"strMeasure2"`
		StrMeasure3     string      `json:"strMeasure3"`
		StrMeasure4     interface{} `json:"strMeasure4"`
		StrMeasure5     interface{} `json:"strMeasure5"`
		StrMeasure6     interface{} `json:"strMeasure6"`
		StrMeasure7     interface{} `json:"strMeasure7"`
		StrMeasure8     interface{} `json:"strMeasure8"`
		StrMeasure9     interface{} `json:"strMeasure9"`
		StrMeasure10    interface{} `json:"strMeasure10"`
	} `json:"drinks"`
}

func main() {
	fmt.Println("What would you like to drink?")
	fmt.Println("There are 635 drinks in the Database! ")
	fmt.Println("Provide at least 3 letters in order to narrow your search or ")
	fmt.Println("if you know the full name of the drink please provide it.")

	initialReader := bufio.NewReader(os.Stdin)
	initialSearch, err := initialReader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	modifyInputToAcceptedFormat := strings.Replace(initialSearch, "\n", "", -1)
	modifyInputToAcceptedFormat = strings.Replace(modifyInputToAcceptedFormat, "\r", "", -1)
	everyFirstLetterCapital := strings.Title(strings.ToLower(string(modifyInputToAcceptedFormat)))

	s := CocktailBartender{
		UserInput: modifyInputToAcceptedFormat,
	}

	newCocktail := Bartender.Start(s)

	var turnInputToRunes []rune

	if smallLetterAfterQuotationMarks(modifyInputToAcceptedFormat) != nil {
		modify := strings.Replace(everyFirstLetterCapital, "'S", string(smallLetterAfterQuotationMarks(modifyInputToAcceptedFormat)[0][0]), -1)
		turnInputToRunes = adjustTheRuneSlice(modify)
	} else {
		turnInputToRunes = adjustTheRuneSlice(everyFirstLetterCapital)
	}

	var secondDraw CocktailBartender

	if everyFirstLetterCapital != "Nothing" && len(newCocktail.FirstPick) > 0 {
		for len(secondDraw.FirstPick) == 0 || len(secondDraw.FirstPick) > 1 {
			if newCocktail.FirstPick[0] != string(turnInputToRunes[0:len(newCocktail.FirstPick[0])]) {

				fmt.Println(`
				<- Now you can choose from the listed drinks! ->
				`)

				reader := bufio.NewReader(os.Stdin)
				drink, err := reader.ReadString('\n')

				if err != nil {
					log.Println(err)
				}

				s := CocktailBartender{
					UserInput: drink,
				}

				secondDraw = Bartender.Start(s)

			} else {
				os.Exit(1)
			}
		}
	}
}

func smallLetterAfterQuotationMarks(str string) [][][]byte {
	re := regexp.MustCompile(`'(.?)`)
	smallLetterAfterQuotationMarks := re.FindAllSubmatch([]byte(str), -1)

	return smallLetterAfterQuotationMarks
}

func adjustTheRuneSlice(str string) []rune {
	turnInputToRunes := []rune(str)
	return turnInputToRunes
}

func (c CocktailBartender) Start() CocktailBartender {
	// drinks with Quotation mark
	modifyToURLAcceptableFormat := strings.Replace(c.UserInput, "'", "%27", -1)

	everyFirstLetterCapital := strings.Title(strings.ToLower(modifyToURLAcceptableFormat))
	modifyToURLAcceptableFormat = strings.Replace(everyFirstLetterCapital, " ", "%20", -1)
	modifyToURLAcceptableFormat = strings.Replace(modifyToURLAcceptableFormat, "\n", "", -1)
	modifyToURLAcceptableFormat = strings.Replace(modifyToURLAcceptableFormat, "\r", "", -1)

	resp, err := http.Get("https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + modifyToURLAcceptableFormat)
	if err != nil {
		fmt.Println("No response from request ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var result CocktailBartender
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	var new []string

	// Find ingredients for requested drink
	for _, rec := range result.Drinks {

		var runesToStrings string
		if smallLetterAfterQuotationMarks(c.UserInput) != nil {
			everyFirstLetterCapital := strings.Title(strings.ToLower(c.UserInput))
			modify := strings.Replace(everyFirstLetterCapital, "'S", string(smallLetterAfterQuotationMarks(c.UserInput)[0][0]), -1)
			adjustTheRuneSlice := adjustTheRuneSlice(modify)[0:len(rec.StrDrink)]
			runesToStrings = string(adjustTheRuneSlice)
		} else {
			everyFirstLetterCapital := strings.Title(strings.ToLower(c.UserInput))
			adjustTheRuneSlice := adjustTheRuneSlice(everyFirstLetterCapital)[0:len(rec.StrDrink)]
			runesToStrings = string(adjustTheRuneSlice)
		}

		if len(result.Drinks) > 1 {
			fmt.Println("--", rec.StrDrink)
		}
		new = append(new, rec.StrDrink)

		if rec.StrDrink == runesToStrings {
			if rec.StrMeasure1 != "" && rec.StrIngredient1 != "" {
				fmt.Println(rec.StrMeasure1, rec.StrIngredient1)
			}
			if rec.StrMeasure2 != "" && rec.StrIngredient2 != "" {
				fmt.Println(rec.StrMeasure2, rec.StrIngredient2)
			}
			if rec.StrMeasure3 != "" && rec.StrIngredient3 != "" {
				fmt.Println(rec.StrMeasure3, rec.StrIngredient3)
			}
			if rec.StrMeasure4 != nil && rec.StrIngredient4 != nil {
				fmt.Println(rec.StrMeasure4, rec.StrIngredient4)
			}
			if rec.StrMeasure5 != nil && rec.StrIngredient5 != nil {
				fmt.Println(rec.StrMeasure5, rec.StrIngredient5)
			}
			if rec.StrMeasure6 != nil && rec.StrIngredient6 != nil {
				fmt.Println(rec.StrMeasure6, rec.StrIngredient6)
			}
			if rec.StrMeasure7 != nil && rec.StrIngredient7 != nil {
				fmt.Println(rec.StrMeasure7, rec.StrIngredient7)
			}
			if rec.StrMeasure8 != nil && rec.StrIngredient8 != nil {
				fmt.Println(rec.StrMeasure8, rec.StrIngredient8)
			}
			if rec.StrMeasure9 != nil && rec.StrIngredient9 != nil {
				fmt.Println(rec.StrMeasure9, rec.StrIngredient9)
			}
			if rec.StrMeasure10 != nil && rec.StrIngredient10 != nil {
				fmt.Println(rec.StrMeasure10, rec.StrIngredient10)
			}

			for _, line := range strings.Split(rec.StrInstructions, ",") {
				fmt.Println(line)
			}
		}
	}

	return CocktailBartender{
		FirstPick: new,
	}
}
