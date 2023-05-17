<template>
  <q-input class="input" @focus="teste" label-slot outlined>
    <template v-slot:prepend v-if="icon && iconLeft">
      <i :class="['ph', `ph-${icon}`]"></i>
    </template>

    <template v-slot:label>
      <span class="label">
        {{ label }}
      </span>
    </template>

    <template v-slot:append v-if="icon && !iconLeft">
      <i :class="['ph', `ph-${icon}`, 'icon']"></i>
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
    label: String | Number
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
