<template>
  <v-flex xs4>
    <v-card>
      <v-card-media
          :src="product.ImgURL"
          height="400px"
      />
      <v-card-title primary-title>
        <div>
          <h3 class="headline">{{ product.Name }}</h3>
        </div>
      </v-card-title>
      <v-card-actions>
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
