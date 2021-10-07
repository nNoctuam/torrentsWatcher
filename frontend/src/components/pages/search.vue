<template>
  <div id="search">
    <h1>{{ $t("search.title") }}</h1>

    <form class="form-group" id="search-form" v-on:submit.prevent="search">
      <div class="input-group">
        <input
          type="text"
          class="form-input"
          name="search"
          :placeholder="$t('search.placeholder')"
          :disabled="searching"
          v-model="searchText"
        />
        <button class="btn" :disabled="searching">
          {{ searching ? $t("search.searching") : $t("search.search") }}
        </button>
      </div>
    </form>

    <table
      class="table table-striped table-hover"
      v-if="searchResults.length > 0"
    >
      <thead>
        <tr>
          <th class="forum">{{ $t("search.table.forum") }}</th>
          <th class="title">{{ $t("search.table.title") }}</th>
          <th class="seeders">{{ $t("search.table.seeders") }}</th>
          <th class="size">{{ $t("search.table.size") }}</th>
          <th class="updated_at">{{ $t("search.table.updated") }}</th>
          <th class="download"></th>
        </tr>
      </thead>

      <tbody>
        <tr
          :class="{ active: selectedRow === i }"
          v-for="(torrent, i) in searchResults"
          @click="selectedRow = i"
          v-bind:key="torrent.pageUrl"
        >
          <td class="forum">{{ torrent.forum }}</td>
          <td class="title">
            <img :src="getFavicon(torrent.pageUrl)" /><a
              class="open"
              :href="torrent.pageUrl"
              target="_blank"
              >{{ torrent.title }}</a
            >
          </td>
          <td class="seeders">{{ torrent.seeders }}</td>
          <td class="size">{{ byteSize(torrent.size) }}</td>
          <td
            class="updated_at"
            :title="timeFormat(torrent.updatedAt.seconds * 1000)"
          >
            {{ timeFromNow(torrent.updatedAt.seconds * 1000) }}
          </td>
          <td>
            <a class="download" v-on:click.prevent="download(torrent)">
              <i class="icon icon-2x icon-download"></i>
            </a>
          </td>
        </tr>
      </tbody>
    </table>

    <div class="modal folders" :class="{ active: showSelectFolder }">
      <div class="modal-overlay"></div>
      <div class="modal-container">
        <div class="modal-header">
          <button
            class="btn btn-clear float-right"
            v-on:click="folderSelectCancel()"
          ></button>
          <h5>В какую папку?</h5>
        </div>
        <div class="modal-body">
          <ul>
            <li v-for="folder in downloadFolders" v-bind:key="folder">
              <button
                class="btn"
                v-on:click="folderSelect(folder)"
                v-text="folder"
              ></button>
            </li>
          </ul>
        </div>
        <div class="modal-footer">
          <button class="btn close" v-on:click="folderSelectCancel()">
            Отмена
          </button>
        </div>
      </div>
    </div>

    <div class="modal modal-sm downloading" :class="{ active: downloading }">
      <div class="modal-overlay"></div>
      <div class="modal-container">
        <div class="modal-header"></div>
        <div class="modal-body">
          <h5>{{ t("search.downloading") }}</h5>
        </div>
        <div class="modal-footer"></div>
      </div>
    </div>

    <div class="modal renaming" :class="{ active: downloadedTorrent.id }">
      <div class="modal-overlay"></div>
      <form
        class="modal-container form-group"
        v-on:submit.prevent="
          renameTorrent(downloadedTorrent.id, downloadedTorrent.name, newName)
        "
      >
        <div class="modal-header">
          <button
            class="btn btn-clear float-right"
            @click.prevent="downloadedTorrent.id = 0"
          ></button>
          <h5>{{ t("search.downloaded-rename") }}</h5>
        </div>
        <div class="modal-body">
          <input type="text" class="form-input" v-model="newName" />
        </div>

        <div class="modal-footer">
          <input
            class="btn float-left"
            :disabled="renaming"
            type="submit"
            :value="t('search.rename')"
          />
          <button
            class="btn"
            :disabled="renaming"
            @click.prevent="downloadedName = null"
          >
            {{ t("search.leave-as-is") }}
          </button>
        </div>
      </form>
    </div>

    <errorModal :message="error" @close="error = null" />
  </div>
</template>

<script lang="ts">
import api from "../../ts/api";
import moment from "moment";
import { PartToRename, Torrent } from "@/pb/baseService_pb";
import { defineComponent } from "vue";
import { useStore, State } from "@/store";
import { Store, mapState } from "vuex";
import errorModal from "@/components/fragments/errorModal.vue";
import { useI18n } from "vue-i18n";

let store: Store<State>;

interface TorrentLocal extends Torrent.AsObject {
  isBeingDownloaded: boolean;
}

class DownloadedTorrent {
  id: number | null = null;
  name: string | null = null;
  hash: string | null = null;
}

class Data {
  searchText = "";
  searching = false;

  selectedRow: number | null = null;

