local lib = import './lib.libsonnet';

local renderPath(manifest, path, config, watch) =
  local libGenerator = if watch then lib.watchGenerators(config) else lib.generators;
  local segments = if std.type(path) == 'string' then std.split(path, '/') else path;
  local file = segments[0];
  if lib.isDir(file)
  then
    renderPath(manifest { directory: manifest.directory[file] }, segments[1:], watch)
  else
    local generate = std.get(
      libGenerator + std.get(manifest, 'generators', {}),
      lib.ext(file),
      function(value) std.manifestJson(value)
    );
    generate(manifest.directory[file]);

renderPath
