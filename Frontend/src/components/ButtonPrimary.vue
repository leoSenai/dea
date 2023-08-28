<template>
  <button
    :class="[outlined ? 'outlined' : '', roundedFull ? 'rounded-full' : '']"
    :disabled="disabled"
    type="button"
  >
    <slot
      name="before-label"
      class="before-label"
    />
    <slot />
    <slot
      name="after-label"
      class="after-label"
    />
  </button>
</template>
<script>
export default {
  props: {
    outlined: {
      type: Boolean,
      default: false,
    },
    roundedFull: {
      type: Boolean,
      default: false,
    },
    disabled: {
      type: Boolean,
      default: false
    },
    size: {
      type: String,
      default: 'md',
    }
  },
};
</script>
<style scoped>
button {
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  gap: 0.5rem;
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

button:has(.before-label) {
  padding-left: 1.5rem;
}

button:has(.after-label) {
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

button[size=xs] {
  padding: .5rem 1.5rem;
  font-size: .5rem;
}

button[size=sm] {
  padding: .75rem 1.75rem;
  font-size: .75rem;
}

button[size=lg] {
  padding: 1.25rem 2.25rem;
  font-size: 1.25rem;
}

button[size=xl] {
  padding: 1.5rem 2.5rem;
  font-size: 1.5rem;
}
</style>
