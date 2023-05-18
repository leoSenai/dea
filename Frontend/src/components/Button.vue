<template>
  <button :class="[outlined ? 'outlined' : '', roundedFull ? 'rounded-full' : '']" type="button">
    <i v-if="iconLeft" :class="['ph', `ph-${iconLeft}`, 'icon-left']"></i>
    <span v-if="label">{{ label }}</span>
    <i v-if="icon" :class="['ph', `ph-${icon}`, 'icon-right']"></i>
</button>
</template>
<script>
export default {
  props: {
    label: String,
    icon: String,
    iconLeft: String,
    outlined: {
      type: Boolean,
      default: false
    },
    roundedFull: {
      type: Boolean,
      default: false
    },
  }
}
</script>
<style scoped>

button {
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  gap: .5rem;
  padding: 1rem 2rem;
  border: none;
  background: var(--primary);
}

button.rounded-full {
  border-radius: 99999px;
}

button.outlined {
  background: none;
  border: 2px solid var(--primary);
  color: var(--primary);
}

button:has(.icon-left) {
  padding-left: 1.5rem;
}

button:has(.icon-right) {
  padding-right: 1.5rem;
}

button:not(.outlined):disabled {
  cursor: not-allowed;
  background: var(--primary-300);
}

button.outlined:disabled {
  cursor: not-allowed;
  border: 2px solid var(--primary-300);
}

button:not(:disabled):hover {
  filter: opacity(0.8);
}

</style>