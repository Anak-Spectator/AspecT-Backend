package shared

import "fmt"

func BeautyCli(mode string) {
	border := ""
	borderTop := " ┌"
	borderBottom := " └"
	lenght := 50
	for i := 0; i <= lenght; {
		border = border + "─"
		i = i + 1
	}
	borderTop = borderTop + border + "┐"
	borderBottom = borderBottom + border + "┘"

	space := lenght - 0
	br := "│"
	textSpace := " " + br
	for i := 0; i <= space; {
		textSpace = textSpace + " "
		i = i + 1
	}
	textSpace = textSpace + br

	mainText := fmt.Sprintf("Apps Run ON \"%s\" ENV", mode)

	semiSpace := ((lenght - len(mainText)) - 2) / 2
	firstSpace := ""
	for i := 0; i < semiSpace; {
		firstSpace = firstSpace + " "
		i = i + 1
	}
	secSpace := ""
	for i := len(firstSpace) + len(mainText); i < lenght; {
		secSpace = secSpace + " "
		i++
	}
	blank := " " + br
	mainText = blank + firstSpace + mainText + secSpace + blank + "\n"

	fmt.Println(borderTop)
	fmt.Println(textSpace)
	fmt.Printf(mainText)
	fmt.Println(textSpace)
	fmt.Println(borderBottom)
}
