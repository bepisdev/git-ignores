# git-ignores

`git-ignores` is a Git plugin designed to streamline the creation of `.gitignore` files for your projects. It leverages GitHub's extensive collection of `.gitignore` templates to generate tailored ignore rules based on your project's language or framework.

## Usage

To use `git-ignores`, follow these simple steps:

1. **Install the Plugin**:
   Ensure you have Go 1.22 or later installed, then run:
   ```shell
   $ make
   $ sudo make install
   ```

2. **Generate `.gitignore`**:
   Run the `ignores` subcommand with the desired options:
   ```shell
   $ git ignores --template [TEMPLATE_NAME] [--force] [--output PATH]
   ```
   - `--template [TEMPLATE_NAME]`: Specifies the name of the gitignore template from the [GitHub repository](https://github.com/github/gitignore). For example, `Python` or `Javascript`.
   - `--force`: (Optional) Forces the replacement of the existing `.gitignore` file with the new template.
   - `--output [PATH]`: (Optional) Specifies a custom path to write the `.gitignore` file.

3. **Example**:
   ```shell
   $ git ignores --template Python --force
   ```

4. **View Help**:
   To view the help message, run:
   ```shell
   $ git-ignores --help
   ```

## Future Enhancements

- **Man-Page Integration**:
  Currently, calling `git ignores --help` returns an error as Git tries to load a man page. A future update will include a dedicated man-page for seamless help access.

## Contributing

Contributions to `git-ignores` are welcome! Whether it's bug fixes, feature enhancements, or documentation improvements, feel free to submit pull requests.

---
