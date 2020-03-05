module github.com/ontio/ontology

go 1.12

require (
	github.com/JohnCGriffin/overflow v0.0.0-20170615021017-4d914c927216
	github.com/ethereum/go-ethereum v1.9.6
	github.com/go-interpreter/wagon v0.6.0
	github.com/gorilla/websocket v1.4.1
	github.com/gosuri/uiprogress v0.0.1
	github.com/hashicorp/golang-lru v0.5.3
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/itchyny/base58-go v0.1.0
	github.com/ontio/ontology-crypto v1.0.7
	github.com/ontio/ontology-eventbus v0.9.1
	github.com/ontio/ontology-go-sdk v1.11.0
	github.com/ontio/wagon v0.4.1
	github.com/pborman/uuid v1.2.0
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/urfave/cli v1.22.1
	github.com/valyala/bytebufferpool v1.0.0
	golang.org/x/crypto v0.0.0-20191219195013-becbf705a915
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191029031824-8986dd9e96cf
	golang.org/x/net => github.com/golang/net v0.0.0-20191028085509-fe3aa8a45271
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190412213103-97732733099d
	golang.org/x/text => github.com/golang/text v0.3.0
)
