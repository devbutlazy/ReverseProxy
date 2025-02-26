# ⚕️ ReverseProxy 

A simple reverse proxy in GoLang to forward HTTP/HTTPS requests to backend services. Aimed at practicing Go concurrency and networking, with basic routing, logging, and error handling features.

## Tasks
- [x] Port availability check
- [x] Startup TCP on port (and log connections)
- [ ] Startup a HTTP/HTTPS reverse-proxy
- [ ] Log the len(conections) & traffic usage
- [ ] Automatic certs configuration via openssl (for HTTPS)
- [ ] Automatic proxy link (.pac)

## Installation (non-docker)

`1.` Download and install Go from [go.dev](https://go.dev/)  
`2.` Download and install git from [git-scm.com](ws.1/Git-2.44.0-64-bit.exe)  
`3.` Clone this repository
```
git clone https://github.com/devbutlazy/ReverseProxy
```
`4.` Run the program (or build it)
```
go run ./main.go // run it

go build ./main.go // build it to main.exe
```


## Installation (Docker):
`-` Not supported yet :)

## Feel free to open [issues](https://github.com/devbutlazy/ReverseProxy/issues) or [pull requests](https://github.com/devbutlazy/ReverseProxy/pulls) if you have encountered any kind of problems.

### (c) LazyWeb License: MIT-LICENSE