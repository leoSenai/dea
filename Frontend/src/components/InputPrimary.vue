<template>
  <q-input
    v-model="model"
    class="input"
    label-slot
    color="primary"
    :outlined="outlined"
    :dark="dark"
    :rules="rulesComputed"
    @update:model-value="(current) => $emit('update:modelValue', current)"
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

    <template #error>
      <span>
        aaa
      </span>
    </template>
  </q-input>
</template>
<script>
export default {
  props: {
    label: {
      type: String,
      default: '0',
    },
    modelValue: {
      type: String,
      default: '',
    },
    outlined: {
      type: Boolean,
      default: true,
    },
    dark: {
      type: Boolean,
      default: false,
    },
    labelColor: {
      type: String,
      default: 'white',
    },
    required: {
      type: Boolean,
      default: false,
    },
    rules: {
      type: Array,
      default: () => [],
    },
  },
  emits: ['update:modelValue'],
  data() {
    return {
      model: this.modelValue,
    };
  },
  computed: {
    rulesComputed() {
      return [
        ...this.rules,
        () => {
          if (this.required && this.model.length <= 0) {
            return 'Campo requirido!'
          }
          return null
        },
      ];
    },
  },
};
</script>
<style>
.input {
  position: relative;
  width: 100%;
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
