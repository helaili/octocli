# octocli
A CLI for GitHub


### Usage:

List every org on a GitHub Enterprise server:

```
octocli org list -s octodemo.com
```

List every org a given user belong to:

```
octocli org list --user monalisaoctocat
```

More information with `octocli --help`

```
Usage:
  octocli [command]

Available Commands:
  help        Help about any command
  org         Base for various GitHub Organizations related commands

Flags:
      --config string   config file (default is $HOME/.octocli.yaml)
  -h, --help            help for octocli
  -s, --server string   Hostname of the GitHub Enterprise server. Using github.com if omitted (default "github.com")
  -t, --token string    personal authentication token to use. Required when environement variable GITHUB_AUTH_TOKEN is not set

Use "octocli [command] --help" for more information about a command.
```

### Contribute


Dependencies:

```
go get github.com/mitchellh/go-homedir
go get github.com/spf13/cobra
go get github.com/spf13/viper
go get github.com/olekukonko/tablewriter
```

Run tests:

```
go test ./...
```


Add a new `dosmthg` command

```
cobra add dosmthg
```

Add a sub command `mysubdosmthg` in `dosmthg`


```
cobra add mysubdosmthg -p 'dosmthgCmd'
```

More info about this [here](https://github.com/spf13/cobra/blob/master/cobra/README.md) and [here](https://github.com/spf13/cobra/blob/master/README.md)
