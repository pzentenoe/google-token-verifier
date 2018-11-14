package googleTokenVerifier

import (
	"strings"
	"testing"
	"time"
)

const (
	validTestToken = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImQxZTg2OWU3YmY0MGRkYzNkM2RlMDgwNDI1OThiYTgzNTA5NzBmMGEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNzk4MTM4NDI2NjY3LWpiZmJ1c3RmaGY3bDI5dmZxYm5jczd0dTR1aTVtZWZoLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiNzk4MTM4NDI2NjY3LWpiZmJ1c3RmaGY3bDI5dmZxYm5jczd0dTR1aTVtZWZoLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTAxMTg2MTk4MTkwMzYyMzgzNzg3IiwiZW1haWwiOiJwemVudGVub2VAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF0X2hhc2giOiJyV2hlS1VfQ21TSG1tZ3lSUHJteWJBIiwibmFtZSI6IlBhYmxvIFplbnRlbm8iLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDYuZ29vZ2xldXNlcmNvbnRlbnQuY29tLy1ReHdocDdBMXQ0RS9BQUFBQUFBQUFBSS9BQUFBQUFBQWtDUS9qdUQ5aVlTS29tMC9zOTYtYy9waG90by5qcGciLCJnaXZlbl9uYW1lIjoiUGFibG8iLCJmYW1pbHlfbmFtZSI6IlplbnRlbm8iLCJsb2NhbGUiOiJlcyIsImlhdCI6MTU0MjIxMjQ1NSwiZXhwIjoxNTQyMjE2MDU1LCJqdGkiOiJhZTkwZjFhODMwYzg3OWVjNWU5MmE1OGI2MTgzYmU5MjZiMDFiODg5In0.YqbdA3uLF71PXH9DoIlS8slo5bHntN09iUmYwAvtmXIcCXPb-GVWkRyRYU7Aqv-63vnYCt_zY1cO6v_qsJu2RdFDbVfa8VDCEgTVXXG3h_LKf_85hgFm7BvJ4Svpt0WfJZLmPg0xY21GJREYrw6WP-hTToUn3aRfXvTvxyrdcA-FZ71UFkKs7enI_Rcn8LBPh_EIPZx7loD0bjqlVsRKmYnEm5x32Ai3ncMrH3hTA3oVnzftDQgGrjHd7WfMWNskGElx_VWpSBp8uDsOceQSrF2akROzUeNTscr2Ou-hITBPY_bv0RChUfpJF2QxAZXaI23fqclMFmOC_1hv7_dBWA"
	wrongSigToken  = validTestToken + "A"
)

func TestParseJWT(t *testing.T) {
	header, claimSet, _ := parseJWT(validTestToken)
	if len(header.KeyID) == 0 {
		t.Errorf("Invalid kid")
	}
	if len(claimSet.Email) == 0 {
		t.Errorf("Invalid Email")
	}
}

func TestVerifier(t *testing.T) {
	v := Verifier{}
	err := v.VerifyIDToken(wrongSigToken, []string{})
	if err != ErrWrongSignature {
		t.Error("Expect ErrWrongSignature")
	}
	err = v.VerifyIDToken(validTestToken, []string{})
	if err != ErrTokenUsedTooLate {
		t.Error("Expect ErrTokenUsedTooLate")
	}

	_, claimSet, _ := parseJWT(validTestToken)

	nowFn = func() time.Time {
		return time.Unix(claimSet.Exp, 0)
	}
	err = v.VerifyIDToken(validTestToken, []string{})
	if !strings.Contains(err.Error(), "Wrong aud:") {
		t.Log(err.Error())
		t.Error("Expect wrong aud error")
	}

	t.Log(claimSet.Aud)

	err = v.VerifyIDToken(validTestToken, []string{
		claimSet.Aud,
	})
	if err != nil {
		t.Error(err)
	}

	nowFn = time.Now
}
