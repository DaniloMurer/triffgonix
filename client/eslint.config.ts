// eslint.config.ts
import js from '@eslint/js';
import tseslint from 'typescript-eslint';
import vue from 'eslint-plugin-vue';
import prettier from 'eslint-plugin-prettier';
import prettierConfig from 'eslint-config-prettier';
import { fileURLToPath } from 'node:url';
import { dirname, resolve } from 'node:path';
import vueParser from 'vue-eslint-parser';

const __dirname = dirname(fileURLToPath(import.meta.url));

export default tseslint.config(
  // Base JavaScript and TypeScript configurations
  js.configs.recommended,
  ...tseslint.configs.recommended,

  // Vue 3 specific configurations
  {
    files: ['**/*.vue', '**/*.ts', '**/*.tsx'],
    plugins: {
      vue,
    },
    languageOptions: {
      parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
      },
      globals: {
        console: 'readonly',
        window: 'readonly',
        document: 'readonly',
      },
    },
    rules: {
      // Vue 3 core rules that are definitely available
      'vue/multi-word-component-names': 'warn',
      'vue/no-v-html': 'warn',
      'vue/require-default-prop': 'off',
      'vue/require-explicit-emits': 'warn',

      // Basic Vue formatting rules that should be available
      'vue/html-indent': ['error', 2],
      'vue/html-quotes': ['error', 'double'],
      'vue/max-attributes-per-line': [
        'error',
        {
          singleline: 3,
          multiline: 1,
        },
      ],
      'vue/no-spaces-around-equal-signs-in-attribute': 'error',
      'vue/v-bind-style': ['error', 'shorthand'],
      'vue/v-on-style': ['error', 'shorthand'],
    },
  },

  // Specific config for .vue files
  {
    files: ['**/*.vue'],
    languageOptions: {
      parser: vueParser,
      parserOptions: {
        parser: tseslint.parser,
        extraFileExtensions: ['.vue'],
        ecmaFeatures: {
          jsx: true,
        },
        sourceType: 'module',
        globals: {
          console: 'readonly',
          window: 'readonly',
          document: 'readonly',
          ref: 'readonly',
          onMounted: 'readonly',
          MessageEvent: 'readonly',
        },
      },
    },
  },

  // TypeScript-specific configuration
  {
    files: ['**/*.ts', '**/*.tsx'],
    languageOptions: {
      parser: tseslint.parser,
      parserOptions: {
        project: [resolve(__dirname, './tsconfig.json')],
        ecmaVersion: 'latest',
        sourceType: 'module',
        ecmaFeatures: {
          jsx: true,
        },
      },
    },
    rules: {
      '@typescript-eslint/no-explicit-any': 'warn',
      '@typescript-eslint/explicit-function-return-type': 'off',
      '@typescript-eslint/no-unused-vars': [
        'error',
        {
          argsIgnorePattern: '^_',
          varsIgnorePattern: '^_',
        },
      ],
      '@typescript-eslint/ban-ts-comment': 'warn',
      '@typescript-eslint/no-non-null-assertion': 'warn',
    },
  },

  // Prettier integration
  {
    plugins: {
      prettier,
    },
    rules: {
      'prettier/prettier': 'error',
    },
  },

  // Apply prettier config
  prettierConfig,

  // Nuxt-specific files
  {
    files: ['pages/**/*.vue', 'layouts/**/*.vue', 'app.vue'],
    rules: {
      'vue/multi-word-component-names': 'off',
    },
  },

  // Files to ignore
  {
    ignores: [
      'node_modules/**',
      '.nuxt/**',
      '.output/**',
      'dist/**',
      '.eslintrc.{js,cjs}',
      'shared/utils/sdk.gen.ts',
      'shared/utils/types.gen.ts',
      'shared/utils/client.gen.ts',
    ],
  }
);
