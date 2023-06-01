import Button from '../components/Button.vue';
import { PhUser } from '@phosphor-icons/vue'

export default {
  title: 'Example/Button',
  component: Button,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { Button },
    setup() {
      return args;
    },
    template: '<Button v-bind="args">Label</Button>',
  }),
};

export const Outlined = {
  render: (args) => ({
    components: { Button },
    setup() {
      return args;
    },
    template: '<Button v-bind="args">Label</Button>',
  }),
  args: {
    outlined: true,
  },
};

export const RoundedFull = {
  render: (args) => ({
    components: { Button },
    setup() {
      return { args };
    },
    template: '<Button v-bind="args">Label</Button>',
  }),
  args: {
    roundedFull: true,
  },
};

export const OutlinedRoundedFull = {
  render: (args) => ({
    components: { Button },
    setup() {
      return { args };
    },
    template: '<Button v-bind="args">Label</Button>',
  }),
  args: {
    outlined: true,
    roundedFull: true,
  },
};

export const IconBefore = {
  render: (args) => ({
    components: { Button, PhUser },
    setup() {
      return { args };
    },
    template: `
    <Button v-bind="args">
      <template v-slot:before-label>
        <PhUser class="before-label" />
      </template>
      Label
    </Button>`,
  }),
  args: {
  },
}

export const IconAfter = {
  render: (args) => ({
    components: { Button, PhUser },
    setup() {
      return { args };
    },
    template: `
    <Button v-bind="args">
      <template v-slot:after-label>
        <PhUser class="after-label" />
      </template>
      Label
    </Button>`,
  }),
  args: {
  },
}


