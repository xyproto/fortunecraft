# FortuneCraft

Output fortunes by using Ollama and a user-defined text generation model.

`fortunecraft` has a wider selection of fortunes than the good old `fortune` program.

One of the goals for this utility is to be one of the first Arch Linux packages that depends on Ollama and `ollama-environment`, to help open up this path for other packages and applications.

### Requirements

Requires an Ollama server to be up and running, and the selected model to be able to run / be evaluated, in terms of CPU and memory. Using a GPU is optional.

* `OLLAMA_HOST` can be used to set a different host address than `localhost:11434` for the Ollama server.

### Installation

    go install github.com/xyproto/fortunecraft@latest

Then place `~/go/bin` in the `PATH`, or install `~/go/bin/fortunecraft` somewhere else, if you want.

* If you want to use a different model than the default one, install [`llm-manager`](https://github.com/xyproto/llm-manager) and select a model for the text generation task (can also be configured by editing `~/.config/llm-manager/llm.conf` or `/etc/llm.conf`).

### Example output

```
./fortunecraft
The best way to predict the future is to invent it.

❯ ./fortunecraft
The future is a tapestry woven with threads of chance and choice.

❯ fortunecraft -giz
Yeet your fears to the wind 🌬️🚀 You're boutta slay.

❯ ./fortunecraft -eNb
Resistance is futile!
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

* Version: 1.8.3
* License: Apache 2
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
