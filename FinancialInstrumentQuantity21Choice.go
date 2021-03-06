package iso20022

// Choice between formats for the quantity of security.
type FinancialInstrumentQuantity21Choice struct {

	// Quantity expressed as a number, for example, a number of shares.
	Unit *RestrictedFINDecimalNumber `xml:"Unit"`

	// Quantity expressed as an amount representing the face amount, that is, the principal, of a debt instrument.
	FaceAmount *RestrictedFINImpliedCurrencyAndAmount `xml:"FaceAmt"`

	// Quantity expressed as an amount representing the current amortised face amount of a bond, for example, a periodic reduction/increase of a bond's principal amount.
	AmortisedValue *RestrictedFINImpliedCurrencyAndAmount `xml:"AmtsdVal"`

	// Quantity expressed as a code.
	Code *Quantity4Code `xml:"Cd"`

}


func (f *FinancialInstrumentQuantity21Choice) SetUnit(value string) {
	f.Unit = (*RestrictedFINDecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity21Choice) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity21Choice) SetAmortisedValue(value, currency string) {
	f.AmortisedValue = NewRestrictedFINImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity21Choice) SetCode(value string) {
	f.Code = (*Quantity4Code)(&value)
}

