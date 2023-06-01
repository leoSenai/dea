import Select from '../components/Select.vue'

export default {
  title: 'Example/Select',
  component: Select,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { Select },
    setup() {
      return { args };
    },
    template: '<Select label="Label" />',
  }),
  args: {
    options: ['Teste 1', 'Teste 2'],
    modelValue: '',
    value: ''
  }
};