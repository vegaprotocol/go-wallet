package wallet_test

import (
	"testing"

	commandspb "code.vegaprotocol.io/go-wallet/internal/proto/commands/v1"
	walletpb "code.vegaprotocol.io/go-wallet/internal/proto/wallet/v1"
	"code.vegaprotocol.io/go-wallet/wallet"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testHandler struct {
	*wallet.Handler
	ctrl  *gomock.Controller
	store *mockedStore
}

func getTestHandler(t *testing.T) *testHandler {
	ctrl := gomock.NewController(t)
	store := newMockedStore()

	h := wallet.NewHandler(store)
	return &testHandler{
		Handler: h,
		ctrl:    ctrl,
		store:   store,
	}
}

func TestHandler(t *testing.T) {
	t.Run("Creating a wallet succeeds", testHandlerCreatingWalletSucceeds)
	t.Run("Creating a HD wallet whereas a legacy wallet exists fails", testHandlerCreatingHDWalletWhereasLegacyWalletExistsFails)
	t.Run("Creating an already existing wallet fails", testHandlerCreatingAlreadyExistingWalletFails)
	t.Run("Importing a wallet succeeds", testHandlerImportingWalletSucceeds)
	t.Run("Importing a wallet with invalid mnemonic fails", testHandlerImportingWalletWithInvalidMnemonicFails)
	t.Run("Importing an already existing wallet fails", testHandlerImportingAlreadyExistingWalletFails)
	t.Run("Verifying wallet existence succeeds", testHandlerVerifyingWalletExistenceSucceeds)
	t.Run("Verifying wallet non existence succeeds", testHandlerVerifyingWalletNonExistenceSucceeds)
	t.Run("Recreating a wallet with same name fails", testHandlerRecreatingWalletWithSameNameFails)
	t.Run("Recreating a wallet with same name and different passphrase fails", testHandlerRecreatingWalletWithSameNameButDifferentPassphraseFails)
	t.Run("Login to existing wallet succeeds", testHandlerLoginToExistingWalletSucceeds)
	t.Run("Login to non-existing wallet fails", testHandlerLoginToNonExistingWalletFails)
	t.Run("Logout logged in wallet succeeds", testHandlerLogoutLoggedInWalletSucceeds)
	t.Run("Logout not-logged in wallet succeeds", testHandlerLogoutNotLoggedInWalletSucceeds)
	t.Run("Generating new key pair securely succeeds", testHandlerGeneratingNewKeyPairSecurelySucceeds)
	t.Run("Generating new key pair securely with invalid name fails", testHandlerGeneratingNewKeyPairSecurelyWithInvalidNameFails)
	t.Run("Generating new key pair securely without wallet fails", testHandlerGeneratingNewKeyPairSecurelyWithoutWalletFails)
	t.Run("Generating new key pair succeeds", testHandlerGeneratingNewKeyPairSucceeds)
	t.Run("Generating new key pair with invalid name fails", testHandlerGeneratingNewKeyPairWithInvalidNameFails)
	t.Run("Generating new key pair without wallet fails", testHandlerGeneratingNewKeyPairWithoutWalletFails)
	t.Run("Listing public keys succeeds", testHandlerListingPublicKeysSucceeds)
	t.Run("Listing public keys with logged out wallet fails", testHandlerListingPublicKeysWithLoggedOutWalletFails)
	t.Run("Listing public keys with invalid name fails", testHandlerListingPublicKeysWithInvalidNameFails)
	t.Run("Listing public keys without wallet fails", testHandlerListingPublicKeysWithoutWalletFails)
	t.Run("Listing key pairs succeeds", testHandlerListingKeyPairsSucceeds)
	t.Run("Listing key pairs with invalid name fails", testHandlerListingKeyPairsWithInvalidNameFails)
	t.Run("Listing key pairs with logged out wallet fails", testHandlerListingKeyPairsWithLoggedOutWalletFails)
	t.Run("Listing key pairs without wallet fails", testHandlerListingKeyPairsWithoutWalletFails)
	t.Run("Getting public key succeeds", testHandlerGettingPublicKeySucceeds)
	t.Run("Getting public key with logged out wallet fails", testHandlerGettingPublicKeyWithLoggedOutWalletFails)
	t.Run("Getting public key without wallet fails", testHandlerGettingPublicKeyWithoutWalletFails)
	t.Run("Getting public key with invalid name fails", testHandlerGettingPublicKeyWithInvalidNameFails)
	t.Run("Getting non-existing public key fails", testGettingNonExistingPublicKeyFails)
	t.Run("Tainting key pair succeeds", testHandlerTaintingKeyPairSucceeds)
	t.Run("Tainting key pair with invalid name fails", testHandlerTaintingKeyPairWithInvalidNameFails)
	t.Run("Tainting key pair without wallet fails", testHandlerTaintingKeyPairWithoutWalletFails)
	t.Run("Tainting key pair that is already tainted fails", testHandlerTaintingKeyThatIsAlreadyTaintedFails)
	t.Run("Updating key pair meta succeeds", testHandlerUpdatingKeyPairMetaSucceeds)
	t.Run("Updating key pair meta with invalid passphrase fails", testHandlerUpdatingKeyPairMetaWithInvalidPassphraseFails)
	t.Run("Updating key pair meta with invalid name fails", testHandlerUpdatingKeyPairMetaWithInvalidNameFails)
	t.Run("Updating key pair meta without wallet fails", testHandlerUpdatingKeyPairMetaWithoutWalletFails)
	t.Run("Updating key pair meta with non-existing public key fails", testHandlerUpdatingKeyPairMetaWithNonExistingPublicKeyFails)
	t.Run("Get wallet path succeeds", testHandlerGettingWalletPathSucceeds)
	t.Run("Signing transaction request (v2) succeeds", testHandlerSigningTxV2Succeeds)
	t.Run("Signing transaction request (v2) with logged out wallet fails", testHandlerSigningTxV2WithLoggedOutWalletFails)
	t.Run("Signing transaction request (v2) with tainted key fails", testHandlerSigningTxV2WithTaintedKeyFails)
	t.Run("Signing and verifying a message succeeds", testHandlerSigningAndVerifyingMessageSucceeds)
	t.Run("Signing a message with logged out wallet fails", testHandlerSigningMessageWithLoggedOutWalletFails)
	t.Run("Verifying a message with logged out wallet fails", testHandlerVerifyingMessageWithLoggedOutWalletSucceeds)
}

func testHandlerCreatingWalletSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)
}

