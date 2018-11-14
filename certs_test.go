package googleTokenVerifier

import (
	"testing"
	"time"
)

func TestGetFederatedSignonCerts(t *testing.T) {
	certs, err := getFederatedSignonCerts()
	if err != nil {
		t.Error(err)
		return
	}

	cacheAge := certs.Expiry.Sub(time.Now()).Seconds()
	t.Logf("cacheAge: %f", cacheAge)
	if cacheAge <= 7200 {
		t.Error("max-age not found")
	}

	key := certs.Keys["aa436c3f63b281ce0d976da0b51a34860ff960eb"]
	if key == nil {
		t.Error("aa436c3f63b281ce0d976da0b51a34860ff960eb should exists")
	}
}

func TestGetFederatedSignonCertsCache(t *testing.T) {
	certs = &Certs{
		Expiry: time.Now(),
	}
	certs, err := getFederatedSignonCerts() // trigger update
	if err != nil {
		t.Error(err)
		return
	}
	key := certs.Keys["aa436c3f63b281ce0d976da0b51a34860ff960eb"]
	if key == nil {
		t.Error("aa436c3f63b281ce0d976da0b51a34860ff960eb should exists")
	}
}
