package info

import "testing"

func TestValidateAssetDecimalsAccordingType_BEP2InvalidDecimals(t *testing.T) {
	if err := ValidateAssetDecimalsAccordingType("BEP2", 9); err == nil {
		t.Fatal("expected error for BEP2 token with decimals != 8")
	}
}

func TestValidateAssetDecimalsAccordingType_OtherTokenValid(t *testing.T) {
	if err := ValidateAssetDecimalsAccordingType("ERC20", 18); err != nil {
		t.Fatalf("unexpected error for non-BEP2 token: %v", err)
	}
}
