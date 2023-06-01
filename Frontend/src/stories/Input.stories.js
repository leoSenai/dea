import Input from '../components/Input.vue'

export default {
  title: 'Example/Input',
  component: Input,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { Input },
    setup () {
      return { args }
    },
    template: '<Input label="Input" />',
  }),
};