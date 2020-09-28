# b64

`b64` is a small tool that fixes the shortcoming of the default `base64` command line on MacOS:

1. it accepts string as parameter (as well as pipe)
2. it allows Base64 URL encoding, not only standard Base64, using the `-u` flag
3. it allows to disable padding from input / output, using the `-p=false` flag

## Installation

You can either download the latest version from the [GitHub release](https://github.com/kdisneur/b64/releases) page or use the Homebrew formula:

```
brew install kdisneur/formulae/b64
```
