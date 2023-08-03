// Overriding CreateReactApp settings, ref: https://github.com/arackaf/customize-cra
const {
  override,
  useEslintRc,
  addDecoratorsLegacy,
  useBabelRc
} = require('customize-cra')

module.exports = override(
  addDecoratorsLegacy(),
  useEslintRc(),
  useBabelRc()
)
