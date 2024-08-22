# fortune9000

Output fortunes by using Ollama and the [`gemma2:2b`](https://ollama.com/library/gemma2) model.

`fortune9000` has a different selection of fortunes from the good old `fortune` program.

### Requirements

Requires a working Ollama server.

* The `NO_COLOR` environment variable is respected.
* `OLLAMA_HOST` can be used to set a different host address than `localhost:11434` for the Ollama server.

### Installation

    go install github.com/xyproto/fortune9000@latest

Then place `~/go/bin` in the `PATH`, or install `~/go/bin/fortune9000` somewhere else, if you want.

### Example output

```
❯ fortune9000
A snail in an elevator is a very fast-moving slowpoke.

❯ fortune9000
To err is human; to forgive, divine.

❯ fortune9000
Why don't scientists trust atoms? Because they make up everything!

❯ fortune9000
The cat sat in a sunbeam.

❯ fortune9000
Re having trouble getting your cat to wear clothes, try using a hairdryer and setting it to

❯ fortune9000
What's red and bad for your teeth? A brick.

❯ fortune9000
Why don't they play poker in the jungle?  Too many cheetahs!
```

### General info

* Version: 1.1.3
* License: Apache 2
* Author: Alexander F. Rødseth
