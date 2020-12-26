import { Torrents } from '../pb/torrentsList_pb'

const api = {

  getTorrents () {
    return fetch('/torrents')
      .then(async (r) => {
        const result = await r.arrayBuffer()
        return Torrents.deserializeBinary(result).toObject().torrentsList
      })
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
