package mongo

type Advt struct {
}

type AdvtTable struct {
	Id                   string   `bson:"_id"`
	SiteCode             string   `bson:"site_code"`
	SaleAccount          string   `bson:"sale_account"`
	FlushInterShipping   int      `bson:"flush_inter_shipping"`
	SaleAccountId        int      `bson:"sale_account_id"`
	AdvtStatus           int      `bson:"advt_status"`
	ProductId            int      `bson:"product_id"`
	PriceType            int      `bson:"price_type"`
	SalePrice            float64  `bson:"sale_price"`
	ProductBasePrice     float64  `bson:"product_base_price"`
	CostPrice            float64  `bson:"cost_price"`
	Profit               float64  `bson:"profit"`
	ProfitRate           float64  `bson:"profit_rate"`
	TotalWeight          float64  `bson:"total_weight"`
	IsSmallAmount        int      `bson:"is_small_amount"`
	AdvtShippingFee      float64  `bson:"advt_shipping_fee"`
	Stock                int      `bson:"stock"`
	EbayCategoryId       int      `bson:"ebay_category_id"`
	SecondEbayCategoryId int      `bson:"second_ebay_category_id"`
	ShippingPolicyId     int      `bson:"shipping_policy_id"`
	OptimizationStatus   int      `bson:"optimization_status"`
	M3Data               []string `bson:"m_3_data"`
	ComboProductKey      []string `bson:"combo_product_key"`
}

func (a *Advt) Collection() string {
	return "sfc_ebay_advt"
}

func (a *Advt) DataBase() string {
	return "db_ebay_advt"
}

func (a *Advt) GetTable() interface{} {
	return AdvtTable{}
}
