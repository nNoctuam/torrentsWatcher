<template>
  <div class="main-menu">
    <router-link to="/search">
      <img :alt="t('mainMenu.search')" src="@/assets/search.png" />
    </router-link>
    <router-link to="/watch">
      <img :alt="t('mainMenu.watch')" src="@/assets/transmission-logo.png" />
    </router-link>
    <router-link to="/mass-rename">
      <img :alt="t('mainMenu.massRename')" src="@/assets/rename.png" />
    </router-link>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useI18n } from "vue-i18n";

export default defineComponent({
  name: "mainMenu",

  setup() {
    const { t } = useI18n({
      inheritLocale: true,
      useScope: "global",
    });
    return { t };
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.main-menu
  text-align center
  a:focus {
    outline: none;
    box-shadow: none;
  }
  > a
    transition: 0.3s all
    cursor pointer
    > img
      height 100px
    &:not(.router-link-active):not(:hover)
      opacity 0.25
</style>
