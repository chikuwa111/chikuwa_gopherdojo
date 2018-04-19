// Made by take0shun1111@gmail.com

// 追加機能
// - 複数ファイルを取れる
// - 引数がない場合は標準出力を使う

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	statusCode = 0
	count      = flag.Int("n", 10, "The number of displayed lines")
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: myhead [option] [file ...]")
	fmt.Fprintln(os.Stderr, "*) If no files are specified, this filter uses the standard input.")
	fmt.Fprintln(os.Stderr, "option:")
	flag.PrintDefaults()
}

func handleErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	statusCode = 1
}

func head(input *os.File) {
	scanner := bufio.NewScanner(input)
	for i := 1; i <= *count; i++ {
		if !scanner.Scan() {
			break
		}
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		handleErr(err)
	}
}

func headFile(path string, printPath bool, index int) {
	file, err := os.Open(path)
	if err != nil {
		handleErr(err)
		return
	}
	defer file.Close()

	if printPath {
		if index > 0 {
			fmt.Println("")
		}
		fmt.Printf("==> %s <==\n", path)
	}
	head(file)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	switch len(args) {
	case 0:
		head(os.Stdin)
	case 1:
		headFile(args[0], false, 0)
	default: // len(args) > 1
		for index, path := range args {
			headFile(path, true, index)
		}
	}

	os.Exit(statusCode)
}
