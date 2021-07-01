package wallet

import (
	"code.vegaprotocol.io/go-wallet/crypto"
	wcrypto "code.vegaprotocol.io/go-wallet/wallet/crypto"
	typespb "github.com/vegaprotocol/api/grpc/clients/go/generated/code.vegaprotocol.io/vega/proto"
	commandspb "github.com/vegaprotocol/api/grpc/clients/go/generated/code.vegaprotocol.io/vega/proto/commands/v1"

	"github.com/golang/protobuf/proto"
)

type LegacyWallet struct {
	Owner   string        `json:"Owner"`
	KeyRing LegacyKeyRing `json:"Keypairs"`
}

func NewLegacyWallet(name string) *LegacyWallet {
	return &LegacyWallet{
		Owner:   name,
		KeyRing: NewLegacyKeyRing(),
	}
}

func (w *LegacyWallet) Name() string {
	return w.Owner
}

func (w *LegacyWallet) DescribePublicKey(pubKey string) (PublicKey, error) {
	keyPair, err := w.KeyRing.FindPair(pubKey)
	if err != nil {
		return nil, err
	}

	return keyPair.ToPublicKey(), nil
}

func (w *LegacyWallet) ListPublicKeys() []PublicKey {
	originalKeys := w.KeyRing.GetPublicKeys()
	keys := make([]PublicKey, len(originalKeys))
	for i, key := range originalKeys {
		keys[i] = key
	}
	return keys
}

func (w *LegacyWallet) GenerateKeyPair() (KeyPair, error) {
	kp, err := GenKeyPair(wcrypto.Ed25519)
	if err != nil {
		return nil, err
	}

	w.KeyRing.Upsert(*kp)

	deepCopy := kp.DeepCopy()
	return deepCopy, nil
}

func (w *LegacyWallet) TaintKey(pubKey string) error {
	keyPair, err := w.KeyRing.FindPair(pubKey)
	if err != nil {
		return err
	}

	if err := keyPair.Taint(); err != nil {
		return err
	}

	w.KeyRing.Upsert(keyPair)

	return nil
}

func (w *LegacyWallet) UpdateMeta(pubKey string, meta []Meta) error {
	keyPair, err := w.KeyRing.FindPair(pubKey)
	if err != nil {
		return err
	}

	keyPair.MetaList = meta

	w.KeyRing.Upsert(keyPair)
	return nil
}

func (w *LegacyWallet) SignAny(pubKey string, data []byte) ([]byte, error) {
	keyPair, err := w.KeyRing.FindPair(pubKey)
	if err != nil {
		return nil, err
	}

	if keyPair.Tainted {
		return nil, ErrPubKeyIsTainted
	}

	return keyPair.SignAny(data)
}

func (w *LegacyWallet) VerifyAny(pubKey string, data, sig []byte) (bool, error) {
	keyPair, err := w.KeyRing.FindPair(pubKey)
	if err != nil {
		return false, err
	}

	return keyPair.VerifyAny(data, sig)
}

func (w *LegacyWallet) SignTxV1(pubKey string, data []byte, blockHeight uint64) (SignedBundle, error) {
	keyPair, err := w.KeyRing.FindPair(pubKey)
	if err != nil {
		return SignedBundle{}, err
	}

	if keyPair.Tainted {
		return SignedBundle{}, ErrPubKeyIsTainted
	}

	txTy := &typespb.Transaction{
		InputData:   data,
		Nonce:       crypto.NewNonce(),
		BlockHeight: blockHeight,
		From: &typespb.Transaction_PubKey{
			PubKey: keyPair.pubBytes,
		},
	}

	rawTxTy, err := proto.Marshal(txTy)
	if err != nil {
		return SignedBundle{}, err
	}

	sig, err := keyPair.Algorithm.Sign(keyPair.privBytes, rawTxTy)
	if err != nil {
		return SignedBundle{}, err
	}

	return SignedBundle{
		Tx: rawTxTy,
		Sig: Signature{
			Sig:     sig,
			Algo:    keyPair.Algorithm.Name(),
			Version: keyPair.Algorithm.Version(),
		},
	}, nil
}

func (w *LegacyWallet) SignTxV2(pubKey string, data []byte) (*commandspb.Signature, error) {
	keyPair, err := w.KeyRing.FindPair(pubKey)
	if err != nil {
		return nil, err
	}

	if keyPair.Tainted {
		return nil, ErrPubKeyIsTainted
	}

	return keyPair.Sign(data)
}

type Meta struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
