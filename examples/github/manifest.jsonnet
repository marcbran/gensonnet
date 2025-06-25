local directory = {
  'dependabot.yml': {
    version: 2,
    updates: [
      {
        'package-ecosystem': 'gomod',
        directory: '/',
        schedule: { interval: 'daily' },
      },
      {
        'package-ecosystem': 'github-actions',
        directory: '/',
        schedule: { interval: 'daily' },
      },
    ],
  },
  workflows: {
    local workflows = self,
    'test.yml': {
      name: 'Tests',
      on: {
        pull_request: {
          'paths-ignore': [
            'README.md',
          ],
        },
        push: {
          branches: [
            'main',
          ],
          'paths-ignore': [
            'README.md',
          ],
        },
      },
      permissions: {
        contents: 'read',
      },
      jobs: {
        build: {
          name: 'Build',
          'runs-on': 'ubuntu-latest',
          'timeout-minutes': 5,
          steps: [
            {
              uses: 'actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683',
            },
            {
              uses: 'actions/setup-go@d35c59abb061a4a6fb18e82ac0862c26744d6ab5',
              with: {
                'go-version-file': 'go.mod',
                cache: true,
              },
            },
            {
              run: 'go mod download',
            },
            {
              run: 'go test -v -cover -timeout=120s -parallel=10 ./...',
            },
            {
              run: 'go build -v .',
            },
            {
              name: 'Run linters',
              uses: 'golangci/golangci-lint-action@4afd733a84b1f43292c63897423277bb7f4313a9',
              with: {
                version: 'latest',
              },
            },
          ],
        },
      },
    },
    'release.yml': {
      name: 'Release',
      on: {
        push: {
          tags: [
            'v*',
          ],
        },
      },
      permissions: {
        contents: 'write',
        packages: 'write',
      },
      jobs: {
        goreleaser: {
          'runs-on': 'ubuntu-latest',
          steps: [
            {
              uses: 'actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683',
              with: {
                'fetch-depth': 0,
              },
            },
            {
              uses: 'actions/setup-go@d35c59abb061a4a6fb18e82ac0862c26744d6ab5',
              with: {
                'go-version-file': 'go.mod',
                cache: true,
              },
            },
            {
              name: 'Login to GitHub Container Registry',
              uses: 'docker/login-action@v3',
              with: {
                registry: 'ghcr.io',
                username: '${{ github.actor }}',
                password: '${{ secrets.GITHUB_TOKEN }}',
              },
            },
            {
              name: 'Run GoReleaser',
              uses: 'goreleaser/goreleaser-action@9c156ee8a17a598857849441385a2041ef570552',
              with: {
                args: 'release --clean --config .goreleaser.yaml',
              },
              env: {
                GITHUB_TOKEN: '${{ secrets.GITHUB_TOKEN }}',
              },
            },
          ],
        },
      },
    },
  },
};

{
  directory: directory,
}