func testHandlerCreatingHDWalletWhereasLegacyWalletExistsFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"
	legacyWallet := wallet.NewLegacyWallet(name)

	// given
	err := h.store.SaveWallet(legacyWallet, passphrase)

	// then
	require.NoError(t, err)

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletAlreadyExists.Error())
	assert.Empty(t, mnemonic)
}

func testHandlerCreatingAlreadyExistingWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	mnemonic, err = h.CreateWallet(name, passphrase)

	// then
	require.Error(t, err, wallet.ErrWalletAlreadyExists)
	assert.Empty(t, mnemonic)
}

func testHandlerImportingWalletSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	err := h.ImportWallet(name, passphrase, TestMnemonic1)

	// then
	require.NoError(t, err)
}

func testHandlerImportingWalletWithInvalidMnemonicFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	err := h.ImportWallet(name, passphrase, "this is not a valid mnemonic")

	// then
	require.EqualError(t, err, wallet.ErrInvalidMnemonic.Error())
}

func testHandlerImportingAlreadyExistingWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	err := h.ImportWallet(name, passphrase, TestMnemonic1)

	// then
	require.NoError(t, err)

	// when
	err = h.ImportWallet(name, passphrase, TestMnemonic2)

	// then
	require.Error(t, err, wallet.ErrWalletAlreadyExists)
}

func testHandlerVerifyingWalletExistenceSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	exists := h.WalletExists(name)

	// then
	assert.True(t, exists)
}

func testHandlerVerifyingWalletNonExistenceSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"

	// when
	exists := h.WalletExists(name)

	// then
	assert.False(t, exists)
}

func testHandlerRecreatingWalletWithSameNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	mnemonic, err = h.CreateWallet(name, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletAlreadyExists.Error())
	assert.Empty(t, mnemonic)
}

func testHandlerRecreatingWalletWithSameNameButDifferentPassphraseFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"
	othPassphrase := "different-passphrase"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	mnemonic, err = h.CreateWallet(name, othPassphrase)

	// then
	require.Error(t, err)
}

func testHandlerLoginToExistingWalletSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// then
	err = h.LoginWallet(name, passphrase)

	require.NoError(t, err)
}

func testHandlerLoginToNonExistingWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	err := h.LoginWallet(name, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
}

func testHandlerLogoutLoggedInWalletSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet("jeremy")
	})
}

func testHandlerLogoutNotLoggedInWalletSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet("jeremy")
	})
}

func testHandlerGeneratingNewKeyPairSecurelySucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	keys, err := h.ListPublicKeys(name)

	// then
	require.NoError(t, err)
	assert.Len(t, keys, 1)
	assert.Equal(t, key, keys[0].Key())
	assert.False(t, keys[0].IsTainted())
}

func testHandlerGeneratingNewKeyPairSecurelyWithInvalidNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"
	otherName := "bad name"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(otherName, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)
}

func testHandlerGeneratingNewKeyPairSecurelyWithoutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)
}

func testHandlerGeneratingNewKeyPairSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	keyPair, err := h.GenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, keyPair.PublicKey())
	assert.NotEmpty(t, keyPair.PrivateKey())
	assert.False(t, keyPair.IsTainted())
	assert.Empty(t, keyPair.Meta())

	// when
	keys, err := h.ListPublicKeys(name)

	// then
	require.NoError(t, err)
	assert.Len(t, keys, 1)
	assert.Equal(t, keyPair.PublicKey(), keys[0].Key())
	assert.False(t, keys[0].IsTainted())
}

func testHandlerGeneratingNewKeyPairWithInvalidNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"
	otherName := "bad name"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	keyPair, err := h.GenerateKeyPair(otherName, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, keyPair)
}

func testHandlerGeneratingNewKeyPairWithoutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"
	passphrase := "Th1isisasecurep@ssphraseinnit"

	// when
	keyPair, err := h.GenerateKeyPair(name, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, keyPair)
}

func testHandlerListingPublicKeysSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	keyPair, err := h.GenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotNil(t, keyPair)

	// when
	publicKeys, err := h.ListPublicKeys(name)

	// then
	require.NoError(t, err)
	assert.Len(t, publicKeys, 1)
	returnedPublicKey := publicKeys[0]
	assert.Equal(t, keyPair.PublicKey(), returnedPublicKey.Key())
	assert.Equal(t, keyPair.IsTainted(), returnedPublicKey.IsTainted())
	assert.Equal(t, keyPair.AlgorithmName(), returnedPublicKey.AlgorithmName())
	assert.Equal(t, keyPair.AlgorithmVersion(), returnedPublicKey.AlgorithmVersion())
	assert.Equal(t, keyPair.Meta(), returnedPublicKey.Meta())
}

func testHandlerListingPublicKeysWithLoggedOutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	keyPair, err := h.GenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotNil(t, keyPair)

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet(name)
	})

	// when
	publicKeys, err := h.ListPublicKeys(name)

	// then
	require.EqualError(t, err, wallet.ErrWalletNotLoggedIn.Error())
	assert.Len(t, publicKeys, 0)
}

func testHandlerListingPublicKeysWithInvalidNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"
	otherName := "bad name"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.ListPublicKeys(otherName)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)
}

func testHandlerListingPublicKeysWithoutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"

	// when
	key, err := h.ListPublicKeys(name)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)
}

func testHandlerListingKeyPairsSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	keyPair, err := h.GenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotNil(t, keyPair)

	// when
	publicKeys, err := h.ListKeyPairs(name)

	// then
	require.NoError(t, err)
	assert.Len(t, publicKeys, 1)
	returnedPublicKey := publicKeys[0]
	assert.Equal(t, keyPair.PublicKey(), returnedPublicKey.PublicKey())
	assert.Equal(t, keyPair.PrivateKey(), returnedPublicKey.PrivateKey())
	assert.Equal(t, keyPair.IsTainted(), returnedPublicKey.IsTainted())
	assert.Equal(t, keyPair.AlgorithmName(), returnedPublicKey.AlgorithmName())
	assert.Equal(t, keyPair.AlgorithmVersion(), returnedPublicKey.AlgorithmVersion())
	assert.Equal(t, keyPair.Meta(), returnedPublicKey.Meta())
}

func testHandlerListingKeyPairsWithLoggedOutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	keyPair, err := h.GenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotNil(t, keyPair)

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet(name)
	})

	// when
	publicKeys, err := h.ListKeyPairs(name)

	// then
	require.Error(t, err, wallet.ErrWalletNotLoggedIn)
	assert.Len(t, publicKeys, 0)
}

func testHandlerListingKeyPairsWithInvalidNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"
	otherName := "bad name"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.ListKeyPairs(otherName)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)
}

func testHandlerListingKeyPairsWithoutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"

	// when
	key, err := h.ListKeyPairs(name)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)
}

func testHandlerGettingPublicKeyWithoutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"

	// when
	key, err := h.GetPublicKey(name, name)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)
}

func testHandlerGettingPublicKeySucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	keyPair, err := h.GetPublicKey(name, key)

	require.NoError(t, err)
	assert.Equal(t, key, keyPair.Key())
}

func testHandlerGettingPublicKeyWithLoggedOutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet(name)
	})

	// when
	keyPair, err := h.GetPublicKey(name, key)

	require.Error(t, err)
	assert.Empty(t, keyPair)
}

func testHandlerGettingPublicKeyWithInvalidNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	otherName := "bad name"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	keyPair, err := h.GetPublicKey(otherName, key)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Nil(t, keyPair)
}

func testGettingNonExistingPublicKeyFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	keyPair, err := h.GetPublicKey(name, "non-existing-pub-key")
	assert.EqualError(t, err, wallet.ErrPubKeyDoesNotExist.Error())
	assert.Nil(t, keyPair)
}

func testHandlerTaintingKeyPairSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	publicKey, err := h.GetPublicKey(name, key)

	// then
	require.NoError(t, err)
	assert.NotNil(t, publicKey)
	assert.False(t, publicKey.IsTainted())

	// when
	err = h.TaintKey(name, key, passphrase)

	// then
	require.NoError(t, err)
	assert.True(t, h.store.GetKey(name, key).IsTainted())
}

func testHandlerTaintingKeyPairWithInvalidNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	otherName := "other name"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	keyPair, err := h.GetPublicKey(name, key)

	// then
	require.NoError(t, err)
	assert.NotNil(t, keyPair)
	assert.False(t, keyPair.IsTainted())

	// when
	err = h.TaintKey(otherName, key, passphrase)

	// then
	assert.Error(t, err)
}

func testHandlerTaintingKeyPairWithoutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	err := h.TaintKey(name, "non-existing-pub-key", passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
}

func testHandlerTaintingKeyThatIsAlreadyTaintedFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	keyPair, err := h.GetPublicKey(name, key)

	// then
	require.NoError(t, err)
	assert.NotNil(t, keyPair)
	assert.False(t, keyPair.IsTainted())

	// when
	err = h.TaintKey(name, key, passphrase)

	// then
	require.NoError(t, err)
	assert.True(t, h.store.GetKey(name, key).IsTainted())

	// when
	err = h.TaintKey(name, key, passphrase)

	// then
	assert.EqualError(t, err, wallet.ErrPubKeyAlreadyTainted.Error())
}

func testHandlerUpdatingKeyPairMetaSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	meta := []wallet.Meta{{Key: "primary", Value: "yes"}}

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	err = h.UpdateMeta(name, key, passphrase, meta)

	// then
	require.NoError(t, err)
	updatedKp := h.store.GetKey(name, key)
	assert.Len(t, updatedKp.Meta(), 1)
	assert.Equal(t, updatedKp.Meta()[0].Key, "primary")
	assert.Equal(t, updatedKp.Meta()[0].Value, "yes")
}

func testHandlerUpdatingKeyPairMetaWithInvalidPassphraseFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	othPassphrase := "other-passphrase"
	name := "jeremy"

	meta := []wallet.Meta{{Key: "primary", Value: "yes"}}

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	err = h.UpdateMeta(name, key, othPassphrase, meta)

	// then
	assert.Error(t, err)
	assert.Len(t, h.store.GetKey(name, key).Meta(), 0)
}

func testHandlerUpdatingKeyPairMetaWithInvalidNameFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"
	otherName := "other name"
	meta := []wallet.Meta{{Key: "primary", Value: "yes"}}

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	err = h.UpdateMeta(otherName, key, passphrase, meta)

	// then
	assert.Error(t, err)
	assert.Len(t, h.store.GetKey(name, key).Meta(), 0)
}

func testHandlerUpdatingKeyPairMetaWithoutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"
	pubKey := "non-existing-public-key"
	meta := []wallet.Meta{{Key: "primary", Value: "yes"}}

	// when
	err := h.UpdateMeta(name, pubKey, passphrase, meta)

	// then
	assert.Error(t, err)
}

