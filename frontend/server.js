const express = require('express');
const cors = require('cors');

const app = express();

app.use(cors({
    origin: 'http://localhost:5173', // Разрешенный источник
    methods: ['GET', 'POST', 'OPTIONS'],
    allowedHeaders: ['Content-Type'],
    credentials: true,
}));

app.post('/auth/login/email', (req, res) => {
    res.send('Hello from Express!');
});

app.listen(8090, () => {
    console.log('Server is running on port 8090');
});