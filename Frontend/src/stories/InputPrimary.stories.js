import InputPrimary from '../components/InputPrimary.vue'

export default {
  title: 'Example/InputPrimary',
  component: InputPrimary,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { InputPrimary },
    setup () {
      return { args }
    },
    template: '<InputPrimary label="InputPrimary" />',
  }),
};