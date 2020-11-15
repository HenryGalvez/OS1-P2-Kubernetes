const { Schema, model } = require('mongoose')

const userSchema = new Schema({
    name: String,
    location: String,
    age: Number,
    infectedtype: String,
    state: String
})

module.exports = model('cases', userSchema)