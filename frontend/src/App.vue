<template>
  <div id="app">
    <router-link to="/search">
      <img alt="Search" src="./assets/search.png" :class="{active:mode==='search'}" v-on:click="mode='search'">
    </router-link>
    <router-link to="/watch">
      <img alt="Watch" src="./assets/transmission-logo.png" :class="{active:mode==='watch'}" v-on:click="mode='watch'">
    </router-link>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    mode: 'search'
  })
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
  > a > img
    transition: 0.3s all
    cursor pointer
    height 100px
    &:not(.active):not(:hover)
      opacity 0.25
</style>
