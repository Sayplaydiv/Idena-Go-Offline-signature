package sdk

import (
	encod "encoding/hex"
	"fmt"
	txApi "github.com/idena-network/idena-go/blockchain/types"
	com "github.com/idena-network/idena-go/common"
	idnaCrypto "github.com/idena-network/idena-go/crypto"
	"github.com/shopspring/decimal"
	"math/big"
)

func SignTransaction(fromAddress string,fromAddrPrivateKey string,toAddress string,amount string,nonce uint32,epoch uint16) (string,error) {

	//PrivateKey to ECDSA_KEY
	dValue, err := encod.DecodeString(fromAddrPrivateKey)
	if err != nil {
		return "", err
	}
	key, err := idnaCrypto.ToECDSA(dValue)
	if err != nil {
		return "", err
	}
	//encoding address
	addr := idnaCrypto.PubkeyToAddress(key.PublicKey)
	if addr.String() != fromAddress {
		return "", fmt.Errorf("sendtx from and address not equal. Got %x want %x", fromAddress, addr)
	}

	//encoding toAddress
	toAddressByte, err := encod.DecodeString(toAddress[2:])
	if err != nil {
		return "", err
	}
	encodToAddress := com.BytesToAddress(toAddressByte)

	//get nonce
	nonceValue :=nonce

	//get epoch
	epochValue:=epoch

	//decimal : 18
	dec := decimal.NewFromFloat(1000000000000000000)
	decAmount, err := decimal.NewFromString(amount)
	if err != nil {
		return "", err
	}
	//float amount --> big.int amount
	strAmount := decAmount.Mul(dec).String()
	bigAmount, ok := new(big.Int).SetString(strAmount, 10)
	if !ok {
		return "", fmt.Errorf("sendTx amount ->bigin error ")
	}

	txData := txApi.Transaction{
		AccountNonce: nonceValue,
		Type:         txApi.SendTx,
		To:           &encodToAddress,
		Amount:       bigAmount,
		MaxFee:       new(big.Int).SetUint64(1000000000000000000), //max fee setting:1 IDNA
		Tips:         new(big.Int),
		Epoch:        epochValue,
	}


	//enciding hexing --->subumit transaction
	//process：signedTx-->ToProto-->proto.Marshal-->hash Encoding
	signedTx, err := txApi.SignTx(&txData, key)
	if err != nil {
		return "", err
	}
	byteSignedTx, err := signedTx.ToBytes()
	if err != nil {
		return "", err
	}
	hexing := "0x" + encod.EncodeToString(byteSignedTx)
	return hexing, nil
}
