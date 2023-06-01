import Checkbox from '../components/Checkbox.vue'

export default {
  title: 'Example/Checkbox',
  component: Checkbox,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { Checkbox },
    setup() {
      return args;
    },
    template: '<Checkbox :label="label" v-model="value" />',
  }),
  args: {
    label: 'Label',
    value: true
  }
};
