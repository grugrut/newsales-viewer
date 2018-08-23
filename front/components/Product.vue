<template>
  <v-flex xs12 sm6 md3>
    <v-card>
      <v-card-media
          :src="product.ImgURL"
          height="200px"
      />
      <v-card-title primary-title>
        <div>
          <h3 class="headline">{{ product.Name }}</h3>
          <span>発売元: {{product.Publisher }}</span>
          <span>発売日: {{product.SaleDate.slice(0, 10)}}</span>
        </div>
      </v-card-title>
      <v-card-actions>
        <v-btn flat color="orange" :href="product.NewsURL" target="_blank">詳細</v-btn>
      </v-card-actions>
    </v-card>
  </v-flex>
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
