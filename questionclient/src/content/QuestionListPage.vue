<template>
    <router-view />
    <NavBar />

    <ul v-for="question in questions" :key="question.ID">
        <QuestionComponent :question="question" />
        
        <li v-for="answer in question.answers" :key="answer.ID">
            <AnswerComponent :answer="answer" />
        </li>
    </ul>

    <FooterComponent />
</template>

<script>
import axios from 'axios';

import QuestionComponent from '../components/QuestionComponent.vue';
import AnswerComponent from '../components/AnswerComponent.vue';
import NavBar from '../components/NavBar.vue';
import FooterComponent from '../components/FooterComponent.vue';

export default {
    components: {
        QuestionComponent,
        AnswerComponent,
        NavBar,
        FooterComponent
    },
    methods: {
        getQuestionPage(sortType, page, numPerPage) {
            axios.get(`/api/questions`, {
                    sortType: sortType,
                    page: page,
                    numPerPage: numPerPage
                })
                .then(response => {
                    this.questions = response.data;
                })
                .catch(error => {
                    console.log(error);
                });
        }
    },
    mounted() {
        this.getQuestionPage(this.$route.params.sortType, 
                             this.$route.params.page, 
                             this.$route.params.numPerPage);
    },
    data() {
        return {
            questions: [],
        };
    },
}
</script>