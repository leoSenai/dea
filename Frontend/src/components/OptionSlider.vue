<template>
  <div class="option-slider-container">
    <input
    type="checkbox"
    class="option-slider"
    ref="input"
    @input="$emit('update:modelValue', $event.target.checked)"
    :checked="modelValue"
    />
    <label v-if="label">{{ label }}</label>
  </div>
</template>
<script>

export default {
  props: {
    modelValue: Boolean,
    label: String,
  },
  emits: ['update:modelValue'],
}
</script>
<style scoped>
.option-slider-container {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: .5rem;
  padding: 1rem;
}

.option-slider {
  position: relative;
  width: 2.5rem;
  height: 1.33rem;
  height: auto;
  transition: 1s ease-in-out;
}

.option-slider::after,
.option-slider::before {
  content: "";
  position: absolute;
  border-radius: 9999px;
  transition: 300ms ease-in-out;
}

.option-slider::before {
  height: 1.33rem;
  width: 2.5rem;
  background: var(--neutral-light-gray);
  transform: translateY(-50%);
  top: 50%;
  left: 0;
}

.option-slider::after {
  height: 1rem;
  width: 1rem;
  background: var(--neutral-dark-gray);
  transform: translate(0%, -50%);
  top: 50%;
  margin-left: 10%;
}

.option-slider:checked::after {
  background: var(--primary);
  margin-left: 50%;
}

label {
  color: var(--neutral-black);
}

@media (prefers-color-scheme: dark) {
    label {
        color: var(--neutral-white);
    }

    .option-slider::before {
      background: var(--neutral-gray);
    }
}

</style>
