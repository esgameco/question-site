<template>
    <div class="question">
        <div class="question-header">
            <h1>{{ title }}</h1>
            
            <p v-if="updatedDate.length > 0">Updated: {{ updatedDate }}</p>
        </div>
        <div class="question-body">
            <p>{{ body }}</p>
        </div>
        <div class="question-footer">
            <p>{{ author }} asked on {{ createdDate }}</p>
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
            title: 'Title',
            createdDate: '1/1/22',
            updatedDate: '',
            author: 'Author',
            body: 'Body',
            likes: 0,
        }
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