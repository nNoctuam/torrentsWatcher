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

}

