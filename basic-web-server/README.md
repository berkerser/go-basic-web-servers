# Basic Web Server
Server returns simple Hello World! message. It will contain basic server setups to understand logic of go web servers.

Currently server runs at https protocol. Since it runs at https protocol by default it runs at http/2 since Go 1.6.
For more information please check [here](https://golang.orgc/pkg/net/http/).

https protocol requires a security certificate to run. We can use [tls](https://en.wikipedia.org/wiki/Transport_Layer_Security) (transport layer security) certificate for that. Go offers a small code that can create self-signed tls certificate.

### Note
> Self signed certificate will give a warning when you ran it from browsers.Because browsers will not know anything about you own certificate. To publish https servers to public you need a certificate that browser will know about it. For that you need to buy a certificate. So you need to research it 😅.

## Creating Self-Signed Certificate
We can use small code google provided to create basic self-signed certificate for developing purposes. That code is inside of "crypto/tls" package and it's called generate_cert.go. This code will create key.pem and cert.pem files in your terminal's current location. Use following command to create certification files:

bash
```bash
go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```

cmd
```bat
go run %GOROOT%\\src\\crypto\\tls\\generate_cert.go --host localhost
```

### Note
> GOROOT is an enviroment variables of your operating system. It stores currently installed go location. If GOROOT enviroment variable isn't available please enter installed location of go.

> Also --host parameter is required. It needs a value as hostname or and IP adress. So tested writing localhost.