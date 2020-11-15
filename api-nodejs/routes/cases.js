var express = require('express');
var router = express.Router();
const cases = require('./models/Case');

var async = require("async");
var redis = require('redis');
var client = redis.createClient({ port: 6379, host: "3.139.106.96", db: 0 });
//const redisScan = require('redisscan');
const ioredis = require('ioredis');


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

router.get('/getLast', async function (req, res) {
    var jobs = [];
    client.keys('*', function (err, keys) {
        if (err) return console.log(err);
        if (keys) {
            async.map(keys, function (key, cb) {
                client.get(key, function (error, value) {
                    if (error) return cb(error);
                    var job = {};
                    job['key'] = new Date(key);
                    job['value'] = value;
                    cb(null, job);
                });
            }, function (error, results) {
                if (error) return console.log(error);
                //console.log(results);
                results.sort((a, b) => a.key - b.key)
                if (results.length > 0) {
                    res.status(200).json({ message: "Succesfully", data: [JSON.parse(results[results.length - 1].value)] });
                } else {
                    res.status(200).json({ message: "Succesfully", data: [{ name: "", location: "", age: 0, infectedtype: "", state: "" }] });
                }
            });
        }
    });
    /*client.keys('*', function (err, keys) {
        //console.log(err);
        //console.log(keys);
        console.log(keys.sort((a, b) => a.fechas > b.fechas));

    });*/
    //res.status(200).json({ message: "Succesfully", data: true });
});

router.get('/getLastMongo', async function (req, res, next) {

    var s = await cases.find().sort({ $natural: -1 }).limit(1);
    //var s = await cases.find()
    res.status(200).json({ message: "Succesfully", data: s });
});

module.exports = router;
