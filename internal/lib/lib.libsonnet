local flattenObject(value, separator='/') =
  if std.type(value) == 'object' then
    std.foldl(function(acc, curr) acc + curr, [
      {
        [std.join(separator, std.filter(function(key) key != '', [child.key, childChild.key]))]: childChild.value
        for childChild in std.objectKeysValues(flattenObject(child.value))
      }
      for child in std.objectKeysValues(value)
    ], {})
  else { '': value };

local isDir(file) = std.length(std.findSubstr('.', file)) == 0;

local ext(file) = '.%s' % std.split(file, '.')[1];

local generators = {
  '.yml'(data): std.manifestYamlDoc(data, indent_array_in_object=true, quote_keys=false),
  '.yaml'(data): self['.yml'](data),
  '.htm'(data): '<!doctype html>' + std.manifestXmlJsonml(data),
  '.html'(data): self['.htm'](data),
};

local watchScript(path, config) = std.manifestXmlJsonml(
  [
    'script',
    |||
      (function() {
          const socket = new WebSocket('ws://localhost:%(port)s/_reload?path=%(path)s');
          socket.addEventListener('message', function(event) {
              if (event.data === 'reload') {
                  window.location.reload();
              }
          });
          socket.addEventListener('error', function(err) {
              console.error('WebSocket error:', err);
          });
          socket.addEventListener('close', function() {
              console.warn('Live-reload connection closed.');
          });
      })();
    ||| % {
      path: path,
      port: config.server.port,
    },
  ]
);

local watchGenerators(path, config) = generators {
  '.htm'(data): super['.htm'](data) + watchScript(path, config),
};

{
  flattenObject: flattenObject,
  isDir: isDir,
  ext: ext,
  generators: generators,
  watchGenerators: watchGenerators,
}
