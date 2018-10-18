const path = require('path');
const fs = require('fs');
const pbjs = require('protobufjs/cli/pbjs');
const pbts = require('protobufjs/cli/pbts');
const { promisify } = require('util');
const _mkdirp = require('mkdirp');

const mkdirp = promisify(_mkdirp);
const copyFile = promisify(fs.copyFile);
const writeFile = promisify(fs.writeFile);
const compile = promisify(pbjs.main.bind(pbjs));
const compileTs = promisify(pbts.main.bind(pbts));

async function build({ include, protos }) {
  const output = path.join(__dirname, '../js');
  await mkdirp(output);

  const json = await compile(['-t', 'json', '-p', include, ...protos]);
  await writeFile(path.join(output, 'api.json'), json);

  const js = await compile([
    '-t', 'static-module',
    '-w', 'commonjs',
    '-p', include,
    '--force-long',
    '--no-delimited',
    '--force-message',
    ...protos
  ]);
  await writeFile(path.join(output, 'api.js'), js);

  await compileTs([
    '-o', path.join(output, 'index.d.ts'),
    path.join(output, 'api.js')
  ]);

  await copyFile(
    path.join(__dirname, '../src/js/index.js'), 
    path.join(output, 'index.js')
  );

  const proto = await compile(['-t', 'proto3', '-p', include, ...protos]);
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
