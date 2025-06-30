local h = import 'html/main.libsonnet';

local directory = {
  'index.html': h.html({ lang: 'en' }, [
    h.body({}, [
      h.div({}, [
        'Hello World!',
      ]),
    ]),
  ]),
  articles: {
    'index.html': h.html({ lang: 'en' }, [
      h.body({}, [
        h.div({}, [
          'Articles',
        ]),
      ]),
    ]),
  },
};

{
  directory: directory,
  config+: {
    serve+: {
      lib+: {
        jpath: ['vendor'],
      },
      server+: {
        staticFiles+: {
          static: 'static',
        },
      },
    },
  },
}
