<template>
  <div class="input-date-container">
    <InputPrimary
      v-model="model"
      class="input"
      :label="label"
      label-color="primary"
      color="primary"
      :outlined="outlined"
      :dark="dark"
      :rules="rulesComputed"
      mask="##/##/####"
      @update:model-value="(current) => $emit('update:modelValue', current)"
      @focus="() => showDatePicker = true"
    />
    <div
      v-show="showDatePicker"
      class="date-modal"
    >
      <q-date
        v-model="model"
        :locale="locale"
        mask="DD/MM/YYYY"
        @update:model-value="() => showDatePicker = false"
      />
    </div>
  </div>
</template>
<script>
import InputPrimary from './InputPrimary.vue';

export default {
  props: {
    label: {
      type: String,
      default: '',
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
      showDatePicker: false
    };
  },
  computed: {
    rulesComputed() {
      return [
        ...this.rules,
        () => {
          if (this.required && this.model.length <= 0) {
            return 'Campo requirido!';
          }
          return null;
        },
      ];
    },
    locale() {
      return {
        months: [
          'Janeiro',
          'Fevereiro',
          'Março',
          'Abril',
          'Maio',
          'Junho',
          'Julho',
          'Agosto',
          'Setembro',
          'Outubro',
          'Novembro',
          'Dezembro',
        ],
        monthsShort: [
          'Jan',
          'Fev',
          'Mar',
          'Abr',
          'Mai',
          'Jun',
          'Jul',
          'Ago',
          'Set',
          'Out',
          'Nov',
          'Dez'
        ],
        daysShort: [
          'Seg',
          'Ter',
          'Qua',
          'Qui',
          'Sex',
          'Sáb',
          'Dom'
        ]
      };
    },
  },
  watch: {
    modelValue() {
      if (this.modelValue.length === 10) {
        this.showDatePicker = false
      }
    }
  },
  methods: {
    toggleShowDatePicker(condition) {
      if (typeof condition !== 'boolean') {
        this.showDatePicker = !this.showDatePicker;
      }
      else {
        this.showDatePicker = condition;
      }
    },
    navi(a) {
      console.log(a);
    }
  },
  components: { InputPrimary }
};
</script>
<style scoped>
.input-date-container {
  position: relative;
}

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

.date-modal {
  position: absolute;
  top: -10%;
  left: 105%;
}
</style>
