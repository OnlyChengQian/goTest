package mongo

type Site struct {
}

type SiteTable struct {
	SiteId     int     `bson:"site_id"`
	SiteCode   string  `bson:"site_code"`
	SiteName   string  `bson:"site_name"`
	CurrencyId string  `bson:"currency_id"`
	MinPayPal  float64 `bson:"min_cost_pay_pal_email"`
}

func (s *Site) Collection() string {
	return "sfc_ebay_site"
}

func (s *Site) DataBase() string {
	return "db_ebay_advt"
}

func (s *Site) GetTable() interface{} {
	return SiteTable{}
}
