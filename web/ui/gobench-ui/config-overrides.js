// Overriding CreateReactApp settings, ref: https://github.com/arackaf/customize-cra
const {
  override,
  // fixBabelImports,
  addLessLoader,
  useEslintRc,
  addDecoratorsLegacy,
  useBabelRc
} = require('customize-cra')

module.exports = override(
  addDecoratorsLegacy(),
  useEslintRc(),
  addLessLoader({
    javascriptEnabled: true
  }),
  useBabelRc()
)
