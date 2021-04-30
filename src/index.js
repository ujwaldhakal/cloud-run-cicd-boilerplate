
const express = require('express')
const app = express()
const port = 8080

app.get('/', (req, res) => {
  res.send('hello world')
})

app.get('/healthcheck', (req, res) => {
    res.send('up!')
  })

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})