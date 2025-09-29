const express = require('express')
const { statfsSync } = require('fs')
const app = express()
const port = 8090

const startTime = Date.now()

app.get('/status', (req, res) => {
  const now = Date.now()
  const uptime = (now - startTime) / 3600000
  const stat = statfsSync('/')
  const freeSpace = stat.bfree * stat.bsize / 1000000
  res.send(`Timestamp2: uptime ${uptime.toFixed(2)} hours, free disk in root: ${freeSpace.toFixed(0)} Mbytes`)
})

app.listen(port, () => {
  console.log(`App listening on port ${port}`)
})
