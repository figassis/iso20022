package iso20022

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification78 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Type *GenericIdentification30 `xml:"Tp"`

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id,omitempty"`

}


func (g *GenericIdentification78) AddType() *GenericIdentification30 {
	g.Type = new(GenericIdentification30)
	return g.Type
}

func (g *GenericIdentification78) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

