import express from 'express'
import os from 'os'
import fetch from 'node-fetch'

const app = express()
const PORT = 3000

app.get("/", (req, res) => {
  const helloMessage = `<h1>Hello from the ${os.hostname()}</h1>`
  console.log(helloMessage)
  res.send(helloMessage)
})

app.get("/nginx", async (req, res) => {
  const url = 'http://nginx' // nginx will be the service name to connect
  const response = await fetch(url);
  const body = await response.text();
  res.send(body)
})

app.listen(PORT, () => {
  console.log(`Web server is listening at port ${PORT}`)
})