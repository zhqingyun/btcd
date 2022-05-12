// Copyright (c) 2014-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	//"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:8332",
		User:         "",
		Pass:         "",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
	txhashstring := "acb31e9da8b9f46a5b7f6f883c45e08cac8362cfdad9ef7af6b656ab3eb3f7c8";
	//txhash, _ := chainhash.NewHashFromStr(txhashstring)
	txhash, _ := chainhash.NewHashFromStr(txhashstring)
	txDetails, err := client.GetRawTransaction(txhash)
	txoutr1, err := client.GetTxOut(txhash, 1, false)
	txoutr2, err := client.GetTxOut(txhash, 1, true)
	for _, txout := range txDetails.MsgTx().TxOut {
		reeee, _ := client.DecodeScript(txout.PkScript)
		fmt.Println(reeee)
	}

	txresult, err := client.GetRawTransactionVerbose(txhash)

	for _, txout := range txDetails.MsgTx().TxOut {
		reeee, _ := client.DecodeScript(txout.PkScript)
		fmt.Println(reeee)
	}

	resp := client.GetRawTransactionVerboseAsync(txhash)
	rest  := client.GetRawTransactionAsync(txhash)


	fmt.Println(txresult)
	fmt.Println(resp)
	fmt.Println(rest)
	fmt.Println(txoutr1)
	fmt.Println(txoutr2)

	for _, oneTxOut := range txDetails.MsgTx().TxOut{
fmt.Println(oneTxOut)
	}

	fmt.Println(txDetails)



}
