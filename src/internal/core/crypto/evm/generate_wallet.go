package crypto_evm

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golibs-starter/golib/log"
	"strings"
)

func GenerateWallet() (privateKey string, address string, err error) {
	ecdsaPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	privateKeyBytes := crypto.FromECDSA(ecdsaPrivateKey)
	privateKey = hexutil.Encode(privateKeyBytes)[2:]
	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKey, strings.ToLower(address), nil

}
