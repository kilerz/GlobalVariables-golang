package model

type Activities struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

type ArrayActivities struct {
	Activities []Activities `json:"activities"`
}

var GlobalActivities []Activities
