module.exports = {
  extends: [
    'airbnb',
    'plugin:@typescript-eslint/recommended',
    'plugin:@typescript-eslint/eslint-recommended',
    'prettier/react',
    'plugin:prettier/recommended',
    'prettier/@typescript-eslint',
    'plugin:import/errors',
    'plugin:import/warnings',
    'plugin:import/typescript',
  ],
  plugins: ['@typescript-eslint'],
  parser: '@typescript-eslint/parser',
  env: { browser: true, node: true, es6: true },
  parserOptions: {
    sourceType: 'module',
  },
  rules: {
    'prettier/prettier': 'error',
    'jsx-a11y/anchor-is-valid': 'off',
    'react/prop-types': 'off',
    'react/jsx-props-no-spreading': 'off',
    'no-shadow': 'off',
    'no-plusplus': [2, { allowForLoopAfterthoughts: true }],
    'no-param-reassign': [2, { props: false }],
    'no-restricted-syntax': ['error', 'ForInStatement', 'LabeledStatement', 'WithStatement'],
    'import/prefer-default-export': 'off',
    'react/jsx-filename-extension': [
      'error',
      {
        extensions: ['.js', '.jsx', '.ts', '.tsx'],
      },
    ],
    'import/extensions': [
      'error',
      'ignorePackages',
      {
        js: 'never',
        jsx: 'never',
        ts: 'never',
        tsx: 'never',
        'd.ts': 'never',
      },
    ],
  },
  settings: {
    'import/extensions': ['.js', '.jsx', '.ts', '.tsx', '.d.ts'],
    'import/parsers': {
      '@typescript-eslint/parser': ['.ts', '.tsx', '.d.ts'],
    },
    'import/resolver': {
      node: {
        moduleDirectory: ['node_modules', 'components'],
        extensions: ['.js', '.jsx', '.ts', '.tsx', '.d.ts'],
      },
    },
  },
};
