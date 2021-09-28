/**
 * @fileoverview gRPC-Web generated client stub for protobuf
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as baseService_pb from './baseService_pb';


export class BaseServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSearch = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.TorrentsResponse,
    (request: baseService_pb.SearchRequest) => {
      return request.serializeBinary();
    },
    baseService_pb.TorrentsResponse.deserializeBinary
  );

  search(
    request: baseService_pb.SearchRequest,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.TorrentsResponse>;

  search(
    request: baseService_pb.SearchRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.TorrentsResponse) => void): grpcWeb.ClientReadableStream<baseService_pb.TorrentsResponse>;

  search(
    request: baseService_pb.SearchRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.TorrentsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/Search',
        request,
        metadata || {},
        this.methodInfoSearch,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/Search',
    request,
    metadata || {},
    this.methodInfoSearch);
  }

  methodInfoGetMonitoredTorrents = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.TorrentsResponse,
    (request: baseService_pb.Empty) => {
      return request.serializeBinary();
    },
    baseService_pb.TorrentsResponse.deserializeBinary
  );

  getMonitoredTorrents(
    request: baseService_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.TorrentsResponse>;

  getMonitoredTorrents(
    request: baseService_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.TorrentsResponse) => void): grpcWeb.ClientReadableStream<baseService_pb.TorrentsResponse>;

  getMonitoredTorrents(
    request: baseService_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.TorrentsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/GetMonitoredTorrents',
        request,
        metadata || {},
        this.methodInfoGetMonitoredTorrents,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/GetMonitoredTorrents',
    request,
    metadata || {},
    this.methodInfoGetMonitoredTorrents);
  }

  methodInfoGetDownloadFolders = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.DownloadFoldersResponse,
    (request: baseService_pb.Empty) => {
      return request.serializeBinary();
    },
    baseService_pb.DownloadFoldersResponse.deserializeBinary
  );

  getDownloadFolders(
    request: baseService_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.DownloadFoldersResponse>;

  getDownloadFolders(
    request: baseService_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.DownloadFoldersResponse) => void): grpcWeb.ClientReadableStream<baseService_pb.DownloadFoldersResponse>;

  getDownloadFolders(
    request: baseService_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.DownloadFoldersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/GetDownloadFolders',
        request,
        metadata || {},
        this.methodInfoGetDownloadFolders,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/GetDownloadFolders',
    request,
    metadata || {},
    this.methodInfoGetDownloadFolders);
  }

  methodInfoAddTorrent = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.TorrentResponse,
    (request: baseService_pb.AddTorrentRequest) => {
      return request.serializeBinary();
    },
    baseService_pb.TorrentResponse.deserializeBinary
  );

  addTorrent(
    request: baseService_pb.AddTorrentRequest,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.TorrentResponse>;

  addTorrent(
    request: baseService_pb.AddTorrentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.TorrentResponse) => void): grpcWeb.ClientReadableStream<baseService_pb.TorrentResponse>;

  addTorrent(
    request: baseService_pb.AddTorrentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.TorrentResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/AddTorrent',
        request,
        metadata || {},
        this.methodInfoAddTorrent,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/AddTorrent',
    request,
    metadata || {},
    this.methodInfoAddTorrent);
  }

  methodInfoDeleteTorrent = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.Empty,
    (request: baseService_pb.DeleteTorrentRequest) => {
      return request.serializeBinary();
    },
    baseService_pb.Empty.deserializeBinary
  );

  deleteTorrent(
    request: baseService_pb.DeleteTorrentRequest,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.Empty>;

  deleteTorrent(
    request: baseService_pb.DeleteTorrentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.Empty) => void): grpcWeb.ClientReadableStream<baseService_pb.Empty>;

  deleteTorrent(
    request: baseService_pb.DeleteTorrentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/DeleteTorrent',
        request,
        metadata || {},
        this.methodInfoDeleteTorrent,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/DeleteTorrent',
    request,
    metadata || {},
    this.methodInfoDeleteTorrent);
  }

  methodInfoDownloadTorrent = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.DownloadTorrentResponse,
    (request: baseService_pb.DownloadTorrentRequest) => {
      return request.serializeBinary();
    },
    baseService_pb.DownloadTorrentResponse.deserializeBinary
  );

  downloadTorrent(
    request: baseService_pb.DownloadTorrentRequest,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.DownloadTorrentResponse>;

  downloadTorrent(
    request: baseService_pb.DownloadTorrentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.DownloadTorrentResponse) => void): grpcWeb.ClientReadableStream<baseService_pb.DownloadTorrentResponse>;

  downloadTorrent(
    request: baseService_pb.DownloadTorrentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.DownloadTorrentResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/DownloadTorrent',
        request,
        metadata || {},
        this.methodInfoDownloadTorrent,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/DownloadTorrent',
    request,
    metadata || {},
    this.methodInfoDownloadTorrent);
  }

  methodInfoRenameTorrentParts = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.Empty,
    (request: baseService_pb.RenameTorrentPartsRequest) => {
      return request.serializeBinary();
    },
    baseService_pb.Empty.deserializeBinary
  );

  renameTorrentParts(
    request: baseService_pb.RenameTorrentPartsRequest,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.Empty>;

  renameTorrentParts(
    request: baseService_pb.RenameTorrentPartsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.Empty) => void): grpcWeb.ClientReadableStream<baseService_pb.Empty>;

  renameTorrentParts(
    request: baseService_pb.RenameTorrentPartsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/RenameTorrentParts',
        request,
        metadata || {},
        this.methodInfoRenameTorrentParts,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/RenameTorrentParts',
    request,
    metadata || {},
    this.methodInfoRenameTorrentParts);
  }

  methodInfoGetActiveTorrents = new grpcWeb.AbstractClientBase.MethodInfo(
    baseService_pb.ActiveTorrentsResponse,
    (request: baseService_pb.GetActiveTorrentsRequest) => {
      return request.serializeBinary();
    },
    baseService_pb.ActiveTorrentsResponse.deserializeBinary
  );

  getActiveTorrents(
    request: baseService_pb.GetActiveTorrentsRequest,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.ActiveTorrentsResponse>;

  getActiveTorrents(
    request: baseService_pb.GetActiveTorrentsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.ActiveTorrentsResponse) => void): grpcWeb.ClientReadableStream<baseService_pb.ActiveTorrentsResponse>;

  getActiveTorrents(
    request: baseService_pb.GetActiveTorrentsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.ActiveTorrentsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protobuf.BaseService/GetActiveTorrents',
        request,
        metadata || {},
        this.methodInfoGetActiveTorrents,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protobuf.BaseService/GetActiveTorrents',
    request,
    metadata || {},
    this.methodInfoGetActiveTorrents);
  }

}

