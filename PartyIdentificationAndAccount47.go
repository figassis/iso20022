package iso20022

// Party and account details.
type PartyIdentificationAndAccount47 struct {

	// Identification of the party.
	Identification *PartyIdentification45Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

}


func (p *PartyIdentificationAndAccount47) AddIdentification() *PartyIdentification45Choice {
	p.Identification = new(PartyIdentification45Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount47) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount47) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount47) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount47) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

