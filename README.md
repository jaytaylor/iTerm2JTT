# iTerm2JTT

iTerm2 Just-The-Tips

I made this because sometimes I've accidentally dismissed an iTerm2 tip of the day, and then not been able to find it again via keyword web search.

This exists in part as a measure to remedy the googleability of these gems :)

## About

Golang package for parsing the [iTerm2 Tips of the Day](https://github.com/gnachman/iTerm2/raw/master/sources/iTermTipData.m) source code file.

Also generates a static website of tips.

## Usage

```bash
go get github.com/jaytaylor/iTerm2JTT

curl -sSL https://github.com/gnachman/iTerm2/raw/master/sources/iTermTipData.m \
    | go run main.go \
    && rsync -azve ssh \
        public/* \
        ${SERVER}:/var/www/${SITE}/public_html/iterm2JTT/
```

## License

[Permissive MIT license.](LICENSE)
