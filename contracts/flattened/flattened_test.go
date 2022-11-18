package flattened

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestMint(t *testing.T) {
	c, err := ethclient.Dial("https://http-testnet.hecochain.com")
	assert.NoError(t, err)

	media, err := NewMedia(common.HexToAddress("0x08544e0bfda7c4699b2182485e75fb5af2ba4a05"), c)
	assert.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	hash, raw, err := media.MintForCreator(context.TODO(), big.NewInt(256), fromAddress, privateKey, common.Address{}, big.NewInt(int64(nonce)), big.NewInt(10001), "fps://123.com", "kfjlajdlfalfjda")
	assert.NoError(t, err)
	t.Logf("%v, %v", hash, raw)

}
