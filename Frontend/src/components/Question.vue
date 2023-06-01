<template>
  <div
    :id="`question-${questionNumber}`"
    class="question-container"
  >
    <div class="question-title">
      <span class="heading-6">Questão
        {{ questionNumber < 10 ? "0" + questionNumber : questionNumber }}:</span>
    </div>
    <div class="question">
      <p>
        <slot />
      </p>
    </div>
    <div class="question-answer">
      <template v-if="isAnswered">
        <p class="heading-subtitle">
          Resposta: {{ modelValue < 10 ? "0" + modelValue : modelValue }}
        </p>
      </template>
      <template v-else>
        <div class="disagree text-caption">
          Discordo
        </div>
        <div class="question-range">
          <template
            v-for="answer in answers"
            :key="answer.value"
          >
            <div class="answer-container">
              <input
                :id="`question-${questionNumber}-option-${answer.value}`"
                ref="option"
                :value="answer.value"
                type="radio"
                name="option-question"
                style="display: none"
                :checked="modelValue === answer.value"
                @input="$emit('update:modelValue', answer.value)"
              >
              <label
                :for="`question-${questionNumber}-option-${answer.value}`"
                class="icon"
                :style="{
                  transform: `scale(${answer.scale})`,
                  background: modelValue === answer.value ? answer.color : '',
                  borderColor: modelValue === answer.value ? answer.color : '',
                }"
              >
                <PhCheck v-if="answer.value === modelValue" />
              </label>
            </div>
          </template>
        </div>
        <div class="agree text-caption">
          Concordo
        </div>
      </template>
    </div>
  </div>
</template>
<script>
import { PhCheck } from '@phosphor-icons/vue'
export default {
  components: {
    PhCheck
  },
  props: {
    modelValue: {
      type: Number,
      default: 0
    },
    questionNumber: {
      type: Number,
      required: true,
    },
    isAnswered: {
      type: Boolean,
      required: true,
      default: false,
    },
    answerRange: {
      type: Number,
      default: 5,
    },
  },
  emits: ['update:modelValue'],
  computed: {
    scaleMax() {
      return this.answerRange;
    },
    scaleMin() {
      return 1;
    },
    answers() {
      const arr = new Array(this.answerRange).fill(1).map((el, i) => el + i);
      const isEven = this.answerRange % 2 === 0;

      const middleArray = Math.ceil(arr.length / 2);
      const middleValues = isEven
        ? arr.slice(middleArray - 1, middleArray + 1)
        : arr.slice(middleArray - 1, middleArray);

      const colors = new Array(this.answerRange);
      const scales = new Array(this.answerRange).fill(0);

      let j = middleArray;
      let k = 2;
      for (let i = 0; i < arr.length; i++) {
        if (i >= middleArray) {
          scales[i] = k;
          colors[i] = `hsl(102, ${i * 10}%, 40%)`;
          k++;
          if (middleValues[1] - 1 === i && middleValues.length == 2) {
            scales[i] = 1;
            colors[i] = '#656565';
          }
        } else {
          scales[i] = j;
          colors[i] = `hsl(0, ${40 / (i + 1)}%, 46%)`;
          j--;
          if (middleValues[0] - 1 === i) {
            scales[i] = 1;
            colors[i] = '#656565';
          }
          if (i === 0) colors[i] = 'hsl(0, 40%, 46%)';
        }
      }

      return arr.map((el, i) => {
        return {
          scale: scales[i] / 10 + 0.9,
          value: el,
          color: colors[i],
        };
      });
    },
  },
};
</script>
<style scoped>
.question-container {
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  flex-direction: column;
  padding: 1rem 1.5rem;

  border-radius: 8px;
  background: var(--primary-700);
}

.question-answer {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.question-answer .question-range {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-evenly;
}

.icon {
  display: flex;
  width: 1.75rem;
  height: 1.75rem;
  text-align: center;
  padding: 0.5rem;
  border-radius: 99999px;
  border: 1px solid;
  align-items: center;
  justify-content: center;
  transition: 0.2s ease;
}
</style>
