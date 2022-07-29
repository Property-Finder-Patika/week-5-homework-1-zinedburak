package ProxyLogic

type User struct {
	admin         Admin
	adminQuarters *AdminQuarters
}

func NewUser(quarters *AdminQuarters) User {
	user := new(User)
	user.adminQuarters = quarters
	user.admin = user.adminQuarters.GiveSiteAdmin()
	return *user
}

func (u *User) CheckLicenseStatus(license *License) {
	u.admin.CheckLicenseStatus(license)
}

func (u *User) BuyLicense(license *License, limit int) {
	license.BuyLicenseLimit(limit)
}
