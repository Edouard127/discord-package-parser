# Discord Package Parser

Simple tool that allows you to parse all the channels ids and all your messages sent in them from your personal package data on discord.

## How to use
Place the executable at the root of your discord package data and run it.
You can exclude servers and channels by using the `-i` or `--ignore` argument like this:
> discord-package-parser -i 2196765115853734098 1978827830145665748 ...

It will generate a CSV file called `messages.csv` with all the channels ids mapped to the messages you sent in them.

The format is taken directly from [Discord's CSV template](https://docs.google.com/spreadsheets/d/1XvVHgET0LYrUiDvRy2cPfBMQTIr3AulYkpLbUVdQjGk/) 

## What can I do with this ?
You can force discord to delete all the channels in your package data by sending a request to the discord support.

See here: https://www.youtube.com/watch?v=g5FbRfwMEuo

This tool should support both legacy and current discord package data.
If not you can always open an issue.