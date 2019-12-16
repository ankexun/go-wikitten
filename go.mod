module go-wikitten

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/net v0.0.0-20181220203305-927f97764cc3 => github.com/golang/net v0.0.0-20181220203305-927f97764cc3
	golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190613124609-5ed2794edfdc
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190613204242-ed0dc450797f
	golang.ort/x/net => github.com/golang/net v0.0.0-20190613194153-d28f0bde5980
	gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday v2.0.0+incompatible
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v2.1.0+incompatible
)

require (
	github.com/davidrjenni/reftools v0.0.0-20190411195930-981bbac422f8
	github.com/dchest/captcha v0.0.0-20170622155422-6a29415a8364
	github.com/fatih/motion v0.0.0-20190527122956-41470362fad4
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.46.0
	github.com/gorilla/context v1.1.1
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.1.3
	github.com/keegancsmith/rpc v1.1.0
	github.com/kisielk/errcheck v1.2.0
	github.com/kisielk/gotool v1.0.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/mattn/go-isatty v0.0.8
	github.com/mdempsky/gocode v0.0.0-20190203001940-7fb65232883f
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/mozillazg/go-pinyin v0.15.0
	github.com/pkg/errors v0.8.1
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518
	github.com/stamblerre/gocode v0.0.0-20190327203809-810592086997
	github.com/stretchr/testify v1.3.0
	go4.org v0.0.0-20190313082347-94abd6928b1d
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sys v0.0.0-20190412213103-97732733099d
	golang.org/x/text v0.3.0
	golang.org/x/tools v0.0.0-20190408220357-e5b8258f4918
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/ini.v1 v1.46.0
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
)
