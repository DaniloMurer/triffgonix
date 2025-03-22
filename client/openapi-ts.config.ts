import { defineConfig } from '@hey-api/openapi-ts';

export default defineConfig({
  input: '../server/docs/swagger.json',
  output: './shared/utils',
  plugins: ['@hey-api/client-fetch'],
});
