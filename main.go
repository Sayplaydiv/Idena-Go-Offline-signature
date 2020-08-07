package main

import (
	"golang_OfflineSignature/sdk"
	"log"
)

func main()  {
	//create newAddress
	newAddress,privateKey,err:=sdk.GetNewAccount()
	if err!=nil {
		log.Println(err)
	}
	log.Println("newAddress:",newAddress)
	log.Println("privateKey:",privateKey)

	//Offline signature
	//fromAddress string,fromAddrPrivateKey string,toAddress string,amount string,nonce uint32,epoch uint16
	fromAddress:="0xcc233A3dbcf9A2a6949d21d1092C010CAA75C2da"
	fromAddrPrivateKey:="b5df47a343b362e2969fe0897da9d6e2d923aa9d06612a417815a5dc7c2e2e66"
	toAddress:="0xcc233A3dbcf9A2a6949d21d1092C010CAA75C2da"
	amount:="1.2345678"
	nonce:=uint32(0)
	epoch:=uint16(0)

	hexing,err:=sdk.SignTransaction(fromAddress,fromAddrPrivateKey,toAddress,amount,nonce,epoch)
	if err!=nil {
		log.Println(err)
	}
	log.Println("Offline signature data：",hexing)
}