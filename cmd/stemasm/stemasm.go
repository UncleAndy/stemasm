package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	mixFile := flag.String("mix", "", "Mix part file name")
	drumFile := flag.String("drum", "", "Drum part file name")
	bassFile := flag.String("bass", "", "Bass part file name")
	melodyFile := flag.String("melody", "", "Melody part file name")
	vocalFile := flag.String("vocal", "", "Vocal part file name")

	isError := false
	if mixFile == nil || strings.TrimSpace(*mixFile) == "" {
		log.Println("Mix file name should by present (-mix)")
		isError = true
	}
	if drumFile == nil || strings.TrimSpace(*drumFile) == "" {
		log.Println("Drum file name should by present (-drum)")
		isError = true
	}
	if bassFile == nil || strings.TrimSpace(*bassFile) == "" {
		log.Println("Bass file name should by present (-bass)")
		isError = true
	}
	if melodyFile == nil || strings.TrimSpace(*melodyFile) == "" {
		log.Println("Melody file name should by present (-melody)")
		isError = true
	}
	if vocalFile == nil || strings.TrimSpace(*vocalFile) == "" {
		log.Println("Vocal file name should by present (-vocal)")
		isError = true
	}
	if isError {
		fmt.Println("\nUsing: stemasm <options>")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
