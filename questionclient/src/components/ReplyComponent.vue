<template>
    <router-view />
    <NavBar />
    <!-- Reply input -->
    <div class="reply-input">
        <input v-model="reply" type="text" placeholder="Reply" />
        <button @click="postReply">Post</button>
    </div>
    <p>{{ posted }}</p>
</template>

<script>
import NavBar from '../components/NavBar.vue';

export default {
    components: {
        NavBar,
    },
    methods: {
        postReply() {
            axios.post(`/api/questions/${this.$route.params.id}/answers`, {
                    reply: this.reply,
                })
                .then(response => {
                    if (response.status == 200) {
                        this.posted = "Reply posted!";
                    }
                })
                .catch(error => {
                    this.posted = "Error in posting reply.";
                    console.log('error', error);
                });
        }
    },
    data() {
        return {
            reply: '',
            posted: '',
        };
    },
}
</script>