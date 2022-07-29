package ProxyLogic

import "fmt"

// RealAdmin which is equivalent of the RealPrime minister in lecture code
type RealAdmin struct {
	Name string
}

// BlockLicense Only real admin has the authority to Block a License
func (r *RealAdmin) BlockLicense(license *License) {
	if !license.Blocked {
		license.Blocked = true
	}
}

// Don't bother the Admin with your stuff

func (r RealAdmin) CheckLicenseStatus(license *License) {
	fmt.Printf("Taking Care of this %s license is not my job dont bother me", license.Name)
}
