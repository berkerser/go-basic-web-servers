# Basic Web Server
Server returns simple Hello World! message. It will contain basic server setups to understand logic of go web servers.

Currently server runs at https protocol. Since it runs at https protocol by default it runs at http/2 since Go 1.6.
For more information please check [here](https://golang.org/pkg/net/http/).

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

## Enviroment Variables
### Warning
> This section is for people who don't know about enviroment variables. If you know about how to set enviroment variables you don't need to look for it.

Enviroment variables (in short env) are dynamic values that will affect processes behaviours. This variables can contain values we want to hide or some shortcut for programs. In this project i try to hide values that i think it should be hide or practicising env values in the code. For example:

go
```go
log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", os.Getenv(port)), certification, key, nil))
```

in this particular code we don't need to hide the port if we will publish this web server in default port. But for example if we are connecting a database with write rights, we should hide how to connect to database. This could lead to data manipulations from unknown sources.

In following subsections i will explain way to assign enviroment variables.

### Note
> Assigning new env, changing currently used env will not change env of currently running processes. That's why these values are dynamic.

> Also assigning new env, changing currently used env will not change env of currently running terminals. You need to open new terminals to affect changes.

> If you are doing assignments in terminal you won't need to open new terminal. Also if you didn't set it as global env or user level env, this values will not be saved so you can't access this values from another terminal. (Note: I could be wrong about this, so it could be need to be changed.).

### 1. Visual Studio Code

If you're using integrated terminal that Visual Studio Code support, you can set env values inside of vscode. You can use instruction below to assign new env values:

1. Open command pallete (View>Command Pallete)
2. Type settings.json
3. Select Preferences: Open Settings (JSON)
4. Please enter your env variables to following json section.

#### Windows
```json
{
    "terminal.integrated.env.windows": {
        "[YOUR ENV VARIABLE]": "[YOUR ENV VALUE]"
    }
}
```

#### MacOS
```json
{
    "terminal.integrated.env.osx": {
        "[YOUR ENV VARIABLE]": "[YOUR ENV VALUE]"
    }
}
```