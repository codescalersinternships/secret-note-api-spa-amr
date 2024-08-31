<template>
    <div class="container">
      <h1>Create a Secret Note</h1>
      <form @submit.prevent="createNote">
        <label for="note_content">Note Content:</label>
        <textarea v-model="noteContent" id="note_content" name="content" placeholder="Enter your note here" rows="4"></textarea>
  
        <label for="max_views">Maximum Views:</label>
        <input v-model="maxViews" type="number" id="max_views" name="maxViews" placeholder="Enter maximum views" value="1">
  
        <label for="expiration">Expiration After:</label>
        <input v-model="expirationAfter" type="number" id="expiration_after" name="expireAfter" placeholder="Enter expiration time in hours" value="1">
  
        <button type="submit">Create Note</button>
      </form>
  
      <div v-if="noteURL" class="note-created">
        <h2>Note Created!</h2>
        <p>
          <router-link :to="noteURL">{{ noteURL }}</router-link>
        </p>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        noteContent: '',
        maxViews: 1,
        expirationAfter: 1,
        noteURL: ''
      };
    },
    methods: {
      async createNote() {
        try {
          const response = await axios.post('http://localhost:8000/create', {
            content: this.noteContent,
            max_views: this.maxViews,
            expire_after: this.expirationAfter,
          });
          console.log('Note created:', response.data);
          this.noteURL = `/${response.data.data.url}`;
        } catch (error) {
          console.error('Error creating note:', error);
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
}

h1 {
    text-align: center;
    color: #333333;
}

form {
    display: flex;
    flex-direction: column;
}

label {
    font-weight: bold;
    margin-bottom: 8px;
    color: #666666;
}

input[type="number"],
textarea {
    padding: 10px;
    margin-bottom: 16px;
    border: 1px solid #dddddd;
    border-radius: 4px;
    font-size: 14px;
}

button[type="submit"] {
    background-color: #007bff;
    color: #ffffff;
    border: none;
    padding: 12px 20px;
    border-radius: 4px;
    cursor: pointer;
    font-size: 16px;
}

button[type="submit"]:hover {
    background-color: #0056b3;
}

.note-created {
    margin-top: 20px;
    padding: 10px;
    border: 1px solid #ccc;
    background-color: #f9f9f9;
}
</style>
