import { Torrent } from "@/pb/torrent_pb";
import { Torrents } from "@/pb/torrentsList_pb";

const api = {
  domain: "http://localhost:8803",

  getTorrents(): Promise<Torrent.AsObject[]> {
    return fetch(this.domain + "/torrents").then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      const result = await r.arrayBuffer();
      return Torrents.deserializeBinary(new Uint8Array(result)).toObject()
        .torrentsList;
    });
  },

  getTransmissionTorrents() {
    return fetch(this.domain + "/transmission-torrents").then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      return r.json();
    });
  },

  getTransmissionTorrentFiles(id: bigint) {
    return fetch(
      this.domain +
        "/transmission-torrent-files?" +
        new URLSearchParams({
          id: id.toString(),
        })
    ).then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      return r.json();
    });
  },

  getDownloadFolders() {
    return fetch(this.domain + "/download-folders").then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      return r.json();
    });
  },

  addTorrent(url: string) {
    return fetch(this.domain + "/torrent", {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify({
        url,
      }),
    }).then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      const result = await r.arrayBuffer();
      return Torrent.deserializeBinary(new Uint8Array(result)).toObject();
    });
  },

  deleteTorrent(id: number) {
    return fetch(this.domain + "/torrent/" + id, {
      method: "DELETE",
    });
  },

  search(text: string) {
    return fetch(this.domain + "/search", {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify({
        text,
      }),
    }).then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      const result = await r.arrayBuffer();
      return Torrents.deserializeBinary(new Uint8Array(result)).toObject()
        .torrentsList;
    });
  },

  downloadTorrent(pageUrl: string, folder: string) {
    return fetch(this.domain + "/download", {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify({
        Url: pageUrl,
        Folder: folder,
      }),
    }).then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      return r.json();
    });
  },

  renameTorrent(hash: string, newName: string) {
    return fetch(this.domain + "/rename", {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify({
        hash,
        newName,
      }),
    }).then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      return r;
    });
  },

  renameTorrentParts(id: number, names: string[][]) {
    return fetch(this.domain + "/rename-parts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json;charset=utf-8",
      },
      body: JSON.stringify({
        id,
        names,
      }),
    }).then(async (r) => {
      if (r.status !== 200 && r.status !== 204) {
        const text = await r.text();
        throw new Error(text);
      }
      return r;
    });
  },
};

export default api;
