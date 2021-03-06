package wallet

import (
	"errors"
	"fmt"
	"sync"

	"code.vegaprotocol.io/go-wallet/commands"
	commandspb "code.vegaprotocol.io/go-wallet/internal/proto/commands/v1"
	walletpb "code.vegaprotocol.io/go-wallet/internal/proto/wallet/v1"

	"github.com/golang/protobuf/proto"
)

var (
	ErrPubKeyIsTainted = errors.New("public key is tainted")
)

type Wallet interface {
	Version() uint32
	Name() string
	DescribePublicKey(pubKey string) (PublicKey, error)
	ListPublicKeys() []PublicKey
	ListKeyPairs() []KeyPair
	GenerateKeyPair() (KeyPair, error)
	TaintKey(pubKey string) error
	UpdateMeta(pubKey string, meta []Meta) error
	SignAny(pubKey string, data []byte) ([]byte, error)
	VerifyAny(pubKey string, data, sig []byte) (bool, error)
	SignTxV1(pubKey string, data []byte, blockHeight uint64) (SignedBundle, error)
	SignTxV2(pubKey string, data []byte) (*commandspb.Signature, error)
}

type KeyPair interface {
	PublicKey() string
	PrivateKey() string
	IsTainted() bool
	Meta() []Meta
	AlgorithmVersion() uint32
	AlgorithmName() string
}

type PublicKey interface {
	Key() string
	IsTainted() bool
	Meta() []Meta
	AlgorithmVersion() uint32
	AlgorithmName() string

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

// Store abstracts the underlying storage for wallet data.
type Store interface {
	WalletExists(name string) bool
	SaveWallet(w Wallet, passphrase string) error
	GetWallet(name, passphrase string) (Wallet, error)
	GetWalletPath(name string) string
}

type Handler struct {
	store         Store
	loggedWallets wallets

	// just to make sure we do not access same file concurrently or the map
	mu sync.RWMutex
}

func NewHandler(store Store) *Handler {
	return &Handler{
		store:         store,
		loggedWallets: newWallets(),
	}
}

func (h *Handler) WalletExists(name string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.store.WalletExists(name)
}

func (h *Handler) CreateWallet(name, passphrase string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.store.WalletExists(name) {
		return "", ErrWalletAlreadyExists
	}

	w, mnemonic, err := NewHDWallet(name)
	if err != nil {
		return "", err
	}

	err = h.saveWallet(w, passphrase)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

func (h *Handler) ImportWallet(name, passphrase, mnemonic string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.store.WalletExists(name) {
		return ErrWalletAlreadyExists
	}

	w, err := ImportHDWallet(name, mnemonic)
	if err != nil {
		return err
	}

	return h.saveWallet(w, passphrase)
}

func (h *Handler) LoginWallet(name, passphrase string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	w, err := h.store.GetWallet(name, passphrase)
	if err != nil {
		return ErrWalletDoesNotExists
	}

	h.loggedWallets.Add(w)

	return nil
}

func (h *Handler) LogoutWallet(name string) {
	h.loggedWallets.Remove(name)
}

func (h *Handler) GenerateKeyPair(name, passphrase string) (KeyPair, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	w, err := h.store.GetWallet(name, passphrase)
	if err != nil {
		return nil, err
	}

	kp, err := w.GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	err = h.saveWallet(w, passphrase)
	if err != nil {
		return nil, err
	}

	return kp, nil
}

func (h *Handler) SecureGenerateKeyPair(name, passphrase string) (string, error) {
	kp, err := h.GenerateKeyPair(name, passphrase)
	if err != nil {
		return "", err
	}

	return kp.PublicKey(), nil
}

func (h *Handler) GetPublicKey(name, pubKey string) (PublicKey, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	w, err := h.getLoggedWallet(name)
	if err != nil {
		return nil, err
	}

	return w.DescribePublicKey(pubKey)
}

func (h *Handler) ListPublicKeys(name string) ([]PublicKey, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	w, err := h.getLoggedWallet(name)
	if err != nil {
		return nil, err
	}

	return w.ListPublicKeys(), nil
}

func (h *Handler) ListKeyPairs(name string) ([]KeyPair, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	w, err := h.getLoggedWallet(name)
	if err != nil {
		return nil, err
	}

	return w.ListKeyPairs(), nil
}

func (h *Handler) SignAny(name string, inputData []byte, pubKey string) ([]byte, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	w, err := h.getLoggedWallet(name)
	if err != nil {
		return nil, err
	}

	return w.SignAny(pubKey, inputData)
}

func (h *Handler) SignTxV2(name string, req *walletpb.SubmitTransactionRequest, height uint64) (*commandspb.Transaction, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	w, err := h.getLoggedWallet(name)
	if err != nil {
		return nil, err
	}

	data := commands.NewInputData(height)
	wrapRequestCommandIntoInputData(data, req)
	marshalledData, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}

	pubKey := req.GetPubKey()
	signature, err := w.SignTxV2(pubKey, marshalledData)
	if err != nil {
		return nil, err
	}

	return commands.NewTransaction([]byte(pubKey), marshalledData, signature), nil
}

func (h *Handler) VerifyAny(name string, inputData, sig []byte, pubKey string) (bool, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	w, err := h.getLoggedWallet(name)
	if err != nil {
		return false, err
	}

	return w.VerifyAny(pubKey, inputData, sig)
}

func (h *Handler) TaintKey(name, pubKey, passphrase string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	w, err := h.store.GetWallet(name, passphrase)
	if err != nil {
		return err
	}

	err = w.TaintKey(pubKey)
	if err != nil {
		return err
	}

	return h.saveWallet(w, passphrase)
}

func (h *Handler) UpdateMeta(name, pubKey, passphrase string, meta []Meta) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	w, err := h.store.GetWallet(name, passphrase)
	if err != nil {
		return err
	}

	err = w.UpdateMeta(pubKey, meta)
	if err != nil {
		return err
	}

	return h.saveWallet(w, passphrase)
}

