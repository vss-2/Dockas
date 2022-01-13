const jwt    = require('jsonwebtoken');
const config = process.env;

const verifyToken = (req, res, next) => {
	const token = req.body.token || req.query.token || req.headers['x-access-token'];

	if(!token){
		return res.status(403).send('Token não identificado');
	}

	try {
		const decoded = jwt.verify(token, config.TOKEN_KEY);
		req.user = decoded;
	} catch (e) {
		return res.status(401).send('Token inválido');
	}

	return next();
}

module.exports = verifyToken;
