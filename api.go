// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factom

import (
<<<<<<< HEAD
	"fmt"
	"encoding/json"
=======
	"encoding/json"
	"fmt"
>>>>>>> development
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	server = "localhost:8088"
)

// GetBlockHeight reports the current Directory Block Height
func GetBlockHeight() (int, error) {
	api := fmt.Sprintf("http://%s/v1/dblockheight/", server)
<<<<<<< HEAD
	
=======

>>>>>>> development
	resp, err := http.Get(api)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
<<<<<<< HEAD
	
=======

>>>>>>> development
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	height, err := strconv.Atoi(string(p))
	if err != nil {
		return 0, err
	}
	return height, nil
}

<<<<<<< HEAD
// GetDBlock gets a Directory Block by the Directory Block Hash. The Directory
// Block should contain a series of Entry Block Hashes.
func GetDBlock(hash string) (DBlock, error) {
	var dblock DBlock
=======
// GetChain gets a Entry Block Chain by the ChainID. The Chain should contain a
// series of Entry Block Hashes.
func GetChain(hash string) (*Chain, error) {
	chain := new(Chain)
	api := fmt.Sprintf("http://%s/v1/chain/%s", server, hash)

	resp, err := http.Get(api)
	if err != nil {
		return chain, err
	}
	defer resp.Body.Close()

	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return chain, err
	}

	err = json.Unmarshal(p, chain)
	if err != nil {
		return chain, err
	}

	return chain, nil
}

// GetChains gets all of the Chains. Each Chain should contain a series of
// Entry Block Hashes
func GetChains() ([]Chain, error) {
	chains := make([]Chain, 0)
	api := fmt.Sprintf("http://%s/v1/chains/", server)

	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	for {
		var chain Chain
		if err := dec.Decode(&chain); err == io.EOF {
			break
		} else if err != nil {
			return chains, err
		}
		chains = append(chains, chain)
	}

	return chains, nil
}

// GetDBlock gets a Directory Block by the Directory Block Hash. The Directory
// Block should contain a series of Entry Block Hashes.
func GetDBlock(hash string) (*DBlock, error) {
	dblock := new(DBlock)
>>>>>>> development
	api := fmt.Sprintf("http://%s/v1/dblock/%s", server, hash)

	resp, err := http.Get(api)
	if err != nil {
		return dblock, err
	}
	defer resp.Body.Close()
<<<<<<< HEAD
	
=======

>>>>>>> development
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dblock, err
	}
<<<<<<< HEAD
	
=======

>>>>>>> development
	err = json.Unmarshal(p, dblock)
	if err != nil {
		return dblock, err
	}

	return dblock, nil
}

// GetDBlocks gets the Directory Blocks whithin the Block Height Range provided
// (inclusive). Each DBlock should contain a series of Entry Block Merkel Roots.
func GetDBlocks(from, to int) ([]DBlock, error) {
	dblocks := make([]DBlock, 0)
	api := fmt.Sprintf("http://%s/v1/dblocksbyrange/%s/%s", server,
		strconv.Itoa(from), strconv.Itoa(to))

	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
<<<<<<< HEAD
	
=======

>>>>>>> development
	dec := json.NewDecoder(resp.Body)
	for {
		var block DBlock
		if err := dec.Decode(&block); err == io.EOF {
			break
		} else if err != nil {
			return dblocks, err
		}
		dblocks = append(dblocks, block)
	}

	return dblocks, nil
}

<<<<<<< HEAD
// GetEBlock gets an entry block specified by the Entry Block Merkel Root. The
// EBlock should contain a series of Entry Hashes.
func GetEBlock(s string) (EBlock, error) {
	var eblock EBlock
	api := fmt.Sprintf("http://%s/v1/eblockbymr/%s", server, s)
=======
// GetDBInfo gets the Directory Block information by the Directory Block Hash.
// The Directory Block Info should contain information about the directory
// block and the BTC transaction containing the Merkle Root.
func GetDBInfo(hash string) (*DBInfo, error) {
	dbinfo := new(DBInfo)
	api := fmt.Sprintf("http://%s/v1/dbinfo/%s", server, hash)

	resp, err := http.Get(api)
	if err != nil {
		return dbinfo, err
	}
	defer resp.Body.Close()

	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dbinfo, err
	}
	err = json.Unmarshal(p, dbinfo)
	if err != nil {
		return dbinfo, err
	}

	return dbinfo, nil
}

// GetEBlock gets an entry block specified by the Entry Block Merkel Root. The
// EBlock should contain a series of Entry Hashes.
func GetEBlock(s string) (*EBlock, error) {
	eblock := new(EBlock)
	api := fmt.Sprintf("http://%s/v1/eblock/%s", server, s)
>>>>>>> development

	resp, err := http.Get(api)
	if err != nil {
		return eblock, err
	}
	defer resp.Body.Close()
<<<<<<< HEAD
	
=======

>>>>>>> development
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return eblock, err
	}
	err = json.Unmarshal(p, eblock)
	if err != nil {
		return eblock, err
	}
	return eblock, nil
}

// GetEntry gets an entry based on the Entry Hash. The Entry should contain a
// hex encoded string of Entry Data and a series of External IDs.
<<<<<<< HEAD
func GetEntry(s string) (Entry, error) {
	var entry Entry
=======
func GetEntry(s string) (*Entry, error) {
	entry := new(Entry)
>>>>>>> development
	api := fmt.Sprintf("http://%s/v1/entry/%s", server, s)

	resp, err := http.Get(api)
	if err != nil {
		return entry, err
	}
	defer resp.Body.Close()
<<<<<<< HEAD
	
=======

>>>>>>> development
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entry, err
	}
	err = json.Unmarshal(p, entry)

	return entry, nil
}

<<<<<<< HEAD


// TODO ...
// ........

// GetChain gets a series of Entry Hashes associated with the Chain ID provided
func GetChain(s string) (Chain, error) {
	var c Chain
	return c, nil
=======
func GetEntriesByExtID(eid string) ([]Entry, error) {
	entries := make([]Entry, 0)
	api := fmt.Sprintf("http://%s/v1/entriesbyeid/%s", server, eid)

	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	for {
		var entry Entry
		if err := dec.Decode(&entry); err == io.EOF {
			break
		} else if err != nil {
			return entries, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
>>>>>>> development
}
