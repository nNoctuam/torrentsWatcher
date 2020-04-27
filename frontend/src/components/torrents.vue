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
          <th>Last upload</th>
          <th>Last check</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="torrent in torrents" v-bind:key="torrent.page_url">
          <td><a :href="torrent.page_url" target="_blank">{{ torrent.title }}</a></td>
          <td :title="timeFormat(torrent.uploaded_at * 1000)">{{ timeFromNow(torrent.uploaded_at * 1000) }}</td>
          <td :title="timeFormat(torrent.updated_at * 1000)">{{ timeFromNow(torrent.updated_at * 1000) }}</td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
import api from '../js/api'
import moment from 'moment'

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
    timeFormat (time, format = 'dddd MMM Do h:mm:ss') {
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
    }
  },

  mounted () {
    api.getTorrents()
      .then(r => r.json())
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

ul
  list-style-type none
  padding 0

li
  display inline-block
  margin 0 10px

a
  color #42b983
</style>
