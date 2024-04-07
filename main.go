package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

func main() {
	nameIn := flag.String("in", "source", "Папка для исходных изображений")
	nameOut := flag.String("out", "converted", "Папка для конвертированных изображений")
	flag.Parse()

	err := compressWebP(*nameIn, *nameOut, "png")
	if err != nil {
		panic(err)
	}
	fmt.Println("Complete!")
}

func compressWebP(srcDir, dstDir, srcExt string) error {
	files, err := filepath.Glob(filepath.Join("./"+srcDir, "*."+srcExt))
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println("Start converted: " + file)
		imgFile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer imgFile.Close()

		img, err := png.Decode(imgFile)
		if err != nil {
			return err
		}

		err = checkAndCreateDir("./" + dstDir)
		if err != nil {
			return err
		}

		outFile, err := os.Create(filepath.Join("./"+dstDir, strings.TrimSuffix(filepath.Base(file), "."+srcExt)+".webp"))
		if err != nil {
			return err
		}
		defer outFile.Close()

		err = webp.Encode(outFile, img, &webp.Options{Lossless: true})
		if err != nil {
			return err
		}
	}

	return nil
}

func checkAndCreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
