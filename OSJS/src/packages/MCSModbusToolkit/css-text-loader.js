// Compile approved donor CSS into text for a ShadowRoot-only stylesheet.
// Never import renderer.css using the global MiniCssExtractPlugin rule.
module.exports = source => `module.exports = ${JSON.stringify(source)};`;
