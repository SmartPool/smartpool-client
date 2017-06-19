mustrun() {
   "$@"
   if [ $? != 0 ]; then
      printf "Error when executing command: '$*'\n"
      exit $ERROR_CODE
   fi
}

export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64

echo "Install dependencies..."
mustrun build/env.sh go get -v golang.org/x/crypto/ssh/terminal
mustrun build/env.sh go get -v gopkg.in/urfave/cli.v1
mustrun build/env.sh go get -v github.com/bmizerany/pat
mustrun build/env.sh go get -v github.com/mitchellh/go-homedir
mustrun build/env.sh go get -v golang.org/x/net/context
mustrun build/env.sh go get -v github.com/gorilla/websocket
mustrun build/env.sh go get -v github.com/ethereum/go-ethereum
echo "Compiling SmartPool client for Windows..."
mustrun build/env.sh go build -o smartpool.exe cmd/ropsten/ropsten.go
echo "Done. You can run SmartPool by ./smartpool --help"
