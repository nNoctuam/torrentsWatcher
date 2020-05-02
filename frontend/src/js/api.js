const api = {

  getTorrents () {
    return fetch('/torrents')
  },

  addTorrent (url) {
    return fetch('/torrent', {
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
