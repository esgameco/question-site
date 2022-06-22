<template>
    <QuestionComponent :question="question" />
    
    <li v-for="answer in question.answers" :key="answer.ID">
        <AnswerComponent :answer="answer" />
    </li>
</template>

<script>
import QuestionComponent from '../components/QuestionComponent.vue';
import AnswerComponent from '../components/AnswerComponent.vue';
import axios from 'axios';

export default {
    components: {
        QuestionComponent,
        AnswerComponent,
    },
    methods: {
        getQuestion(questionID) {
            axios.get(`/api/questions/${questionID}`)
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