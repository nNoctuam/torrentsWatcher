<template>
  <div id="search">

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
          <td><a :class="{download: true, pending: torrent.isBeingDownloaded}" v-on:click="download(torrent)"><img src="../assets/transmission-logo.png" :alt="torrent.title"></a></td>
        </tr>
      </tbody>
    </table>

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

    torrents: []
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
          r.forEach(torrent => {
            torrent.isBeingDownloaded = false
          })
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
      if (torrent.isBeingDownloaded) {
        return
      }
      console.log('downloading torrent', torrent.pageUrl, torrent.isBeingDownloaded)
      torrent.isBeingDownloaded = true
      api.downloadTorrent(torrent.pageUrl)
        .then(() => {
          alert("torrent '" + torrent.title + "' send to download")
        })
        .catch(e => {
          alert('download failed: ' + e)
        })
        .then(() => {
          torrent.isBeingDownloaded = false
        })
      return false
    }
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

</style>
