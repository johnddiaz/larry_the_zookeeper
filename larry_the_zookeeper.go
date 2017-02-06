package main

import (
	"time"
	"fmt"
	"os"
	"io"
	"bufio"
)

func load_file(filename string) (s string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(file)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 { break }
		s += string(buf[:n])
	}

	return s
}

func print_delayed_text(text string, delay time.Duration) {
	tlen := len(text)
	for i := 0; i < tlen; i++ {
		fmt.Print(string(text[i]))
		time.Sleep(delay * time.Millisecond)
	}
}

func main() {
	if l := len(os.Args); l == 1 {
		fmt.Println("You must provide a filename my friend.")
		return
	} else if l > 2 {
		fmt.Println("Sorry, I only work with one file.")
		return
	}

	dialog := load_file(os.Args[1])
	print_delayed_text(dialog, 1)
}
