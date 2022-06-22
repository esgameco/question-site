<template>
    <nav>
        <ul>
            <li><router-link to="/">Home</router-link></li>
            <li><router-link to="/login">Login</router-link></li>
            <li><router-link to="/register">Register</router-link></li>
        </ul>
    </nav>
</template>

<style scoped>
ul {
    display: flex;
    justify-content: space-around;
}
</style>