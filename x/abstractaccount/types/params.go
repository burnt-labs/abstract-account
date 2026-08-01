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

func NewParamsWithRegistrationCodeIDs(
	allowAllCodeIDs bool,
	allowedCodeIDs []uint64,
	maxGasBefore, maxGasAfter, bootstrapCodeID, implementationCodeID uint64,
) (*Params, error) {
	params, err := NewParams(allowAllCodeIDs, allowedCodeIDs, maxGasBefore, maxGasAfter)
	if err != nil {
		return nil, err
	}

	params.BootstrapCodeID = bootstrapCodeID
	params.ImplementationCodeID = implementationCodeID
	params.RegistrationEnabled = true

	return params, params.Validate()
}

func DefaultParams() *Params {
	params, _ := NewParams(true, []uint64{}, DefaultMaxGas, DefaultMaxGas)

	return params
}

func (p *Params) Validate() error {
	if err := p.validateRegistration(); err != nil {
		return err
	}

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
	prevCodeID := uint64(0)
	for _, codeID := range p.AllowedCodeIDs {
		if prevCodeID >= codeID {
			return ErrMalformedAllowList
		}

		prevCodeID = codeID
	}

	return nil
}

func (p *Params) validateRegistration() error {
	if (p.BootstrapCodeID == 0) != (p.ImplementationCodeID == 0) {
		return ErrMismatchedCodeIDs
	}
	if p.RegistrationEnabled && !p.RegistrationConfigured() {
		return ErrRegistrationNotConfigured
	}
	if p.ImplementationCodeID != 0 && !p.IsAllowed(p.ImplementationCodeID) {
		return ErrNotAllowedCodeID.Wrap("implementation code ID")
	}

	return nil
}

func (p *Params) RegistrationConfigured() bool {
	return p.BootstrapCodeID != 0 && p.ImplementationCodeID != 0
}

// IsAllowed returns whether a code ID is allowed for a raw Wasm migration of
// an AbstractAccount.
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
