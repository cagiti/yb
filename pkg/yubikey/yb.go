package yubikey

import (
	"fmt"

	"github.com/cagiti/yb/pkg/cmd"
	"github.com/enescakir/emoji"
	"github.com/go-piv/piv-go/piv"
	"github.com/pkg/errors"
)

// Stores yubikey card information.
type Yubikey struct {
	cmd.CommonOptions
	yb     *piv.YubiKey
	chosen string
}

// NewYubikey.
func NewYubikey() *Yubikey {
	yb := &Yubikey{}
	return yb
}

func ListYubikeys() ([]string, error) {
	cards, err := piv.Cards()
	if err != nil {
		return []string{}, err
	}
	return cards, nil
}

func (y *Yubikey) SelectYubikey() error {
	var err error
	y.yb, err = y.chooseYubikey()
	if err != nil {
		return errors.Wrap(err, "failed to select yubikey.")
	}
	return nil
}

func (y *Yubikey) chooseYubikey() (*piv.YubiKey, error) {
	ybs, err := ListYubikeys()
	if err != nil {
		return nil, errors.Wrap(err, "while listing connected yubikeys.")
	}
	prompter := y.Prompter()
	y.chosen, err = prompter.SelectFromOptions("Select a yubikey:", ybs)
	if err != nil {
		return nil, errors.Wrap(err, "when selecting yubikey device.")
	}
	yb, err := piv.Open(y.chosen)
	if err != nil {
		return nil, errors.Wrap(err, "when connecting to selected yubikey device.")
	}
	return yb, nil
}

func (y *Yubikey) ok() bool {
	// All yubikeys have a certificate, which is unique to the key
	// and signed by Yubico. If we're unable to retrieve it then the
	// yubikey is not ok.
	_, err := y.yb.AttestationCertificate()
	return err == nil
}

func (y *Yubikey) printYubikeyCertificate() error {
	cert, err := y.yb.AttestationCertificate()
	if err != nil {
		return errors.Wrap(err, "while accessing attestation certificate.")
	}
	// TODO: need to tidy this mess up - having a play around!
	if cert.Issuer.String() == "" {
		fmt.Printf("%v %s\n", emoji.CrossMarkButton, "No Issuer")
	} else {
		fmt.Printf("%v %s\n", emoji.CheckMarkButton, cert.Issuer)
	}
	if cert.PublicKeyAlgorithm.String() == "" {
		fmt.Printf("%v %s\n", emoji.Unlocked, "No Algorithm")
	} else {
		fmt.Printf("%v %s\n", emoji.Locked, cert.PublicKeyAlgorithm)
	}
	if !cert.IsCA {
		fmt.Printf("%v %s\n", emoji.NoEntry, "Not a valid CA")
	} else {
		fmt.Printf("%v  %s\n", emoji.Shield, cert.Subject)
	}
	return nil
}

func (y *Yubikey) Check() error {
	var err error
	if y.yb == nil || !y.ok() {
		if y.yb != nil {
			y.yb.Close()
		}
		y.yb, err = piv.Open(y.chosen)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("when connecting to yubkey: %s", y.chosen))
		}
	}
	err = y.printYubikeyCertificate()
	if err != nil {
		return err
	}
	return nil
}
