package main

//A Passwords represents the standard password payload for key rotation on a car.
type Passwords struct {
	Name          string `json:"name,omitempty"`          //The name of the account
	Password      string `json:"password,omitempty"`      //The password for the named account
	Name2         string `json:"name2,omitempty"`         //The second named account
	Password2     string `json:"password2,omitempty"`     //The password of the second named account
	EffectiveDate string `json:"effectiveDate,omitempty"` //The date and time of when this change took effect
}
