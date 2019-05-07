module github.com/seamounts/essh

go 1.12

require (
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/manifoldco/promptui v0.3.2
	github.com/nicksnyder/go-i18n v1.10.0 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20180810215634-df19058c872c // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/crypto v0.0.0-20190426145343-a29dc8fdc734 => github.com/golang/crypto v0.0.0-20190426145343-a29dc8fdc734
)

replace google.golang.org/grpc v0.0.0-20190429225008-d5973a91700d => github.com/grpc/grpc-go v0.0.0-20190429225008-d5973a91700d

replace cloud.google.com/go v0.26.0 => github.com/googleapis/google-cloud-go v0.26.0

replace google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.1.0

replace golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be

replace google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8

replace golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3

replace (
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/sys v0.0.0-20190412213103-97732733099d => github.com/golang/sys v0.0.0-20190412213103-97732733099d
)

replace (
	golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
)

replace golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f

replace golang.org/x/tools v0.0.0-20190311212946-11955173bddd => github.com/golang/tools v0.0.0-20190311212946-11955173bddd

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
