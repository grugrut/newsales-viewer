module.exports = {
    entry: './front/index.js',
    devtool: 'source-map',
    output: {
        filename: 'bundle.js',
        path: `${__dirname}/app`,
    },
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.esm.js'
        }
    }
};
