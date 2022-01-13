const mongoose      = require('mongoose');
const { MONGO_URI } = process.env;

exports.connect = () => {
	mongoose.connect(
			MONGO_URI,
			{
				useNewUrlParser: true,
				useUnifiedTopology: false
			}).then( () => {
				console.log('Conectado ao banco de dados');
			}).catch( (e) => {
				console.log('Erro: \n');
				console.error(e);
			});
};

