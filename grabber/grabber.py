import grpc
import configparser
import re
import proto.trader_pb2 as trader_pb2
import proto.trader_pb2_grpc as trader_pb2_grpc
import time
from datetime import datetime, timezone, timedelta
from telethon.sync import TelegramClient
from telethon.tl.functions.messages import GetHistoryRequest


config = configparser.ConfigParser()
config.read("config.ini")

api_id      = config['Telegram']['api_id']
api_hash    = config['Telegram']['api_hash']
username    = config['Telegram']['username']

channel_url = config['Channel']['url']
msg_pattern = config['Channel']['msg_pattern']

trader_host = config['Trader']['trader_host']
trader_port = config['Trader']['trader_port']
scale_factor = config['Trader']['scale_factor']

channel = grpc.insecure_channel(trader_host + ':' + trader_port)
stub = trader_pb2_grpc.TraderStub(channel)

client = TelegramClient(username, api_id, api_hash)
client.start()

def round_scale(x):
    return int(x / scale_factor)

def parse_text(pattern, text):
    regex = re.compile(pattern)
    result = regex.findall(text)
    return result

def transfer_to_trader(op_type, ticker, qty):
    try:
        request = trader_pb2.CreateMarketOrderRequest(opType=op_type, ticker=ticker, qty=qty)
        response = stub.CreateMarketOrder(request)
        print('[Ответ] %s' % response)
    except Exception:
        return False

    return True

async def get_messages(channel):

    service_time = datetime.now(timezone.utc)
    limit_msg = 5

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

        for message in messages[::-1]:
            msg = message.to_dict()

            if msg['date'] > service_time:
                service_time = msg['date']
                if 'message' in msg:
                    result = parse_text(msg_pattern, msg['message'])
                    if len(result) > 0:
                        print('[Принято сообщение] %s' % msg['message'])
                        op_type = result[0][0]
                        ticker = result[0][1]
                        # price = float(result[0][2])
                        qty = int(result[0][3])
                        scaled_qty = round_scale(qty)

                        if scaled_qty == 0:
                            scaled_qty = 1

                        ok = transfer_to_trader(op_type, ticker, scaled_qty)
                        if not ok:
                            print('[Ошибка] Ошибка вызова сервиса Trader')
        time.sleep(2)

if __name__ == "__main__":
    async def main():
        channel = await client.get_entity(channel_url)
        await get_messages(channel)

    with client:
        client.loop.run_until_complete(main())
