package mongo

type AdvtChild struct {
}

type AdvtChildTable struct {
	ID               int      `bson:"_id"`
	AdvtId           int      `bson:"advt_id"`
	SaleAccountId    string   `bson:"sale_account_id"`
	ProductId        int      `bson:"product_id"`
	CostPrice        string   `bson:"cost_price"`
	Profit           string   `bson:"profit"`
	ProfitRate       string   `bson:"profit_rate"`
	TotalWeight      string   `bson:"total_weight"`
	Stock            int      `bson:"stock"`
	IsSmallAmount    int      `bson:"is_small_amount"`
	ProductBasePrice string   `bson:"product_base_price"`
	M3Data           []string `bson:"m_3_data"`
}

func (a *AdvtChild) Collection() string {
	return "sfc_ebay_advt_child"
}

func (a *AdvtChild) DataBase() string {
	return "db_ebay_advt"
}

func (a *AdvtChild) GetTable() interface{} {
	return AdvtChildTable{}
}
