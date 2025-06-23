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

{
  flattenObject: flattenObject,
  isDir: isDir,
  ext: ext,
}
