local md = import './main.libsonnet';
local p = import 'jsonnet-pkg/main.libsonnet';

p.ex({
  example:
    local g = import 'gensonnet/main.libsonnet';
    g.manifestMarkdown(
      md.Document([
        md.Heading1('Title'),
        md.Paragraph(['Hello World!']),
      ])
    ),
}, {
  ThematicBreak: p.ex([{
    name: 'JSON format',
    inputs: [],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.ThematicBreak()
      ),
  }]),
  Heading: p.ex([{
    name: 'JSON format',
    inputs: [1, 'Title'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Heading(1, 'Title')
      ),
  }]),
  Heading1: p.ex([{
    name: 'JSON format',
    inputs: ['Title'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Heading1('Title')
      ),
  }]),
  Heading2: p.ex([{
    name: 'JSON format',
    inputs: ['Title'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Heading2('Title')
      ),
  }]),
  Heading3: p.ex([{
    name: 'JSON format',
    inputs: ['Title'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Heading3('Title')
      ),
  }]),
  Heading4: p.ex([{
    name: 'JSON format',
    inputs: ['Title'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Heading4('Title')
      ),
  }]),
  Heading5: p.ex([{
    name: 'JSON format',
    inputs: ['Title'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Heading5('Title')
      ),
  }]),
  Heading6: p.ex([{
    name: 'JSON format',
    inputs: ['Title'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Heading6('Title')
      ),
  }]),
  CodeBlock: p.ex([{
    name: 'JSON format',
    inputs: [|||
      func main() {
        fmt.Println("Hello World!")
      }
    |||],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.CodeBlock(|||
          func main() {
            fmt.Println("Hello World!")
          }
        |||)
      ),
  }]),
  FencedCodeBlock: p.ex([{
    name: 'JSON format',
    inputs: [|||
      func main() {
        fmt.Println("Hello World!")
      }
    |||, 'go'],
  }]),
  HTMLBlock: p.ex([{
    name: 'JSON format',
    inputs: [|||
      <marquee>Welcome to my website</marquee>
    |||],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.HTMLBlock(|||
          <marquee>Welcome to my website</marquee>
        |||)
      ),
  }]),
  Paragraph: p.ex([{
    name: 'JSON format',
    inputs: [['Hello World!']],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Paragraph(['Hello World!']),
      ),
  }]),
  Blockquote: p.ex([{
    name: 'JSON format',
    inputs: [[md.Paragraph(['Intelligent quote here'])]],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Blockquote([md.Paragraph(['Intelligent quote here'])])
      ),
  }]),
  ListItem: p.ex([{
    name: 'JSON format',
    inputs: [[md.Paragraph(['Do dishes'])]],
  }]),
  List: p.ex([{
    name: 'JSON format',
    inputs: ['-', 0, [
      md.ListItem('Do this'),
      md.ListItem('Do that'),
      md.ListItem('Do this again'),
    ]],
  }]),
  Emphasis: p.ex([{
    name: 'JSON format',
    inputs: [1, 'Emphasised text'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Paragraph([
          md.Emphasis(1, 'Emphasised text'),
        ])
      ),
  }]),
  Em: p.ex([{
    name: 'JSON format',
    inputs: ['Emphasised text'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Paragraph([
          md.Em('Emphasised text'),
        ])
      ),
  }]),
  Strong: p.ex([{
    name: 'JSON format',
    inputs: ['Bold text'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Paragraph([
          md.Strong('Bold text'),
        ])
      ),
  }]),
  Link: p.ex([{
    name: 'JSON format',
    inputs: ['jsonnet', 'https://github.com/marcbran/jsonnet'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Paragraph([
          md.Link('jsonnet', 'https://github.com/marcbran/jsonnet'),
        ])
      ),
  }]),
  Image: p.ex([{
    name: 'JSON format',
    inputs: ['illustrative diagram', './diag.png'],
  }, {
    name: 'Markdown format with gensonnet',
    example:
      local g = import 'gensonnet/main.libsonnet';
      g.manifestMarkdown(
        md.Paragraph([
          md.Image('illustrative diagram', './diag.png'),
        ])
      ),
  }]),
})
