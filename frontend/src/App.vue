<template>
<NavBarVue />
<router-view />
</template>

<script>
import NavBarVue from './components/NavBar.vue'

export default {
    name: 'app',
    components: {
        NavBarVue,
    }
}
</script>

<style>
</style>
