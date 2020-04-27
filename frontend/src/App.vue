<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/transmission-logo.png">
    <h1>Torrents Watcher</h1>
    <torrents/>
  </div>
</template>

<script>
import torrents from './components/torrents.vue'

export default {
  name: 'App',
  components: {
    torrents
  }
}
</script>

<style lang="stylus">
#app
  font-family Avenir, Helvetica, Arial, sans-serif
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale
  text-align center
  color #2c3e50
  margin-top 60px
  > img
    height 100px
</style>
