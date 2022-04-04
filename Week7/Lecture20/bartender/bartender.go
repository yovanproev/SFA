package cocktails

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Bartender interface {
	Start() CocktailBartender
}

type CocktailBartender struct {
	UserInput   string
	DrinksFound []string
	Drinks      []struct {
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

func (c CocktailBartender) Start() CocktailBartender {
	result := makeQueryToURL(c)

	var drinks []string

	// Find ingredients for requested drink
	for _, rec := range result.Drinks {

		onlyOneDrinkMatch := RunesToStrings(c.UserInput, rec.StrDrink)

		// list drinks DB if more than 1 match
		if len(result.Drinks) > 1 {
			fmt.Println("--", rec.StrDrink)
		}

		drinks = append(drinks, rec.StrDrink)

		// if the input matches only 1 drink
		if rec.StrDrink == onlyOneDrinkMatch {
			ingredientsInTheStruct := reflect.ValueOf(rec)

			listAllMeasuresAndIngredients(ingredientsInTheStruct)

			// list the instructions
			for _, line := range strings.Split(rec.StrInstructions, ",") {
				fmt.Println(line)
			}
		}
	}

	_, everyFirstLetterCapital := TurnInputToRunes(c.UserInput)

	if everyFirstLetterCapital == "Nothing" {
		os.Exit(1)
	} else if drinks == nil {
		fmt.Println("No matches in the DB.")
		os.Exit(1)
	}

	return CocktailBartender{
		DrinksFound: drinks,
	}
}

func makeQueryToURL(c CocktailBartender) CocktailBartender {
	modifyToURLAcceptableFormat := TurnInputToURLFormat(c.UserInput)

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

	return result
}

func HandleConsoleInput() string {
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	return query
}

func DrinksHandler(c CocktailBartender) {
	runes, _ := TurnInputToRunes(c.UserInput)

	input := HandleConsoleInput()
	cocktail := CocktailBartender{
		UserInput: input,
	}

	c = Bartender.Start(cocktail)

	for {
		fmt.Println(`
					<- Now you can choose from the listed drinks! ->
					`)

		input = HandleConsoleInput()

		if c.DrinksFound[0] != string(runes) {
			cocktail := CocktailBartender{
				UserInput: input,
			}
			c = Bartender.Start(cocktail)
		} else {
			os.Exit(1)
		}
	}

}

func listAllMeasuresAndIngredients(ingredientsInTheStruct reflect.Value) {
	var lastKey string
	typeOfS := ingredientsInTheStruct.Type()

	maps := make(map[string]string)
	// There are up to 15 items as ingredients and measures, but not all
	// drinks have 15 items. Find the ones populated and map it.
	for i := 0; i < ingredientsInTheStruct.NumField(); i++ {
		if ingredientsInTheStruct.Field(i).Interface() != nil {
			maps[typeOfS.Field(i).Name] = ingredientsInTheStruct.Field(i).Interface().(string)
			lastKey = typeOfS.Field(i).Name
		}
	}

	// find the number of the last populated item and take the number from the back
	re := regexp.MustCompile("[0-9]+")
	numberInLastKey := re.FindAllString(lastKey, -1)
	toNumber, err := strconv.Atoi(numberInLastKey[0])
	if err != nil {
		fmt.Println(err)
	}

	// take the last populated item number and list all igredients and measures up to that number
	for i := 0; i < toNumber; i++ {
		fmt.Println(maps[`StrMeasure`+strconv.Itoa(i+1)], maps[`StrIngredient`+strconv.Itoa(i+1)])
	}
}
