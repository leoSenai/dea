<template>
  <q-input @focus="teste" ref="input">
    <template v-slot:prepend v-if="icon && iconLeft">
      <i :class="['ph', `ph-${icon}`]"></i>
    </template>

    <template v-slot:append v-if="icon && !iconLeft">
      <i :class="['ph', `ph-${icon}`]"></i>
    </template>

    <template v-slot:hint v-if="caption">
      {{ caption }}
    </template>
  </q-input>
</template>
<script>
export default {
  props: {
    icon: String,
    iconLeft: {
      type: Boolean,
      default: false,
    },
    caption: String,
  },
  computed: {
    focused() {
      console.log(this.$refs.input.focus);
      return this.$refs.input.focus;
    },
  },
  methods: {
    teste() {
      console.log(this.$refs.input);
    },
  },
};
</script>
