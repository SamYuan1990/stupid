package infra_test

import (
	"io/ioutil"
	"os"

	"github.com/guoger/stupid/pkg/infra"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Context("config", func() {
		It("successful load", func() {
			var configText = `
org1peer0: &org1peer0
  addr: peer0.org1.example.com:7051
org2peer0: &org2peer0
  addr: peer0.org2.example.com:7051
org0orderer0: &org0orderer0
  addr: orderer.example.com:7050

endorsers:
  - *org1peer0
  - *org2peer0
committer: *org2peer0
orderer: *org0orderer0

channel: mychannel
chaincode: mycc
args:
  - invoke
  - a
  - b
  - 1
mspid: Org1MSP
private_key: ../../test/private.key
sign_cert: ../../test/sign.cert
num_of_conn: 20
client_per_conn: 40`

			f, _ := ioutil.TempFile("", "config-*.yaml")
			defer os.Remove(f.Name())
			f.WriteString(configText)
			f.Close()

			c := infra.LoadConfig(f.Name())

			Expect(c).To(Equal(infra.Config{
				Endorsers: []infra.Node{
					{Addr: "peer0.org1.example.com:7051"},
					{Addr: "peer0.org2.example.com:7051"},
				},
				Committer:     infra.Node{Addr: "peer0.org2.example.com:7051"},
				Orderer:       infra.Node{Addr: "orderer.example.com:7050"},
				Channel:       "mychannel",
				Chaincode:     "mycc",
				Version:       "",
				Args:          []string{"invoke", "a", "b", "1"},
				MSPID:         "Org1MSP",
				PrivateKey:    "../../test/private.key",
				SignCert:      "../../test/sign.cert",
				NumOfConn:     20,
				ClientPerConn: 40,
			}))
		})
	})
})
