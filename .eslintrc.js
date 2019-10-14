module.exports = {
  parser: '@typescript-eslint/parser',
  plugins: [ 'import', 'prettier', '@typescript-eslint'],
  extends: [
    'airbnb-base',
    'plugin:@typescript-eslint/recommended',
    'prettier',
    'prettier/@typescript-eslint'
  ],
  rules: {
    'prettier/prettier': ['error'],
    '@typescript-eslint/no-explicit-any': 'error',
    '@typescript-eslint/no-unused-vars': 'error',
    'no-underscore-dangle': 'off',
    'import/prefer-default-export': 'off'
  },
  settings: {
    'import/extensions': ['.js','.ts'],
    'import/parsers': {
      '@typescript-eslint/parser': ['.ts']
      },
      'import/resolver': {
          node: {
              extensions: ['.js','.ts']
          }
      }
  },
  globals: {}
}