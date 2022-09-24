# tg-antibot
### Bot for protect against bots-spammers.

# Hosted
### Hosted version: [@againstspam_bot](http://t.me/againstspam_bot)

### Please notice that bot doesn't respond to commands, it just doesn't have them.

### You need to add bot to your chat and give to him "Delete Message" and "Restrict Members" rights.

# Self-Hosted
## Prepare for build

### You need to have installed GoLang 1.18 at least to build that, also you need to have git installed.

### Clone source code

`git clone https://github.com/ti-bone/tg-antibot.git`

### Enter into sources directory

`cd tg-antibot/src`

## Building
#### Windows: 
`go build -o tg-antibot.exe .`

#### Unix-based:
`go build -o tg-antibot .`

## Create a bot.
### Go to [@BotFather](https://t.me/BotFather).
### Write `/newbot`.
### Write desired name for your bot.
### Write desired username for your bot.
### Copy Bot Token.

# Obtain your ID
### Go to [@myidbot](https://t.me/myidbot).
### Write `/getid`.
### Copy your ID.

## Configuring

### Go to the config directory

`cd config`

### Rename\Copy config.json.example file to config.json

`cp config.json.example config.json`

### Open `config.json` in your favorite editor.
### Paste Bot's Token into `Token` field.
### Paste your ID into `AdminID` field.
### Finnaly, your config should looks like that:
```
{
  "Bot": {
    "Token": "1234567890:Abcdefgh1245678aCGHDee9",
    "AdminID": 777000
  }
}
```

## Run
### Windows:
`.\tg-antibot.exe`

### Unix-based
```
chmod +x tg-antibot
./tg-antibot
```

