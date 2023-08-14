import SelectPrimary from '../components/SelectPrimary.vue'

export default {
  title: 'Example/SelectPrimary',
  component: SelectPrimary,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { SelectPrimary },
    setup() {
      return { args };
    },
    template: '<SelectPrimary label="Label" />',
  }),
  args: {
    options: ['Teste 1', 'Teste 2'],
    modelValue: '',
    value: ''
  }
};