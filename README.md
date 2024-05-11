# git-ignores

Git plugin that generates a .gitignore for your project based on Githubs gitignore templates.

## Usage

The plugin provides a new git subcommand `ignores`. which takes the following options.

- `--template [NAME] ` - The name of the gitignore template from 
  [this repo](https://github.com/github/gitignore) (_i.e_ `Python` or `Javascript`)

- `--force` - Replace the .gitignore file by force with the new template.

- `--output [PATH]` - Write `.gitignore` to a custom path

> You can also run `git-ignores --help` to view the help message. Note that `git ignores --help` returns an error as git tries to load a man page when --help is called. A man-page will be shipped in a future update.

### Example

```
$ git ignores -t Python --force
```

## Installation

*Requires Go 1.22*

``` shell
$ make
$ sudo make install
```
