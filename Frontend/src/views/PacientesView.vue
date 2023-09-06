<template>
  <div>
    <div class="row">
      <div class="col-12">
        <div class="pacientes">
          <div class="info-pacientes flex justify-around items-start">
            <h3 class="titulo-pacientes">Pacientes</h3>
            <buttonPrimary>
              Adicionar
              <PhPlus class="icon-color"></PhPlus>
            </buttonPrimary>
          </div>
          <div class="q-mt-lg flex" style="gap: 1rem;">
            <div class="lista-pacientes q-pa-sm">
              <span>Paciente</span>
            </div>
            <div class="lista-pacientes q-pa-sm">
              <span>Paciente</span>
            </div>
            <div class="lista-pacientes q-pa-sm">
              <span>Paciente</span>
            </div>
            <div class="lista-pacientes q-pa-sm">
              <span>Paciente</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import buttonPrimary from '../components/ButtonPrimary.vue';
import { PhPlus } from '@phosphor-icons/vue';

export default {
  components: {
    buttonPrimary,
    PhPlus,
  },
};
</script>

<style scoped>
.pacientes {
  height: 89.5vh;
  background: #fff;
}

.titulo-pacientes {
  color: black;
  font-size: 3rem;
}

.lista-pacientes {
  border: 0.094rem solid black;
  width: 100%;
}

.lista-pacientes span {
  color: #000;
  font-weight: bold;
}

.lista-item {
  color: #000;
  font-weight: 500;
}
</style>