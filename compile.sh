echo "Install dependencies..."
go get -v github.com/ethereum/go-ethereum
go get -v golang.org/x/crypto/ssh/terminal
go get -v gopkg.in/urfave/cli.v1
echo "Compiling SmartPool client..."
go build -o smartpool cmd/kovan/main.go
echo "Done. You can run SmartPool by ./smartpool --help"
