package ProxyLogic

type Admin interface {
	BlockLicense(license *License)
	CheckLicenseStatus(license *License)
}
