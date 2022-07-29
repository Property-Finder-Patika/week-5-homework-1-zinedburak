package ProxyLogic

import "fmt"

type SiteAdmin struct {
	realAdmin *RealAdmin
}

func (s *SiteAdmin) New(realAdmin *RealAdmin) SiteAdmin {
	s.realAdmin = realAdmin
	return *s
}

// The communucation with the real admin
func (s SiteAdmin) BlockLicense(license *License) {
	s.realAdmin.BlockLicense(license)
}

// Site admin will communicate with real admin if the license is out of limit

func (s SiteAdmin) CheckLicenseStatus(license *License) {
	if license.Limit >= license.UserCount {
		fmt.Printf("License is okay thera are %d users %d limit", license.UserCount, license.Limit)
	} else {
		fmt.Printf("License has exceeded limit user ")
		s.BlockLicense(license)
	}
}
