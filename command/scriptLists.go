package command

var scriptList = map[string]ScriptProvider{
	"producer:computeShipping": &ComputeShippingScript{},
}
