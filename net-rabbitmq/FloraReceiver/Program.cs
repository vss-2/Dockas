using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;

ConnectionFactory factory = new();
factory.Uri = new Uri("amqp://flora_admin:totallynot123");
factory.ClientProvidedName = "Flora Receiver App";

IConnection conn = factory.CreateConnection();
IModel channel = conn.CreateModel();

string exchangeName = "FloraExchange";
string routingKey = "flora-routing-key";
string queueName = "FloraQueue";
// TODO: Make credentials be set in other file

channel.ExchangeDeclare(exchangeName, ExchangeType.Direct);
channel.QueueDeclare(queueName, durable: false, exclusive: false, arguments: null);
channel.QueueBind(queueName, exchangeName, routingKey, arguments: null);
channel.BasicQos(prefetchSize: 0, prefetchCount: 1, global: false);

var consumer = new EventingBasicConsumer(channel);
consumer.Received += (sender, args) => 
{
	var body = args.Body.ToArray();

	string message = Encoding.UTF8.GetString(body);

	Console.WriteLine(value: $"Message Received: {message}");
	// TODO: Make something useful with the message
	channel.BasicAck(args.DeliveryTag, multiple: false);
}

string consumerTag = channel.BasicConsume(queueName, autoAck: false, consumer);
Console.ReadLine();

channel.BasicCancel(consumerTag);

channel.Close();
conn.Close();

