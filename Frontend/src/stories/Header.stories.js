import Header from '../components/Header.vue'

export default {
  title: 'Example/Header',
  component: Header,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { Header },
    setup() {
      return { args };
    },
    template: '<Header />',
  }),
};