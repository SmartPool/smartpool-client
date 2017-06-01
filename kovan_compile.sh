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
mustrun build/env.sh go get -v github.com/mitchellh/go-homedir
mustrun build/env.sh go get -v golang.org/x/net/context
echo "Compiling SmartPool client..."
mustrun build/env.sh go build -ldflags -s -o kovan cmd/kovan/main.go
echo "Done. You can run SmartPool by ./kovan --help"
