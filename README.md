[![Build Status](https://travis-ci.org/helaili/octocli.svg?branch=master)](https://travis-ci.org/helaili/octocli)


# octocli
A CLI for GitHub

Current commands include:

- Orgs
  - List every org a user belongs to
  - List every orgs (GHE)
- Teams
  - List every teams of an org
  - Create a team
  - Add a member to a team
  - List every member of a team
  - List every repository of a team


### Usage:

List every org on a GitHub Enterprise server:

```
octocli org list -s octodemo.com
```

List every org user `monalisa` belongs to:

```
octocli org list --login monalisa
```

Add users `monalisa` and `hubot` as maintainers to the `developers` team within the `OctoCheese` organization

```
octocli team members add -o OctoCheese -t developers -r maintainer monalisa hubot
```

More information with `octocli --help`

```
Usage:
  octocli [command]

Available Commands:
  help        Help about any command
  org         Base for various GitHub Organizations related commands
  team        Base for various team related commands

Flags:
      --config string   config file (default is $HOME/.octocli.yaml)
  -h, --help            help for octocli
  -s, --server string   Hostname of the GitHub Enterprise server. Using github.com if omitted (default "github.com")
  -k, --token string    personal authentication token to use. Required when environement variable GITHUB_AUTH_TOKEN is not set

Use "octocli [command] --help" for more information about a command.
```

### Contribute

Have a look at `release.sh`

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
