go get github.com/mitchellh/go-homedir
go get github.com/spf13/cobra
go get github.com/spf13/viper
go get github.com/olekukonko/tablewriter
go get github.com/inconshreveable/mousetrap


rm -rf release
mkdir release

OS_LIST=(windows darwin linux)

for os in "${OS_LIST[@]}"
do
  echo "Building for $os"
  mkdir release/$os
  env GOOS=$os GOARCH=amd64 go build -ldflags="-s -w" -o release/$os/octocli
  echo "Generating archive release/octocli-$os-amd64.zip"
  if [ $os = "windows" ]
  then
    mv release/$os/octocli release/$os/octocli.exe
    tar -czf "release/octocli-$os-amd64.zip" -C release/$os/ .
  else
    tar -cf "release/octocli-$os-amd64.zip" -C release/$os/ .
  fi
  rm -rf release/$os
done
