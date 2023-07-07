# .NET RabbitMQ

## Description
### Using RabbitMQ AMQP implementation to implement some services

## Requirements
* .NET 7
* RabbitMQ-management docker running

## Usage 
### #Deploy a docker container `docker run -d --hostname rmq --name flora-rabbitmq -p 8080:15672 -p 5672:5672 rabbitmq:3-management`.
Then run an instance of FloraSender to send a message and run an instance of FloraReceiver to read the message.
