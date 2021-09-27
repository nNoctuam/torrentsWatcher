import { SearchRequest, TorrentsResponse, Torrent } from "@/pb/baseService_pb";
import { BaseServiceClient } from "@/pb/BaseServiceServiceClientPb";

const domainRPC = "http://localhost:8805";

const api = {
  domain: "http://localhost:8803",
  domainRPC: domainRPC,
  rpcClient: new BaseServiceClient(domainRPC),

  getTorrents(): Promise<Torrent.AsObject[]> {
    return fetch(this.domain + "/torrents").then(async (r) => {
      if (r.status !== 200) {
        const text = await r.text();
        throw new Error(text);
      }
      const result = await r.arrayBuffer();
      return TorrentsResponse.deserializeBinary(
        new Uint8Array(result)
      ).toObject().torrentsList;
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

  getTransmissionTorrentFiles(id: number) {
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

  search(text: string): Promise<Torrent.AsObject[]> {
    const request = new SearchRequest();
    request.setText(text);

    return this.rpcClient.search(request, null).then((r) => {
      return r.getTorrentsList().map((torrent) => torrent.toObject());
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
