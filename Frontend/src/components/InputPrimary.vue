<template>
  <q-input
    v-model="inputValue"
    class="input"
    label-slot
    color="primary"
    :outlined="outlined"
    :dark="dark"
  >
    <template #prepend>
      <slot name="before-label" />
    </template>

    <template #label>
      <span
        class="label"
        :style="`color:${labelColor};`"
      >
        {{ label }}
      </span>
    </template>

    <template #append>
      <slot name="after-label" />
    </template>
  </q-input>
</template>
<script>
export default {
  data() {
    return {
      inputValue: '',
    }
  },  
  props: {
    label: {
      type: String,
      default: '0'
    },
    modelValue: {
      type: String,
      default: ''
    },
    outlined: {
      type: Boolean,
      default: true
    },
    dark: {
      type: Boolean,
      default: false
    },
    labelColor: {
      type: String,
      default: 'white'
    }
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
