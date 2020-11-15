FROM python:2
COPY main.py .
RUN pip install pika
RUN pip install pymongo
RUN pip install redis
CMD ["python", "main.py"]