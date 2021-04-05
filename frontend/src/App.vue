<template>
  <div id="app">
    <div class="main-menu">
      <router-link to="/search">
        <img alt="Search" src="./assets/search.png" :class="{active:mode==='search'}" v-on:click="mode='search'">
      </router-link>
      <router-link to="/watch">
        <img alt="Watch" src="./assets/transmission-logo.png" :class="{active:mode==='watch'}" v-on:click="mode='watch'">
      </router-link>
    </div>
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

<style lang="scss">

$primary-color: #1abc9c !default;
@import "../node_modules/spectre.css/src/variables";
@import "../node_modules/spectre.css/src/mixins";

@import "../node_modules/spectre.css/src/base";
@import "../node_modules/spectre.css/src/normalize";
@import "../node_modules/spectre.css/src/buttons";
@import "../node_modules/spectre.css/src/forms";
@import "../node_modules/spectre.css/src/media";
@import "../node_modules/spectre.css/src/typography";
@import "../node_modules/spectre.css/src/modals";
@import "../node_modules/spectre.css/src/utilities";
@import "../node_modules/spectre.css/src/tables";
@import "../node_modules/spectre.css/src/animations";
@import "../node_modules/spectre.css/src/icons";
</style>

<style lang="stylus">
//@import "../node_modules/spectre.css/dist/spectre.min.css";
#app
  font-family Avenir, Helvetica, Arial, sans-serif
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale
  color #2c3e50
  margin-top 60px
  h1
    text-align: center
  .main-menu
    text-align center
    > a > img
      transition: 0.3s all
      cursor pointer
      height 100px
      &:not(.active):not(:hover)
        opacity 0.25
</style>
