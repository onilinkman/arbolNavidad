package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	clearWindow()

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			drawColor = colorRed
		} else {
			drawColor = colorYellow
		}
		makeTree(12)

		time.Sleep(1 * time.Second)
		clearWindow()

	}
}

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorWhite = "\033[37m"

var colorActual string
var drawColor string

func makeTree(l int) {
	c := l*2 + 1
	z := c / 2

	for i := 0; i < l; i++ {
		cont := 0
		sw := true
		for j := 0; j < c; j++ {
			if z-i <= j && z+i >= j {
				if sw {

					fmt.Print(string(colorActual), modelTree(i, cont))
					sw = false
					cont++

				} else {
					fmt.Print(" ")
					sw = true
				}

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

}

func modelTree(i, cont int) string {
	num := binomio(i, cont)
	if num == 1 {
		colorActual = colorWhite
		return "*"
	} else if num%2 == 0 {
		colorActual = colorGreen
		return "#"
	} else {
		colorActual = drawColor
		return "@"
	}
}

func binomio(n, k int) int {
	s := n - k
	resu := 0
	if k > s {
		resu = desfract(n, k) / factorial(s)
	} else {
		resu = desfract(n, s) / factorial(k)
	}
	return resu

}

func desfract(n, u int) int {
	res := 1
	for n > u {
		res *= n
		n--
	}
	return res
}

func factorial(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}

func clearWindow() {
	operativeSistem := runtime.GOOS
	switch operativeSistem {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		break
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		break
	}

}
