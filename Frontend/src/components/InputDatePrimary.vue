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
      @update:model-value="updateModel"
    >
      <template #after-label>
        <PhCalendarBlank
          class="calendar-icon"
          :size="30"
          @click="() => showDatePicker = !showDatePicker"
        />
      </template>
    </InputPrimary>
    <div
      v-show="showDatePicker"
      class="date-modal"
      @blur="showDatePicker = false"
    >
      <q-date
        v-model="model"
        :locale="locale"
        mask="DD/MM/YYYY"
        @update:model-value="updateModel"
      />
    </div>
  </div>
</template>
<script>
import InputPrimary from './InputPrimary.vue';
import { PhCalendarBlank } from '@phosphor-icons/vue'

export default {
  components: {
    InputPrimary,
    PhCalendarBlank
  },
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
  methods: {
    updateModel(current) {
      this.model = current
      this.$emit('update:modelValue', this.model)
      if (current.length === 10 || current.length === 0) {
        this.showDatePicker = false
      } else {
        this.showDatePicker = true
      }
    }
  },
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

.q-field--highlighted .calendar-icon,
.q-field--highlighted .label {
  color: var(--primary);
}

.date-modal {
  position: absolute;
  top: -10%;
  left: 105%;
  z-index: 1;
}

.calendar-icon {
  background: white;
  border-radius: 99999px;
  padding: .25rem;
  transition: .2s;
  cursor: pointer;
}

.calendar-icon:hover {
  filter: brightness(0.8);
}
</style>
