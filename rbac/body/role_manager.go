package body

type Policy struct {
	Params []string `json:"params,omitempty" binding:"required"`
}

func (p *Policy) GetPolcy() []string {
	return p.Params
}