func (h *Handler) GetWalletPath(name string) (string, error) {
	return h.store.GetWalletPath(name), nil
}

func (h *Handler) saveWallet(w Wallet, passphrase string) error {
	err := h.store.SaveWallet(w, passphrase)
	if err != nil {
		return err
	}

	h.loggedWallets.Add(w)

	return nil
}

func wrapRequestCommandIntoInputData(data *commandspb.InputData, req *walletpb.SubmitTransactionRequest) {
	switch cmd := req.Command.(type) {
	case *walletpb.SubmitTransactionRequest_OrderSubmission:
		data.Command = &commandspb.InputData_OrderSubmission{
			OrderSubmission: req.GetOrderSubmission(),
		}
	case *walletpb.SubmitTransactionRequest_OrderCancellation:
		data.Command = &commandspb.InputData_OrderCancellation{
			OrderCancellation: req.GetOrderCancellation(),
		}
	case *walletpb.SubmitTransactionRequest_OrderAmendment:
		data.Command = &commandspb.InputData_OrderAmendment{
			OrderAmendment: req.GetOrderAmendment(),
		}
	case *walletpb.SubmitTransactionRequest_VoteSubmission:
		data.Command = &commandspb.InputData_VoteSubmission{
			VoteSubmission: req.GetVoteSubmission(),
		}
	case *walletpb.SubmitTransactionRequest_WithdrawSubmission:
		data.Command = &commandspb.InputData_WithdrawSubmission{
			WithdrawSubmission: req.GetWithdrawSubmission(),
		}
	case *walletpb.SubmitTransactionRequest_LiquidityProvisionSubmission:
		data.Command = &commandspb.InputData_LiquidityProvisionSubmission{
			LiquidityProvisionSubmission: req.GetLiquidityProvisionSubmission(),
		}
	case *walletpb.SubmitTransactionRequest_ProposalSubmission:
		data.Command = &commandspb.InputData_ProposalSubmission{
			ProposalSubmission: req.GetProposalSubmission(),
		}
	case *walletpb.SubmitTransactionRequest_NodeRegistration:
		data.Command = &commandspb.InputData_NodeRegistration{
			NodeRegistration: req.GetNodeRegistration(),
		}
	case *walletpb.SubmitTransactionRequest_NodeVote:
		data.Command = &commandspb.InputData_NodeVote{
			NodeVote: req.GetNodeVote(),
		}
	case *walletpb.SubmitTransactionRequest_NodeSignature:
		data.Command = &commandspb.InputData_NodeSignature{
			NodeSignature: req.GetNodeSignature(),
		}
	case *walletpb.SubmitTransactionRequest_ChainEvent:
		data.Command = &commandspb.InputData_ChainEvent{
			ChainEvent: req.GetChainEvent(),
		}
	case *walletpb.SubmitTransactionRequest_OracleDataSubmission:
		data.Command = &commandspb.InputData_OracleDataSubmission{
			OracleDataSubmission: req.GetOracleDataSubmission(),
		}
	default:
		panic(fmt.Errorf("command %v is not supported", cmd))
	}
}

func (h *Handler) getLoggedWallet(name string) (Wallet, error) {
	exists := h.store.WalletExists(name)
	if !exists {
		return nil, ErrWalletDoesNotExists
	}

	w, loggedIn := h.loggedWallets.Get(name)
	if !loggedIn {
		return nil, ErrWalletNotLoggedIn
	}
	return w, nil
}

type wallets map[string]Wallet

func newWallets() wallets {
	return map[string]Wallet{}
}

func (w wallets) Add(wallet Wallet) {
	w[wallet.Name()] = wallet
}

func (w wallets) Get(name string) (Wallet, bool) {
	wallet, ok := w[name]
	return wallet, ok
}

func (w wallets) Remove(name string) {
	delete(w, name)
}
