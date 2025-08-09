import { defineConfig } from 'orval';

export default defineConfig({
  api: {
    input: './openapi.yml',
    output: {
      mode: 'tags-split',
      target: './src/libs/api/generated/orval',
      schemas: './src/libs/api/generated/orval/model',
      client: 'react-query',
      prettier: true,
      override: {
        mutator: {
          path: './src/libs/api/custom-instance.ts',
          name: 'customInstance',
        },
        query: {
          useQuery: true,
          useMutation: true,
          signal: true,
        },
      },
    },
    hooks: {
      afterAllFilesWrite: 'prettier --write',
    },
  },
});