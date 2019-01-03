module e.coding.net/aliens/aliensboot_toolkit

replace (
	cloud.google.com/go v0.26.0 => github.com/GoogleCloudPlatform/gcloud-golang v0.32.0
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/lint v0.0.0-20180702182130-06c8688daad7 => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20181108082009-03003ca0c849
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20181106182150-f42d05182288
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181107165924-66b7b1311ac8
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52 => github.com/golang/tools v0.0.0-20181111003725-6d71ab8aade0
	google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.3.0
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20181109154231-b5d43981345b
	google.golang.org/grpc v1.16.0 => github.com/grpc/grpc-go v1.16.0
)

require (
	github.com/KylinHe/aliensboot v0.0.1
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/gogo/protobuf v1.2.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/klauspost/reedsolomon v1.8.0 // indirect
	github.com/pkg/errors v0.8.0
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20181023030647-4e92f724b73b // indirect
	github.com/tjfoc/gmsm v1.0.1 // indirect
	github.com/urfave/cli v1.20.0
	github.com/xtaci/kcp-go v5.0.2+incompatible
	github.com/xtaci/smux v1.1.0
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
)
