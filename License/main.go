package main

import (
	"license/License/ProxyLogic"
)

func main() {
	admin := new(ProxyLogic.RealAdmin)
	adminQuarters := ProxyLogic.NewAdminQuarters(admin)
	user := ProxyLogic.NewUser(adminQuarters)
	license := &ProxyLogic.License{
		Name:      "Goland",
		UserCount: 10,
		Blocked:   false,
		Limit:     1000,
	}
	user.CheckLicenseStatus(license)
}
