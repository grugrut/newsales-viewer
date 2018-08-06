<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title>New Sales Viewer</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <Product
          v-for="product in products"
          :product="product"
      />
    </v-content>
  </v-app>
</template>

<script>
 import Product from './Product.vue'
 import axios from 'axios'
 export default {
     data() {
         return {
             products: []
         }
     },
     components: {
         Product
     },
     created() {
         axios.get('/api/product')
              .then(response => {
                  this.products = response.data;
              });
     }
 }
</script>
