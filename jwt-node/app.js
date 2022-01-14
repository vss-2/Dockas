require('dotenv').config();
require('./config/database.js').connect();
const express  = require('express');
const app      = express();
const jwt      = require('jsonwebtoken');
const auth     = require('./middleware/auth.js');
const User     = require('./model/user.js');
const mongoose = require('mongoose');
const bcrypt   = require('bcryptjs');
app.use(express.json());

// const Routes = {
// 	"register": "POST /register",
//	"login": "POST, GET /login"
// }

app.post('/register', async (req, res) => {
	try {

		const { name_first, name_last, email, password } = req.body;

		if(!(email && password && name_first && name_last)){
			res.status(400).send('Favor preencher todos os campos');
		}

		const oldUser = await User.findOne({ email });

		if(oldUser){
			return res.status(409).send('Usuário já está registrado. Favor fazer login ou recupere sua senha.');
		}

		encryptedPassword = await bcrypt.hash(password, 10);
		
		const user = await User.create({                              
			name_first,                           
			name_last,                            
			email: email.toLowerCase(),           
			password: encryptedPassword,
			expiresAfterSeconds: 60*60*12         
		});

		const token = jwt.sign(
			{
				user_id: user._id,
				email
			},
			process.env.TOKEN_KEY,
			{
				expiresIn: '12h',
			}
		);

		await User.updateOne({_id: user._id}, {$set: {"token": token}});
		res.status(201).json(`Bem-vindo(a) ${name_first}`);

	} catch (e) {
		console.error(e);
	}
});

app.post('/welcome', auth,  (req, res) => {
	res.status(200).send('Olá mundo');
});

app.post('/logout', auth, async (req, res) => {
	const token = req.body.token || req.query.token || req.headers['x-access-token'];
	if(token){
			await User.findOneAndUpdate(
					{ token: token },
					{ token: null },
					{ new: false }
			);
			const user = await User.findOne({ token });console.log(user);
			res.status(200).send('Logout');
	} else {
			res.status(200).send('Logout');
	}

});

app.post('/login', async (req, res) => {
	try {

		const { email, password } = req.body;
		if(!(email && password)){
			res.status(400).send('Favor preencher todos os campos.');
		}

        const user = await User.findOne({ email });

        if(user && (await bcrypt.compare(password, user.password))) {

		    const token = jwt.sign(
				    {
					    user_id: user._id,
					    email
				    },
				    process.env.TOKEN_KEY,
				    {
					    expiresIn: '12h',
				    }
			);

			await User.updateOne({_id: user._id}, {$set: {"token": token}});
    		res.status(201).json(`Olá novamente ${user.name_first}`);

        } else {
        	res.status(400).send('Email ou senha inválidos.');
		}
	} catch (e) {
		console.error(e);
	}

});

module.exports = app;
