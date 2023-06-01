import OptionSlider from '../components/OptionSlider.vue'

export default {
  title: 'Example/OptionSlider',
  component: OptionSlider,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { OptionSlider },
    setup() {
      return { args };
    },
    template: '<OptionSlider v-bind="args" />',
  }),
  args: {
    modelValue: '',
    label: 'Label'
  }
};