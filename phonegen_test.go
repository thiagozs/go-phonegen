package phonegen

import (
	"strings"
	"testing"
)

// TestRandom checks if the Random function generates the correct number of phone numbers.
func TestRandom(t *testing.T) {
	pg := New()
	limit := 10
	phones := pg.Random(limit)
	if len(phones) != limit {
		t.Errorf("Random() generated %d phone numbers, expected %d", len(phones), limit)
	}
}

// TestRandomE164 checks if the RandomE164 function generates phone numbers in E.164 format.
func TestRandomE164(t *testing.T) {
	pg := New()
	limit := 5
	countryCode := "55" // Brazil country code
	phones, err := pg.RandomE164(limit, countryCode)
	if err != nil {
		t.Errorf("RandomE164() returned an error: %v", err)
	}
	if len(phones) != limit {
		t.Errorf("RandomE164() generated %d phone numbers, expected %d", len(phones), limit)
	}
	for _, phone := range phones {
		if !strings.HasPrefix(phone, "+"+countryCode) {
			t.Errorf("Phone number %s does not start with the country code +%s", phone, countryCode)
		}
	}
}

// TestRandomMobile checks if the RandomMobile function generates mobile phone numbers.
func TestRandomMobile(t *testing.T) {
	pg := New()
	limit := 5
	phones := pg.RandomMobile(limit)
	if len(phones) != limit {
		t.Errorf("RandomMobile() generated %d phone numbers, expected %d", len(phones), limit)
	}
	for _, phone := range phones {
		// Assuming that the second digit should be '9' for mobile numbers
		if phone[2] != '9' {
			t.Errorf("Mobile phone number %s does not have '9' as the second digit", phone)
		}
	}
}

// TestRandomLandline checks if the RandomLandline function generates landline phone numbers.
func TestRandomLandline(t *testing.T) {
	pg := New()
	limit := 10 // or any other number of phone numbers you want to generate
	phones := pg.RandomLandline(limit)
	for _, phone := range phones {
		if len(phone) != 10 {
			t.Errorf("Generated phone number %s does not have the correct length", phone)
		}
		if phone[2] == '9' {
			t.Errorf("Landline phone number %s incorrectly contains '9' after the area code", phone)
		}
	}
}

// TestApplyMask checks if the applyMask function correctly applies the mask to phone numbers.
func TestApplyMask(t *testing.T) {
	phone := "11977569420"
	expected := "(11) 97756-9420"
	maskedPhone := applyMask(phone)
	if maskedPhone != expected {
		t.Errorf("applyMask() returned %s, expected %s", maskedPhone, expected)
	}
}

// TestGetNumberPattern checks if the getNumberPattern function returns a valid pattern.
func TestGetNumberPattern(t *testing.T) {
	areaCode := "11"
	pattern, err := getNumberPattern(areaCode)
	if err != nil {
		t.Errorf("getNumberPattern() returned an error for area code %s: %v", areaCode, err)
	}
	if pattern == "" {
		t.Errorf("getNumberPattern() returned an empty pattern for area code %s", areaCode)
	}
}

// Run all the tests
func TestMain(m *testing.M) {
	m.Run()
}
