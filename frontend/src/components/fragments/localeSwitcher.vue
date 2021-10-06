<template>
  <div class="locale-switcher">
    <div
      v-for="locale in $i18n.availableLocales"
      :key="locale"
      @click="$i18n.locale = locale"
      :class="{ selected: $i18n.locale === locale }"
    >
      <img :src="require('../../assets/' + locale + '-flag.png')" alt="" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useI18n } from "vue-i18n";

export default defineComponent({
  name: "localeSwitcher",

  watch: {
    "$i18n.locale"(value: string) {
      localStorage.setItem("locale", value);
    },
  },

  setup() {
    const { t } = useI18n({
      inheritLocale: true,
      useScope: "local",
    });
    console.log(t);
    return { t };
  },
  mounted() {
    console.log(this);
  },
});
</script>

<style scoped lang="stylus">
.locale-switcher
  position: absolute
  top: 0
  right: 0
  div
    cursor: pointer
    float: left
    transition: 0.2s all
    &:not(.selected):not(:hover)
      opacity: 0.5
  img
    width: 50px
    height: 33px
</style>
