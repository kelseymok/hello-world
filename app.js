const express = require('express');
const winston = require('./winston-config')
const heartbeats = require('heartbeats')
const app = express()
const port = 8080

app.get('/', (req, res) => res.send('Hello World!'))
app.get('/health', (req, res) => res.send('Ok'))


app.listen(port, () => winston.info(`Example app listening on port ${port}!`))

var heart = heartbeats.createHeart(1000);
heart.createEvent(5, function(count, last){
  winston.info("Meow");
});
