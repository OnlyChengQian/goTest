package mongo

type Test struct {
}

type TestTable struct {
	Id                 string `bson:"_id"`
	SiteCode           string `bson:"site_code"`
	SaleAccount        string `bson:"sale_account"`
	FlushInterShipping int    `bson:"flush_inter_shipping"`
	SaleAccountId      int    `bson:"sale_account_id"`
	ShippingPolicyId   int    `bson:"shipping_policy_id"`
	OptimizationStatus int    `bson:"optimization_status"`
}

func (t *Test) Collection() string {
	return "test"
}

func (t *Test) DataBase() string {
	return "db_ebay_advt"
}

func (t *Test) GetTable() interface{} {
	return TestTable{}
}
