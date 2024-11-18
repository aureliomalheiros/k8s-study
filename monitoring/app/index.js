var express = require('express');
var app = express();

const client = require('prom-client');

const register = new client.Registry();

client.collectDefaultMetrics({ register });

const counter = new client.Counter({
    name: 'aula_request_total',
    help: 'Request Counter',
    registers: [register] 
});

app.get('/', function (req, res) {
    counter.inc();
    res.send('Hello World!');
});
app.get('/metrics', async function (req, res) {
    try {
        res.set('Content-Type', register.contentType);
        const metrics = await register.metrics(); 
        res.end(metrics);
    } catch (err) {
        res.status(500).end(err.message); 
    }
});

const PORT = 3000;

app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});