  showSelectFolder = false;
  folderSelect: null | ((folder: string) => void) = null;
  folderSelectCancel: null | (() => void) = null;

  downloading = false;

  downloadedTorrent: DownloadedTorrent = new DownloadedTorrent();

  newName: string | null = null;
  renaming = false;

  error: string | null = null;
}

export default defineComponent({
  name: "search",

  components: {
    errorModal,
  },

  data: (): Data => ({
    searchText: "",
    searching: false,

    selectedRow: null,

    showSelectFolder: false,
    folderSelect: null,
    folderSelectCancel: null,

    downloading: false,
    downloadedTorrent: {
      id: null,
      name: null,
      hash: null,
    },

    newName: null,
    renaming: false,

    error: null,
  }),

  setup() {
    const { t } = useI18n({
      inheritLocale: true,
      useScope: "global",
    });
    store = useStore();
    return { t };
  },

  computed: {
    ...mapState(["searchResults", "downloadFolders"]),
  },

  methods: {
    timeFromNow(time: string | number): string {
      return time < 0 ? '' : moment(time).fromNow();
    },

    byteSize(bytes: number): string {
      const postfixes = ["", "K", "M", "G", "T", "P", "Y", "Z"];
      let i = 0;
      while (bytes > 1024) {
        bytes = Math.round((bytes / 1024) * 100) / 100;
        i++;
      }
      return bytes + " " + postfixes[i] + "B";
    },

    timeFormat(time: string | number, format = "llll"): string {
      return moment(time).format(format);
    },

    getFavicon(url: string): string {
      const a = document.createElement("a");
      a.href = url;
      return a.protocol + "//" + a.hostname + "/favicon.ico";
    },

    search(): void {
      this.searching = true;
      api
        .search(this.searchText)
        .then((r: any) => {
          const torrents = r.map((torrent: Torrent.AsObject): TorrentLocal => {
            const t: TorrentLocal = torrent as unknown as TorrentLocal;
            t.isBeingDownloaded = false;
            return t;
          });
          store.commit("setSearchResults", torrents);
        })
        .catch((e: any) => {
          this.error = e;
        })
        .then(() => {
          this.searching = false;
        });
    },

    download(torrent: TorrentLocal & Torrent.AsObject): void {
      if (this.downloading || this.showSelectFolder) {
        return;
      }
      this.folderSelect = (folder) => {
        console.log(
          "downloading torrent",
          torrent.pageUrl,
          torrent.isBeingDownloaded,
          folder
        );
        this.showSelectFolder = false;
        this.downloading = true;
        api
          .downloadTorrent(torrent.pageUrl, folder)
          .then((r) => {
            console.log(r);
            this.downloadedTorrent.id = r.id;
            this.downloadedTorrent.name = r.name;
            this.downloadedTorrent.hash = r.hash;

            this.newName = r.name;
          })
          .catch((e) => {
            this.error = "Не удалось скачать: " + JSON.stringify(e);
          })
          .then(() => {
            this.downloading = false;
          });
      };
      this.folderSelectCancel = () => {
        this.showSelectFolder = false;
      };

      this.showSelectFolder = true;

      return;
    },

    renameTorrent(id: number, oldName: string, newName: string): void {
      console.log(`renaming #${id} ${oldName} to ${newName}`);
      this.renaming = true;
      const part = new PartToRename();
      part.setOldname(oldName);
      part.setNewname(newName);
      api
        .renameTorrentParts(id, [part])
        .catch((e) => {
          this.error = "Не удалось переименовать: " + e;
        })
        .then(() => {
          this.renaming = false;
          this.downloadedTorrent.id = 0;
        });
    },
  },

  mounted(): void {
    if (this.$route.query.s) {
      this.searchText = this.$route.query.s as string;
      this.search();
    }
    if (this.downloadFolders.length > 0) {
      return;
    }

    api
      .getDownloadFolders()
      .then((folders) => {
        store.commit("setDownloadFolders", folders.sort());
      })
      .catch((e) => {
        this.error = e;
      });
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
h3
  margin 40px 0 0

a
  color #16a085
  display inline
  transition: 0.2s all
  padding-top: 7px
  //background: red
  text-decoration underline rgba(0,0,0,0)
  text-underline-color white
  &:hover
    text-decoration underline darken(#16a085, 20%)
    color darken(#1abc9c, 20%)

#search-form
  width: 500px
  text-align: center
  margin: 0 auto

table
  margin-top: 30px
  width: 100%

.table tbody tr.active
  background: saturation(rgba(#16a085, 20%), 50%)

thead th
  white-space nowrap

td.forum
  max-width 15%

td.size, td.updated_at
  white-space nowrap

td.title img
  width: 16px
  height: 16px
  margin-right: 10px

.download
  cursor: pointer
  padding-left: 10px
  padding-right: 20px
  img
    height: 25px
    cursor: pointer

.download.pending img
  cursor: default
  opacity 0.25

.modal-header h5
  text-align: center

.downloading h5
  text-align: center

.folders
    li
      display inline-block
      margin-right: 5px
</style>
