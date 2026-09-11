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
      {test: /\.js$/, exclude: /node_modules/, use: {loader: 'babel-loader'}}
    ]
  }
};
