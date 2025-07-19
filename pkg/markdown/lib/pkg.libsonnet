local p = import 'jsonnet-pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/gensonnet/tree/main/pkg/markdown/lib',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'markdown',
  path: 'markdown',
  target: 'md',
}, |||
  DSL for creating Markdown documents.

  Creating Markdown documents with this library is a two-step process.
  This library itself doesn't output any Markdown strings.
  Instead, it outputs a format that is similar to [JsonML](http://www.jsonml.org/), but for Markdown elements.
  The unofficial name for this format is "JsonMD".

  The second step, creating actual Markdown documents, will require the usage of [gensonnet](https://github.com/marcbran/gensonnet) itself.
  The gensonnet project's `manifestMarkdown` native function takes any value that is valid JsonMD and outputs a string in Markdown format.
  It does so by relying on another Go library ([goldmark](https://github.com/yuin/goldmark)), converting the JsonMD value to a goldmark AST before rendering out the final string.

  So strictly speaking this here Jsonnet library is nothing but syntactic sugar on top of what is provided by gensonnet.
|||, {
  ThematicBreak: p.desc(|||
    https://spec.commonmark.org/0.31.2/#thematic-breaks
  |||),
  Heading: p.desc(|||
    https://spec.commonmark.org/0.31.2/#atx-headings
  |||),
  Heading1: p.desc(|||
    Level 1 heading
  |||),
  Heading2: p.desc(|||
    Level 2 heading
  |||),
  Heading3: p.desc(|||
    Level 3 heading
  |||),
  Heading4: p.desc(|||
    Level 4 heading
  |||),
  Heading5: p.desc(|||
    Level 5 heading
  |||),
  Heading6: p.desc(|||
    Level 6 heading
  |||),
  CodeBlock: p.desc(|||
    https://spec.commonmark.org/0.31.2/#indented-code-blocks
  |||),
  FencedCodeBlock: p.desc(|||
    https://spec.commonmark.org/0.31.2/#fenced-code-blocks
  |||),
  HTMLBlock: p.desc(|||
    https://spec.commonmark.org/0.31.2/#html-blocks
  |||),
  Paragraph: p.desc(|||
    https://spec.commonmark.org/0.31.2/#paragraphs
  |||),
  Blockquote: p.desc(|||
    https://spec.commonmark.org/0.31.2/#block-quotes
  |||),
  ListItem: p.desc(|||
    https://spec.commonmark.org/0.31.2/#list-items
  |||),
  List: p.desc(|||
    https://spec.commonmark.org/0.31.2/#lists
  |||),
  Emphasis: p.desc(|||
    https://spec.commonmark.org/0.31.2/#emphasis-and-strong-emphasis
  |||),
  Em: p.desc(|||
    &lt;em&gt; emphasis
  |||),
  Strong: p.desc(|||
    &lt;strong&gt; emphasis
  |||),
  Link: p.desc(|||
    https://spec.commonmark.org/0.31.2/#links
  |||),
  Image: p.desc(|||
    https://spec.commonmark.org/0.31.2/#images
  |||),
})
