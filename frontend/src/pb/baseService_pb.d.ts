import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Torrent extends jspb.Message {
  getId(): number;
  setId(value: number): Torrent;

  getTitle(): string;
  setTitle(value: string): Torrent;

  getPageUrl(): string;
  setPageUrl(value: string): Torrent;

  getFileUrl(): string;
  setFileUrl(value: string): Torrent;

  getForum(): string;
  setForum(value: string): Torrent;

  getAuthor(): string;
  setAuthor(value: string): Torrent;

  getSize(): number;
  setSize(value: number): Torrent;

  getSeeders(): number;
  setSeeders(value: number): Torrent;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Torrent;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Torrent;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Torrent;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Torrent;

  getUploadedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUploadedAt(value?: google_protobuf_timestamp_pb.Timestamp): Torrent;
  hasUploadedAt(): boolean;
  clearUploadedAt(): Torrent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Torrent.AsObject;
  static toObject(includeInstance: boolean, msg: Torrent): Torrent.AsObject;
  static serializeBinaryToWriter(message: Torrent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Torrent;
  static deserializeBinaryFromReader(message: Torrent, reader: jspb.BinaryReader): Torrent;
}

export namespace Torrent {
  export type AsObject = {
    id: number,
    title: string,
    pageUrl: string,
    fileUrl: string,
    forum: string,
    author: string,
    size: number,
    seeders: number,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    uploadedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class PartToRename extends jspb.Message {
  getOldname(): string;
  setOldname(value: string): PartToRename;

  getNewname(): string;
  setNewname(value: string): PartToRename;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PartToRename.AsObject;
  static toObject(includeInstance: boolean, msg: PartToRename): PartToRename.AsObject;
  static serializeBinaryToWriter(message: PartToRename, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PartToRename;
  static deserializeBinaryFromReader(message: PartToRename, reader: jspb.BinaryReader): PartToRename;
}

export namespace PartToRename {
  export type AsObject = {
    oldname: string,
    newname: string,
  }
}

export class ActiveTorrent extends jspb.Message {
  getId(): number;
  setId(value: number): ActiveTorrent;

  getName(): string;
  setName(value: string): ActiveTorrent;

  getHash(): string;
  setHash(value: string): ActiveTorrent;

  getComment(): string;
  setComment(value: string): ActiveTorrent;

  getDownloaddir(): string;
  setDownloaddir(value: string): ActiveTorrent;

  getDatecreated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDatecreated(value?: google_protobuf_timestamp_pb.Timestamp): ActiveTorrent;
  hasDatecreated(): boolean;
  clearDatecreated(): ActiveTorrent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActiveTorrent.AsObject;
  static toObject(includeInstance: boolean, msg: ActiveTorrent): ActiveTorrent.AsObject;
  static serializeBinaryToWriter(message: ActiveTorrent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActiveTorrent;
  static deserializeBinaryFromReader(message: ActiveTorrent, reader: jspb.BinaryReader): ActiveTorrent;
}

export namespace ActiveTorrent {
  export type AsObject = {
    id: number,
    name: string,
    hash: string,
    comment: string,
    downloaddir: string,
    datecreated?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class ActiveTorrentPart extends jspb.Message {
  getName(): string;
  setName(value: string): ActiveTorrentPart;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActiveTorrentPart.AsObject;
  static toObject(includeInstance: boolean, msg: ActiveTorrentPart): ActiveTorrentPart.AsObject;
  static serializeBinaryToWriter(message: ActiveTorrentPart, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActiveTorrentPart;
  static deserializeBinaryFromReader(message: ActiveTorrentPart, reader: jspb.BinaryReader): ActiveTorrentPart;
}

export namespace ActiveTorrentPart {
  export type AsObject = {
    name: string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class SearchRequest extends jspb.Message {
  getText(): string;
  setText(value: string): SearchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SearchRequest): SearchRequest.AsObject;
  static serializeBinaryToWriter(message: SearchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchRequest;
  static deserializeBinaryFromReader(message: SearchRequest, reader: jspb.BinaryReader): SearchRequest;
}

export namespace SearchRequest {
  export type AsObject = {
    text: string,
  }
}

export class AddTorrentRequest extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): AddTorrentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddTorrentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddTorrentRequest): AddTorrentRequest.AsObject;
  static serializeBinaryToWriter(message: AddTorrentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddTorrentRequest;
  static deserializeBinaryFromReader(message: AddTorrentRequest, reader: jspb.BinaryReader): AddTorrentRequest;
}

export namespace AddTorrentRequest {
  export type AsObject = {
    url: string,
  }
}

export class DeleteTorrentRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteTorrentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteTorrentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteTorrentRequest): DeleteTorrentRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteTorrentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteTorrentRequest;
  static deserializeBinaryFromReader(message: DeleteTorrentRequest, reader: jspb.BinaryReader): DeleteTorrentRequest;
}

export namespace DeleteTorrentRequest {
  export type AsObject = {
    id: number,
  }
}

export class DownloadTorrentRequest extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): DownloadTorrentRequest;

  getFolder(): string;
  setFolder(value: string): DownloadTorrentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DownloadTorrentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DownloadTorrentRequest): DownloadTorrentRequest.AsObject;
  static serializeBinaryToWriter(message: DownloadTorrentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DownloadTorrentRequest;
  static deserializeBinaryFromReader(message: DownloadTorrentRequest, reader: jspb.BinaryReader): DownloadTorrentRequest;
}

export namespace DownloadTorrentRequest {
  export type AsObject = {
    url: string,
    folder: string,
  }
}

export class RenameTorrentPartsRequest extends jspb.Message {
  getId(): number;
  setId(value: number): RenameTorrentPartsRequest;

  getNamesList(): Array<PartToRename>;
  setNamesList(value: Array<PartToRename>): RenameTorrentPartsRequest;
  clearNamesList(): RenameTorrentPartsRequest;
  addNames(value?: PartToRename, index?: number): PartToRename;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RenameTorrentPartsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RenameTorrentPartsRequest): RenameTorrentPartsRequest.AsObject;
  static serializeBinaryToWriter(message: RenameTorrentPartsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RenameTorrentPartsRequest;
  static deserializeBinaryFromReader(message: RenameTorrentPartsRequest, reader: jspb.BinaryReader): RenameTorrentPartsRequest;
}

export namespace RenameTorrentPartsRequest {
  export type AsObject = {
    id: number,
    namesList: Array<PartToRename.AsObject>,
  }
}

export class GetActiveTorrentsRequest extends jspb.Message {
  getOnlyregistered(): boolean;
  setOnlyregistered(value: boolean): GetActiveTorrentsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActiveTorrentsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetActiveTorrentsRequest): GetActiveTorrentsRequest.AsObject;
  static serializeBinaryToWriter(message: GetActiveTorrentsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActiveTorrentsRequest;
  static deserializeBinaryFromReader(message: GetActiveTorrentsRequest, reader: jspb.BinaryReader): GetActiveTorrentsRequest;
}

export namespace GetActiveTorrentsRequest {
  export type AsObject = {
    onlyregistered: boolean,
  }
}

export class GetActiveTorrentPartsRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetActiveTorrentPartsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetActiveTorrentPartsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetActiveTorrentPartsRequest): GetActiveTorrentPartsRequest.AsObject;
  static serializeBinaryToWriter(message: GetActiveTorrentPartsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetActiveTorrentPartsRequest;
  static deserializeBinaryFromReader(message: GetActiveTorrentPartsRequest, reader: jspb.BinaryReader): GetActiveTorrentPartsRequest;
}

export namespace GetActiveTorrentPartsRequest {
  export type AsObject = {
    id: number,
  }
}

export class TorrentResponse extends jspb.Message {
  getTorrent(): Torrent | undefined;
  setTorrent(value?: Torrent): TorrentResponse;
  hasTorrent(): boolean;
  clearTorrent(): TorrentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TorrentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TorrentResponse): TorrentResponse.AsObject;
  static serializeBinaryToWriter(message: TorrentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TorrentResponse;
  static deserializeBinaryFromReader(message: TorrentResponse, reader: jspb.BinaryReader): TorrentResponse;
}

export namespace TorrentResponse {
  export type AsObject = {
    torrent?: Torrent.AsObject,
  }
}

export class TorrentsResponse extends jspb.Message {
  getTorrentsList(): Array<Torrent>;
  setTorrentsList(value: Array<Torrent>): TorrentsResponse;
  clearTorrentsList(): TorrentsResponse;
  addTorrents(value?: Torrent, index?: number): Torrent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TorrentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TorrentsResponse): TorrentsResponse.AsObject;
  static serializeBinaryToWriter(message: TorrentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TorrentsResponse;
  static deserializeBinaryFromReader(message: TorrentsResponse, reader: jspb.BinaryReader): TorrentsResponse;
}

export namespace TorrentsResponse {
  export type AsObject = {
    torrentsList: Array<Torrent.AsObject>,
  }
}

export class DownloadFoldersResponse extends jspb.Message {
  getFoldersList(): Array<string>;
  setFoldersList(value: Array<string>): DownloadFoldersResponse;
  clearFoldersList(): DownloadFoldersResponse;
  addFolders(value: string, index?: number): DownloadFoldersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DownloadFoldersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DownloadFoldersResponse): DownloadFoldersResponse.AsObject;
  static serializeBinaryToWriter(message: DownloadFoldersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DownloadFoldersResponse;
  static deserializeBinaryFromReader(message: DownloadFoldersResponse, reader: jspb.BinaryReader): DownloadFoldersResponse;
}

export namespace DownloadFoldersResponse {
  export type AsObject = {
    foldersList: Array<string>,
  }
}

export class DownloadTorrentResponse extends jspb.Message {
  getId(): number;
  setId(value: number): DownloadTorrentResponse;

  getName(): string;
  setName(value: string): DownloadTorrentResponse;

  getHash(): string;
  setHash(value: string): DownloadTorrentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DownloadTorrentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DownloadTorrentResponse): DownloadTorrentResponse.AsObject;
  static serializeBinaryToWriter(message: DownloadTorrentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DownloadTorrentResponse;
  static deserializeBinaryFromReader(message: DownloadTorrentResponse, reader: jspb.BinaryReader): DownloadTorrentResponse;
}

export namespace DownloadTorrentResponse {
  export type AsObject = {
    id: number,
    name: string,
    hash: string,
  }
}

export class ActiveTorrentsResponse extends jspb.Message {
  getTorrentsList(): Array<ActiveTorrent>;
  setTorrentsList(value: Array<ActiveTorrent>): ActiveTorrentsResponse;
  clearTorrentsList(): ActiveTorrentsResponse;
  addTorrents(value?: ActiveTorrent, index?: number): ActiveTorrent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActiveTorrentsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ActiveTorrentsResponse): ActiveTorrentsResponse.AsObject;
  static serializeBinaryToWriter(message: ActiveTorrentsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActiveTorrentsResponse;
  static deserializeBinaryFromReader(message: ActiveTorrentsResponse, reader: jspb.BinaryReader): ActiveTorrentsResponse;
}

export namespace ActiveTorrentsResponse {
  export type AsObject = {
    torrentsList: Array<ActiveTorrent.AsObject>,
  }
}

export class ActiveTorrentPartsResponse extends jspb.Message {
  getPartsList(): Array<ActiveTorrentPart>;
  setPartsList(value: Array<ActiveTorrentPart>): ActiveTorrentPartsResponse;
  clearPartsList(): ActiveTorrentPartsResponse;
  addParts(value?: ActiveTorrentPart, index?: number): ActiveTorrentPart;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActiveTorrentPartsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ActiveTorrentPartsResponse): ActiveTorrentPartsResponse.AsObject;
  static serializeBinaryToWriter(message: ActiveTorrentPartsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActiveTorrentPartsResponse;
  static deserializeBinaryFromReader(message: ActiveTorrentPartsResponse, reader: jspb.BinaryReader): ActiveTorrentPartsResponse;
}

export namespace ActiveTorrentPartsResponse {
  export type AsObject = {
    partsList: Array<ActiveTorrentPart.AsObject>,
  }
}

