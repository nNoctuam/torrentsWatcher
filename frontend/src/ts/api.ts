import {
  SearchRequest,
  Torrent,
  Empty,
  AddTorrentRequest,
  DeleteTorrentRequest,
  DownloadTorrentRequest,
  DownloadTorrentResponse,
  RenameTorrentPartsRequest,
  PartToRename,
  ActiveTorrent,
  GetActiveTorrentsRequest,
  GetActiveTorrentPartsRequest,
  ActiveTorrentPart,
} from "@/pb/baseService_pb";
import { BaseServiceClient } from "@/pb/BaseServiceServiceClientPb";

class API {
  domainRPC = "";
  // @ts-ignore
  rpcClient: BaseServiceClient = null;

  setRpcDomain(domain: string): void {
    this.domainRPC = domain;
    this.rpcClient = new BaseServiceClient(domain);
  }

  getTorrents(): Promise<Torrent.AsObject[]> {
    return this.rpcClient.getMonitoredTorrents(new Empty(), null).then((r) => {
      return r.getTorrentsList().map((torrent: Torrent) => torrent.toObject());
    });
  }

  getTransmissionTorrents(
    onlyRegistered = false
  ): Promise<ActiveTorrent.AsObject[]> {
    const request = new GetActiveTorrentsRequest();
    request.setOnlyregistered(onlyRegistered);
    return this.rpcClient.getActiveTorrents(request, null).then((r) => {
      return r.getTorrentsList().map((torrent: ActiveTorrent) => {
        return torrent.toObject();
      });
    });
  }

  getTransmissionTorrentFiles(
    id: number
  ): Promise<ActiveTorrentPart.AsObject[]> {
    const request = new GetActiveTorrentPartsRequest();
    request.setId(id);
    return this.rpcClient.getActiveTorrentParts(request, null).then((r) => {
      return r.getPartsList().map((part: ActiveTorrentPart) => {
        return part.toObject();
      });
    });
  }

  getDownloadFolders(): Promise<string[]> {
    return this.rpcClient
      .getDownloadFolders(new Empty(), null)
      .then(async (r) => {
        return r.getFoldersList();
      });
  }

  addTorrent(url: string): Promise<Torrent.AsObject | undefined> {
    const request = new AddTorrentRequest();
    request.setUrl(url);
    return this.rpcClient.addTorrent(request, null).then((r) => {
      return r.getTorrent()?.toObject();
    });
  }

  deleteTorrent(id: number): Promise<Empty> {
    const request = new DeleteTorrentRequest();
    request.setId(id);
    return this.rpcClient.deleteTorrent(request, null);
  }

  search(text: string): Promise<Torrent.AsObject[]> {
    const request = new SearchRequest();
    request.setText(text);

    return this.rpcClient.search(request, null).then((r) => {
      return r.getTorrentsList().map((torrent: Torrent) => torrent.toObject());
    });
  }

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
  }

  renameTorrentParts(id: number, names: Array<PartToRename>): Promise<Empty> {
    const request = new RenameTorrentPartsRequest();
    request.setId(id);
    request.setNamesList(names);
    return this.rpcClient.renameTorrentParts(request, null);
  }
}

export default new API();
