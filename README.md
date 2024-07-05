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
Verify | verfing transaction signature | :heavy_check_mark:
MessageEncode | encoding message uft-8 (plain) | :heavy_minus_sign: 
MessageEncode | encoding message uft-8 (encrypt) | :heavy_minus_sign: 
CoSign | transaction co-signing | :heavy_minus_sign: 

### Transaction
ID | Name | Description | Supported
:-- |:------------ | :-------------| :-------------
16724 | TransferTransactionV1 | transfer mosaic from signer to receipient | :heavy_check_mark:
16717 | MosaicDefinitionV1 | definition new a mosaic | :heavy_minus_sign:
16973 | MosaicSupplyChange | change mosaic's supply | :heavy_minus_sign:

### Network Utilities
 Name | Description | Supported
:------------ | :-------------| :-------------
NetworkEstimateFee | get network estimated tx fee | :heavy_minus_sign:
NetworkProperties | get network properties(/network/properties) | :heavy_minus_sign:

## Contribute