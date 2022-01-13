const mongoose = require('mongoose');
const userSchema = new mongoose.Schema({
	name_first: { type: String, default: null },
	name_last: { type: String, default: null },
	email: { type: String, unique: true },
	password: { type: String },
	token: { type: String }
});

module.exports = mongoose.model('user', userSchema);
