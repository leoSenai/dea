<template>
  <q-input @focus="teste" ref="input">
    <template v-slot:before></template>

    <template v-slot:prepend v-if="icon && !iconLeft">
      <i :class="['ph', `ph-${icon}`]"></i>
    </template>

    <template v-slot:append v-if="icon && !iconLeft">
      <i :class="['ph', `ph-${icon}`]"></i>
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
  },
};
</script>
