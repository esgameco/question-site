<template>
    <div class="answer">
        <div class="answer-body">
            <p>{{ body }}</p>
        </div>
        <div class="question-footer">
            <p>{{ author }} answered on {{ createdDate }}</p>
            <LikeComponent :likes="likes" :id="ID" />
        </div>
    </div>
</template>

<script>
import LikeComponent from './LikeComponent.vue'

export default {
    props: ['answer'],
    components: {
        LikeComponent,
    },
    data() {
        return {
            ID: 0,
            body: '',
            updatedDate: '',
            createdDate: '',
            author: '',
            likes: 0,
        };
    },
    mounted() {
        this.ID = this.question.ID;
        this.body = this.question.body;
        this.updatedDate = this.question.updatedDate;
        this.createdDate = this.question.createdDate;
        this.author = this.question.author;
        if (this.question.likes) {
            this.likes = this.question.likes.amount;
        }
    },
}
</script>

<style lang="css" scoped>
.answer {
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    margin-bottom: 20px;
    max-width: 40%;
}

.answer-body {
    border-bottom: 1px solid #ccc;
    padding-bottom: 10px;
}

.answer-footer {
    padding-top: 10px;
}
</style>