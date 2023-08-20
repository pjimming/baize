<template>
<NavBar />
<router-view />
</template>

<script>
import NavBar from './components/NavBar.vue'

export default {
    name: 'app',
    components: {
        NavBar,
    }
}
</script>

<style>
body {
    background-image: url('./assets/background.jpg');
    background-size: cover;
}
</style>
