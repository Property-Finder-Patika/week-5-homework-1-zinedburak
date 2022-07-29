package ProxyLogic

type AdminQuarters struct {
	SiteAdmin *SiteAdmin
}

func NewAdminQuarters(realAdmin *RealAdmin) *AdminQuarters {
	adminQuarters := new(AdminQuarters)
	adminQuarters.SiteAdmin = new(SiteAdmin)
	adminQuarters.SiteAdmin.New(realAdmin)
	return adminQuarters
}

func (aq *AdminQuarters) GiveSiteAdmin() Admin {
	return aq.SiteAdmin
}
