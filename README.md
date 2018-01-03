##
- ch1: use openssh to generate key and cert
- ch2: use go tool to generate key and cert
- ch3: server uses tls and clien uses InsecureSkipVerify
- ch4: server and client use cert
- ch5: server uses CA and client doesn't work
- ch6: server and client use CA
- ch7: http2
- ch8: h2c
- ch9: autocert


## 参考文档

1. https://github.com/golang/net/tree/master/http2/h2demo
1. https://ericchiang.github.io/post/go-tls/
1. https://gist.github.com/denji/12b3a568f092ab951456
1. https://github.com/denji/golang-tls
1. https://github.com/jcbsmpsn/golang-https-example
1. https://hacked.work/blog/writing-simple-http2-server-go/
1. https://www.cnblogs.com/E7868A/archive/2012/11/16/2772240.html
1. https://support.ssl.com/Knowledgebase/Article/View/19/0/der-vs-crt-vs-cer-vs-pem-certificates-and-how-to-convert-them
1. https://stackoverflow.com/questions/38582937/how-to-create-a-tls-client-with-ca-certificates-in-go
1. http://www.levigross.com/2015/11/21/mutual-tls-authentication-in-go/
1. https://fale.io/blog/2017/06/05/create-a-pki-in-golang/
1. https://godoc.org/golang.org/x/crypto/acme/autocert#example-NewListener 
1. https://github.com/golang/go/issues/14141
1. https://github.com/golang/go/issues/14391
1. https://www.jianshu.com/p/ff16b0308e7c