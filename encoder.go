package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)
var DATAFILE = "/tmp/dataFile.gob"

func save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}
	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	defer loadFrom.Close()
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&DATA)
	return nil
}

func ADD(n myElement) bool {
	if LOOKUP(n.Id) == nil {
		DATA[n.Id] = n
		return true
	}
	return false
}

func LOOKUP(k string) *myElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	} else {
		return nil
	}
}

func PRINT() {
	for _, d := range DATA {
		fmt.Printf("%s\n", d)
	}
}

func _init() {
	for i := 0; i < 50; i++ {
		e := myElement{Id: strconv.Itoa(i), Name: "Name " + strconv.Itoa(i), Surname: "Surname " + strconv.Itoa(i)}
		ADD(e)
	}
	err := save()
	if err != nil {
		fmt.Println(err)
	}
}

func goEncoder() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}
	PRINT()
}
