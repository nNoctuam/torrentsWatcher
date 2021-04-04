import { Torrent } from '../pb/torrent_pb'
import { Torrents } from '../pb/torrentsList_pb'

const api = {
  domain: '',

  getTorrents () {
    return fetch(this.domain + 'torrents')
      .then(async (r) => {
        if (r.status !== 200) {
          const text = await r.text()
          throw new Error(text)
        }
        const result = await r.arrayBuffer()
        return Torrents.deserializeBinary(result).toObject().torrentsList
      })
  },

  getDownloadFolders () {
    return fetch(this.domain + '/download-folders')
      .then(async (r) => {
        if (r.status !== 200) {
          const text = await r.text()
          throw new Error(text)
        }
        return r.json()
      })
  },

  addTorrent (url) {
    return fetch(this.domain + '/torrent', {
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
    return fetch(this.domain + '/torrent/' + id, {
      method: 'DELETE'
    })
  },

  search (text) {
    return fetch(this.domain + '/search', {
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
  },

  downloadTorrent (pageUrl, folder) {
    return fetch(this.domain + '/download', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json;charset=utf-8'
      },
      body: JSON.stringify({
        Url: pageUrl,
        Folder: folder
      })
    })
      .then(async (r) => {
        if (r.status !== 200) {
          const text = await r.text()
          throw new Error(text)
        }
        return r.json()
      })
  },

  renameTorrent (hash, newName) {
    return fetch(this.domain + '/rename', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json;charset=utf-8'
      },
      body: JSON.stringify({
        hash,
        newName
      })
    })
      .then(async (r) => {
        if (r.status !== 200) {
          const text = await r.text()
          throw new Error(text)
        }
        return r
      })
  }
}

export default api
