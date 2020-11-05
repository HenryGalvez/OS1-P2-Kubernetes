import sys
import time
import os
import grpc
import decimal
import cases_pb2
import cases_pb2_grpc
from concurrent import futures
from pymongo import MongoClient
from bson import json_util

mongo = MongoClient('mongodb://localhost:27017')


class Cases(cases_pb2_grpc.InsertServicer):
    def InsertCase(self, request, content):
        print("asdasd")
        print(request)
        case = request.case
        #product = request.product
        # codigo para insertar en mongo y en redis
        new = {
            "name": case.name,
            "location": case.location,
            "age": case.age,
            "infectedtype": case.infectedtype,
            "state": case.state
        }
        id = mongo.cases.case.insert_one(new)
        case.id = str(id.inserted_id)
        print("id")
        print(str(id.inserted_id))
        return cases_pb2.CaseRes(case = case)


def get_server(host):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    #keys_dir = os.path.abspath(os.path.join('.', os.pardir, 'keys'))
    #with open('%s/private.key' % keys_dir, 'rb') as f:
    #    private_key = f.read()
    #with open('%s/cert.pem' % keys_dir, 'rb') as f:
    #    certificate_chain = f.read()
    #server_credentials = grpc.ssl_server_credentials(
    #    ((private_key, certificate_chain),))
    #server.add_secure_port(host, server_credentials)
    cases_pb2_grpc.add_InsertServicer_to_server(Cases(), server)
    return server


if __name__ == '__main__':
    port = sys.argv[1] if len(sys.argv) > 1 else 4000
    host = '[::]:%s' % port
    server = get_server(host)
    try:
        server.add_insecure_port(f'[::]:{port}')
        server.start()
        print('Running Server on %s' % host)
        server.wait_for_termination()
        #while True:
        #    time.sleep(1)
    except Exception as e:
        print('[error] %s' % e)
        server.stop(0)
