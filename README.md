## split

A Slack chat bot that meows.

![](http://i.imgur.com/bn5b8U9.jpg)

### Running

#### Prerequisites
- Go 1.7

#### Setting up
```sh
make setup
cp config.json.example config.json # then fill in SlackToken
```

#### Running
```sh
make
```

Split will automatically restart whenever a Go file is changed.
