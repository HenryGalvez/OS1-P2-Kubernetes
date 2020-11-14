const mongoose = require('mongoose')

mongoose.connect('mongodb://3.139.106.96:27017/db', {
    useNewUrlParser: true,
    useUnifiedTopology: true
})
.then(db => console.log('Database is Connected'))
.catch(err => console.log(err))
