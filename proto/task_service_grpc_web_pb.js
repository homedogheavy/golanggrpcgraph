/**
 * @fileoverview gRPC-Web generated client stub for task
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.task = require('./task_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.task.TaskServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.task.TaskServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.GetTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodDescriptor_TaskService_GetTasks = new grpc.web.MethodDescriptor(
  '/task.TaskService/GetTasks',
  grpc.web.MethodType.UNARY,
  proto.task.GetTaskRequest,
  proto.task.Tasks,
  /**
   * @param {!proto.task.GetTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.task.GetTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodInfo_TaskService_GetTasks = new grpc.web.AbstractClientBase.MethodInfo(
  proto.task.Tasks,
  /**
   * @param {!proto.task.GetTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @param {!proto.task.GetTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.task.Tasks)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.Tasks>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.getTasks =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/GetTasks',
      request,
      metadata || {},
      methodDescriptor_TaskService_GetTasks,
      callback);
};


/**
 * @param {!proto.task.GetTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.Tasks>}
 *     A native promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.getTasks =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/GetTasks',
      request,
      metadata || {},
      methodDescriptor_TaskService_GetTasks);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.PostTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodDescriptor_TaskService_PostTask = new grpc.web.MethodDescriptor(
  '/task.TaskService/PostTask',
  grpc.web.MethodType.UNARY,
  proto.task.PostTaskRequest,
  proto.task.Tasks,
  /**
   * @param {!proto.task.PostTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.task.PostTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodInfo_TaskService_PostTask = new grpc.web.AbstractClientBase.MethodInfo(
  proto.task.Tasks,
  /**
   * @param {!proto.task.PostTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @param {!proto.task.PostTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.task.Tasks)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.Tasks>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.postTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/PostTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_PostTask,
      callback);
};


/**
 * @param {!proto.task.PostTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.Tasks>}
 *     A native promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.postTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/PostTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_PostTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.PutTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodDescriptor_TaskService_PutTask = new grpc.web.MethodDescriptor(
  '/task.TaskService/PutTask',
  grpc.web.MethodType.UNARY,
  proto.task.PutTaskRequest,
  proto.task.Tasks,
  /**
   * @param {!proto.task.PutTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.task.PutTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodInfo_TaskService_PutTask = new grpc.web.AbstractClientBase.MethodInfo(
  proto.task.Tasks,
  /**
   * @param {!proto.task.PutTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @param {!proto.task.PutTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.task.Tasks)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.Tasks>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.putTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/PutTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_PutTask,
      callback);
};


/**
 * @param {!proto.task.PutTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.Tasks>}
 *     A native promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.putTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/PutTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_PutTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.DeleteTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodDescriptor_TaskService_DeleteTask = new grpc.web.MethodDescriptor(
  '/task.TaskService/DeleteTask',
  grpc.web.MethodType.UNARY,
  proto.task.DeleteTaskRequest,
  proto.task.Tasks,
  /**
   * @param {!proto.task.DeleteTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.task.DeleteTaskRequest,
 *   !proto.task.Tasks>}
 */
const methodInfo_TaskService_DeleteTask = new grpc.web.AbstractClientBase.MethodInfo(
  proto.task.Tasks,
  /**
   * @param {!proto.task.DeleteTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.Tasks.deserializeBinary
);


/**
 * @param {!proto.task.DeleteTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.task.Tasks)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.Tasks>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.deleteTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/DeleteTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_DeleteTask,
      callback);
};


/**
 * @param {!proto.task.DeleteTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.Tasks>}
 *     A native promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.deleteTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/DeleteTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_DeleteTask);
};


module.exports = proto.task;

