package types

const DefaultMaxGas = 2_000_000

func NewParams(allowAllCodeIDs bool, allowedCodeIDs []uint64, maxGasBefore, maxGasAfter uint64) (*Params, error) {
	params := &Params{
		AllowAllCodeIDs: allowAllCodeIDs,
		AllowedCodeIDs:  allowedCodeIDs,
		MaxGasBefore:    maxGasBefore,
		MaxGasAfter:     maxGasAfter,
	}

	return params, params.Validate()
}

func DefaultParams() *Params {
	params, _ := NewParams(true, []uint64{}, DefaultMaxGas, DefaultMaxGas)

	return params
}

func (p *Params) Validate() error {
	if p.MaxGasBefore <= 0 {
		return ErrZeroMaxGas
	}

	if p.MaxGasAfter <= 0 {
		return ErrZeroMaxGas
	}

	// if all code IDs are allowed, then the allowed list must be empty
	if p.AllowAllCodeIDs && len(p.AllowedCodeIDs) != 0 {
		return ErrNonEmptyAllowList
	}

	// allowed list must contain non-zero, unique, and sorted code IDs
	prev := uint64(0)
	for _, codeID := range p.AllowedCodeIDs {
		if codeID == 0 {
			return ErrMalformedAllowList
		}

		if prev >= codeID {
			return ErrMalformedAllowList
		}

		prev = codeID
	}

	return nil
}

// IsAllowed returns whether a code ID is allowed to be used to register
// AbstractAccounts.
func (p *Params) IsAllowed(codeID uint64) bool {
	if p.AllowAllCodeIDs {
		return true
	}

	for _, allowedCodeID := range p.AllowedCodeIDs {
		if codeID == allowedCodeID {
			return true
		}
	}

	return false
}
