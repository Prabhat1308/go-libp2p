// +build !openssl

package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"

	btcec "github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ed25519"
)

// KeyPairFromKey generates a new private and public key from an input private key
func KeyPairFromKey(priv crypto.PrivateKey) (PrivKey, PubKey, error) {
	if priv == nil {
		return nil, nil, ErrNilPrivateKey
	}

	switch p := priv.(type) {
	case *rsa.PrivateKey:
		return &RsaPrivateKey{*p}, &RsaPublicKey{p.PublicKey}, nil

	case *ecdsa.PrivateKey:
		return &ECDSAPrivateKey{p}, &ECDSAPublicKey{&p.PublicKey}, nil

	case *ed25519.PrivateKey:
		pubIfc := p.Public()
		pub, _ := pubIfc.(ed25519.PublicKey)
		return &Ed25519PrivateKey{*p}, &Ed25519PublicKey{pub}, nil

	case *btcec.PrivateKey:
		sPriv := Secp256k1PrivateKey(*p)
		sPub := Secp256k1PublicKey(*p.PubKey())
		return &sPriv, &sPub, nil

	default:
		return nil, nil, ErrBadKeyType
	}
}
