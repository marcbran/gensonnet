local lib = import './lib.libsonnet';

local defaults(manifestDir) = {
  render: {
    targetDir: manifestDir,
    lib: {
      manifestDir: manifestDir,
    },
  },
  serve: {
    server: {
      port: 8000,
      directoryIndex: 'index.html',
    },
    lib: {
      manifestDir: manifestDir,
    },
  },
};

local readConfig(manifest, manifestDir) = defaults(manifestDir) + std.get(manifest, 'config', {});

readConfig
