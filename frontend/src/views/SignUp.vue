<template>
    <div class="container">
      <h1>Sign Up</h1>
      <form @submit.prevent="signUp">
        <label for="name">Name:</label>
        <input v-model="name" id="name" type="text" placeholder="Enter your name" />
  
        <label for="password">Password:</label>
        <input v-model="password" id="password" type="password" placeholder="Enter your password" />
  
        <button type="submit">Sign Up</button>
      </form>
  
      <p>Already have an account? <router-link to="/signin">Sign In</router-link></p>
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
      async signUp() {
        try {
          const response = await axios.post('http://localhost:8000/signup', {
            name: this.name,
            password: this.password,
          });
          console.log('User signed up:', response.data);
          this.$router.push('/signin');
        } catch (error) {
          console.error('Error signing up:', error);
        }
      },
    },
  };
  </script>