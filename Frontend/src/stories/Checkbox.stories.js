import CheckboxPrimary from '../components/CheckboxPrimary.vue'

export default {
  title: 'Example/CheckboxPrimary',
  component: CheckboxPrimary,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { CheckboxPrimary },
    setup() {
      return args;
    },
    template: '<CheckboxPrimary :label="label" v-model="value" />',
  }),
  args: {
    label: 'Label',
    value: true
  }
};
