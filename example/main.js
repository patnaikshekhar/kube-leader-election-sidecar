const request = require('request-promise')

const LE_SERVER = process.env.LEADER_ELECTION_SERVER || 'localhost:8080'

const runLogic = () => {
  setInterval(() => console.log("I'm now doing work"), 3000)
}

const run = async () => {
  while(true) {
    try {
      const result = await request.get(`http://${LE_SERVER}/isLeader`)
      const resultJSON = JSON.parse(result)
      if (resultJSON.isLeader == true) {
        runLogic()
      }
    } catch(e) {
      console.error(e)
    }
    
    await waitFor(3)
  }
}

const waitFor = (seconds) => new Promise((resolve, reject) => {
  setTimeout(resolve, seconds * 1000)
})

run()
