const express = require("express")
const app = express()

// use the express-static middleware
app.use(express.static("dist"))

// define the first route
app.get('/*', function(req, res) {
    res.sendFile(__dirname + '/dist/index.html');
});

// start the server listening for requests
app.listen(process.env.PORT || 8080);