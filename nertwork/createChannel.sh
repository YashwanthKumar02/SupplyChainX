#!/bin/bash
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel.tx -channelID mychannel
peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel.tx
peer channel join -b mychannel.block