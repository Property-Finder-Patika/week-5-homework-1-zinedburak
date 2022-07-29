package ProxyLogic

// Admin interface to trick users to thinking they are communicating with the admin
type Admin interface {
	BlockLicense(license *License)
	CheckLicenseStatus(license *License)
}
