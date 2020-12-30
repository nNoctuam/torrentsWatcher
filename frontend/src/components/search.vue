<template>
  <div id="search">

    <form v-on:submit.prevent="addTorrent">
      <input type="text" name="search" :disabled="searching" v-model="searchText">
      <button :disabled="searching">{{ newTorrentAdding ? 'Searching...' : 'Search' }}</button>
    </form>

    <table v-if="torrents">
      <thead>
        <tr>
          <th>Title</th>
          <th>Forum</th>
          <th>Seeders</th>
          <th>Size</th>
          <th>Last upload</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="torrent in torrents" v-bind:key="torrent.id">
          <td><a class="open" :href="torrent.pageUrl" target="_blank">{{ torrent.title }}</a></td>
          <td>{{ torrent.forum }}</td>
          <td>{{ torrent.seeders }}</td>
          <td>{{ byteSize(torrent.size) }}</td>
          <td :title="timeFormat(torrent.updatedAt.seconds * 1000)">{{ timeFromNow(torrent.updatedAt.seconds * 1000) }}</td>
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
  name: 'torrents',

  data: () => ({
    searchText: '',
    searching: false,

    torrents: Array
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
    addTorrent () {
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

td:nth-child(n+1)
  white-space nowrap

</style>
