<template>
  <div class="product">
    <span>{{product.Name}}</span>
  </div>
</template>

<script>
  export default {
      props: {
          product: {
              type: Object,
              required: true
          }
      }
  }
</script>
