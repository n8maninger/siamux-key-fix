package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gitlab.com/NebulousLabs/Sia/crypto"
	"gitlab.com/NebulousLabs/Sia/types"
	nebLog "gitlab.com/NebulousLabs/log"
	"gitlab.com/NebulousLabs/siamux"
	"gitlab.com/NebulousLabs/siamux/mux"
)

// loadHostKeyPair loads the host's public and private keys from the config file ignoring versioning or validation
func loadHostKeyPair(path string) (muxPub mux.ED25519PublicKey, muxPriv mux.ED25519SecretKey, err error) {
	f, err := os.Open(path)
	if err != nil {
		return mux.ED25519PublicKey{}, mux.ED25519SecretKey{}, fmt.Errorf("unable to open config file: %w", err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	// config files are prefixed with 3 lines for versioning and validation that can be ignored for our purposes
	var header string
	if err := dec.Decode(&header); err != nil {
		return mux.ED25519PublicKey{}, mux.ED25519SecretKey{}, fmt.Errorf("unable to decode header type: %w", err)
	}
	if err := dec.Decode(&header); err != nil {
		return mux.ED25519PublicKey{}, mux.ED25519SecretKey{}, fmt.Errorf("unable to decode header version: %w", err)
	}
	if err := dec.Decode(&header); err != nil {
		return mux.ED25519PublicKey{}, mux.ED25519SecretKey{}, fmt.Errorf("unable to decode header sig: %w", err)
	}

	// loads the keys from the config
	hostKey := struct {
		PublicKey types.SiaPublicKey `json:"publickey"`
		SecretKey crypto.SecretKey   `json:"secretkey"`
	}{}

	if err := dec.Decode(&hostKey); err != nil {
		return mux.ED25519PublicKey{}, mux.ED25519SecretKey{}, fmt.Errorf("unable to decode config: %w", err)
	}

	copy(muxPub[:], hostKey.PublicKey.Key)
	copy(muxPriv[:], hostKey.SecretKey[:])

	return
}

func main() {
	hostConfigPath := os.Args[1]

	muxPub, muxPriv, err := loadHostKeyPair(hostConfigPath)
	if err != nil {
		log.Fatalf("unable to load existing keypair: %s", err)
	}

	// let the siamux package initialize its own persistence because I'm lazy
	_, err = siamux.CompatV1421NewWithKeyPair(":0", ":0", nebLog.DiscardLogger, ".", muxPriv, muxPub)
	if err != nil {
		log.Fatalf("unable to generate config: %s", err)
	}
}
