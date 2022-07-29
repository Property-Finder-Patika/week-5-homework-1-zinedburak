package ProxyLogic

import "fmt"

type RealAdmin struct {
	Name string
}

func (r *RealAdmin) BlockLicense(license *License) {
	if !license.Blocked {
		license.Blocked = true
	}
}

func (r RealAdmin) CheckLicenseStatus(license *License) {
	fmt.Printf("Taking Care of this %s license is not my job dont bother me", license.Name)
}
