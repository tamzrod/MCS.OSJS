// Nameless Classic Icons — package build (SHELL-004).
// Mirrors the @osjs/gnome-icons convention: the SCSS references each SVG via a
// relative `url("./icons/<name>.svg")`; file-loader emits those SVGs into
// `dist/icons/`, and MiniCssExtractPlugin emits `dist/main.css` whose `url()`
// rules then resolve against `dist/icons/` at serve time.
const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const mode = process.env.NODE_ENV || 'development';
const minimize = mode === 'production';

module.exports = {
  mode,
  devtool: 'source-map',
  entry: [path.resolve(__dirname, 'index.js')],
  output: {path: path.resolve(__dirname, 'dist'), filename: 'main.js'},
  optimization: {minimize},
  externals: {osjs: 'OSjs'},
  plugins: [new MiniCssExtractPlugin({filename: 'main.css', chunkFilename: '[id].css'})],
  module: {
    rules: [
      {
        test: /\.(sa|sc|c)ss$/,
        exclude: /node_modules/,
        use: [
          MiniCssExtractPlugin.loader,
          {loader: 'css-loader', options: {sourceMap: true}},
          {loader: 'sass-loader', options: {sourceMap: true}}
        ]
      },
      {
        // Emit each referenced SVG to dist/icons/<name>.svg (preserve path).
        test: /\.svg$/,
        use: [{loader: 'file-loader', options: {name: 'icons/[name].[ext]'}}]
      },
      {test: /\.js$/, exclude: /node_modules/, use: {loader: 'babel-loader'}}
    ]
  }
};
