# Symbol SDK for Go
Symbol Blockchain Network SDK for Go  

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/karriz-dev/symbol-sdk)
![GitHub Release](https://img.shields.io/github/v/release/karriz-dev/symbol-sdk)

## Installation
```bash
go get github.com/karriz-dev/symbol-sdk
```
## Usage
Check `example` folder this repo. [go to example](https://github.com/karriz-dev/symbol-sdk/tree/main/example)

## Supported List
### Transaction Factory
 Function Name | Description | Supported 
:------------ | :-------------| :------------- 
Sign | transaction signing | :heavy_check_mark: 
Verify | verfying transaction signature | :heavy_check_mark:
MessageEncode | encoding message uft-8 (plain) | :heavy_check_mark: 
MessageEncode | encoding message uft-8 (encrypt) | :heavy_minus_sign: 
CoSign | transaction co-signing | :heavy_minus_sign: 

### Transaction
Type (Decimal) | Name | Description | Supported | Embedded
:-- |:------------ | :-------------| :------------- | :-------------
0x4154 | TransferTransactionV1 | transfer mosaic from signer to receipient | :heavy_check_mark: | :heavy_check_mark:
0x4148 | HashLockTransactionV1 | locked mosaic with sha256 hash | :heavy_check_mark: | :heavy_check_mark:
0x4241 | AggregateBondedV2 | aggregate bonded some transactions | :heavy_check_mark: | :heavy_minus_sign:


### Network Utilities
 Name | Description | Supported
:------------ | :-------------| :-------------
NetworkEstimateFee | get network estimated tx fee | :heavy_check_mark:
NetworkProperties | get network properties(/network/properties) | :heavy_check_mark:

## Contribute