<template>
  <div>
    <div id="downloads" v-if="!selected">
      <ul class="downloads">
        <li
          v-for="torrent in downloads"
          :key="torrent.id"
          @click="selected = torrent"
        >
          <span class="chip">{{ torrent.id }}</span>
          <span class="path"
            >{{ torrent.downloaddir.replace(/\/$/, "") }}/</span
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

<script lang="ts">
import { defineComponent } from "vue";
import api from "@/ts/api";
import convertNamesList from "@/ts/renameNamesConverter";
import {
  ActiveTorrent,
  ActiveTorrentPart,
  PartToRename,
} from "@/pb/baseService_pb";

class Data {
  downloads: Array<ActiveTorrent.AsObject> = [];
  selected: ActiveTorrent.AsObject | null = null;
  newNamesList = "";
  files: ActiveTorrentPart.AsObject[] = [];
}

export default defineComponent({
  name: "mass-rename",

  data: (): Data => ({
    downloads: [],
    selected: null,
    newNamesList: "",
    files: [],
  }),

  mounted(): void {
    api.getTransmissionTorrents().then((r) => {
      this.downloads = r.sort((a, b) => b.id - a.id);
    });
  },

  watch: {
    selected(value: ActiveTorrent.AsObject | null): void {
      if (value !== null) {
        api.getTransmissionTorrentFiles(value.id).then((r2) => {
          this.files = r2;
          this.newNamesList = this.getFilesList();
        });
      }
    },
  },

  computed: {
    valid(): boolean {
      return this.newNamesList.trim().split("\n").length === this.files.length;
    },
  },

  methods: {
    getFilesList(): string {
      return this.files.map((f) => f.name).join("\n");
    },

    rename(): void {
      if (this.selected === null) {
        throw new Error("selected cannot be null on rename");
      }
      const basicNamesList: PartToRename[] = [];
      this.newNamesList.split("\n").forEach((name, i) => {
        if (this.files[i].name !== name) {
          const part = new PartToRename();
          part.setOldname(this.files[i].name);
          part.setNewname(name);
          basicNamesList.push(part);
        }
      });
      api
        .renameTorrentParts(this.selected.id, convertNamesList(basicNamesList))
        .then(() => {
          if (this.selected === null) {
            throw new Error("selected cannot be null right after rename");
          }
          api.getTransmissionTorrentFiles(this.selected.id).then((r2) => {
            this.files = r2;
            this.newNamesList = this.getFilesList();
          });
        });
    },
  },
});
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
