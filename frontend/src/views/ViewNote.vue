<template>
<div class="container">
    <div v-if="error">
        <p>{{ error }}</p>
    </div>
    <div v-else>
        <p>{{ note.content }}</p>
    </div>
</div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            note: {},
            error: ''
        };
    },
    created() {
        this.fetchNote();
    },
    methods: {
        async fetchNote() {
            const url = this.$route.params.url;
            try {
                const response = await axios.get(`http://localhost:8000/${url}`);
                this.note = response.data.data;
            } catch (error) {
                if (error.response && error.response.status === 410) {
                    this.error = 'this note has expired or reached its view limit.';
                } else {
                    this.error = 'record not found!';
                }
            }
        }
    }
};
</script>

<style>
body {
    font-family: Arial, sans-serif;
    background-color: #f0f0f0;
    margin: 0;
    padding: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.container {
    background-color: #ffffff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    padding: 20px;
    width: 400px;
    text-align: center;
}

h1 {
    color: #333333;
}

p {
    color: #666666;
}
</style>
