package types

import "cosmossdk.io/errors"

var (
	ErrMalformedAllowList        = errors.Register(ModuleName, 2, "code ID allow list must contain non-zero, unique, and sorted code IDs")
	ErrNonEmptyAllowList         = errors.Register(ModuleName, 3, "code ID allow list must be empty when AllowAllCodeIDs is true")
	ErrNotAllowedCodeID          = errors.Register(ModuleName, 4, "not an allowed wasm code ID")
	ErrNotBaseAccount            = errors.Register(ModuleName, 5, "account is not an authtypes.BaseAccount")
	ErrNotSingleSignature        = errors.Register(ModuleName, 6, "signature is not a txsigning.SingleSignatureData")
	ErrParsingParams             = errors.Register(ModuleName, 7, "failed to marshal or unmarshal module params")
	ErrZeroMaxGas                = errors.Register(ModuleName, 8, "max gas cannot be zero")
	ErrNoBlockTime               = errors.Register(ModuleName, 9, "block time can not be zero")
	ErrRegistrationDisabled      = errors.Register(ModuleName, 10, "abstract account registration is disabled")
	ErrMismatchedCodeIDs         = errors.Register(ModuleName, 11, "bootstrap and implementation code IDs must both be zero or both be non-zero")
	ErrImmutableBootstrap        = errors.Register(ModuleName, 12, "bootstrap code ID is immutable once configured")
	ErrCodeIDNotFound            = errors.Register(ModuleName, 13, "wasm code ID does not exist")
	ErrAccountAlreadyRegistered  = errors.Register(ModuleName, 14, "account namespace is already registered")
	ErrNotAbstractAccount        = errors.Register(ModuleName, 15, "account is not an AbstractAccount")
	ErrRegistrationNotConfigured = errors.Register(ModuleName, 16, "abstract account registration code IDs are not configured")
)
