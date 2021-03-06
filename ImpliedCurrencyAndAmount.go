package iso20022

// Number of monetary units specified in a currency where the unit of currency is implied by the context and compliant with ISO 4217. The decimal separator is a dot.
	// Note: a zero amount is considered a positive amount.
type ImpliedCurrencyAndAmount struct {
	Value string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewImpliedCurrencyAndAmount (value, currency string) *ImpliedCurrencyAndAmount {
	return &ImpliedCurrencyAndAmount{Value: value, Currency: currency}
}

