<template>
    <router-view />
    <div class="">
        <input v-model="username" type="text" placeholder="Username" />
        <input v-model="email" type="text" placeholder="Email" />
        <input v-model="password" type="password" placeholder="Password" />
        <button @click="register">Register</button>
        <p v-if="error">Error: {{ error }}</p>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    methods: {
        register() {
            axios.post(`/api/register`, {
                    username: this.username,
                    email: this.email,
                    password: this.password,
                })
                .then(response => {
                    if (response.status == 200) {
                        this.$router.push("/");
                    }
                })
                .catch(error => {
                    console.log('error', error);
                    this.error = error.response.data.message;
                });
        }
    },
    data() {
        return {
            username: '',
            email: '',
            password: '',
            error: null,
        };
    },
}
</script>