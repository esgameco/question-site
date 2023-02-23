<template>
    <router-view />
    <NavBar />

    <QuestionComponent :question="question" />
    
    <li v-for="answer in question.answers" :key="answer.ID">
        <AnswerComponent :answer="answer" />
    </li>
</template>

<script>
import axios from 'axios';

import QuestionComponent from '../components/QuestionComponent.vue';
import AnswerComponent from '../components/AnswerComponent.vue';
import NavBar from '../components/NavBar.vue';

export default {
    components: {
        QuestionComponent,
        AnswerComponent,
        NavBar
    },
    methods: {
        getQuestion(questionID) {
            axios.get(`/api/question/${questionID}`)
                .then(response => {
                    this.question = response.data;
                })
                .catch(error => {
                    console.log(error);
                });
        }
    },
    mounted() {
        this.getQuestion(this.$route.params.id);
    },
    data() {
        return {
            question: {},
        };
    },
}
</script>