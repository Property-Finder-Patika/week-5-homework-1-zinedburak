package ProxyLogic

import "fmt"

// License object to create simple licenses and check its limit ext...
type License struct {
	Name      string
	UserCount int
	Blocked   bool
	Limit     int
}

func (l *License) AddUser() {
	if l.Limit > l.UserCount {
		l.UserCount++
	} else {
		fmt.Println(" There are no limit in this license please but new limit")
	}
}

func (l *License) BuyLicenseLimit(limit int) {
	l.Limit = limit
}
