<template>
    <div class="container">
      <h1>Sign In</h1>
      <form @submit.prevent="signIn">
        <label for="name">Name:</label>
        <input v-model="name" id="name" type="text" placeholder="Enter your name" />
  
        <label for="password">Password:</label>
        <input v-model="password" id="password" type="password" placeholder="Enter your password" />
  
        <button type="submit">Sign In</button>
      </form>
  
      <p>Don't have an account? <router-link to="/signup">Sign Up</router-link></p>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        name: '',
        password: '',
      };
    },
    methods: {
      async signIn() {
        try {
          const response = await axios.post('http://localhost:8000/signin', {
            name: this.name,
            password: this.password,
          });
          console.log('User signed in:', response.data);
  
          localStorage.setItem('user', JSON.stringify(response.data));
  
          this.$router.push('/create');
        } catch (error) {
          console.error('Error signing in:', error);
        }
      },
    },
  };
  </script>
  