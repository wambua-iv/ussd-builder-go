package ussdbuilder

import (
	//"fmt"
	"strings"
)

// define the methods to be implemented by the State

type StateImpl interface {
	run () string
	next () string
}

type State struct {
	Name string
}

var States map[string] State



type Args struct {
	SessionID   string
	PhoneNumber string
	ServiceCode string
	Text        string
	NetworkCode string
}





func (s *State) GetName() string {
	return s.Name
}

type UssdMenu struct {
	session string
	Args 		Args
	States     mapping
	result     string
	routeParts []string
}

func (ussd *UssdMenu) CON(text string) string {
	ussd.Text = "CON " + text
	return ussd.Text
}

func (ussd *UssdMenu) END(text string) string {
	ussd.result = "END" + text
	return ussd.result
}

//func (ussd *UssdMenu) BuildState(text string, next mapping, a ...value) {}

func (ussd *UssdMenu) BuildState(states mapping, a ...value) mapping {
	ussd.States = states

	return ussd.States
}



func (ussd *UssdMenu) GoToState(state int) string {
	return ussd.States[state]
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

func (ussd *UssdMenu) GetCurrentRoute(route string) string {
	chars := "*"
	routes := strings.ReplaceAll(route, chars, "")
	for key := range routes {
		if string(routes[key]) != ussd.routeParts[key] {
			return string(routes[key])
		}
	}
	return ussd.routeParts[1]
}

func (ussd *UssdMenu) GetValue() {}

func (UssdMenu) Run(run func(any ...interface{}) string) string { return run() }

// func main() {
// 	var sess UssdMenu

// 	var sss = sess
// 	mapping1 := map[int]string{
// 		1: "here",
// 	}
// 	fmt.Print(sss.StartState(mapping1, sss.CON("heey")))
// 	fmt.Print(sss.CON("heey"))
// }
