// package: protobuf
// file: torrent.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Torrent extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getTitle(): string;
  setTitle(value: string): void;

  getPageUrl(): string;
  setPageUrl(value: string): void;

  getFileUrl(): string;
  setFileUrl(value: string): void;

  getForum(): string;
  setForum(value: string): void;

  getAuthor(): string;
  setAuthor(value: string): void;

  getSize(): number;
  setSize(value: number): void;

  getSeeders(): number;
  setSeeders(value: number): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUploadedAt(): boolean;
  clearUploadedAt(): void;
  getUploadedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUploadedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Torrent.AsObject;
  static toObject(includeInstance: boolean, msg: Torrent): Torrent.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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

