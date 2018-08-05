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
 export default {
     components: {
         Product
     },
     created: function () {
         
     }
 }
</script>
