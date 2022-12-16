<template>
    <div class="question">
        <div class="question-header">
            <h2>Create Question</h2>
        </div>
        <div class="question-body">
            Title: <input type="text" v-model="title" placeholder="Title" />
            Body: <input type="textarea" v-model="body" placeholder="Body" />
        </div>
        <div class="question-footer">
            <button @click="submit">Submit</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    methods: {
        submit() {
            axios.post(`/api/questions/create`, {
                    title: this.title,
                    body: this.body
                })
                .then(response => {
                    console.log('Created question: ', response)
                    if (response.status == 200) {
                        this.$router.push(`/questions/${response.data.ID}`);
                    }
                })
                .catch(error => {
                    console.log(error);
                });
        },
    },
    data() {
        return {
            title: '',
            body: ''
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