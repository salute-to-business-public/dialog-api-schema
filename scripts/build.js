const path = require('path');
const fs = require('fs');
const pbjs = require('protobufjs/cli/pbjs');
const { promisify } = require('util');
const _mkdirp = require('mkdirp');

const mkdirp = promisify(_mkdirp);
const writeFile = promisify(fs.writeFile);
const compile = promisify(pbjs.main.bind(pbjs));

async function build({ include, protos }) {
  const output = path.join(__dirname, '../js');
  await mkdirp(output);

  const json = await compile(['--target', 'json', '--path', include, ...protos]);
  await writeFile(path.join(output, 'api.json'), json);

  const js = await compile(['--target', 'static-module', '-w', 'commonjs', '--path', include, ...protos]);
  await writeFile(path.join(output, 'api.js'), js);

  const proto = await compile(['--target', 'proto3', '--path', include, ...protos]);
  await writeFile(path.join(output, 'api.proto'), proto);
}

const protoRoot = path.resolve(__dirname, '../proto');

build({
  include: path.resolve(__dirname, '../include'),
  protos: fs.readdirSync(protoRoot)
    .filter((name) => name.endsWith('.proto'))
    .map((name) => path.resolve(protoRoot, name))
}).catch((error) => {
  console.error(error);
  process.exit(1);
});
