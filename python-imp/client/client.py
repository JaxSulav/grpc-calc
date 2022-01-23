import grpc
import sys
sys.path.append("..")
from libs import calculator_pb2_grpc, calculator_pb2



channel = grpc.insecure_channel("localhost:50051")
stub = calculator_pb2_grpc.calculatorStub(channel)

response = stub.SumService(calculator_pb2.SumRequest(a=10, b=20))
print("SUM: ", response.result)