const base = 'http://localhost:8080'

const api = {

  getTorrents () {
    return fetch(`${base}/torrents`)
  },

  addTorrent (url) {
    return fetch(`${base}/torrent`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json;charset=utf-8'
      },
      body: JSON.stringify({
        url
      })
    })
  }
}

export default api
