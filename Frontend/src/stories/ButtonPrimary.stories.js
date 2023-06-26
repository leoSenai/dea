import ButtonPrimary from '../components/ButtonPrimary.vue';
import { PhUser } from '@phosphor-icons/vue'

export default {
  title: 'Example/ButtonPrimary',
  component: ButtonPrimary,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { ButtonPrimary },
    setup() {
      return args;
    },
    template: '<ButtonPrimary v-bind="args">Label</ButtonPrimary>',
  }),
};

export const Outlined = {
  render: (args) => ({
    components: { ButtonPrimary },
    setup() {
      return args;
    },
    template: '<ButtonPrimary v-bind="args">Label</ButtonPrimary>',
  }),
  args: {
    outlined: true,
  },
};

export const RoundedFull = {
  render: (args) => ({
    components: { ButtonPrimary },
    setup() {
      return { args };
    },
    template: '<ButtonPrimary v-bind="args">Label</ButtonPrimary>',
  }),
  args: {
    roundedFull: true,
  },
};

export const OutlinedRoundedFull = {
  render: (args) => ({
    components: { ButtonPrimary },
    setup() {
      return { args };
    },
    template: '<ButtonPrimary v-bind="args">Label</ButtonPrimary>',
  }),
  args: {
    outlined: true,
    roundedFull: true,
  },
};

export const IconBefore = {
  render: (args) => ({
    components: { ButtonPrimary, PhUser },
    setup() {
      return { args };
    },
    template: `
    <ButtonPrimary v-bind="args">
      <template v-slot:before-label>
        <PhUser class="before-label" />
      </template>
      Label
    </ButtonPrimary>`,
  }),
  args: {
  },
}

export const IconAfter = {
  render: (args) => ({
    components: { ButtonPrimary, PhUser },
    setup() {
      return { args };
    },
    template: `
    <ButtonPrimary v-bind="args">
      <template v-slot:after-label>
        <PhUser class="after-label" />
      </template>
      Label
    </ButtonPrimary>`,
  }),
  args: {
  },
}


