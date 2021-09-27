<template>
  <div>
    <div id="downloads" v-if="!selected">
      <ul class="downloads">
        <li
          v-for="torrent in downloads"
          :key="torrent.ID"
          @click="selected = torrent"
        >
          <span class="chip">{{ torrent.ID }}</span>
          <span class="path"
            >{{ torrent.downloadDir.replace(/\/$/, "") }}/</span
          >
          <span class="name">{{ torrent.name }}</span>
        </li>
      </ul>
    </div>

    <div id="rename" v-if="files && files.length">
      <button
        class="btn"
        @click="
          selected = null;
          files = [];
        "
      >
        &larr; назад
      </button>

      <div id="mapping" class="input-group">
        <textarea
          :rows="files.length + 1"
          style="width: 50%; overflow-x: auto"
          :value="getFilesList()"
          readonly="true"
        ></textarea>

        <textarea
          :rows="files.length + 1"
          style="width: 50%; overflow-x: auto"
          v-model="newNamesList"
        ></textarea>
      </div>

      <button class="btn" @click="rename" :disabled="!valid">
        Переименовать
      </button>
    </div>
  </div>
</template>

<script>
import api from "../js/api";
import convertNamesList from "../js/renameNamesConverter";

export default {
  name: "mass-rename",
  data: () => ({
    downloads: [],
    selected: null,
    newNamesList: "",
    files: [],
  }),
  mounted() {
    api.getTransmissionTorrents().then((r) => {
      this.downloads = r.sort((a, b) => b.ID - a.ID);
    });
  },
  watch: {
    selected(value) {
      if (value !== null && value !== undefined) {
        api.getTransmissionTorrentFiles(value.ID).then((r2) => {
          this.files = r2;
          this.newNamesList = this.getFilesList();
        });
      }
    },
  },
  computed: {
    valid() {
      return this.newNamesList.trim().split("\n").length === this.files.length;
    },
  },
  methods: {
    getFilesList() {
      return this.files.map((f) => f.name).join("\n");
    },
    rename() {
      const basicNamesList = [];
      this.newNamesList.split("\n").forEach((name, i) => {
        if (this.files[i].name !== name) {
          basicNamesList.push([this.files[i].name, name]);
        }
      });
      api
        .renameTorrentParts(this.selected.ID, convertNamesList(basicNamesList))
        .then(() => {
          const selected = this.selected;
          this.selected = null;
          this.selected = selected;
        });
    },
  },
};
</script>

<style scoped lang="scss">
ul.downloads {
  list-style: none;
  li {
    cursor: pointer;
    transition: 0.2s all;
    &:hover {
      background-color: #eee;
    }
  }
  .chip {
    margin-right: 10px;
    padding-top: 4px;
    width: 50px;
    text-align: center;
  }
  .path {
    color: #bbb;
    margin-right: 15px;
  }
}

#mapping {
  margin: 20px 0;
}
</style>
