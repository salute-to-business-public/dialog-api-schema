/*
 * Copyright 2018 Dialog LLC <info@dlg.im>
 */

const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const packageDefinition = protoLoader.loadSync(
  require.resolve('./api.proto'),
  { keepCase: false, arrays: true }
);

const api = grpc.loadPackageDefinition(packageDefinition);

module.exports = api;