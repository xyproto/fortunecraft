# FortuneCraft

Output fortunes by using Ollama and the [`gemma2:2b`](https://ollama.com/library/gemma2) model.

`fortunecraft` has a very different selection of fortunes compared to the good old `fortune` program.

One of the goals for this utility is to be a proof of concept for an Arch Linux package that depends on a model that in turn depends on Ollama.

### Requirements

Requires an Ollama server to be up and running, and the `gemma2:2b` model to be able to be evaluated, in terms of CPU and memory. Using a GPU is optional.

* `OLLAMA_HOST` can be used to set a different host address than `localhost:11434` for the Ollama server.

### Installation

    go install github.com/xyproto/fortunecraft@latest

Then place `~/go/bin` in the `PATH`, or install `~/go/bin/fortunecraft` somewhere else, if you want.

### Example output

```
❯ fortunecraft
A snail in an elevator is a very fast-moving slowpoke.

❯ fortunecraft
To err is human; to forgive, divine.

❯ fortunecraft
Why don't scientists trust atoms? Because they make up everything!

❯ fortunecraft
The cat sat in a sunbeam.

❯ fortunecraft
What's red and bad for your teeth? A brick.

❯ fortunecraft
Why don't they play poker in the jungle? Too many cheetahs!

❯ fortunecraft
I once knew a guy who thought he was a penguin.

❯ fortunecraft
My cat sits on my keyboard and types.

❯ fortunecraft
I'm pretty sure it's just running out of RAM.

❯ fortunecraft
Shell shocked and amazed by an accidental catnip-powered reboot.

❯ fortunecraft
Your systemd service is ready to rock and roll.

❯ fortunecraft
The kernel's got too many flags, man.

❯ fortunecraft
50% chance of mild system crashes, 25% chance of cat memes, and 25% chance you forgot what you were doing.

❯ fortunecraft
The path to enlightenment is paved with broken shells and spilled coffee.

❯ fortunecraft
Time to start eating your own cat memes.
```

### Flags

```
Available Flags:
-a, --absurd           Be absurd
-B, --boomer           Boomer style
-b, --borg             Make it about Borg
-c, --cats             Make it about cats
-o, --computer         Make it about computers
-D, --delusional       Be delusional
-d, --dogs             Make it about dogs
-e, --evil             Be evil
-f, --fantasy          Make it about fantasy
-z, --genz             Make it more Gen Z
-g, --good             Be good
-N, --inappropriate    Be inappropriate
-i, --inspire          Be inspirational
-t, --international    Be international
-I, --ironic           Be ironic
-k, --keyword string   Specify a custom keyword
-1, --leet             1337 style
-l, --logical          Make it more logical
-n, --ninja            Make it about ninjas
-O, --old              Use language from 100 years ago
-p, --pirate           Write like a pirate
-P, --political        Be political
-y, --pony             Make it about ponies
-A, --praise           Fill it with praise
-r, --robot            Make it about robots
-R, --romantic         Add a romantic touch to the fortune
-s, --sarcastic        Generate a sarcastic fortune
-C, --scifi            Make it sci-fi related
-u, --user             Make it about the current user
-V, --version          Output the current version
-w, --weird            Be weird

Examples:
fortunecraft -giz      - Generate good inspirational GenZ fortunes
fortunecraft -eNb      - Generate evil inappropriate Borg fortunes
fortunecraft -gaCy     - Generate good absurd sci-fi fortunes about ponies
fortunecraft -iep      - Generate inspirational evil pirate fortunes
fortunecraft -sPB      - Generate sarcastic political boomer fortunes
fortunecraft -I -k AI  - Generate ironic fortunes about AI
```

### General info

* Version: 1.7.0
* License: Apache 2
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
