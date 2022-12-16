<template>
    <router-view />
    <NavBar />

    <CreateQuestionComponent />
</template>

<script>
import NavBar from '../components/NavBar.vue';
import CreateQuestionComponent from '../components/CreateQuestionComponent.vue';

export default {
    components: {
        NavBar,
        CreateQuestionComponent
    },
};
</script>