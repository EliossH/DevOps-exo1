const { error } = require('console')
const express = require('express')
const { statfsSync, writeFileSync } = require('fs')
const { createBrotliCompress } = require('zlib')
const app = express()
const port = 8080
const ip = '192.168.50.12'

const startTime = Date.now()

app.get('/status', (req, res) => {
  const status = ownStatus()
  writeStatus(status)
  res.send(status)
})

app.listen(port, ip, () => {
  console.log(`App listening on port ${port}`)
})

writeStatus = (status) => {
  writeFileSync('/vStorage', status + "\n", { flag: 'a' })
}

ownStatus = () => {
  const now = Date.now()
  const uptime = (now - startTime) / 3600000
  const stat = statfsSync('/')
  const freeSpace = stat.bfree * stat.bsize / 1000000
  return `Timestamp2: uptime ${uptime.toFixed(2)} hours, free disk in root: ${freeSpace.toFixed(0)} Mbytes`
}