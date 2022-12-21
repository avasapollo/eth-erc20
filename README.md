# eth-erc20
## Introduction
This is a playground for the Ethereum [ERC-20](https://ethereum.org/en/developers/docs/standards/tokens/erc-20/) standard.
There are 3 applications:
- owner-creator: it creates an account on the keystore that will be used to deploy the contract. 
- contract-deploy: it deploys the contract to Ethereum network
- erc-20: it exposes GRPC/REST endpoint to interact with the DApp

## owner-creator
This application will create an EOA(Externally Owned Account). 
You can find more clarification on the Ethereum account [here](https://ethereum.org/en/developers/docs/accounts/#types-of-account).
We will use this account to deploy the contract, and it will be the owner of the DApp.
The application require these env variables
```
OWNER_PASSWORD=password                                 // password for the EOA account
OWNER_BALANCE=1000000000000                             // the initial token balance of the owner
KEY_DIR=./accounts                                      // where you want to store the accounts
```
you can use/change the script that I added. You should see something like this.
```
➜  eth-erc20 git:(master) ./scripts/owner-creator.sh
INFO[0001] owner address: 0x3F15cb553FAA92aD4fBde47E6CA727A4A0d49d85  app=owner-creator
```
You should see a new file inside your accounts directory.