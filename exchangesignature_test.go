package exchangesignature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrbix(t *testing.T) {
	t.Run("Test NewOrbix", func(t *testing.T) {
		secret := "fc8fa6ef2a9e4949bdf72d38208803657659ff67f2a74486a04a64b0bf1f2e6f"
		payload := "amount=1&nonce=2731832&pair=usdt_thb&price=31&side=buy&type=limit"
		sig, err := NewOrbix(secret, payload)
		if err != nil {
			t.Error("Error should be nil")
		}
		if sig == "" {
			t.Error("Signature should not be empty")
		}
		assert.Equal(t, sig, "5959460f890d9dad1fe1cdaf73bea955eef8c38da6a0b3139dbbe0d7e5fabfb3d0d3a4786767e759502ebd6d8878ac875441909f3c5232fa842c9349c03988bf")
	})

	t.Run("Test NewOrbix with invalid payload", func(t *testing.T) {
		secret := "fc8fa6ef2a9e4949bdf72d38208803657659ff67f2a74486a04a64b0bf1f2e6f"
		payload := "amount=1;nonce=2731832;pair=usdt_thb;price=31;side=buy;type=limit"
		_, err := NewOrbix(secret, payload)
		assert.Error(t, err)
	})
}

func TestGenerateNonceInt(t *testing.T) {
	t.Run("Test GenerateNonceInt", func(t *testing.T) {
		nonce := GenerateNonceInt()
		if nonce == 0 {
			t.Error("Nonce should not be 0")
		}
	})
}

func BenchmarkNewOrbix(b *testing.B) {
	secret := "fc8fa6ef2a9e4949bdf72d38208803657659ff67f2a74486a04a64b0bf1f2e6f"
	payload := "amount=1&nonce=2731832&pair=usdt_thb&price=31&side=buy&type=limit"
	for i := 0; i < b.N; i++ {
		_, _ = NewOrbix(secret, payload)
	}
}

func BenchmarkGenerateNonceInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateNonceInt()
	}
}

func TestNewBinance(t *testing.T) {
	t.Run("Test NewBinance", func(t *testing.T) {
		secret := "fc8fa6"
		payload := "amount=1&nonce=2731832&pair=usdt_thb&price=31&side=buy&type=limit"
		sig, err := NewBinance(secret, payload)
		if err != nil {
			t.Error("Error should be nil")
		}
		if sig == "" {
			t.Error("Signature should not be empty")
		}
	})
}

func BenchmarkNewBinance(b *testing.B) {
	secret := "fc8fa6"
	payload := "amount=1&nonce=2731832&pair=usdt_thb&price=31&side=buy&type=limit"
	for i := 0; i < b.N; i++ {
		_, _ = NewBinance(secret, payload)
	}
}
