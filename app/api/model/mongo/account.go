package mongo

type Account struct {
}

type AccountTable struct {
	Id               string  `bson:"_id"`
	Account          string  `bson:"account"`
	SaleAccount      string  `bson:"sale_account"`
	SaleAccountId    int     `bson:"sale_account_id"`
	SiteCode         string  `bson:"add_prefix_site_code"`
	SiteID           int     `bson:"site_id"`
	ProfitRate       float64 `bson:"gross_profit_rate"`
	ShopLevel        string  `bson:"shop_level"`
	M3               float64 `bson:"m3"`
	MinPrice         float64 `bson:"min_price"`
	MaxPrice         float64 `bson:"max_price"`
	MaxStock         int     `bson:"max_stock"`
	MinM3            float64 `bson:"min_m3"`
	IsIndiaUsAccount int     `bson:"is_india_us_account"`
}

func (a *Account) Collection() string {
	return "sfc_ebay_account"
}

func (a *Account) DataBase() string {
	return "db_ebay_advt"
}

func (a *Account) GetTable() interface{} {
	return AccountTable{}
}
