local lib = import './lib.libsonnet';

local render(manifest) =
  lib.flattenObject({
    [kv.key]:
      if lib.isDir(file)
      then
        render(manifest { directory: kv.value })
      else
        local generate = std.get(manifest.generators, lib.ext(file), function(value) value);
        generate(kv.value)
    for kv in std.objectKeysValues(manifest.directory)
  });

render
