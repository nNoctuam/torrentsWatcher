import { Torrent } from '../pb/torrent_pb'
import { Torrents } from '../pb/torrentsList_pb'

const api = {

  getTorrents () {
    return fetch('/torrents')
      .then(async (r) => {
        if (r.status !== 200) {
          const text = await r.text()
          throw new Error(text)
        }
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
      .then(async (r) => {
        if (r.status !== 200) {
          const text = await r.text()
          throw new Error(text)
        }
        const result = await r.arrayBuffer()
        return Torrent.deserializeBinary(result).toObject()
      })
  },

  deleteTorrent (id) {
    return fetch('/torrent/' + id, {
      method: 'DELETE'
    })
  },

  search (text) {
    return fetch('/search', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json;charset=utf-8'
      },
      body: JSON.stringify({
        text
      })
    })
      .then(async (r) => {
        if (r.status !== 200) {
          const text = await r.text()
          throw new Error(text)
        }
        const result = await r.arrayBuffer()
        return Torrents.deserializeBinary(result).toObject().torrentsList
      })
  }

}

export default api
