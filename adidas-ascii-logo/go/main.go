package main

import (
	"fmt"
	"math"
	"strings"
)

func createAdidasAsciiLogo(width int) string {

	//Just to accomplish some tests
	if width < 2 {
		return "Error, minimum width is 2"
	} else {

		//Rules to calculate values
		//blank Space between stripes is the same size of low stripe height
		lowStripeHeight := int(math.Round(math.Sqrt(float64(width))))
		mediumStripeHeight := int(lowStripeHeight * 2)
		highestStripeHeight := int(mediumStripeHeight + lowStripeHeight)
		return GetStripes(lowStripeHeight, mediumStripeHeight, highestStripeHeight, width, lowStripeHeight)

	}

}

//Support function to generate the stripes
func GetStripes(lowStripeH int, medStripeH int, highStripeH int, stripesWidth int, spaceBetweenStripes int) string {

	var (
		lowStripePattern  = ""
		medStripePattern  = ""
		highStripePattern = ""
		fullLine          = ""
		blankCellLineLow  = 0

		//used to adjust the size of small Stripe according medium
		factorDecreaseLow = (stripesWidth - 1) - lowStripeH

		//used to adjust the size of medium Stripe according high
		factorDecreaseMed = (stripesWidth) - lowStripeH
	)

	//Generates the blank space betwwen stripes
	blankSpace := RetrieveBlankSpace(spaceBetweenStripes, " ")

	//Working in a reverse loop to make easey the drawing
	for line := highStripeH; line > 0; line-- {

		//Small Stripe
		if line <= lowStripeH && line <= medStripeH && line <= highStripeH {
			lowStripePattern = RetrieveBlankSpace(stripesWidth, "@")
			//used to decrease the lines on firt stripe
			blankCellLineLow = lowStripeH - line
		}

		//Medium Stripe
		if line > lowStripeH && line <= medStripeH && line <= highStripeH {
			medStripePattern = RetrieveBlankSpace(stripesWidth, "@")

			//Adjust the content of small Stripe
			factorDecreaseLow += 1
			lowStripePattern = RetrieveBlankSpace(factorDecreaseLow, " ")
		}

		//High Stripe
		if line > lowStripeH && line > medStripeH && line <= highStripeH {
			highStripePattern = strings.Repeat("@", stripesWidth)

			//Adjust the content of medium Stripe and repeat for small
			factorDecreaseMed += 1
			medStripePattern = RetrieveBlankSpace(factorDecreaseMed, " ")
			lowStripePattern = RetrieveBlankSpace(factorDecreaseLow, " ")
		}

		fullLine += fmt.Sprintf("%s%s%s%s%s%s\n", RetrieveBlankSpace(blankCellLineLow, " "), lowStripePattern, blankSpace, medStripePattern, blankSpace, highStripePattern)
	}

	//Remove the last Char
	return strings.TrimRight(fullLine, "\n")
}

///Function that Prints the blank space of @@@ from a stripe
func RetrieveBlankSpace(numSpace int, symbol string) string {
	return strings.Repeat(symbol, numSpace)
}

func main() {

	widths := []int{2, 3, 5, 7, 9, 16, 21}

	for _, width := range widths {
		fmt.Println(fmt.Sprintf("\nadidas (width %d)", width))
		fmt.Println("\n-------------------------------------------------------------")
		fmt.Println(fmt.Sprintf("\n%s\n\n", createAdidasAsciiLogo(width)))
	}
}
