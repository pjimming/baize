<template>
    <NavBar />
    <router-view />
</template>

<script>
import NavBar from './components/NavBar.vue'

const debounce = (fn, delay) => {
  let timer = null;
  return function () {
    let context = this;
    let args = arguments;
    clearTimeout(timer);
    timer = setTimeout(function () {
      fn.apply(context, args);
    }, delay);
  }
}

const _ResizeObserver = window.ResizeObserver;
window.ResizeObserver = class ResizeObserver extends _ResizeObserver {
  constructor(callback) {
    callback = debounce(callback, 16);
    super(callback);
  }
}


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
    background-repeat: no-repeat;
    background-attachment: fixed
}
</style>
