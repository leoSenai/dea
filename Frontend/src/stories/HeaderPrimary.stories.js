import HeaderPrimary from '../components/HeaderPrimary.vue'

export default {
  title: 'Example/HeaderPrimary',
  component: HeaderPrimary,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { HeaderPrimary },
    setup() {
      return { args };
    },
    template: '<HeaderPrimary />',
  }),
};