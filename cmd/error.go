package cmd

import "errors"

// Error messages
var (
	ErrorMissingKeyAsArg    = errors.New("key not provided")
	ErrorMissingDecryptArgs = errors.New("expected key and RSA private key path")
	ErrorValuesConfirmation = errors.New("value has not been confirmed")
)
