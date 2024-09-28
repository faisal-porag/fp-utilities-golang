package fp_cm_utils

import "testing"

func TestSecureEncodeDecodeUtilsFunc(t *testing.T) {
	key := "e2f8c35b9cd447b2e29f3738b54c4f2e"
	text := "Sensitive data to be encoded/decoded. Contact with porag2619@gmail.com"

	encoded, err := SecureEncodeUtilsFunc(text, key)
	if err != nil {
		t.Errorf("Encoding failed: %v", err)
	}

	decoded, err := SecureDecodeUtilsFunc(encoded, key)
	if err != nil {
		t.Errorf("Decoding failed: %v", err)
	}

	if decoded != text {
		t.Errorf("Expected %v, got %v", text, decoded)
	}
}
