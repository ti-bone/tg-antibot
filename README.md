## tg-antibot
Bot to defend against spammers

There are a large number of spam accounts in Telegram at the moment, which looks like a regular users, they join into your chat, solves captcha(if it present), and after some time start using your chat as a place for advertising, thereby spoiling the situation in the chat. This bot allows you to block these brazen advertisers, thereby preserving the chat from the essential and often deceptive advertising. Please keep in mind that the bot itself does not distribute advertising, but only fights it.

If a user with a name matching the filters joins the chat - it will be banned.

If the user sends a message that matches the filters - it will be banned.

## Hosted
Hosted version: [@againstspam_bot](http://t.me/againstspam_bot)

Note that the bot does not respond to commands, it just does not have them.

You need to add a bot to the chat and give it "Delete messages" and "Restrict users" permissions

## Self-Hosted
### Prepare for build

You will need [GoLang](https://go.dev/) version 1.18 pre-installed at least, as well as [Git](https://git-scm.com)

Clone source code:

`git clone https://github.com/ti-bone/tg-antibot.git`

Enter into `src` directory:

`cd tg-antibot/src`

### Building
Windows: 
`go build -o tg-antibot.exe .`

Unix-based:
`go build -o tg-antibot .`

### Create a bot.
Go to [@BotFather](https://t.me/BotFather).

Write `/newbot`.

Write desired name for your bot.

Write desired username for your bot.

Copy Bot Token.

### Obtain your ID
Go to [@myidbot](https://t.me/myidbot).

Write `/getid`.

Copy your ID.

### Configuring

Go to the `config` directory:

`cd config`

Copy `config.json.example` file to `config.json`:

`cp config.json.example config.json`

Open `config.json` in your favorite editor.

Paste Bot's Token into `Token` field.

Paste your ID into `AdminID` field.

Finnaly, your config should looks like that:
```
{
  "Bot": {
    "Token": "1234567890:Abcdefgh1245678aCGHDee9",
    "AdminID": 777000
  }
}
```

### Run
Windows:
`.\tg-antibot.exe`

Unix-based:
```
chmod +x tg-antibot
./tg-antibot
```

