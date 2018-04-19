package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"imgconverter"
)

var (
	statusCode = 0
	targetExt  string
	outputExt  string
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: imgconvert [options] [path ...]")
	fmt.Fprintln(os.Stderr, "Convert image files. The path of converted files are like \"[path].[ext]\".")
	fmt.Fprintln(os.Stderr, "If directory path is passed, files below that directory are converted.")
	fmt.Fprintln(os.Stderr, "options:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "-i and -o must be different.")
}

func handleErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	statusCode = 1
}

func convert(path string) error {
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	convertImg := imgconverter.Image{img}
	dest := path + "." + outputExt
	err = convertImg.Convert(dest)
	return err
}

func init() {
	flag.StringVar(&targetExt, "i", "jpg", "The format of target files (jpg or jpeg or png)")
	flag.StringVar(&outputExt, "o", "png", "The format of converted files (jpg or jpeg or png)")
	flag.Usage = usage
	flag.Parse()
}

func main() {
	isValid := (targetExt == "jpg" || targetExt == "jpeg" || targetExt == "png") &&
		(outputExt == "jpg" || outputExt == "jpeg" || outputExt == "png") &&
		(targetExt != outputExt)
	if !isValid {
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()
	for _, argPath := range args {
		err := filepath.Walk(argPath, func(path string, info os.FileInfo, err error) error {
			if strings.ToLower(filepath.Ext(path)) != "."+targetExt {
				return nil
			}
			convertErr := convert(path)
			return convertErr
		})
		if err != nil {
			handleErr(err)
		}
	}
	os.Exit(statusCode)
}
