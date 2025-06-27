local lib = import './lib.libsonnet';

local renderPath(manifest, path, config, watch, segmentIndex=0) =
  local libGenerator = if watch then lib.watchGenerators(path, config) else lib.generators;
  local file = std.split(path, '/')[segmentIndex];
  if lib.isDir(file)
  then
    renderPath(manifest { directory: manifest.directory[file] }, path, config, watch, segmentIndex + 1)
  else
    local generate = std.get(
      libGenerator + std.get(manifest, 'generators', {}),
      lib.ext(file),
      function(value) std.manifestJson(value)
    );
    generate(manifest.directory[file]);

renderPath
