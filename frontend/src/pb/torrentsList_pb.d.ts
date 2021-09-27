// package: protobuf
// file: torrentsList.proto

import * as jspb from "google-protobuf";
import * as torrent_pb from "./torrent_pb";

export class Torrents extends jspb.Message {
  clearTorrentsList(): void;
  getTorrentsList(): Array<torrent_pb.Torrent>;
  setTorrentsList(value: Array<torrent_pb.Torrent>): void;
  addTorrents(value?: torrent_pb.Torrent, index?: number): torrent_pb.Torrent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Torrents.AsObject;
  static toObject(includeInstance: boolean, msg: Torrents): Torrents.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Torrents, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Torrents;
  static deserializeBinaryFromReader(message: Torrents, reader: jspb.BinaryReader): Torrents;
}

export namespace Torrents {
  export type AsObject = {
    torrentsList: Array<torrent_pb.Torrent.AsObject>,
  }
}

