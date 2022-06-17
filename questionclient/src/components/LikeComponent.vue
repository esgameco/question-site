<template>
    <p>Likes: {{ likes }}</p>
    <button @click="like">Like</button>
    <button @click="dislike">Dislike</button>
</template>

<script>
export default {
    methods: {
        like() {
            this.likes++;
        },
        dislike() {
            this.likes--;
        },
    },
    data() {
        return {
            likes: 0,
        }
    },
}
</script>