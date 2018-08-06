const VueLoaderPlugin = require('vue-loader/lib/plugin');

module.exports = {
    entry: './front/index.js',
    devtool: 'source-map',
    output: {
        filename: 'bundle.js',
        path: `${__dirname}/app`,
    },
    module: {
        rules: [
            {test: /\.vue$/, loader: 'vue-loader'},
            {test: /\.css$/, loader: ['style-loader', 'css-loader']},
        ]
    },
    plugins: [
        new VueLoaderPlugin()
    ]

};
