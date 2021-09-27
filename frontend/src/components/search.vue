<template>
  <div id="search">
    <h1>Торрент-поисковик</h1>

    <form class="form-group" id="search-form" v-on:submit.prevent="search">
      <div class="input-group">
        <input
          type="text"
          class="form-input"
          name="search"
          placeholder="Что ищем?"
          :disabled="searching"
          v-model="searchText"
        />
        <button class="btn" :disabled="searching">
          {{ searching ? "Ищем..." : "Искать" }}
        </button>
      </div>
    </form>

    <table class="table table-striped table-hover" v-if="torrents.length > 0">
      <thead>
        <tr>
          <th class="forum">Раздел</th>
          <th class="title">Название</th>
          <th class="seeders">Раздают</th>
          <th class="size">Размер</th>
          <th class="updated_at">Обновлен</th>
          <th class="download"></th>
        </tr>
      </thead>

      <tbody>
        <tr
          :class="{ active: selectedRow === i }"
          v-for="(torrent, i) in torrents"
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
            <li v-for="folder in folders" v-bind:key="folder">
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
          <h5>Загружается...</h5>
        </div>
        <div class="modal-footer"></div>
      </div>
    </div>

    <div class="modal renaming" :class="{ active: downloadedName }">
      <div class="modal-overlay"></div>
      <form
        class="modal-container form-group"
        v-on:submit.prevent="renameTorrent(downloadedHash, newName)"
      >
        <div class="modal-header">
          <button
            href="#close"
            class="btn btn-clear float-right"
            @click.prevent="downloadedName = null"
          ></button>
          <h5>Торрент загружается. Переименовать?</h5>
        </div>
        <div class="modal-body">
          <input type="text" class="form-input" v-model="newName" />
        </div>

        <div class="modal-footer">
          <input
            class="btn float-left"
            :disabled="renaming"
            type="submit"
            value="Переименовать"
          />
          <button
            class="btn"
            :disabled="renaming"
            @click.prevent="downloadedName = null"
          >
            Оставить как есть
          </button>
        </div>
      </form>
    </div>

    <div class="modal error" :class="{ active: error }">
      <div class="modal-overlay"></div>
      <div class="modal-container">
        <div class="modal-header">
          <h5>Что-то пошло не так</h5>
        </div>
        <div class="modal-body">
          <span>{{ error }}</span>
        </div>
        <div class="modal-footer">
          <button class="btn close" v-on:click="error = null">Закрыть</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import api from "../js/api";
import moment from "moment";
import { Torrent } from "@/pb/torrent_pb";
import { defineComponent } from "vue";

interface TorrentLocal extends Torrent.AsObject {
  isBeingDownloaded: boolean;
}

class Data {
  searchText = "";
  searching = false;

  selectedRow: number | null = null;

  folders: Map<string, string> = new Map();
  showSelectFolder = false;
  folderSelect: null | ((folder: string) => void) = null;
  folderSelectCancel: null | (() => void) = null;

  downloading = false;
  downloadedName: string | null = null;
  downloadedHash: string | null = null;

  newName: string | null = null;
  renaming = false;

  error: string | null = null;

  torrents: TorrentLocal[] = [];
}

export default defineComponent({
  name: "search",

  data: (): Data => ({
    searchText: "",
    searching: false,

    selectedRow: null,

    folders: new Map(),
    showSelectFolder: false,
    folderSelect: null,
    folderSelectCancel: null,

    downloading: false,
    downloadedName: null,
    downloadedHash: null,

    newName: null,
    renaming: false,

    error: null,

    torrents: [],
  }),

  methods: {
    timeFromNow(time: string | number): string {
      return moment(time).fromNow();
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
        .then((r) => {
          console.log(r);
          this.torrents = r.map((torrent): TorrentLocal => {
            const t: TorrentLocal = torrent as TorrentLocal;
            t.isBeingDownloaded = false;
            return t;
          });
        })
        .catch((e) => {
          this.error = e;
        })
        .then(() => {
          this.searching = false;
        });
    },

    download(torrent: TorrentLocal): void {
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
            this.downloadedName = r.name;
            this.newName = r.name;
            this.downloadedHash = r.hash;
          })
          .catch((e) => {
            this.error = "Не удалось скачать: " + e;
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

    renameTorrent(downloadedHash: string, newName: string): void {
      console.log("renaming " + downloadedHash + " to " + newName);
      this.renaming = true;
      // setTimeout(() => {
      //   this.downloadedName = null
      // }, 1000)
      api
        .renameTorrent(downloadedHash, newName)
        .catch((e) => {
          this.error = "Не удалось переименовать: " + e;
        })
        .then(() => {
          this.renaming = false;
          this.downloadedName = null;
        });
    },
  },

  mounted(): void {
    if (this.$route.query.s) {
      this.searchText = this.$route.query.s as string;
      this.search();
    }
    api.getDownloadFolders().then((folders) => {
      this.folders = folders.sort();
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
