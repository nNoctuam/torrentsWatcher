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
    baseService_pb.Torrents,
    (request: baseService_pb.SearchRequest) => {
      return request.serializeBinary();
    },
    baseService_pb.Torrents.deserializeBinary
  );

  search(
    request: baseService_pb.SearchRequest,
    metadata: grpcWeb.Metadata | null): Promise<baseService_pb.Torrents>;

  search(
    request: baseService_pb.SearchRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: baseService_pb.Torrents) => void): grpcWeb.ClientReadableStream<baseService_pb.Torrents>;

  search(
    request: baseService_pb.SearchRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: baseService_pb.Torrents) => void) {
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

}

