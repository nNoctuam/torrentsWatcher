<template>
  <div id="app">
    <img alt="Watch" src="./assets/transmission-logo.png" :class="{active:mode==='watch'}" v-on:click="mode='watch'">
    <img alt="Search" src="./assets/search.png" :class="{active:mode==='search'}" v-on:click="mode='search'">
    <h1 v-if="mode==='watch'">Torrents Watcher</h1>
    <h1 v-if="mode==='search'">Torrents Seeker</h1>
    <torrents v-if="mode==='watch'"/>
    <search v-if="mode==='search'"/>
  </div>
</template>

<script>
import torrents from './components/torrents.vue'
import search from './components/search'

export default {
  name: 'App',
  data: () => ({
    mode: 'search'
  }),
  components: {
    torrents,
    search
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
    transition: 0.3s all
    cursor pointer
    height 100px
    &:not(.active):not(:hover)
      opacity 0.25
</style>
