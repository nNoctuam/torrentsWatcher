<template>
  <div id="app">
    <localeSwitcher />
    <mainMenu />
    <router-view />
  </div>
</template>

<script lang="ts">
import mainMenu from "./components/fragments/mainMenu.vue";
import localeSwitcher from "./components/fragments/localeSwitcher.vue";
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
  components: {
    mainMenu,
    localeSwitcher,
  },
});
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

</style>
