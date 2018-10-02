HOME_DIR=$(cd "$(pwd)/.."; pwd)
export GOPATH=$HOME_DIR

#Build the code
go env
cd $HOME_DIR/src/

#Run FMT on the Code base
echo "Running fmt on code"
go fmt $(go list ./... )

#Run Lint on the Code base
echo "Running lint on code"
golint $(go list ./... )

#Run vet on the Code base
echo "Running vet on code"
go vet $(go list ./... )
