package main

import (
	"bufio"
	"mp"
	"oop/fmt"
	"oop/mlib"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, singal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len; i++ {

		}
	}
}
