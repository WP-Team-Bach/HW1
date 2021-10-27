const express = require("express")
const http = require("http")
const crypto = require('crypto');
const redis = require("redis")
const client = redis.createClient({host: "redis"});
const { promisify } = require("util");
const getAsync = promisify(client.get).bind(client);
const setAsync = promisify(client.set).bind(client);

const PORT = 8080

let app = express()
app.use(express.urlencoded({extended: true}));

app.post('/sha256', putHandler)
app.get('/sha256', getHandler)
app.get('*', notFoundHandler)

let httpServer = http.createServer(app)
httpServer.listen(PORT)
console.log(`HTTP Server listening on port: ${PORT}`)

async function putHandler(req, res){
    let str = req.body.str
    if (str == null || str.length < 8) {
        res.status(400)
        res.json({error: "400"})
        return
    }
    let sha = crypto.createHash('sha256').update(str).digest("base64")
    await setAsync(sha, str)
    res.json({result: sha})
}

async function getHandler(req, res){
    let sha = req.query.sha
    if (sha == null || sha.length < 44) {
        res.status(400)
        res.json({error: "400"})
        return
    }
    let str = await getAsync(sha)
    res.json({result: str})
}

function notFoundHandler(req, res) {
    res.status(404)
    res.json({error: "404"})
}

