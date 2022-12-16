<template>
    <div class="question">
        <div class="question-header">
            <h1>{{ title }}</h1>
            
            <p v-if="updatedDate && updatedDate.length > 0">Updated: {{ updatedDate }}</p>
        </div>
        <div class="question-body">
            <p>{{ body }}</p>
        </div>
        <div class="question-footer">
            <p>{{ author.username }} asked on {{ createdDate }}</p>
            <LikeComponent :likes="likes" :id="ID" />
        </div>
    </div>
</template>

<script>
import LikeComponent from './LikeComponent.vue'

export default {
    props: ['question'],
    watch: {
        question: function(newQuestion) {
            this.title = newQuestion.title;
            this.body = newQuestion.body;
            this.author = newQuestion.author;
            this.createdDate = newQuestion.CreatedAt.substring(0, newQuestion.CreatedAt.indexOf("T"));
            this.updatedDate = newQuestion.UpdatedAt.substring(0, newQuestion.UpdatedAt.indexOf("T"));
            if (this.question.likes) {
                this.likes = this.question.likes.amount;
            }
            this.ID = newQuestion.ID;
        }
    },
    components: {
        LikeComponent,
    },
    data() {
        return {
            ID: 0,
            title: '',
            body: '',
            updatedDate: '',
            createdDate: '',
            author: {username: ''},
            likes: 0,
        };
    },
}
</script>

<style lang="css" scoped>
.question {
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    margin-bottom: 20px;
    max-width: 40%;
}

.question-header {
    border-bottom: 1px solid #ccc;
    padding-bottom: 10px;
}

.question-body {
    border-bottom: 1px solid #ccc;
    padding: 10px 0px;
}

.question-footer {
    padding-top: 10px;
}
</style>