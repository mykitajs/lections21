package main

import "strings"

import "fmt"

type row []string

type figure []row

func (f figure) String() string { 
	var res strings.Builder
	for i := range f {
		var subStr strings.Builder
  		for j := range f[i] {
			subStr.WriteString(f[i][j])
  		}
		subStr .WriteString("\n")

		res.WriteString(subStr.String())
 	}

	return res.String()
}

type Modifier func(intArgs map[string]int, strArgs map[string]string)

type figureFiller func(fig figure, intFinalArgs map[string]int, strFinalArgs map[string]string) figure 

const (
	defaultSize = 15
	defaultColor = 34
	defaultChar = "!"
)

const (
	sizeModifierArg = "size"
	colorMogifierArg = "color"
	charModifierArg = "char"
)

func charModifier(char string, color int) Modifier {
	return func(intArgs map[string]int, strArgs map[string]string) {
		intArgs[colorMogifierArg] = color
		strArgs[charModifierArg] = fmt.Sprintf("\033[%dm%s\033[0m", color, char)
	}
}

func sizeModifier(size int) Modifier {
	return func(intArgs map[string]int, strArgs map[string]string) {
		intArgs[sizeModifierArg] = size
	}
}

func colorModifier(color int) Modifier {
	return func(intArgs map[string]int, strArgs map[string]string) {
		intArgs[colorMogifierArg] = color
	}
}

func makeBlankFigure(size int) figure {
	blankFigure := make(figure, size)
	for i := range blankFigure {
		blankFigure[i] = make([]string, size)
		for j := range blankFigure[i] {
			blankFigure[i][j] = " "
		}
	}

	return blankFigure
}

func sandglassFigureFiller(fig figure, intFinalArgs map[string]int, strFinalArgs map[string]string) figure {
	size := intFinalArgs[sizeModifierArg]
	char := strFinalArgs[charModifierArg]

	for i := range fig { 
		fig[0][i] = char
		fig[size-1][i] = char
		fig[i][i] = char
		fig[i][size-1-i] = char
	}
 
	return fig
}

func setDefaultValues() (map[string]int, map[string]string) {
	intArgs := map[string]int {
		sizeModifierArg: defaultSize,
		colorMogifierArg: defaultColor,
	}
	strArgs := map[string]string {
  		charModifierArg: defaultChar,
 	}

	return  intArgs, strArgs
}

func constructfigure(figFiller figureFiller, mods ...Modifier) figure {
	intArgs, strArgs := setDefaultValues()

 	for _, mod := range mods {
  		mod(intArgs, strArgs)
 	}

	fig := makeBlankFigure(intArgs[sizeModifierArg])
	figFiller(fig, intArgs, strArgs)

	return fig
}

func main() {
	fmt.Println(constructfigure(sandglassFigureFiller, sizeModifier(15), charModifier("!", 34)))
}