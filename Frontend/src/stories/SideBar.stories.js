import SideBar from '../components/SideBar.vue'
import { PhClipboardText, PhUsers } from '@phosphor-icons/vue'

export default {
  title: 'Example/SideBar',
  component: SideBar,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { SideBar },
    setup() {
      return { args };
    },
    template: '<SideBar v-bind="args" />',
  }),
  args: {
    links: [
      { path: 'anamnese', name: 'Anamnese', icon: PhClipboardText },
      { path: 'usuarios', name: 'Médicos', icon: PhUsers }
    ]
  }
};