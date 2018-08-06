<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title>New Sales Viewer</v-toolbar-title>
    </v-toolbar>
    <v-card>
      <v-container
          fluid
          grid-list-lg>
        <v-layout row wrap>
          <Product
              v-for="product in products"
              :product="product"
          />
        </v-layout>
      </v-container>
    </v-card>
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
