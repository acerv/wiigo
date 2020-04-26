# wiigo

Telegram bot written in Go for Warsow.ita.

# Build

To build the bot:

    go build -o wii *.go

# Configure
To configure the bot, create a file called `config.ini` inside the binary folder.

    [telegram]
    token = <your bot token>

    [imgur]
    client_id = <imgur client id>
