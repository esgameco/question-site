<template>
    <div class="answer">
        <div class="answer-body">
            <p>{{ body }}</p>
        </div>
        <div class="question-footer">
            <p>{{ author }} answered on {{ createdDate }}</p>
            <LikeComponent />
        </div>
    </div>
</template>

<script>
import LikeComponent from './LikeComponent.vue'

export default {
    components: {
        LikeComponent,
    },
    data() {
        return {
            body: 'I agree',
            author: 'Author',
            createdDate: '7/11/22',
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