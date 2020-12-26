<template>
  <div id="torrents">

    <form v-on:submit.prevent="addTorrent">
      <input type="url" name="url" :disabled="newTorrentAdding" v-model="newTorrentUrl">
      <button :disabled="newTorrentAdding">{{ newTorrentAdding ? 'Adding...' : 'Add' }}</button>
    </form>

    <table v-if="torrents">
      <thead>
        <tr>
          <th>Title</th>
          <th></th>
          <th>Last upload</th>
          <th>Last check</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="torrent in torrents" v-bind:key="torrent.id">
          <td><a class="open" :href="torrent.pageUrl" target="_blank">{{ torrent.title }}</a></td>
          <td><a class="download" v-if="torrent.fileUrl" :href="'/torrent/' + torrent.id + '/download'"><img src="../assets/transmission-logo.png" alt=""></a></td>
          <td :title="timeFormat(torrent.uploadedAt.seconds * 1000)">{{ timeFromNow(torrent.uploadedAt.seconds * 1000) }}</td>
          <td :title="timeFormat(torrent.updatedAt.seconds * 1000)">{{ timeFromNow(torrent.updatedAt.seconds * 1000) }}</td>
          <td><div class="delete" v-on:click="deleteTorrent(torrent)"><img src="../assets/delete.png" alt=""></div></td>
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
    newTorrentUrl: '',
    newTorrentAdding: false,

    torrents: Array
  }),

  methods: {
    timeFromNow (time) {
      return moment(time).fromNow()
    },
    timeFormat (time, format = 'llll') {
      return moment(time).format(format)
    },
    addTorrent () {
      this.newTorrentAdding = true
      api.addTorrent(this.newTorrentUrl)
        .then(async r => {
          if (r.status !== 200) {
            const text = await r.text()
            throw new Error(text)
          }
          return r.json()
        })
        .then(r => {
          this.torrents.push(r)
        })
        .catch(e => {
          alert(e)
        })
        .then(() => {
          this.newTorrentUrl = ''
          this.newTorrentAdding = false
        })
    },
    deleteTorrent (torrent) {
      api.deleteTorrent(torrent.id)
        .then(async r => {
          if (r.status !== 200) {
            const text = await r.text()
            throw new Error(text)
          }
          this.torrents.splice(torrent, 1)
        })
        .catch(e => {
          alert("failed to delete torrent:" + e)
        })
    }
  },

  mounted () {
    api.getTorrents()
      .then(r => {
        this.torrents = r
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

.download img
  height: 25px

.delete img
  height: 25px

.delete
  cursor pointer

table
  margin-top: 30px
  width: 100%

thead th
  white-space nowrap
  border-bottom: 1px solid gray
  padding-bottom: 10px

</style>
