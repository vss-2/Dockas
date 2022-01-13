const http   = require('http');
const app    = require('./app.js');
const server = http.createServer(app);

const { API_PORT } = process.env;
const port = process.env.PORT || API_PORT;

server.listen(port, () => {
	console.log(`Servidor rodando na porta ${port}`);
});
