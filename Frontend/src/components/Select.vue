<template>
    <q-select hideDropdownIcon ref="select" @click="clickSelect" @blur="blurSelect">
        <template v-slot:append>
            <i class="ph ph-caret-down" ref="icon" style="transition: 200ms;"></i>
        </template>
    </q-select>
</template>
<script>
export default {
    props: {

    },
    computed: {
        active () {
            return this.$refs.select
        }
    },
    methods: {
        clickSelect () {
            if (this.$refs.icon.style.transform === "rotate(180deg)" ) {
                this.$refs.icon.style.transform = "rotate(0deg)"
            } else {
                this.$refs.icon.style.transform = "rotate(180deg)"
            }
        },
        blurSelect () {
            this.$refs.icon.style.transform = "rotate(0deg)"
        }    
    }
}
</script>