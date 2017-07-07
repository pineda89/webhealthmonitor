create bot:
1st - join to web.telegram.org
2nd - /newbot to @BotFather

get chat id:
1st - send from receiver a chat to bot
2nd - https://api.telegram.org/bot<YOUR_TOKEN_HERE>/getUpdates
3th - in this reply, you has an "chat":{"id":265646950,"first_name":"********","username":"********","type":"private"}

Environment vars:
* WEBSITE_TO_MONITOR
* FREQUENCY_CHECK   (in seconds)
* TELEGRAM_TOKEN
* TELEGRAM_CHAT_ID
