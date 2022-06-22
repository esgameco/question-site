<template>
    <p>Likes: {{ likes }}</p>
    <button @click="like">Like</button>
    <button @click="dislike">Dislike</button>
</template>

<script>
import axios from 'axios';

export default {
    props: ['likes_init', 'object_id'],
    methods: {
        like() {
            axios.post(`/api/like/${this.object_id}`)
                .then(response => {
                    this.likes = response.data.likes;
                })
                .catch(error => {
                    console.log(error);
                });
        },
        dislike() {
            axios.post(`/api/dislike/${this.object_id}`)
                .then(response => {
                    this.likes = response.data.likes;
                })
                .catch(error => {
                    console.log(error);
                });
        },
    },
    data() {
        return {
            likes: 0,
        };
    },
    mounted() {
        this.likes = this.likes_init;
    },
}
</script>