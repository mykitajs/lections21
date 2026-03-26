package main

import "fmt"

type figureFiller func(intsFig []int, runesFig []rune, strsFig []string) [][]rune  

type Modifier func(ints []int, runes []rune, strings []string)

func charModifier(char rune) Modifier {
	return func(ints []int, runes []rune, strings []string) {
		runes[0] = char
	}
}

func sizeModifier(size int) Modifier {
	return func(ints []int, runes []rune, strings []string) {
		ints[0] = size
	}
}

func colorModifier(color int) Modifier {
	return func(ints []int, runes []rune, strings []string) {
		ints[1] = color
		//Чёрный	30
		//Красный	31
		//Зелёный	32
		//Жёлтый	33
		//Синий		34
		//Фиол.		35
		//Голуб.	36
		//Белый		37
	}
}

func blankFigureFiller(numsFig []int, runesFig []rune, strsFig []string) [][]rune {
	size := numsFig[0]

	image := make([][]rune, size)
	for i := range image {
		image[i] = make([]rune, size)
		for j := range image[i] {
			image[i][j] = ' '
		}
	}

	return image
}

func sandglassFigureFiller(numsFig []int, runesFig []rune, strsFig []string) [][]rune {
	size := numsFig[0]
	char := runesFig[0]

	image := blankFigureFiller(numsFig, runesFig, strsFig)

	for i := 0; i < size; i++ {
		image[0][i] = char
		image[size-1][i] = char
		image[i][i] = char
		image[i][size-1-i] = char
	}

	return image
}

func printlnFigure(figure figureFiller, mods ...Modifier) {
	ints := make([]int, 2)
	runes := make([]rune, 1)
	strings := make([]string, 1)
	ints[0] = 15
	ints[1] = 34
	runes[0] = '!'
	strings[0] = "blue"
	color := ints[1]
	for _, mod := range mods {
		mod(ints, runes, strings)
	}
	image := figure(ints, runes, strings)
	for i := range image {
		for j := range image[i] {
			fmt.Printf("\033[%dm%c\033[0m", color, image[i][j])
		}
		fmt.Println()
	}
}

//fmt.Println("\033[31mКрасный текст\033[0m")

func main() {
	printlnFigure(sandglassFigureFiller, sizeModifier(15), charModifier('!'), colorModifier(34))
}