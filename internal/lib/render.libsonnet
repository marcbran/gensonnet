local lib = import './lib.libsonnet';

local render(manifest) =
  lib.flattenObject({
    [kv.key]:
      if lib.isDir(kv.key)
      then
        render(manifest { directory: kv.value })
      else
        local generate = std.get(
          std.get(lib, 'generators', {}) + std.get(manifest, 'generators', {}),
          lib.ext(kv.key),
          function(value) std.manifestJson(value)
        );
        generate(kv.value)
    for kv in std.objectKeysValues(manifest.directory)
  });

render
