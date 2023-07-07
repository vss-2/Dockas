using RabbitMQ.Client;
using System.Text;

ConnectionFactory factory = new();
factory.Uri = new Uri("amqp://flora_admin:totallynot123");
factory.ClientProvidedName = "Flora Sender App";

IConnection conn = factory.CreateConnection();
IModel channel = conn.CreateModel();

string exchangeName = "FloraExchange";
string routingKey = "flora-routing-key";
string queueName = "FloraQueue";
// TODO: Make credentials be set in other file

channel.ExchangeDeclare(exchangeName, ExchangeType.Direct);
channel.QueueDeclare(queueName, durable: false, exclusive: false, arguments: null);
channel.QueueBind(queueName, exchangeName, routingKey, arguments: null);

byte[] messageBodyBytes = Encoding.UTF8.GetBytes(s: "Hello Flora!");
channel.BasicPublish(exchangeName, routingKey, basicProperties: null, messageBodyBytes);

channel.Close();
conn.Close();

