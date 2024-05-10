# Summary

Uses Google Gemini to explain and help you to debug some piece of code

## How to setup the project

There are two main components in this project:
- Go server, which works like a middleware between your vim and Gemini
- Vim plugin, which sends the code to the server and shows the results

### Go server

Go server is on `server` folder. It expect a TOKEN as environment variable. You can run it with:

```bash
$ GEMINI_API_KEY=YOUR_TOKEN_HERE go run server/main.go
```

It will wait for connections on port 32000 or `PORT` environment variable.

### Vim plugin

Vim plugin can be installed using your favorite plugin manager.
It expect a `g:gemini_server` variable in your `.vimrc` which points to your just started server

```vim
let g:gemini_server = 'localhost:32000'
```

## How to use it

Select your piece of code using visual mode and hit `:GMExplain` or `:GMDebug`

