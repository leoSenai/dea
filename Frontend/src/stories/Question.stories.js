import Question from '../components/Question.vue'

export default {
  title: 'Example/Question',
  component: Question,
  tags: ['autodocs'],
  argTypes: {
    onClick: {},
  },
};

export const Primary = {
  render: (args) => ({
    components: { Question },
    setup() {
      return { args };
    },
    template: '<Question v-bind="args">There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomise</Question>',
  }),
  args: {
    modelValue: 0,
    label: 'Label',
    questionNumber: 5,
    isAnswered: false,
    answerRange: 5
  }
};

export const Answered = {
  render: (args) => ({
    components: { Question },
    setup() {
      return { args };
    },
    template: '<Question v-bind="args">There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomise</Question>',
  }),
  args: {
    modelValue: 5,
    label: 'Label',
    questionNumber: 5,
    isAnswered: true,
    answerRange: 5
  }
};