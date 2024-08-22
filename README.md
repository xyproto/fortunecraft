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
What's red and bad for your teeth? A brick.

❯ fortune9000
Why don't they play poker in the jungle? Too many cheetahs!

❯ fortune9000
I once knew a guy who thought he was a penguin.

❯ fortune9000
My cat sits on my keyboard and types.

❯ fortune9000
I'm pretty sure it's just running out of RAM.

❯ fortune9000
Shell shocked and amazed by an accidental catnip-powered reboot.

❯ fortune9000
Your systemd service is ready to rock and roll.

❯ fortune9000
The kernel's got too many flags, man.

❯ fortune9000
50% chance of mild system crashes, 25% chance of cat memes, and 25% chance you forgot what you were doing.

❯ fortune9000
The path to enlightenment is paved with broken shells and spilled coffee.

❯ fortune9000
Time to start eating your own cat memes.
```

### General info

* Version: 1.2.0
* License: Apache 2
* Author: Alexander F. Rødseth
