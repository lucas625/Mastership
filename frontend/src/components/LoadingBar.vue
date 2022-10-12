<template>
  <v-dialog
    v-model="isLoading"
    persistent
  >
    <v-card
      color="primary"
      dark
    >
      <v-card-title>
        {{ loadingMessage }}
      </v-card-title>
      <v-card-text>
        <v-progress-linear
          indeterminate
          color="white"
        />
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'LoadingBar',
  props: {
    isLoadingProps: {
      type: Boolean,
      required: true
    },
    loadingMessage: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      isLoading: this.isLoadingProps
    }
  },
  watch: {
    isLoadingProps (newIsLoadingProps) {
      this.isLoading = newIsLoadingProps
    }
  }
}
</script>
