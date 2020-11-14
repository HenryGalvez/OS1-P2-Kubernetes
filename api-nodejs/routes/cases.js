var express = require('express');
var router = express.Router();
const cases = require('./models/Case');

var redis = require('redis');
var client = redis.createClient(6379, "3.139.106.96");


/* GET users listing. */
router.get('/', function (req, res, next) {
    res.send('respond with a resource');
});

router.get('/getCasesPerLocation', async function (req, res, next) {
    var s = await cases.aggregate(
        [
            {
                $group:
                {
                    _id: "$location",
                    number:
                        { $sum: 1 }
                }
            }
        ]
    ).sort({ number: -1 })
    //var s = await cases.find()
    res.status(200).json({ message: "Succesfully", data: s });
});

router.get('/getCasesPerRange', async function (req, res, next) {
    var s = await cases.aggregate(
        [
            {
                $group:
                {
                    _id: "$age",
                    number:
                        { $sum: 1 }
                }
            }
        ]
    ).sort({ _id: 1 })
    //var s = await cases.find()
    res.status(200).json({ message: "Succesfully", data: s });
});

router.get('/getLast', async function (req, res, next) {


    client.select(1, function (err, keys) {
        if (err) return console.log(err);
        console.log(keys)
        for (var i = 0; i < keys.length; i++) {
            console.log(keys[i]);
        }
    });
    res.status(200).json({ message: "Succesfully" });
});

router.get('/getLastMongo', async function (req, res, next) {

    var s = await cases.find().sort({ $natural: -1 }).limit(1);
    //var s = await cases.find()
    res.status(200).json({ message: "Succesfully", data: s });
});

module.exports = router;
