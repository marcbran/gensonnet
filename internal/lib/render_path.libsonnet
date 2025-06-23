local lib = import './lib.libsonnet';

local renderPath(manifest, path) =
  local segments = if std.type(path) == 'string' then std.split(path, '/') else path;
  local file = segments[0];
  if lib.isDir(file)
  then
    renderPath(manifest { directory: manifest.directory[file] }, segments[1:])
  else
    local generate = std.get(manifest.generators, lib.ext(file), function(value) value);
    generate(manifest.directory[file]);

renderPath
