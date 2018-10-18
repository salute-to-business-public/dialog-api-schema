/*
 * Copyright 2018 Dialog LLC <info@dlg.im>
 */

const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const api = require('./api.js');

const packageDefinition = protoLoader.loadSync(
  require.resolve('./api.proto'),
  { keepCase: false, arrays: true }
);

const services = grpc.loadPackageDefinition(packageDefinition);

Object.assign(api.dialog, services.dialog);

module.exports = api;