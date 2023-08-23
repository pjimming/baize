<template>
    <div>
        <figure class="text-center">
            <blockquote class="blockquote h1">
                <h3>天呐🙀，页面📃被外星人👽的UFO🛸击落了，😵</h3>
            </blockquote>
            <figcaption class="blockquote-footer">
                <cite title="Source Title">404 | Page Not Found</cite>
            </figcaption>
        </figure>
        <img class="img-fluid rounded" :src='pageNotFoundImg'>
    </div>
</template>

    
<script>
import PNF404 from '@/assets/404.png';

export default {
    name: "PageNotFound",
    data() {
        return {
            pageNotFoundImg: PNF404,
        }
    },
}
</script>

    
<style scoped></style>
