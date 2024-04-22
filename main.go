package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var s []string = []string{
		"test ",
		"halo ",
	}
	var c []byte

	highlightStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)

	for _, v := range s {
		for i := 0; i < len(v); i++ {
			c = append(c, v[i])
		}
	}

	fmt.Println(s)
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			for _, char := range c {

				if ev.Ch == rune(char) {
					print(highlightStyle.Render(string(ev.Ch)))
				}

				if ev.Ch == '~' {
					os.Exit(1)
				}
			}
		case termbox.EventError:
			panic(ev.Err)

		}
	}
}
