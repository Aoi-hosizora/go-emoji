# go-emoji

[![Go Report Card](https://goreportcard.com/badge/github.com/Aoi-hosizora/go-emoji)](https://goreportcard.com/report/github.com/Aoi-hosizora/go-emoji)
[![License](http://img.shields.io/badge/license-mit-blue.svg)](./LICENSE)
[![Release](https://img.shields.io/github/v/release/Aoi-hosizora/go-emoji)](https://github.com/Aoi-hosizora/go-emoji/releases)

+ An emoji characters library for golang. This library includes all emojis from [Emoji List, v13.1](https://unicode.org/emoji/charts/emoji-list.html) and [Full Emoji Modifier Sequences, v13.1](https://unicode.org/emoji/charts/full-emoji-modifiers.html).

### Usage

```go
package main

import (
	"fmt"
	emoji "github.com/Aoi-hosizora/go-emoji"
)

func main() {
	fmt.Println(emoji.GrinningFace)
}
```

### Generate

```bash
# Do generate for the first time.
sh generate.sh

# Do generate after emoji.go has been generated.
go generate
```

### Reference

+ [Unicode: Emoji List, v13.1](https://unicode.org/emoji/charts/emoji-list.html)
+ [Unicode: Full Emoji Modifier Sequences, v13.1](https://unicode.org/emoji/charts/full-emoji-modifiers.html)
+ [Emojipedia: Emoji Version 13.1](https://emojipedia.org/emoji-13.1/)
+ [kyokomi/emoji](https://github.com/kyokomi/emoji)
