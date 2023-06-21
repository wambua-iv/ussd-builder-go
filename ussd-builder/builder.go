package ussdbuilder

import (
	//"fmt"
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

func (ussd *UssdMenu) CON(text string) string {
	ussd.text = "CON " + text
	return ussd.text
}

func (ussd *UssdMenu) END(text string) string {
	ussd.result = text
	return ussd.result
}

func (ussd *UssdMenu) BuildState(text string, option func() string) {}

func (ussd *UssdMenu) StartState(next mapping, a ...value) string {
	ussd.states[1] = "Start State"
	for key, value := range next {
		ussd.states[key] = value
	}
	return ussd.states[1]
}

func (ussd *UssdMenu) GoToState(state int) string {
	return ussd.states[state]
}

func (ussd *UssdMenu) GetRoute(route string) []string {
	var routeParts = make([]string, len(route), len(route)+2)

	chars := "*"
	for key, value := range strings.ReplaceAll(route, chars, "") {
		routeParts[key] += string(value)
	}
	return routeParts
}

// func main() {
// 	var sess UssdMenu

// 	var sss = sess
// 	mapping1 := map[int]string{
// 		1: "here",
// 	}
// 	fmt.Print(sss.StartState(mapping1, sss.CON("heey")))
// 	fmt.Print(sss.CON("heey"))
// }
