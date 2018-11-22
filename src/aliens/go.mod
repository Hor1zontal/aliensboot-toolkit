module aliens

replace (
	cloud.google.com/go v0.26.0 => github.com/GoogleCloudPlatform/gcloud-golang v0.32.0
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/lint v0.0.0-20180702182130-06c8688daad7 => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net v0.0.0-20180724234803-3673e40ba225 => github.com/golang/net v0.0.0-20181108082009-03003ca0c849
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => github.com/golang/net v0.0.0-20181108082009-03003ca0c849
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20181108082009-03003ca0c849
	golang.org/x/net v0.0.0-20181106065722-10aee1819953 => github.com/golang/net v0.0.0-20181108082009-03003ca0c849
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20181106182150-f42d05182288
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 => github.com/golang/sys v0.0.0-20181107165924-66b7b1311ac8
	golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33 => github.com/golang/sys v0.0.0-20181107165924-66b7b1311ac8
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20181107165924-66b7b1311ac8
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52 => github.com/golang/tools v0.0.0-20181111003725-6d71ab8aade0
	google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.3.0
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20181109154231-b5d43981345b
	google.golang.org/grpc v1.16.0 => github.com/grpc/grpc-go v1.16.0
)

require (
	github.com/DataDog/datadog-go v0.0.0-20180822151419-281ae9f2d895 // indirect
	github.com/GoogleCloudPlatform/gcloud-golang v0.32.0 // indirect
	github.com/Shopify/sarama v1.19.0
	github.com/Shopify/toxiproxy v2.1.3+incompatible // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/circonus-labs/circonus-gometrics v2.2.4+incompatible // indirect
	github.com/circonus-labs/circonusllhist v0.1.0 // indirect
	github.com/coreos/bbolt v1.3.0 // indirect
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/coreos/go-semver v0.2.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20181031085051-9002847aa142 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/gogo/protobuf v1.1.1
	github.com/golang/appengine v1.3.0 // indirect
	github.com/golang/groupcache v0.0.0-20181024230925-c65c006176ff // indirect
	github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3 // indirect
	github.com/golang/oauth2 v0.0.0-20181106182150-f42d05182288 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/golang/tools v0.0.0-20181111003725-6d71ab8aade0 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/btree v0.0.0-20180813153112-4030bb1f1f0c // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.5.1 // indirect
	github.com/grpc/grpc-go v1.16.0 // indirect
	github.com/hashicorp/consul v1.3.0
	github.com/hashicorp/go-cleanhttp v0.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v0.0.0-20150518234257-fa3f63826f7c // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.0.0-20180718195005-e651d75abec6 // indirect
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90 // indirect
	github.com/hashicorp/go-sockaddr v0.0.0-20180320115054-6d291a969b86 // indirect
	github.com/hashicorp/memberlist v0.1.0 // indirect
	github.com/hashicorp/serf v0.8.1 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/klauspost/reedsolomon v0.0.0-20180704173009-925cb01d6510 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/magicsea/behavior3go v0.0.0-20180426084336-88b2ae018949
	github.com/magicsea/ganet v0.0.0-20180425082344-2e8b357526bd
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/miekg/dns v1.0.15 // indirect
	github.com/mitchellh/go-homedir v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/name5566/leaf v0.0.0-20181103040206-1364c176dfbd
	github.com/onsi/gomega v1.4.2 // indirect
	github.com/pascaldekloe/goe v0.0.0-20180627143212-57f6aae5913c // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/errors v0.8.0
	github.com/prometheus/client_golang v0.9.1 // indirect
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910 // indirect
	github.com/prometheus/common v0.0.0-20181109100915-0b1957f9d949 // indirect
	github.com/prometheus/procfs v0.0.0-20181005140218-185b4288413d // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/sirupsen/logrus v1.2.0
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/tebeka/strftime v0.0.0-20140926081919-3f9c7761e312 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20181023030647-4e92f724b73b // indirect
	github.com/tjfoc/gmsm v1.0.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20171017195756-830351dc03c6 // indirect
	github.com/tv42/httpunix v0.0.0-20150427012821-b75d8614f926 // indirect
	github.com/ugorji/go/codec v0.0.0-20181022190402-e5e69e061d4f // indirect
	github.com/urfave/cli v1.20.0
	github.com/xiang90/probing v0.0.0-20160813154853-07dd2e8dfe18 // indirect
	github.com/xiaonanln/goworld v0.1.6
	github.com/xtaci/kcp-go v4.3.1+incompatible
	github.com/xtaci/smux v1.0.7
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1 // indirect
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/net v0.0.0-20181106065722-10aee1819953
	google.golang.org/grpc v1.16.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	gopkg.in/olivere/elastic.v5 v5.0.76
	gopkg.in/sohlich/elogrus.v2 v2.0.2
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1 // indirect
)
