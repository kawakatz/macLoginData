# macLoginData🔑
<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/kawakatz/grepfiles/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://twitter.com/kawakatz"><img src="https://img.shields.io/twitter/follow/kawakatz.svg?logo=twitter"></a>
</p>

macLoginData decrypt login data stored in macOS browsers for pentesters.<br>
This tool is intended to be used with C2.

# Installation
```sh
➜  ~ go install -v github.com/kawakatz/macLoginData/cmd/macLoginData@latest
```

## Usage
### Google Chrome, Microsoft Edge
- login-keychain password is required to decrypt login-keychain

```sh
# extract Chrome Safe Storage value
➜  ~ ./chainbreaker.py --dump-all login.keychain-db --password=<login-keychain password>
➜  ~ macLoginData Chrome ~/Library/Application\ Support/Google/Chrome/Default/Login\ Data <Chrome Safe Storage>
```

#### Option
It is also possible to decrypt Login Data retrieved from Windows.
In that case, use <a href="https://github.com/crypt0p3g/bof-collection/tree/main/ChromiumKeyDump">ChromiumKeyDump</a> to retrieve a masterkey.<br>
```sh
➜  ~ macLoginData -win Chrome Cookies <masterkey>
```

## References
- https://github.com/moonD4rk/HackBrowserData (MIT License)<br>
    decryption logic for FIrefox, Google Chrome, Microsoft Edge, etc...
