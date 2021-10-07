<template>
  <div class="modal active error" v-if="message !== null">
    <div class="modal-overlay"></div>
    <div class="modal-container">
      <div class="modal-header">
        <h5>{{ t("error.title") }}</h5>
      </div>
      <div class="modal-body">
        <span>{{ message }}</span>
      </div>
      <div class="modal-footer">
        <button class="btn close" v-on:click="$emit('close')">{{ t("error.close") }}</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import {useI18n} from "vue-i18n";

export default defineComponent({
  name: "errorModal",

  props: ["message"],

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

.modal-header h5
  text-align: center
</style>
