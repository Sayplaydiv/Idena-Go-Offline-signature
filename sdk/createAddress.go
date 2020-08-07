package sdk

import (
	encod "encoding/hex"
	idnaCrypto "github.com/idena-network/idena-go/crypto"
)

func GetNewAccount() (string,string,error) {
	key, err := idnaCrypto.GenerateKey()
	if err != nil {
		return "","", err
	}
	//get privateKey
	privateKey := encod.EncodeToString(key.D.Bytes())
	//PublicKe to address
	addr := idnaCrypto.PubkeyToAddress(key.PublicKey)
	return addr.String(), privateKey, nil
}

