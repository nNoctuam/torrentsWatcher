<template>
  <div id="search">
      <h1>Torrents search</h1>

    <form v-on:submit.prevent="search">
      <input type="text" name="search" :disabled="searching" v-model="searchText">
      <button :disabled="searching">{{ searching ? 'Searching...' : 'Search' }}</button>
    </form>

    <table v-if="torrents.length > 0">
      <thead>
        <tr>
          <th class="forum">Forum</th>
          <th class="title">Title</th>
          <th class="seeders">Seeders</th>
          <th class="size">Size</th>
          <th class="updated_at">Updated at</th>
          <th class="download"></th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="torrent in torrents" v-bind:key="torrent.pageUrl">
          <td class="forum">{{ torrent.forum }}</td>
          <td class="title"><img :src="getFavicon(torrent.pageUrl)" alt="Tracker"><a class="open" :href="torrent.pageUrl" target="_blank">{{ torrent.title }}</a></td>
          <td class="seeders">{{ torrent.seeders }}</td>
          <td class="size">{{ byteSize(torrent.size) }}</td>
          <td class="updated_at" :title="timeFormat(torrent.updatedAt.seconds * 1000)">{{ timeFromNow(torrent.updatedAt.seconds * 1000) }}</td>
          <td><a class="download" v-on:click="download(torrent)"><img src="../assets/transmission-logo.png" :alt="torrent.title"></a></td>
        </tr>
      </tbody>
    </table>

    <div class="popup folders" v-if="showSelectFolder">
      <div class="content">
        <h4>В какую папку?</h4>
        <ul>
          <li v-for="folder in folders" v-bind:key="folder">
            <button v-on:click="folderSelect(folder)" v-text="folder"></button>
          </li>
        </ul>
        <div class="bottom">
          <button class="close" v-on:click="folderSelectCancel()">Отмена</button>
        </div>
      </div>
    </div>

    <div class="popup downloading" v-show="downloading">
      <div class="content">
        <h5>Загружается...</h5>
      </div>
    </div>

    <div class="popup renaming" v-show="downloadedName">
      <div class="content">
        <p>Загружено как "{{ downloadedName }}". Переименовать?</p>
        <form v-on:submit="renameTorrent(downloadedHash, newName)">
          <input type="text" v-model="newName">

        <div class="bottom">
          <input :disabled="renaming" type="submit" value="Переименовать" />
          <button :disabled="renaming" @click.prevent="downloadedName = null">Отмена</button>
        </div>
        </form>
      </div>
    </div>

  </div>
</template>

<script>
import api from '../js/api'
import moment from 'moment'
// import { Torrents } from '../pb/torrentsList_pb'

export default {
  name: 'search',

  data: () => ({
    searchText: '',
    searching: false,

    folders: [],
    showSelectFolder: false,
    folderSelect: null,
    folderSelectCancel: null,

    downloading: false,
    downloadedName: 'gflfkdlfk.mkv',
    downloadedHash: 'sdsdsds',

    newName: 'gflfkdlfk.mkv',
    renaming: false,

    torrents: [{
      id: 1,
      title: 'test',
      author: 't',
      createdAt: null,
      fileUrl: null,
      forum: null,
      pageUrl: null,
      seeders: null,
      size: null,
      updatedAt: 1093493434,
      uploadedAt: null
    }]
  }),

  methods: {
    timeFromNow (time) {
      return moment(time).fromNow()
    },
    byteSize (bytes) {
      var posfixes = ['', 'K', 'M', 'G', 'T', 'P', 'Y', 'Z']
      var i = 0
      while (bytes > 1024) {
        bytes = Math.round(bytes / 1024 * 100) / 100
        i++
      }
      return bytes + ' ' + posfixes[i] + 'B'
    },
    timeFormat (time, format = 'llll') {
      return moment(time).format(format)
    },
    getFavicon (url) {
      var a = document.createElement('a')
      a.href = url
      return a.protocol + '//' + a.hostname + '/favicon.ico'
    },
    search () {
      this.searching = true
      api.search(this.searchText)
        .then(r => {
          console.log(r)
          this.torrents = r
        })
        .catch(e => {
          alert(e)
        })
        .then(() => {
          this.searching = false
        })
    },
    download (torrent) {
      if (this.downloading || this.showSelectFolder) {
        return
      }
      this.folderSelect = (folder) => {
        console.log('downloading torrent', torrent.pageUrl, torrent.isBeingDownloaded, folder)
        this.showSelectFolder = false
        this.downloading = true
        api.downloadTorrent(torrent.pageUrl, folder)
          .then((r) => {
            console.log(r)
            this.downloadedName = r.name
            this.newName = r.name
            this.downloadedHash = r.hash
          })
          .catch(e => {
            alert('download failed: ' + e)
          })
          .then(() => {
            this.downloading = false
          })
      }
      this.folderSelectCancel = () => {
        this.showSelectFolder = false
      }

      this.showSelectFolder = true

      return false
    },
    renameTorrent (downloadedHash, newName) {
      console.log('renaming ' + downloadedHash + ' to ' + newName)
      this.renaming = true
      // setTimeout(() => {
      //   this.downloadedName = null
      // }, 1000)
      api.renameTorrent(downloadedHash, newName)
        .catch(e => {
          alert('Не удалось переименовать: ' + e)
        })
        .then(() => {
          this.renaming = false
          this.downloadedName = null
        })
    }
  },

  mounted () {
    if (this.$route.query.s) {
      this.searchText = this.$route.query.s
      this.search()
    }
    api.getDownloadFolders()
      .then(folders => {
        this.folders = folders.sort()
      })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
h3
  margin 40px 0 0

a
  color #42b983

table
  margin-top: 30px
  width: 100%

thead th
  white-space nowrap
  border-bottom: 1px solid gray
  padding-bottom: 10px
  text-align left

td
  text-align left

td.forum
  max-width 15%

td.size, td.updated_at
  white-space nowrap

td.title img
  width: 16px
  height: 16px
  margin-right: 10px

.download img
  height: 25px
  cursor: pointer

.download.pending img
  cursor: default
  opacity 0.25

.popup
  position fixed
  top 0
  right 0
  left 0
  bottom 0
  background-color: rgba(#64798a, 0.25)

  .content
    border-radius 5px
    box-shadow 0 0 10px 1px gray
    background-color: #fff
    width: 400px
    margin: 20% auto 0
    padding 20px

.folders
    li
      display inline-block
      margin-right: 5px
      margin-bottom: 5px

    .bottom
      padding-top: 20px
      text-align right

//.downloading

</style>
