<template>
  <div>
    <Product
        v-for="product in products"
        :product="product"
    />
  </div>
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
