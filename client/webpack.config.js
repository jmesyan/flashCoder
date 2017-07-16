var webpack = require('webpack');
var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin');

const PATHS = {
    app: path.join(__dirname, 'src'),
    dist: path.join(__dirname, 'dist'),
    tpl: path.join(__dirname,  'tpl')
};

const merge = require('webpack-merge');
const plugins = require('./webpack.plugins');

const common = merge(
    plugins.clean(PATHS.dist),
    {
      entry: {
        index:path.join(PATHS.app, 'index.js'),
        react:['react', 'react-dom', 'react-router', 'react-redux', 'react-bootstrap'],
        redux:['redux',  'redux-saga'],
        plugins:['jquery']
      },
      output:{
        filename: '[name].[chunkhash:8].js',
        path: PATHS.dist,
        publicPath:'/'
      },
      plugins:[
        new webpack.optimize.CommonsChunkPlugin({
          name:['react', 'redux', 'plugins', 'manifest']
        }),
        new HtmlWebpackPlugin({
          filename:'index.html',
          favicon:path.join(PATHS.app,'assets', 'images', 'flash.ico'),
          template:path.join(PATHS.tpl, 'index.html')
        }),
        new webpack.ProvidePlugin({
          $: 'jquery',
          jQuery: 'jquery'
        })
      ],
      module: {
        rules: [
          {
            test: /\.css$/,
            use: [
              'style-loader',
              'css-loader'
            ]
          },
          {
            test: /\.(png|svg|jpg|gif)$/,
            use: [
              'file-loader'
            ]
          },
         {
           test: /\.(woff|woff2|eot|ttf|otf)$/,
          use: [
             'file-loader'
          ]
         },
         {
           test: /\.(js|jsx)$/,
           exclude: /node_modules/, 
           loader: 'babel-loader',
           query: {
             presets:['es2015','react', 'stage-2'],
           }
          }
        ]
      }
    },
    plugins.copy()
  );

var config = null;
process.env.NODE_ENV = 'dev'
// Detect the branch where npm is running on
console.log(process.env.NODE_ENV)
switch(process.env.NODE_ENV) {
    case 'prod':
        config = merge(
            common,
            plugins.minify()
        );
        break;

    case 'dev':
    default:
        config = merge(
            common,
            {
              entry:{
                redux:['redux-logger']
              },
              devtool: 'source-map'
            }
        );
        break;
}
module.exports = config;
