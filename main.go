package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// printSmth()
	pl("-----------")
	typeCheck()
	pl("-----------")
	typeConvert()
	pl("-----------")
	conditions()
	pl("-----------")
	stringsOps()
	pl("-----------")
	runesOps()
	pl("-----------")
	timeOps()
	pl("-----------")
	mathOps()
	pl("-----------")
	formatedPrint()
	pl("-----------")
	forLoops()
	pl("-----------")
	// whileLoops()
	pl("-----------")
	rangeOps()
}

// Alias (shortcut)
var pl = fmt.Println

func printSmth() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		pl("Hello", name)
	}
}

//Variables

//var name type

/* var vName string = "Hi"
var v1,v2 = 1.2, 2.4
var v3 = "hello"

v4 := 2.4
v5 := false */

//Data types

/* int, float64, bool, string, rune */

func typeCheck() {
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(2.45))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
}

//Casting

func typeConvert() {
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "5000000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 3333333
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}

//Conditional Operators: > < >= <= == !=
//Logical Orerators: && || !

func conditions() {
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not Important Birthday")
	}

	pl("!true =", !true)
}

//Strings

func stringsOps() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length:", len(sV2))
	pl("Contains Another:", strings.Contains(sV2, "Another"))
	pl("o index:", strings.Index(sV2, "o"))
	pl("Replace:", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)

	pl("Split:", strings.Split("a-b-c-d", "-"))
	pl("lower:", strings.ToLower(sV2))
	pl("Upper:", strings.ToUpper(sV2))
	pl("Prefix:", strings.HasPrefix("tacocat", "taco"))
	pl("Prefix:", strings.HasSuffix("tacocat", "cat"))
}

//Runes

func runesOps() {
	rStr := "abcdefg"
	pl("Rune Count:", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d: %#U: %c\n", i, runeVal, runeVal)
	}
}

//Time

func timeOps() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}

//Math

func mathOps() {
	pl("5 + 4 =", 5+4)
	pl("5 - 4 =", 5-4)
	pl("5 * 4 =", 5*4)
	pl("5 / 4 =", 5/4)
	pl("5 % 4 =", 5%4)

	mInt := 1
	mInt = mInt + 1 //or mInt += 1 or mInt++

	pl("Float Precison =", 0.11111111111111111111111+0.1111111111111111111111)

	seedSecs := time.Now().Unix()
	//depricated method
	// rand.Seed(seedSecs)
	//should use rand.NewSource(seedSecs) instead
	rand.NewSource(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)

	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	//Convert 90 degrees to radius
	r90 := 90 * math.Pi / 180
	//backwards
	d90 := r90 * (180 / math.Pi)
	pl(r90, d90)

	pl("Sin(90) =", math.Sin(r90))

	// There are also functions for: Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos, Htpot
}

//Formated Print

func formatedPrint() {
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}

//For loop

func forLoops() {
	for x := 1; x <= 5; x++ {
		pl(x)
	}

	for x := 5; x >= 1; x-- {
		pl(x)
	}
}

// While loop
func whileLoops() {
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	seedSecs := time.Now().Unix()
	rand.NewSource(seedSecs)
	randNum := rand.Intn(50) + 1

	//infinite loop
	for true {
		fmt.Print("Guess a number between 0 and 50:")
		pl("Random Number is:", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guessed It")
			break //or "continue" to skip current loop iteration
		}

	}
}

//Range (cycle through an array)

func rangeOps() {
	aNums := []int{1, 2, 3}

	for _, num := range aNums {
		pl(num)
	}
}

// Arrays
