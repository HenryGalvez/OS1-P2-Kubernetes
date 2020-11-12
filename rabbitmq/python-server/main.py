import pika
from pymongo import MongoClient
import json
from datetime import datetime
import redis

client = MongoClient('mongodb://3.139.106.96:27017/')
red = redis.Redis(host='3.139.106.96', port=6379, db=0)

connection = pika.BlockingConnection(
    pika.ConnectionParameters(host='localhost'))

channel = connection.channel()

channel.queue_declare(queue='rpc_queue')

def on_request(ch, method, props, body):
    db = client.db
    cases = db.cases
    cases.insert_one(json.loads(body))

    now = datetime.now()
    dt_string = now.strftime("%Y-%m-%d %H:%M:%S.%f")
    red.set(dt_string, body)
    
    response = '{\"status\": true, \"message\": \"Se creo correctamente\"}'

    ch.basic_publish(exchange='',
                     routing_key=props.reply_to,
                     properties=pika.BasicProperties(correlation_id = \
                                                         props.correlation_id),
                     body=str(response))
    ch.basic_ack(delivery_tag=method.delivery_tag)

channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue='rpc_queue', on_message_callback=on_request)

print(" [x] Awaiting RPC requests")
channel.start_consuming()