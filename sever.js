const express = require('express')
const app = express()
// const bodyParser= require("body-parser");
let multer = require('multer');
let upload = multer();
const port = 3000


// app.use(bodyParser.urlencoded({ extended: true }));
// app.use(bodyParser.json());
// app.use(upload.array()); 

app.get('/', (req, res) => res.send('Hello World!'))
app.post('/login', upload.none(), function (req, res) {
  var user = req.body.user
  var password = req.body.password
  var dataresponse = {}
    if(user && password && user == "admin" && password == "123"){
      dataresponse = {
        status : 'ok'
      }
      res.send(JSON.stringify(dataresponse))
    }
    dataresponse = {
      status : 'error'
    }
    res.send(JSON.stringify(dataresponse))
})
app.listen(port, () => console.log(`Example app listening on port ${port}!`))