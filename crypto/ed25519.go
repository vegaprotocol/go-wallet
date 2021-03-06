package crypto

import (
	"crypto"
	"errors"

	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
)

var (
	// ErrBadED25519PrivateKeyLength is returned if a private key with incorrect length is supplied.
	ErrBadED25519PrivateKeyLength = errors.New("bad ed25519 private key length")

	// ErrBadED25519PublicKeyLength is returned if a public key with incorrect length is supplied.
	ErrBadED25519PublicKeyLength = errors.New("bad ed25519 public key length")
)

// TODO Rethink this struct to hold a ed25519 key pair
type ed25519Sig struct{}

func newEd25519() *ed25519Sig {
	return &ed25519Sig{}
}

// TODO Remove once the LegacyWallet is removed since the key generation will
//   	be handled by the slip10 library
func (e *ed25519Sig) GenKey() (crypto.PublicKey, crypto.PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, nil, err
	}

	return []byte(pub), []byte(priv), nil
}

func (e *ed25519Sig) Sign(priv crypto.PrivateKey, buf []byte) ([]byte, error) {
	privBytes := priv.([]byte)
	// Avoid panic by checking key length
	if len(privBytes) != ed25519.PrivateKeySize {
		return nil, ErrBadED25519PrivateKeyLength
	}
	return ed25519.Sign(privBytes, Hash(buf)), nil
}

func (e *ed25519Sig) Verify(pub crypto.PublicKey, message, sig []byte) (bool, error) {
	pubBytes := pub.([]byte)
	// Avoid panic by checking key length
	if len(pubBytes) != ed25519.PublicKeySize {
		return false, ErrBadED25519PublicKeyLength
	}
	return ed25519.Verify(pubBytes, Hash(message), sig), nil
}

func (e *ed25519Sig) Name() string {
	return "vega/ed25519"
}

func (e *ed25519Sig) Version() uint32 {
	return 1
}
