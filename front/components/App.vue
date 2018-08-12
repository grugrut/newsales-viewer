<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title>New Sales Viewer</v-toolbar-title>
    </v-toolbar>
    <v-card>
      <v-container
          grid-list-lg>
        <div v-for="dateOfProduct in dateOfProducts">
          <h3>{{dateOfProduct.SaleDate.slice(0, 10)}} 発売</h3>
          <v-layout row wrap>
            <Product
                v-for="product in dateOfProduct.Products"
                :product="product"
            />
          </v-layout>
        </div>
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
             dateOfProducts: []
         }
     },
     components: {
         Product
     },
     created() {
         axios.get('/api/product')
              .then(response => {
                  this.dateOfProducts = response.data;
              });
     }
 }
</script>
