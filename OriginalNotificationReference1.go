package iso20022

// Provides details on the account notification.
type OriginalNotificationReference1 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount16 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	OriginalItem []*OriginalItem2 `xml:"OrgnlItm"`

}


func (o *OriginalNotificationReference1) AddAccount() *CashAccount16 {
	o.Account = new(CashAccount16)
	return o.Account
}

func (o *OriginalNotificationReference1) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalNotificationReference1) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalNotificationReference1) AddRelatedAccount() *CashAccount16 {
	o.RelatedAccount = new(CashAccount16)
	return o.RelatedAccount
}

func (o *OriginalNotificationReference1) SetTotalAmount(value, currency string) {
	o.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalNotificationReference1) SetExpectedValueDate(value string) {
	o.ExpectedValueDate = (*ISODate)(&value)
}

func (o *OriginalNotificationReference1) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalNotificationReference1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalNotificationReference1) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalNotificationReference1) AddOriginalItem() *OriginalItem2 {
	newValue := new (OriginalItem2)
	o.OriginalItem = append(o.OriginalItem, newValue)
	return newValue
}

