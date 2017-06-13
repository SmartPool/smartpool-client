mustrun() {
   "$@"
   if [ $? != 0 ]; then
      printf "Error when executing command: '$*'\n"
      exit $ERROR_CODE
   fi
}

echo "Install dependencies..."
mustrun build/env.sh go get -v golang.org/x/crypto/ssh/terminal
mustrun build/env.sh go get -v gopkg.in/urfave/cli.v1
mustrun build/env.sh go get -v github.com/bmizerany/pat
mustrun build/env.sh go get -v github.com/mitchellh/go-homedir
mustrun build/env.sh go get -v golang.org/x/net/context
mustrun build/env.sh go get -v github.com/gorilla/websocket
mustrun build/env.sh go get -v github.com/ethereum/go-ethereum
echo "Compiling SmartPool experiment..."
mustrun build/env.sh go build -ldflags -s -o experiment cmd/experiments/main.go
