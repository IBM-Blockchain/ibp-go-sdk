module github.com/IBM-Blockchain/ibp-go-sdk

go 1.14

require (
	github.com/IBM/go-sdk-core/v4 v4.0.5
	github.com/Knetic/govaluate v3.0.0+incompatible // indirect
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cloudflare/cfssl v1.5.0 // indirect
	github.com/go-openapi/strfmt v0.19.5
	github.com/google/certificate-transparency-go v1.1.1 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grantae/certinfo v0.0.0-20170412194111-59d56a35515b // indirect
	github.com/hyperledger/fabric v2.1.1+incompatible
	github.com/hyperledger/fabric-amcl v0.0.0-20200424173818-327c9e2cf77a // indirect
	github.com/hyperledger/fabric-ca v1.4.9
	github.com/hyperledger/fabric-lib-go v1.0.0 // indirect
	github.com/jmhodges/clock v0.0.0-20160418191101-880ee4c33548 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/kisielk/sqlstruct v0.0.0-20201105191214-5f3e10d3ab46 // indirect
	github.com/lib/pq v1.9.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.5 // indirect
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/prometheus/common v0.15.0 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/ldap.v2 v2.5.1 // indirect
)

replace (
	github.com/cloudflare/cfssl => github.com/cloudflare/cfssl v0.0.0-20190409034051-768cd563887f
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0 // Needed for fabric-ca
	github.com/gorilla/mux => github.com/gorilla/mux v1.7.3 // Needed for fabric-ca
	github.com/hyperledger/fabric => github.com/hyperledger/fabric v0.0.0-20191027202024-115c7a2205a6
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.0 // Needed for fabric-ca
)
