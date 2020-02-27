#!/bin/bash
echo
echo "##########################################################"
echo "##### Generate certificates using cryptogen tool #########"
echo "##########################################################"
bin/cryptogen generate --config=./artifacts/channel/cryptogen.yaml --output="./artifacts/channel/crypto-config"
bin/configtxgen -profile ThreeOrgsOrdererGenesis --configPath=./artifacts/channel -outputBlock ./artifacts/channel/genesis.block
bin/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel -outputCreateChannelTx ./artifacts/channel/mychannel.tx -channelID "mychannel"
bin/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel -outputAnchorPeersUpdate ./artifacts/channel/Org1MSPanchors.tx -channelID "mychannel" -asOrg Org1MSP
bin/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel -outputAnchorPeersUpdate ./artifacts/channel/Org2MSPanchors.tx -channelID "mychannel" -asOrg Org2MSP
bin/configtxgen -profile ThreeOrgsChannel --configPath=./artifacts/channel -outputAnchorPeersUpdate ./artifacts/channel/Org3MSPanchors.tx -channelID "mychannel" -asOrg Org3MSP
