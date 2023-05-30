<template>
  <q-input class="input" label-slot color="primary" :dark="false">
    <template v-slot:prepend>
      <slot name="before-label"></slot>
    </template>

    <template v-slot:label>
      <span class="label">
        {{ label }}
      </span>
    </template>

    <template v-slot:append>
      <slot name="after-label"></slot>
    </template>
  </q-input>
</template>
<script>
export default {
  props: {
    label: String | Number,
    modelValue: String
  },
};
</script>
<style scoped>
.input {
  position: relative;
}

.input::before {
  content: "";
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
}

.label {
  color: var(--neutral-dark-gray);
}

.icon {
  color: var(--neutral-dark-gray);
}

.q-field--highlighted .icon,
.q-field--highlighted .label {
  color: var(--primary);
}
</style>
