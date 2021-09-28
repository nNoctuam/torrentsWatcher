<template>
  <div id="app">
    <div class="main-menu">
      <router-link to="/search">
        <img alt="Search" src="./assets/search.png" />
      </router-link>
      <router-link to="/watch">
        <img alt="Watch" src="./assets/transmission-logo.png" />
      </router-link>
      <router-link to="/mass-rename">
        <img alt="Mass rename" src="" style="color: transparent" />
      </router-link>
    </div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
// import HelloWorld from './components/HelloWorld.vue'

@Options({})
export default class App extends Vue {}
</script>

<style lang="scss">
$primary-color: #16a085 !default;
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
@import "../node_modules/spectre.css/src/chips";
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
    > a
      transition: 0.3s all
      cursor pointer
      > img
        height 100px
      &:not(.router-link-active):not(:hover)
        opacity 0.25
</style>
