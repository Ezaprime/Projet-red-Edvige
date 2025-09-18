package piscine

import (
	"fmt"
	"strings"
	"time"
)

var bannerLines = []string{
    " ####   #   #  ####   #####  ####   #   #  #   #   ###   #   #",
	"#         #    #   #  #      #   #   # #   ##  #  #   #  #   #",
	"#       #   #  ####   ####   ####   #   #  # # #  #   #   # # ",
	"#       #   #  #   #  #      #  #   #   #  #  ##  #   #    #  ",
	" ####   #   #  ####   #####  #   #  #   #  #   #   ###     #  ",
	"                 C  Y  B  E  R  Y  N  O  V",
}

func clear()            { fmt.Print("\033[H\033[2J") }
func hideCursor()       { fmt.Print("\033[?25l") }
func showCursor()       { fmt.Print("\033[?25h") }
func rgb(r, g, b int) string { return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b) }
func reset()            { fmt.Print("\033[0m") }

func lerp(a, b int, t float64) int { return int(float64(a) + (float64(b)-float64(a))*t) }

func printGradientBanner() {
	left := "  "
	start := [3]int{255, 80, 200} 
	end := [3]int{0, 255, 255}  
	for i, line := range bannerLines {
		t := 0.0
		if len(bannerLines) > 1 {
			t = float64(i) / float64(len(bannerLines)-1)
		}
		r := lerp(start[0], end[0], t)
		g := lerp(start[1], end[1], t)
		b := lerp(start[2], end[2], t)
		fmt.Println(rgb(r, g, b) + left + line + "\033[0m")
		time.Sleep(12 * time.Millisecond)
	}
}

func neonFlicker(c1, c2 string, repeats int, hold time.Duration) {
	for i := 0; i < repeats; i++ {
		clear()
		fmt.Print("\033[2m") 
		for _, line := range bannerLines {
			fmt.Println(c1 + "  " + line + "\033[0m")
		}
		time.Sleep(90 * time.Millisecond)

		clear()
		fmt.Print("\033[1m") 
		for _, line := range bannerLines {
			fmt.Println(c2 + "  " + line + "\033[0m")
		}
		time.Sleep(hold)
	}
	fmt.Print("\033[22m") 
}

func spinner(label string, steps int, delay time.Duration) {
	frames := []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
	for i := 0; i < steps; i++ {
		fmt.Printf("%s %c\r", label, frames[i%len(frames)])
		time.Sleep(delay)
	}
	fmt.Print(strings.Repeat(" ", len(label)+2) + "\r")
}

func progress(label string, width int, delay time.Duration) {
	for i := 0; i <= width; i++ {
		t := 0.0
		if width > 0 {
			t = float64(i) / float64(width)
		}
		col := rgb(lerp(255, 0, t), lerp(80, 255, t), lerp(200, 255, t))
		filled := strings.Repeat("█", i)
		empty := strings.Repeat("░", width-i)
		pct := int(t*100 + 0.5)
		fmt.Printf("%s %s[%s%s]\033[0m %3d%%\r", label, col, filled, empty, pct)
		time.Sleep(delay)
	}
	fmt.Print("\n")
}

func AnimateNeonBanner() {
	clear()
	hideCursor()
	defer func() { reset(); showCursor() }()

	cyan := rgb(0, 255, 255)
	magenta := rgb(255, 80, 200)

	neonFlicker(cyan, magenta, 2, 110*time.Millisecond)

	clear()
	printGradientBanner()
	fmt.Println()

	spinner("Initialisation du noyau", 24, 35*time.Millisecond)
	progress("Chargement des modules", 26, 18*time.Millisecond)
	time.Sleep(250 * time.Millisecond)

	clear()
}
