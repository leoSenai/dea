<template>
    <header class="bg-primary-700">
        <div class="logo">
            <q-img :src="LogoSrc" />
            <h6>Clínica Motivar</h6>
        </div>
        <div class="user">
            <ph-user-circle />
            <span class="text-body">Teste</span>    
        </div>
    </header>
</template>

<script>
import { PhUserCircle } from '@phosphor-icons/vue'

export default {
    data () {
        return {
            LogoSrc: 'src/assets/imgs/Logo.png'
        }
    },
}
</script>

<style scoped>
header {
    position: sticky;
    min-width: 100%;

    display: flex;
    align-items: center;
    justify-content: space-between;

    padding: .5rem 1rem;
}

.logo {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 2rem;
}

.q-img {
    height: 100%;
    width: 4rem;
}
</style>