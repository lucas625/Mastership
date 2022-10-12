<template>
  <v-alert
      dense
      :type="messageType"
      :value="active"
  >
    {{ message }}
  </v-alert>
</template>

<script>
export default {
  name: 'Alert',
  props: {
    active: {
      type: Boolean,
      required: true
    },
    message: {
      type: String,
      required: true
    },
    messageType: {
      type: String,
      required: true
    }
  },
  watch: {
    active (newActive) {
      if (newActive) {
        setTimeout(()=>{
          this.$emit('close')
        },5000)
      }
    }
  }
}
</script>
