package main

import (
	"fmt"
	"math"
	"strings"
)

func createAdidasAsciiLogo(width int) string {

	//Rules to calculate values
	lowStripeHeight := math.Round(math.Sqrt(float64(width)))
	mediumStripeHeight := lowStripeHeight * 2
	highestStripeHeight := mediumStripeHeight + lowStripeHeight

	//blank Space between stripes is the same size of low stripe height
	Stripes := GetStripes(int(lowStripeHeight), int(mediumStripeHeight), int(highestStripeHeight), width, int(lowStripeHeight))

	return fmt.Sprintf("%s", Stripes)
}

//Support function
func GetStripes(lowStripeH int, medStripeH int, highStripeH int, stripesWidth int, spaceBetweenStripes int) string {

	var (
		lowStripePattern  = ""
		medStripePattern  = ""
		highStripePattern = ""
		fullLine          = ""
		blankCellLineLow  = 0

		//used to adjust the size of small sprite according medium
		factorDecreaseLow = (stripesWidth - 1) - lowStripeH

		//used to adjust the size of medium sprite according high
		factorDecreaseMed = stripesWidth - lowStripeH
	)

	blankSpace := RetrieveBlankSpace(spaceBetweenStripes, " ")

	for line := highStripeH; line > 0; line-- {

		if line <= lowStripeH && line <= medStripeH && line <= highStripeH {
			lowStripePattern = RetrieveBlankSpace(stripesWidth, "@")
			//used to decrease the lines on firt stripe
			blankCellLineLow = lowStripeH - line

		}

		if line > lowStripeH && line <= medStripeH && line <= highStripeH {

			medStripePattern = RetrieveBlankSpace(stripesWidth, "@")

			//Adjust the content of small sprite
			factorDecreaseLow += 1
			lowStripePattern = RetrieveBlankSpace(factorDecreaseLow, " ")

		}

		if line > lowStripeH && line > medStripeH && line <= highStripeH {
			highStripePattern = strings.Repeat("@", stripesWidth)

			//Adjust the content of medium sprite and repeat for small
			factorDecreaseMed += 1
			medStripePattern = RetrieveBlankSpace(factorDecreaseMed, " ")
			lowStripePattern = RetrieveBlankSpace(factorDecreaseLow, " ")
		}

		fullLine += fmt.Sprintf("%s%s%s%s%s%s\n", RetrieveBlankSpace(blankCellLineLow, " "), lowStripePattern, blankSpace, medStripePattern, blankSpace, highStripePattern)
	}

	return fullLine

}

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
