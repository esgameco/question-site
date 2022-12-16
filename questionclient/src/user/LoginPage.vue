<template>
    <router-view />
    <NavBar />

    <div class="">
        <input v-model="username" type="text" placeholder="Username" />
        <input v-model="password" type="password" placeholder="Password" />
        <button @click="login">Login</button>
        <p v-if="error">Error: {{ error }}</p>
    </div>
</template>

<script>
import axios from 'axios';

import NavBar from '../components/NavBar.vue';

export default {
    components: {
        NavBar,
    },
    methods: {
        login() {
            axios.post(`/api/login`, {
                    username: this.username,
                    password: this.password,
                })
                .then(response => {
                    if (response.status == 200) {
                        this.$router.push('/');
                    }
                })
                .catch(error => {
                    this.error = error.response.data.message;
                });
        }
    },
    mounted() {
    },
    data() {
        return {
            username: '',
            password: '',
            error: null,
        };
    },
}
</script>