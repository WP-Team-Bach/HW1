const express = require("express")
const http = require("http")

let app = express()

function requestHandler(req, res){
    res.json({result: "hello"})
}

function notFoundHandler(req, res) {
    res.status(404)
    res.json({error: "404"})
}

app.get('/', requestHandler)
app.get('*', notFoundHandler)

let httpServer = http.createServer(app)
httpServer.listen(8080)