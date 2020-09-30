const express = require('express')
const path = require('path')
const app = express()

app.use(express.static(path.join(__dirname, 'build')))

app.get('/', function (req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'))
})
app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'build', 'index.html'))
})
console.log('Server starting on port 8080')
app.listen(6789)
