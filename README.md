# Fortuna

Output fortunes by using Ollama and the [`gemma2:2b`](https://ollama.com/library/gemma2) model.

`fortuna` has a different selection of fortunes from the good old `fortune` program.

### Requirements

Requires a working Ollama server.

* The `NO_COLOR` environment variable is respected.
* `OLLAMA_HOST` can be used to set a different host address than `localhost:11434` for the Ollama server.

### Installation

    go install github.com/xyproto/fortuna@latest

Then place `~/go/bin` in the `PATH`, or install `~/go/bin/fortuna` somewhere else, if you want.

### Example output

```
❯ fortuna
A snail in an elevator is a very fast-moving slowpoke.

❯ fortuna
To err is human; to forgive, divine.

❯ fortuna
Why don't scientists trust atoms? Because they make up everything!

❯ fortuna
The cat sat in a sunbeam.

❯ fortuna
What's red and bad for your teeth? A brick.

❯ fortuna
Why don't they play poker in the jungle? Too many cheetahs!

❯ fortuna
I once knew a guy who thought he was a penguin.

❯ fortuna
My cat sits on my keyboard and types.

❯ fortuna
I'm pretty sure it's just running out of RAM.

❯ fortuna
Shell shocked and amazed by an accidental catnip-powered reboot.

❯ fortuna
Your systemd service is ready to rock and roll.

❯ fortuna
The kernel's got too many flags, man.

❯ fortuna
50% chance of mild system crashes, 25% chance of cat memes, and 25% chance you forgot what you were doing.

❯ fortuna
The path to enlightenment is paved with broken shells and spilled coffee.

❯ fortuna
Time to start eating your own cat memes.
```

### General info

* Version: 1.3.1
* License: Apache 2
* Author: Alexander F. Rødseth
