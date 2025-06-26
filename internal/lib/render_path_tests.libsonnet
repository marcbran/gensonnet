local renderPath = import './render_path.libsonnet';

local exampleTests = {
  name: 'examples',
  tests: [
    {
      name: 'github/dependabot.yml',
      input:: [import '../../examples/github/manifest.jsonnet', 'dependabot.yml', {}, false],
      expected: |||
        updates:
          - directory: "/"
            package-ecosystem: "gomod"
            schedule:
              interval: "daily"
          - directory: "/"
            package-ecosystem: "github-actions"
            schedule:
              interval: "daily"
        version: 2
      |||,
    },
    {
      name: 'github/workflows/test.yml',
      input:: [import '../../examples/github/manifest.jsonnet', 'workflows/test.yml', {}, false],
      expected: |||
        jobs:
          build:
            name: "Build"
            runs-on: "ubuntu-latest"
            steps:
              - uses: "actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683"
              - uses: "actions/setup-go@d35c59abb061a4a6fb18e82ac0862c26744d6ab5"
                with:
                  cache: true
                  go-version-file: "go.mod"
              - run: "go mod download"
              - run: "go test -v -cover -timeout=120s -parallel=10 ./..."
              - run: "go build -v ."
              - name: "Run linters"
                uses: "golangci/golangci-lint-action@4afd733a84b1f43292c63897423277bb7f4313a9"
                with:
                  version: "latest"
            timeout-minutes: 5
        name: "Tests"
        "on":
          pull_request:
            paths-ignore:
              - "README.md"
          push:
            branches:
              - "main"
            paths-ignore:
              - "README.md"
        permissions:
          contents: "read"
      |||,
    },
  ],
};

{
  output(input): renderPath(input[0], input[1], input[2], input[3]) + '\n',
  tests: [
    exampleTests,
  ],
}
