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
	states     mapping
	result     string
	routeParts []string
}

func (ussd *UssdMenu) CON(text string) string {
	ussd.text = "CON " + text
	return ussd.text
}

func (ussd *UssdMenu) END(text string) string {
	ussd.result = "END" + text
	return ussd.result
}

//func (ussd *UssdMenu) BuildState(text string, next mapping, a ...value) {}

func (ussd *UssdMenu) BuildState(next mapping, a ...value) mapping {
	ussd.states = next

	return ussd.states
}

func (ussd *UssdMenu) GoToState(state int) string {
	return ussd.states[state]
}

func (ussd *UssdMenu) GetRoutes(route string) []string {
	chars := "*"
	arr := strings.ReplaceAll(route, chars, "")
	ussd.routeParts = make([]string, len(arr), len(arr)+2)
	for key, value := range strings.ReplaceAll(route, chars, "") {
		ussd.routeParts[key] += string(value)
	}
	return ussd.routeParts
}

func (ussd *UssdMenu) GetCurrentRoute(route string) string{
	chars := "*"
	routes:= strings.ReplaceAll(route, chars, "")
	for key := range routes {
		if string(routes[key]) != ussd.routeParts[key]{
			return string(routes[key])
		}
	}
	return ussd.routeParts[1]
}

func (ussd *UssdMenu) GetValue() {}

// func main() {
// 	var sess UssdMenu

// 	var sss = sess
// 	mapping1 := map[int]string{
// 		1: "here",
// 	}
// 	fmt.Print(sss.StartState(mapping1, sss.CON("heey")))
// 	fmt.Print(sss.CON("heey"))
// }
