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

