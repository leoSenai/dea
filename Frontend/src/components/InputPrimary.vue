<template>
  <q-input
    v-model="model"
    class="input"
    label-slot
    color="primary"
    :outlined="outlined"
    :dark="dark"
    :mask="mask"
    :hint="hint"
    :rules="rulesComputed"
    :type="isPwd ? 'password' : 'text'"
    :mask="mask"
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
      <q-icon
        v-if="password"
        :name="isPwd ? 'visibility_off' : 'visibility'"
        class="cursor-pointer"
        @click="isPwd = !isPwd"
      />
      <slot name="after-label" />
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
    mask: {
      type: String,
      default: '',
    },
    hint: {
      type: String,
      default: '',
    },
    password: {
      type: Boolean,
      default: false,
    },
    format: {
      type: String,
      default: ''
    }
  },
  emits: ['update:modelValue'],
  data() {
    return {
      model: this.modelValue,
      isPwd: this.password,
    }
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
    mask() {
      let mask;
      switch (this.format) {
        case 'cpf':
          mask = '###.###.###-##'
          break
        case 'cnpj':
          mask = '##.###.###/####-##'
          break;
        case 'phone':
          mask = '(##) #####-####'
          break;
        default:
          mask = ''
          break;
      }
      return mask
    }
  },
  watch: {
    modelValue(newValue) {
      this.model = newValue;
    },
    password(newValue) {
      this.isPwd = newValue;
    }
  },
  watch: {
    modelValue() {
      this.model = this.modelValue
    }
  }
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
