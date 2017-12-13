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
mustrun build/env.sh go get -v github.com/ethereum/go-ethereum
echo "Compiling epoch tool..."
mustrun build/env.sh go build -o epoch cmd/epoch/main.go
