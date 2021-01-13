import psycopg2
import configparser
import re

from datetime import datetime, timezone, timedelta
from telethon.sync import TelegramClient
from telethon.tl.functions.messages import GetHistoryRequest

config = configparser.ConfigParser()
config.read("config.ini")

database    = config['Database']['name']
user        = config['Database']['user']
password    = config['Database']['password']
host        = config['Database']['host']
port        = config['Database']['port']

api_id      = config['Telegram']['api_id']
api_hash    = config['Telegram']['api_hash']
username    = config['Telegram']['username']

channel_url = config['Channel']['url']

con = psycopg2.connect(
    database=database,
    user=user,
    password=password,
    host=host,
    port=port
)

cur = con.cursor()
cur.execute('''CREATE TABLE IF NOT EXISTS actions
     (id INTEGER PRIMARY KEY,
     type TEXT NOT NULL,
     ticker TEXT NOT NULL,
     price REAL NOT NULL,
     qty INTEGER NOT NULL,
     profit REAL NOT NULL,
     closed BOOLEAN NOT NULL DEFAULT FALSE);''')

client = TelegramClient(username, api_id, api_hash)
client.start()


async def get_messages(channel):
    start_time = datetime.now(timezone.utc) # - timedelta(days=2)

    offset_msg = 0
    limit_msg = 10

    while True:
        history = await client(GetHistoryRequest(
            peer=channel,
            offset_id=offset_msg,
            offset_date=None, add_offset=0,
            limit=limit_msg, max_id=0, min_id=0,
            hash=0))

        if not history.messages:
            continue

        messages = history.messages

        for message in messages:
            msg = message.to_dict()
            if 'message' in msg:
                result = re.findall(r'(SELL|BUY){1} #(\w+)\s?.*\$(\d+\.?\d*)\s?x\s?(\d+)\s?;.*\+(\d+\.?\d*)', msg['message'])
                if len(result) > 0:
                    if msg['date'] > start_time:
                        print(result, msg['date'], start_time, "\n")

                        action_id = msg['id']
                        action_type = result[0][0]
                        ticker = result[0][1]
                        price = result[0][2]
                        qty = result[0][3]
                        profit = result[0][4]

                        cur.execute(
                            "INSERT INTO actions (id,type,ticker,price,qty,profit) VALUES (%s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING", (action_id, action_type, ticker, price, qty, profit)
                        )

        con.commit()
        #offset_msg = messages[len(messages) - 1].id


async def main():
    channel = await client.get_entity(channel_url)
    await get_messages(channel)
    con.close()


with client:
    client.loop.run_until_complete(main())