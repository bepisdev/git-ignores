# git-ignores

Git plugin that generates a .gitignore for your project based on Githubs gitignore templates.

## Usage

The plugin provides a new git subcommand `ignores`. which takes the following options.

- `--template` - The name of the gitignore template from 
  [this repo](https://github.com/github/gitignore) (_i.e_ `Python` or `Javascript`)

- `--append` - Instead of failing or replacing content, `--append` will tell the 
script to simply append the gitignore entries instead of replacing the file wholesale.

- `--force` - Replace the .gitignore file by force with the new template.

> You can also run `git-ignores --help` to view the help message. Note that `git ignores --help` returns an error as git tries to load a man page when --help is called. A man-page will be shipped in a future update.

### Example

```
$ git ignores -t Python --force
```

## Installation

### Pip

Simply install the plugin via `pip`.

``` 
$ pip install git-ignores
```

### From source

To install the plugin from source you will need to have [Poetry](https://python-poetry.org/) installed.

1. clone the repo
  - `$ git clone https://github.com/joshburnsxyz/git-ignores`
2. cd into the repo 
  - `$ cd ./git-ignores` 
3. build wheel
  - `$ poetry build`
4. install wheel with pip
  - `$ pip install ./dist/*.whl`

## Contributing

If you wish to contribute to the project. Here a few things to note.

- Python version used in development: 3.11
- This project uses the [Poetry](https://python-poetry.org/) build tool.
