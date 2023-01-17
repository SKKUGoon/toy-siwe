# Proof of Concept: Login with SIWE

## SIWE

<b>SIWE</b> stands for Sign in with Ethereum. 

## Start go server

`go run .` will automatically initiate gonic-gin server in debug status. 
Concept webserver servers at localhost:3000 

## Check if it works

```console
curl 'http://localhost:3000/api/login/verify' \
  -H 'Content-Type: application/json' \
  --data-raw '{"message":"localhost:8080 wants you to sign in with your Ethereum account:\n0x9D85ca56217D2bb651b00f15e694EB7E713637D4\n\nSign in with Ethereum to the app.\n\nURI: http://localhost:8080\nVersion: 1\nChain ID: 1\nNonce: spAsCWHwxsQzLcMzi\nIssued At: 2022-01-29T03:22:26.716Z","signature":"0xe117ad63b517e7b6823e472bf42691c28a4663801c6ad37f7249a1fe56aa54b35bfce93b1e9fa82da7d55bbf0d75ca497843b0702b9dfb7ca9d9c6edb25574c51c"}'
```