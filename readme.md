## Idena-Go Offline signature Documentation：

#### GetNewAccount() ->> address + privateKey

#### SignTransaction() ---(Offline signature)-->> hexing


#### get nonce api：

```
HTTP PSOT:
req: 
{
  "method": "dna_getBalance",
  "params": [
    "0xcc233A3dbcf9A2a6949d21d1092C010CAA75C2da"
  ],
  "id": 1
}

resp:
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "stake": "0",
    "balance": "0",
    "nonce": 0
  }
}

PS: transaction use should  nonce+1
```

#### get epoch api:

```
HTTP PSOT:
req: 
{
  "method": "dna_epoch",
  "params": [],
  "id": 3,
  "key": "apikey"
}

resp:
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "epoch": 51,
    "nextValidation": "2020-08-15T13:30:00Z",
    "currentPeriod": "None",
    "currentValidationStart": "2020-08-15T13:30:00Z"
  }
}
```

#### submit transaction

```
//use SignTransaction() --->> get hexing
//use hexing submit transaction

req: 
{
  "method": "bcn_sendRawTx",
  "params": [hexing],
  "id": 3,
  "key": "apikey"
}
resp:
{
  "jsonrpc": "2.0",
  "id": 4,
  "result": "txHash"
}


```
