package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

var (
	pubkey  = "/home/akshith/testuser-gpg.asc"
	seckey  = "/home/akshith/testuser-secret.asc"
	passkey = "Ir0nm@nr0ck$"
	message = "test message"
)

func main() {
	/*
		err := encr()
		if err != nil {
			fmt.Printf("err : %+v\n", err)
			os.Exit(1)
		}
	*/
	decrMsg, err := decr()
	if err != nil {
		fmt.Printf("err : %+v\n", err)
		os.Exit(1)
	}
	fmt.Println("********************")
	fmt.Println(string(decrMsg))
	fmt.Println("********************")
}

func readEntity(filename string) (*openpgp.Entity, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	block, err := armor.Decode(f)
	if err != nil {
		return nil, err
	}

	return openpgp.ReadEntity(packet.NewReader(block.Body))

}

func encr() error {
	recepient, err := readEntity(pubkey)
	if err != nil {
		return err
	}

	signer, err := readEntity(seckey)
	if err != nil {
		return err
	}

	err = signer.PrivateKey.Decrypt([]byte(passkey))
	if err != nil {
		return err
	}

	r := bytes.NewReader([]byte(message))
	w := bytes.Buffer{}

	wc, err := openpgp.Encrypt(&w, []*openpgp.Entity{recepient}, signer, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = io.Copy(wc, r)
	if err != nil {
		return err
	}

	encrMsg := base64.StdEncoding.EncodeToString(w.Bytes())

	file, err := os.Create("encrMsg.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	err = ioutil.WriteFile("encrMsg.txt", []byte(encrMsg), fs.FileMode(os.O_RDONLY))
	if err != nil {
		return err
	}

	return nil

}

func decr() ([]byte, error) {
	/*
		encrEncMsg, err := ioutil.ReadFile("encrMsg.txt")
		if err != nil {
			return nil, err
		}
	*/
	encrMsg, err := base64.StdEncoding.DecodeString("wcDMA8O9phgYTw71AQwAVkNaDygR07ZnDHYoYKKvoe8drjYHrVSRZKK2xR1mVAc5VVUfQ++j6BvokZkwOD1qBrkFRbftjubcUj2fCACK1jrL+S9cLl08i+wImkVZ5Ot92vOjEunAddm2JetQ01U90NtmFW0YtZhnlFGID1GyGE7sKAsbnbHeAY21PN+6m/ID90rgpmS6ASqQh4in74ax50LUKHjYXC2YJ3prIZsCOCJtTq6JALypbJed0+gHopkl+x8Z3PL1ijpZ3am3UFalupQQw9MJ5wS89l+Zy35TV/LxUQBpmoW6ESxsV0ZbORddaycjk9KuLRYuqJa78uTi714PPhXSgtxjqBxtUHW1lNptGFQDC7IBrl0J3N/Bwer2QXjMnTg3X+/yqnMoiku5TPIs6F9mPjNxv9NOmgMT8hOo7uzcURI8nbRU+UTTFJZimi9SGvI58r+j2wDMxjuFEF8yKdG8zk5AT+aVNT3NfuZhzS42hO3luOYLOP30ftqwSMW59jFK/FV4CGMAT2fe0g==")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, bytes.NewReader(encrMsg))
	if err != nil {
		return nil, err
	}

	f, err := os.Open(seckey)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	el, err := openpgp.ReadArmoredKeyRing(f)
	if err != nil {
		return nil, err
	}

	e := el[0]

	err = e.PrivateKey.Decrypt([]byte(passkey))
	if err != nil {
		return nil, err
	}

	for _, subkey := range e.Subkeys {
		err = subkey.PrivateKey.Decrypt([]byte(passkey))
		if err != nil {
			return nil, err
		}
	}

	md, err := openpgp.ReadMessage(&buf, el, nil, nil)
	if err != nil {
		fmt.Println("err reading openpgp readmsg : ", err)
		return nil, err
	}

	decrMsg, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}

	return decrMsg, nil
}
