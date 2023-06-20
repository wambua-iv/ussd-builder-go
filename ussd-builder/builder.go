package main

import (
	"fmt"
	"strings"
)

type value interface{}

type option interface {
	run() string
}

type args struct {
	sessionID   string
	phoneNumber string
	serviceCode string
	text        string
}

type mapping map[int]string

type UssdMenu struct {
	session string
	args
	states mapping
	result string
}

func (ussd *UssdMenu) CON(text string) UssdMenu {
	ussd.result = "CON " + text
	return *ussd
}

func (ussd *UssdMenu) END(text string) string {
	ussd.result = text
	return ussd.result
}

func (ussd *UssdMenu) BuildState(text string, option func() string) {}

func (ussd *UssdMenu) StartState(next mapping, a ...value) string {

	for key, value := range(next) {
		ussd.states[key] = value
	}
	ussd.states[1] = "Start State"
	return ussd.states[1]
}

func (ussd *UssdMenu) GoToState(state int) string {
	return ussd.states[state]
}

func (ussd *UssdMenu) GetRoute() {
	var routeParts []string

	chars := ",!"
    for key, value := range strings.ReplaceAll(ussd.text, chars, ""){	
            routeParts[key] += string(value)
    }
}

func main() {
	var sess UssdMenu

	var sss = sess
	mapping1 := map[int]string{
		1: "here",
	}
	fmt.Print(sss.StartState(mapping1, sss.CON("heey")))
	fmt.Print(sss.CON("heey"))
}
