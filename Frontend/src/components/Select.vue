<template>
    <q-select class="select-input" label-slot hideDropdownIcon ref="select" :options="options" :behavior="$q.platform.is.ios === true ? 'dialog' : 'menu'" borderless>
      <template v-slot:label>
        <span class="label">
          {{ label }}
        </span>
      </template>
      <template v-slot:append>
        <PhCaretDown class="icon" />
      </template>
      <template v-slot:option="option">
        <q-item class="option" v-bind="option.itemProps"
        ><p :class="option.selected ? 'selected' : ''">
          {{ option.label }}
        </p></q-item
        >
      </template>
    </q-select>
</template>
<script>
import { PhCaretDown } from '@phosphor-icons/vue'

export default {
  props: {
    label: String | Number,
    options: Array,
  },
  components: {
    PhCaretDown
  }
};
</script>
<style scoped>
.select-input {
  position: relative;
}

.select-input::before {
  content: "";
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  border: 1px solid var(--neutral-dark-gray);
  border-radius: 4px;
}

.select-input:hover{
  opacity: .8;
}

.label {
  margin-left: 1rem;
}

.icon {
  margin-right: 1rem;
}

.select-input:has([aria-expanded="true"]) .icon {
  transform: rotate(180deg);
  color: var(--primary);
}

.select-input:has([aria-expanded="true"]) .label {
  color: var(--primary);
}

.select-input:has([aria-expanded="true"])::before {
  border-color: var(--primary);
}

.icon {
  transition: 200ms;
  color: var(--neutral-dark-gray);
}

.option {
  color: var(--neutral-black);
  display: flex;
  align-items: center;
  justify-content: flex-start;
}

.option.selected {
  color: var(--primary);
}

:has(.option) {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  text-align: center;
}

:has(.option):hover {
  background: rgba(0, 0, 0, 0.1);
}
.label {
  color: var(--neutral-dark-gray);
}

@media (prefers-color-scheme: dark) {
  .select-input {
    color: var(--neutral-white);
  }
  .option {
    color: var(--neutral-white);
    background: rgba(0, 0, 0, 0.8);
  }
}
</style>
