import grpc
import configparser
import re
import proto.trader_pb2 as trader_pb2
import proto.trader_pb2_grpc as trader_pb2_grpc

from datetime import datetime, timezone, timedelta
from telethon.sync import TelegramClient
from telethon.tl.functions.messages import GetHistoryRequest

config = configparser.ConfigParser()
config.read("config.ini")

trader_host = config['Trader']['trader_host']
trader_port = config['Trader']['trader_port']

api_id      = config['Telegram']['api_id']
api_hash    = config['Telegram']['api_hash']
username    = config['Telegram']['username']

channel_url = config['Channel']['url']

# открываем канал и создаем grpc клиент
channel = grpc.insecure_channel(trader_host + ':' + trader_port)
stub = trader_pb2_grpc.TraderStub(channel)

client = TelegramClient(username, api_id, api_hash)
client.start()

async def get_messages(channel):
    start_time = datetime.now(timezone.utc) # - timedelta(days=2)
    limit_msg = 5
    last_msg_id = 306

    while True:
        history = await client(GetHistoryRequest(
            peer=channel,
            offset_id=0,
            offset_date=None, add_offset=0,
            limit=limit_msg, max_id=0, min_id=0,
            hash=0))

        if not history.messages:
            continue

        messages = history.messages

        for message in messages:
            msg = message.to_dict()

            print('message: ', msg['message'])
            print('last_msg_id: ', last_msg_id)

            if msg['id'] <= last_msg_id:
                continue

            if 'message' in msg:
                result = re.findall(r'(BUY){1} #(\w+)\s?.*\$(\d+\.?\d*)\s?x\s?(\d+)\s?;.*\+(\d+\.?\d*)', msg['message'])
                if len(result) > 0:
                    if msg['date'] > start_time:
                        print(result, msg['date'], start_time, "\n")

                        op_type = result[0][0]
                        ticker = result[0][1]
                        # price = result[0][2]
                        qty = result[0][3]
                        scaled_qty = int(qty / 10)
                        if scaled_qty == 0:
                            scaled_qty = 1

                        request = trader_pb2.CreateMarketOrderRequest(opType=op_type, ticker=ticker, qty=scaled_qty)
                        response = stub.CreateMarketOrder(request)
                        print(response.data)

            last_msg_id = msg['id']

async def main():
    channel = await client.get_entity(channel_url)
    await get_messages(channel)

with client:
    client.loop.run_until_complete(main())
