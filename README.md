# Symbol SDK for Go
Symbol Blockchain Network SDK for Go  

[![Quality gate](http://xym0.kalee.land:10001/api/project_badges/quality_gate?project=karriz-dev_symbol-sdk_322a0165-b934-4a5d-b199-84f5bbd42784&token=sqb_547ac2ec0a0df30ca9e4b15e8067841831119b0c)](http://xym0.kalee.land:10001/dashboard?id=karriz-dev_symbol-sdk_322a0165-b934-4a5d-b199-84f5bbd42784)

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
Type | Name | Description | Supported | Embedded
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