func testHandlerUpdatingKeyPairMetaWithNonExistingPublicKeyFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"
	pubKey := "non-existing-public-key"
	meta := []wallet.Meta{{Key: "primary", Value: "yes"}}

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	key, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, key)

	// when
	err = h.UpdateMeta(name, pubKey, passphrase, meta)

	// then
	assert.Error(t, err)
}

func testHandlerGettingWalletPathSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	name := "jeremy"

	// when
	path, err := h.GetWalletPath(name)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, path)
}

func testHandlerSigningTxV2Succeeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	pubKey, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, pubKey)

	// given
	req := &walletpb.SubmitTransactionRequest{
		PubKey: pubKey,
		Command: &walletpb.SubmitTransactionRequest_OrderCancellation{
			OrderCancellation: &commandspb.OrderCancellation{},
		},
	}

	// when
	tx, err := h.SignTxV2(name, req, 42)

	// then
	require.NoError(t, err)
	assert.Equal(t, uint32(1), tx.Version)
	assert.NotEmpty(t, tx.From)
	assert.NotEmpty(t, tx.InputData)
	assert.NotNil(t, tx.Signature)
	key := h.store.GetKey(name, pubKey)
	assert.Equal(t, key.AlgorithmVersion(), tx.Signature.Version)
	assert.Equal(t, key.AlgorithmName(), tx.Signature.Algo)
	assert.NotEmpty(t, tx.Signature.Value)
}

func testHandlerSigningTxV2WithLoggedOutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	pubKey, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, pubKey)

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet(name)
	})

	// given
	req := &walletpb.SubmitTransactionRequest{
		PubKey: pubKey,
		Command: &walletpb.SubmitTransactionRequest_OrderCancellation{
			OrderCancellation: &commandspb.OrderCancellation{},
		},
	}

	// when
	tx, err := h.SignTxV2(name, req, 42)

	// then
	require.EqualError(t, err, wallet.ErrWalletNotLoggedIn.Error())
	assert.Nil(t, tx)
}

func testHandlerSigningTxV2WithTaintedKeyFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	pubKey, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, pubKey)

	// when
	err = h.TaintKey(name, pubKey, passphrase)

	// then
	require.NoError(t, err)
	assert.True(t, h.store.GetKey(name, pubKey).IsTainted())

	// given
	req := &walletpb.SubmitTransactionRequest{
		PubKey: pubKey,
		Command: &walletpb.SubmitTransactionRequest_OrderCancellation{
			OrderCancellation: &commandspb.OrderCancellation{},
		},
	}

	// when
	tx, err := h.SignTxV2(name, req, 42)

	// then
	assert.Error(t, err)
	assert.Nil(t, tx)
}

func testHandlerSigningAndVerifyingMessageSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	pubKey, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, pubKey)

	// given
	data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit. La peur est la petite mort qui conduit à l'oblitération totale.")

	// when
	sig, err := h.SignAny(name, data, pubKey)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, sig)

	// when
	verified, err := h.VerifyAny(name, data, sig, pubKey)

	// then
	require.NoError(t, err)
	assert.True(t, verified)
}

func testHandlerSigningMessageWithLoggedOutWalletFails(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	pubKey, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, pubKey)

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet(name)
	})

	// given
	data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit. La peur est la petite mort qui conduit à l'oblitération totale.")

	// when
	sig, err := h.SignAny(name, data, pubKey)

	// then
	require.EqualError(t, err, wallet.ErrWalletNotLoggedIn.Error())
	assert.Empty(t, sig)
}

func testHandlerVerifyingMessageWithLoggedOutWalletSucceeds(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// given
	passphrase := "Th1isisasecurep@ssphraseinnit"
	name := "jeremy"

	// when
	mnemonic, err := h.CreateWallet(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, mnemonic)

	// when
	pubKey, err := h.SecureGenerateKeyPair(name, passphrase)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, pubKey)

	// given
	data := []byte("Je ne connaîtrai pas la peur car la peur tue l'esprit. La peur est la petite mort qui conduit à l'oblitération totale.")

	// when
	sig, err := h.SignAny(name, data, pubKey)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, sig)

	// when
	assert.NotPanics(t, func() {
		h.LogoutWallet(name)
	})

	// when
	verified, err := h.VerifyAny(name, data, sig, pubKey)

	// then
	require.EqualError(t, err, wallet.ErrWalletNotLoggedIn.Error())
	assert.False(t, verified)
}
