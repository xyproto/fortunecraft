# fortune9000

Output fortunes by using Ollama and the [`tinydolphin`](https://ollama.com/library/tinydolphin) model.

`fortune9000` has a wider variety and more innovative fortunes than the good old `fortune` program, but the quality is all over the place.

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
When you feel lost, just look up at the stars and remember that you're alive.

❯ fortune9000
Why don't you give me your name? Because then I would know what to say!

❯ fortune9000
Every day is a new beginning, so don't forget to make one.

❯ fortune9000
Always be prepared to face the inevitable, as you will never face it alone.

❯ fortune9000
Why did the chicken cross the road? Because the answer was 'Option 2', and I wanted to try new things.
```

### General info

* Version: 1.1.0
* License: Apache 2
* Author: Alexander F. Rødseth
