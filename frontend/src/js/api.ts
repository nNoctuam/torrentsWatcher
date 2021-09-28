import {
  SearchRequest,
  Torrent,
  Empty,
  AddTorrentRequest,
  DeleteTorrentRequest,
  DownloadTorrentRequest,
  DownloadTorrentResponse, RenameTorrentPartsRequest, PartToRename,
} from "@/pb/baseService_pb";
import { BaseServiceClient } from "@/pb/BaseServiceServiceClientPb";

const domainRPC = "http://localhost:8805";

const api = {
  domain: "http://localhost:8803",
  domainRPC: domainRPC,
  rpcClient: new BaseServiceClient(domainRPC),

  getTorrents(): Promise<Torrent.AsObject[]> {
    return this.rpcClient.getMonitoredTorrents(new Empty(), null).then((r) => {
      return r.getTorrentsList().map((torrent) => torrent.toObject());
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

  getDownloadFolders(): Promise<string[]> {
    return this.rpcClient
      .getDownloadFolders(new Empty(), null)
      .then(async (r) => {
        return r.getFoldersList();
      });
  },

  addTorrent(url: string): Promise<Torrent.AsObject | undefined> {
    const request = new AddTorrentRequest();
    request.setUrl(url);
    return this.rpcClient.addTorrent(request, null).then((r) => {
      return r.getTorrent()?.toObject();
    });
  },

  deleteTorrent(id: number): Promise<Empty> {
    const request = new DeleteTorrentRequest();
    request.setId(id);
    return this.rpcClient.deleteTorrent(request, null);
  },

  search(text: string): Promise<Torrent.AsObject[]> {
    const request = new SearchRequest();
    request.setText(text);

    return this.rpcClient.search(request, null).then((r) => {
      return r.getTorrentsList().map((torrent) => torrent.toObject());
    });
  },

  downloadTorrent(
    pageUrl: string,
    folder: string
  ): Promise<DownloadTorrentResponse.AsObject> {
    const request = new DownloadTorrentRequest();
    request.setUrl(pageUrl);
    request.setFolder(folder);
    return this.rpcClient.downloadTorrent(request, null).then((r) => {
      return r.toObject();
    });
  },

  renameTorrentParts(id: number, names: Array<PartToRename>): Promise<Empty> {
    const request = new RenameTorrentPartsRequest();
    request.setId(id);
    request.setNamesList(names);
    return this.rpcClient.renameTorrentParts(request, null);
  },
};

export default api;
