package scanossapi

import (
	"encoding/json"
	"testing"
)

func TestReachabilityCompatibilityFailureModel(t *testing.T) {
	const payload = `{
		"info_code":"UNSUPPORTED_SCHEMA",
		"info_message":"unsupported findings schema version 1.4",
		"status":{"status":"SUCCESS"}
	}`

	var response ComponentReachabilityResponse
	if err := json.Unmarshal([]byte(payload), &response); err != nil {
		t.Fatalf("unmarshal generated response: %v", err)
	}
	if response.InfoCode != "UNSUPPORTED_SCHEMA" || response.InfoMessage == nil {
		t.Fatalf("response = %+v, want explicit compatibility failure", response)
	}
}

func TestCryptoMetadataRoundTrip(t *testing.T) {
	const payload = `{"provider":"BC","futureKey":{"nested":true}}`

	var metadata CryptoFinderMetadata
	if err := json.Unmarshal([]byte(payload), &metadata); err != nil {
		t.Fatalf("unmarshal generated metadata: %v", err)
	}
	if metadata.Provider == nil || *metadata.Provider != "BC" {
		t.Fatalf("provider = %v, want BC", metadata.Provider)
	}
	if _, ok := metadata.Get("futureKey"); !ok {
		t.Fatal("unknown metadata key was not retained")
	}
	if _, ok := metadata.Get("cryptoModule"); ok {
		t.Fatal("cryptoModule was synthesized")
	}

	roundTrip, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("marshal generated metadata: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(roundTrip, &raw); err != nil {
		t.Fatalf("decode round trip: %v", err)
	}
	if raw["futureKey"] == nil {
		t.Fatal("unknown metadata key was lost on round trip")
	}
}
