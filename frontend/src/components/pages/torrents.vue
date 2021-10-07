<template>
  <div id="torrents">
    <h1>{{ t("watch.title") }}</h1>

    <form class="form-group" id="add-form" v-on:submit.prevent="addTorrent">
      <div class="input-group">
        <input
          type="url"
          class="form-input"
          name="url"
          :placeholder="t('watch.input-placeholder')"
          :disabled="newTorrentAdding"
          v-model="newTorrentUrl"
        />
        <button class="btn" :disabled="newTorrentAdding">
          {{ newTorrentAdding ? t("watch.adding") : t("watch.add") }}
        </button>
      </div>
    </form>

    <table class="table table-striped table-hover" v-if="torrents.length > 0">
      <thead>
        <tr>
          <th>{{ t("watch.table.title") }}</th>
          <th></th>
          <th>{{ t("watch.table.updated") }}</th>
          <th>{{ t("watch.table.checked") }}</th>
          <th></th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="torrent in torrents" v-bind:key="torrent.id">
          <td>
            <a class="open" :href="torrent.pageUrl" target="_blank">{{
              torrent.title
            }}</a>
          </td>
          <td>
            <a
              class="download"
              v-if="torrent.fileUrl"
              :href="'/torrent/' + torrent.id + '/download'"
              ><img src="../../assets/transmission-logo.png" alt=""
            /></a>
          </td>
          <td :title="timeFormat(torrent.uploadedAt.seconds * 1000)">
            {{ timeFromNow(torrent.uploadedAt.seconds * 1000) }}
          </td>
          <td :title="timeFormat(torrent.updatedAt.seconds * 1000)">
            {{ timeFromNow(torrent.updatedAt.seconds * 1000) }}
          </td>
          <td>
            <div class="delete" v-on:click="deleteTorrent(torrent)">
              <img src="../../assets/delete.png" alt="" />
            </div>
          </td>
        </tr>
      </tbody>
    </table>

    <errorModal :message="error" @close="error = null" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import moment from "moment";
import api from "@/ts/api";
import { Torrent } from "@/pb/baseService_pb";
import errorModal from "@/components/fragments/errorModal.vue";
import {useI18n} from "vue-i18n";

class Data {
  newTorrentUrl = "";
  newTorrentAdding = false;
  torrents: Torrent.AsObject[] = [];
  error: string | null = null;
}

export default defineComponent({
  name: "torrents",

  components: {
    errorModal,
  },

  setup() {
    const { t } = useI18n({
      inheritLocale: true,
      useScope: "global",
    });
    return { t };
  },

  data: (): Data => ({
    newTorrentUrl: "",
    newTorrentAdding: false,
    torrents: [],
    error: null,
  }),

  methods: {
    timeFromNow(time: string | number): string {
      return moment(time).fromNow();
    },

    timeFormat(time: string | number, format = "llll"): string {
      return moment(time).format(format);
    },

    addTorrent(): void {
      this.newTorrentAdding = true;
      api
        .addTorrent(this.newTorrentUrl)
        .then((r) => {
          if (!r) {
            throw new Error(this.t("watch.error.empty-response"));
          }
          this.torrents.push(r);
        })
        .catch((e) => {
          this.error = e;
        })
        .then(() => {
          this.newTorrentUrl = "";
          this.newTorrentAdding = false;
        });
    },

    deleteTorrent(torrent: Torrent.AsObject): void {
      api
        .deleteTorrent(torrent.id)
        .then(async () => {
          this.torrents.splice(this.torrents.indexOf(torrent), 1);
        })
        .catch((e) => {
          this.error = "failed to delete torrent: " + e;
        });
    },
  },

  mounted(): void {
    api.getTorrents().then((r) => {
      this.torrents = r;
    });
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
h3
  margin 40px 0 0

a
  color #42b983

#add-form
  width: 500px
  text-align: center
  margin: 0 auto

.download img
  height: 25px

.delete img
  height: 25px

.delete
  cursor pointer

table
  margin-top: 30px
  width: 100%

thead th
  white-space nowrap
  border-bottom: 1px solid gray
  padding-bottom: 10px
</style>
